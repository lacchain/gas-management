// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.6.0;

import "./lib/SafeMath.sol";

contract GasLimit{

    using SafeMath for uint256;

    uint256 constant MAX_GASBLOCK_LIMIT = 200000000;

    address[] private writerNodes;

    uint256 private blockNumber;

    uint256 private currentGasLimit;

    mapping(address => uint) private gasLimits;

    uint256 private gasUsedLastBlocks;

    uint256 private blockCalculateExecuted;

    uint8 private blocksFrequency;

    constructor(uint8 _blocksFrequency) public{
        blockNumber = block.number;
        blockCalculateExecuted = block.number;
        blocksFrequency = _blocksFrequency;
    }

    modifier evaluateCurrencyBlock () {
        if (recalcuteGasLimit()){
            calculateGasLimit();
        }
        _;
    }

    function setBlocksFrequency(uint8 _blocksFrequency) external evaluateCurrencyBlock{
        blocksFrequency = _blocksFrequency;
    }

    function _hasEnoughGas(uint256 gas) internal returns (bool){
        if (block.number > blockNumber){
            blockNumber = block.number;
            gasLimits[msg.sender] = currentGasLimit;
        }

        if (gas <= gasLimits[msg.sender])
            return true;
        else
            return false;
    }

    function calculateGasLimit() internal  {
        uint256 newGasLimit = (80*MAX_GASBLOCK_LIMIT)/100;
        if (gasUsedLastBlocks<=((20*MAX_GASBLOCK_LIMIT)/100)){
            newGasLimit = 5*(MAX_GASBLOCK_LIMIT/writerNodes.length);
        }else if(gasUsedLastBlocks<=((40*MAX_GASBLOCK_LIMIT)/100)){
            newGasLimit = 4*(MAX_GASBLOCK_LIMIT/writerNodes.length);
        }else if(gasUsedLastBlocks<=((60*MAX_GASBLOCK_LIMIT)/100)){
            newGasLimit = 3*(MAX_GASBLOCK_LIMIT/writerNodes.length);
        }else if(gasUsedLastBlocks<=((80*MAX_GASBLOCK_LIMIT)/100)){
            newGasLimit = 2*(MAX_GASBLOCK_LIMIT/writerNodes.length);
        }else{
            newGasLimit = MAX_GASBLOCK_LIMIT/writerNodes.length;
        }

        if (newGasLimit > (80*MAX_GASBLOCK_LIMIT)/100){
            newGasLimit = (80*MAX_GASBLOCK_LIMIT)/100;
        }
        _setGasLimit(newGasLimit);
    }

    function _setGasLimit(uint256 _newGasLimit) private {
         for (uint16 id = 0; id < writerNodes.length; id++){
            gasLimits[writerNodes[id]] = _newGasLimit;
         }
         currentGasLimit = _newGasLimit;
         gasUsedLastBlocks = 0;
    }

    function getGasLimit() external view returns (uint256){
        return gasLimits[msg.sender];
    }

    function _addGasUsed(uint256 gasUsed) internal returns (uint256,uint256){
        gasLimits[msg.sender] = gasLimits[msg.sender].sub(gasUsed);
        gasUsedLastBlocks = gasUsedLastBlocks.add(gasUsed);
        return (gasLimits[msg.sender],gasUsedLastBlocks);
    }

    function getGasUsedLastBlocks() external view returns (uint256){
        return gasUsedLastBlocks;
    }

    function setGasUsedLastBlocks(uint256 newGasUsed) external {
        gasUsedLastBlocks = newGasUsed;
    }

    function recalcuteGasLimit() internal returns (bool){
        uint blocks = block.number.sub(blockCalculateExecuted);
        if (blocks == blocksFrequency){
            return true;
        } else if (blocks > blocksFrequency){
            if (blocks.sub(blocksFrequency)>=blocksFrequency){
                gasUsedLastBlocks = 0;
                return true;
            }else{
                gasUsedLastBlocks = blocks.sub(blocksFrequency).mul(gasUsedLastBlocks.div(blocksFrequency));
                return true;
            }
        }
        return false;
    }

    function addNode(address newNode) external {
        writerNodes.push(newNode);
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
}