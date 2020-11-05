// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.6.0;

import "./lib/SafeMath.sol";

contract GasLimit{

    using SafeMath for uint256;

    uint256 constant MAX_GASBLOCK_LIMIT = 20000000;//200000000;
    uint256 constant GASUSED_RELAYHUB = 200000;
    
    mapping(address => uint8) private countExceeded; 

    address[] private writerNodes;

    uint256 private blockNumber;

    uint256 private currentGasLimit;

    mapping(address => uint) private gasLimits;

    uint256 private gasUsedLastBlocks;

    uint256 private averageLastBlocks;

    uint256 private blockCalculateExecuted;

    uint8 private blocksFrequency;

    bool private nodeAdded;

    address private accountIngress;
    
    mapping(address => bool) private nodeWasPenalized;

    mapping(address => uint256) private blockLastTranx;

    constructor(uint8 _blocksFrequency, address _accountIngress) public{
        blockNumber = block.number;
        blockCalculateExecuted = block.number;
        blocksFrequency = _blocksFrequency;
        accountIngress = _accountIngress;
    }

    modifier evaluateCurrencyBlock () {
        if (block.number > blockLastTranx[msg.sender]){
            blockLastTranx[msg.sender] = block.number;
            gasLimits[msg.sender] = currentGasLimit;
            nodeWasPenalized[msg.sender] = false;
        }
        if (recalcuteGasLimit() || nodeAdded){
            emit Recalculated(true);
            calculateGasLimit();
        }
        _;
    }

    function setBlocksFrequency(uint8 _blocksFrequency) external evaluateCurrencyBlock{
        blocksFrequency = _blocksFrequency;
    }

    function _hasEnoughGas(uint256 gas) internal view returns (bool){
    /*    if (block.number > blockNumber){
            blockNumber = block.number;
            gasLimits[msg.sender] = currentGasLimit;
        }*/

        if (gas <= gasLimits[msg.sender])
            return true;
        else
            return false;
    }

    function calculateGasLimit() internal  {
        uint256 newGasLimit = (80*MAX_GASBLOCK_LIMIT)/100;
        if (averageLastBlocks<=((20*MAX_GASBLOCK_LIMIT)/100)){
            newGasLimit = 5*(MAX_GASBLOCK_LIMIT/writerNodes.length);
        }else if(averageLastBlocks<=((40*MAX_GASBLOCK_LIMIT)/100)){
            newGasLimit = 4*(MAX_GASBLOCK_LIMIT/writerNodes.length);
        }else if(averageLastBlocks<=((60*MAX_GASBLOCK_LIMIT)/100)){
            newGasLimit = 3*(MAX_GASBLOCK_LIMIT/writerNodes.length);
        }else if(averageLastBlocks<=((80*MAX_GASBLOCK_LIMIT)/100)){
            newGasLimit = 2*(MAX_GASBLOCK_LIMIT/writerNodes.length);
        }else{
            newGasLimit = MAX_GASBLOCK_LIMIT/writerNodes.length;
        }

        if (newGasLimit > (80*MAX_GASBLOCK_LIMIT)/100){
            newGasLimit = (80*MAX_GASBLOCK_LIMIT)/100;
        }
        _setGasLimit(gasUsedLastBlocks,newGasLimit);
    }

    function _setGasLimit(uint256 _gasUsedLastBlocks, uint256 _newGasLimit) private {
         for (uint16 id = 0; id < writerNodes.length; id++){
            gasLimits[writerNodes[id]] = _newGasLimit;
         }
         currentGasLimit = _newGasLimit;
         gasUsedLastBlocks = 0;
         blockCalculateExecuted = block.number;
         nodeAdded = false;
         emit GasLimitSet(block.number,_gasUsedLastBlocks,averageLastBlocks,_newGasLimit);
    }

    function getGasLimit() external view returns (uint256){
        return gasLimits[msg.sender];
    }

    function _addGasUsed(uint256 gasUsed) internal returns (uint256,uint256){
        if (gasUsed > gasLimits[msg.sender]){
            gasLimits[msg.sender] = 0;
            gasUsedLastBlocks = gasUsedLastBlocks.add(gasUsed);
            _penalizeNode();
        }else{
            gasLimits[msg.sender] = gasLimits[msg.sender].sub(gasUsed);
            gasUsedLastBlocks = gasUsedLastBlocks.add(gasUsed);
        }

        return (gasLimits[msg.sender],gasUsedLastBlocks);
    }

    function getGasUsedLastBlocks() external view returns (uint256){
        return gasUsedLastBlocks;
    }

    function setGasUsedLastBlocks(uint256 newGasUsed) external {
        gasUsedLastBlocks = newGasUsed;
    }

    function recalcuteGasLimit() internal returns (bool){
        if (block.number == blockNumber){
            return false;
        }
        uint blocks = block.number.sub(blockCalculateExecuted);
        if (blocks == blocksFrequency){
            averageLastBlocks = gasUsedLastBlocks.div(blocksFrequency);
            return true;
        } else if (blocks > blocksFrequency){
            if (blocks.sub(blocksFrequency)>=blocksFrequency){
                gasUsedLastBlocks = 0;
                averageLastBlocks = 0;
                return true;
            }else{
                gasUsedLastBlocks = blocks.sub(blocksFrequency).mul(gasUsedLastBlocks.div(blocks));
                averageLastBlocks = gasUsedLastBlocks.div(blocksFrequency);
                return true;
            }
        }
        return false;
    }

    function addNode(address newNode) external returns(bool){
        if(gasLimits[newNode]==0){
            writerNodes.push(newNode);
            nodeAdded = true;
        }
        return true;
    }

    function deleteNode(uint16 index) external {
        if (index >= writerNodes.length) return;

        for (uint i = index; i<writerNodes.length-1; i++){
            writerNodes[i] = writerNodes[i+1];
        }
        delete writerNodes[writerNodes.length-1];
        writerNodes.pop();
    }

    function getNodes() external view returns (uint){
        return writerNodes.length;
    }

    function _penalizeNode() private {
        if (!nodeWasPenalized[msg.sender]){
            countExceeded[msg.sender]++;
            nodeWasPenalized[msg.sender] = true;
        }
        if (countExceeded[msg.sender]<=2){
            emit GasLimitExceeded(msg.sender, block.number, countExceeded[msg.sender]);
        }else{
            bytes memory payload = abi.encodeWithSignature("removeAccount(address)",msg.sender);
            (bool cResponse, bytes memory result) = accountIngress.call(payload);
            countExceeded[msg.sender]=0;
            emit NodeBlocked(msg.sender, block.number);
        }
    }

    function setAccounIngress(address _accountIngress) public {
        accountIngress = _accountIngress;
    }

    /*function getTransactionss(address _address) public view returns(uint8,uint8,uint256){
        return (tranxSent[_address],successfulTranx[_address],totalTranxLastBlocks);
    }*/

    event Recalculated(bool result);
    event GasLimitSet(uint256 blockNumber, uint256 gasUsedLastBlocks, uint256 averageLastBlocks, uint256 newGasLimit);
    event GasLimitExceeded(address node, uint256 blockNumber, uint8 countExceeded);
    event NodeBlocked(address node, uint256 blockNumber);
}