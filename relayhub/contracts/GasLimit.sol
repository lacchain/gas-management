// SPDX-License-Identifier: UNLICENSED

pragma solidity >=0.8.0 <0.9.0;

import "./access/AccessControl.sol";

contract GasLimit is AccessControl{

//    bytes32 public constant ACCOUNT_CONTRACT_ROLE = keccak256("ACCOUNT_CONTRACT_ROLE");
    uint256 internal maxGasBlockLimit = 150000000;
    uint256 internal gasUsedRelayHub = 300000;
    uint256 internal basicGasLimit = 500000;
    uint256 internal standardGasLimit = 1000000;

    address[] private writerNodes;

    address[] private premiumNodes;

    uint256 private blockNumber;

    uint256 private currentGasLimit;

    mapping(address => uint) private gasLimits;

    mapping (address => uint256) private indexOf; //1 based indexing. 0 means non-existent

    mapping (address => uint8) private typeNodes; //1-basic,2-standard,3-premium

    uint256 private gasUsedLastBlocks;

    uint256 private averageLastBlocks;

    uint256 private blockCalculateExecuted;

    uint16 private blocksFrequency;

    bool private nodeAdded;

    address internal trustedAccountIngress;
    
    mapping(address => bool) private nodeWasPenalized;

    mapping(address => uint256) private blockLastTranx;

    constructor(uint16 _blocksFrequency, address _accountIngress){
        currentGasLimit = maxGasBlockLimit;
        blockNumber = block.number;
        blockCalculateExecuted = block.number;
        blocksFrequency = _blocksFrequency;
        trustedAccountIngress = _accountIngress;
        _setupRole(DEFAULT_ADMIN_ROLE, msg.sender);
    }

    modifier evaluateCurrentBlock () {
        if (block.number > blockLastTranx[msg.sender]){
            blockLastTranx[msg.sender] = block.number;
            if (typeNodes[msg.sender]==1){
                gasLimits[msg.sender] = basicGasLimit;    
            }else if(typeNodes[msg.sender]==2){
                gasLimits[msg.sender] = standardGasLimit;    
            }else{
                gasLimits[msg.sender] = currentGasLimit;  //dependiendo del tipo de nodo setear el gas limite
            }
        }
        if (premiumNodes.length > 0){ 
            if (nodeAdded || recalcuteGasLimit()){
                emit Recalculated(true);
                calculateGasLimit();
            }
        }
        _;
    }

    function setBlocksFrequency(uint8 _blocksFrequency) onlyAdmin external evaluateCurrentBlock{
        blocksFrequency = _blocksFrequency;
        emit BlockFrequencyChanged(msg.sender, blocksFrequency);
    }

    function setBasicGasLimit(uint256 _basicGasLimit) onlyAdmin external evaluateCurrentBlock{
        basicGasLimit = _basicGasLimit;
        emit BasicGasLimitChanged(msg.sender, basicGasLimit);
    }

    function setStandardGasLimit(uint256 _standardGasLimit) onlyAdmin external evaluateCurrentBlock{
        standardGasLimit = _standardGasLimit;
        emit StandardGasLimitChanged(msg.sender, standardGasLimit);
    }

    function _hasEnoughGas(uint256 gas) internal view returns (bool){
        if (gas <= gasLimits[msg.sender])
            return true;
        else
            return false;
    }

    function calculateGasLimit() internal  {
        uint256 newGasLimit = (80*maxGasBlockLimit)/100;
        if (averageLastBlocks<=((20*maxGasBlockLimit)/100)){
            newGasLimit = 5*(maxGasBlockLimit/premiumNodes.length);
        }else if(averageLastBlocks<=((40*maxGasBlockLimit)/100)){
            newGasLimit = 4*(maxGasBlockLimit/premiumNodes.length);
        }else if(averageLastBlocks<=((60*maxGasBlockLimit)/100)){
            newGasLimit = 3*(maxGasBlockLimit/premiumNodes.length);
        }else if(averageLastBlocks<=((80*maxGasBlockLimit)/100)){
            newGasLimit = 2*(maxGasBlockLimit/premiumNodes.length);
        }else{
            newGasLimit = maxGasBlockLimit/premiumNodes.length;
        }

        if (newGasLimit > (80*maxGasBlockLimit)/100){
            newGasLimit = (80*maxGasBlockLimit)/100;
        }
        _setGasLimit(gasUsedLastBlocks,newGasLimit);
    }

    function _setGasLimit(uint256 _gasUsedLastBlocks, uint256 _newGasLimit) private {
         for (uint16 id = 0; id < premiumNodes.length; id++){
            gasLimits[writerNodes[id]] = _newGasLimit;
         }
         currentGasLimit = _newGasLimit;
         gasUsedLastBlocks = 0;
         blockCalculateExecuted = block.number;
         nodeAdded = false;
         emit GasLimitSet(block.number,_gasUsedLastBlocks,averageLastBlocks,_newGasLimit);
    }

    function getGasLimit() external view returns (uint256){
        require(exists(msg.sender), "Node is not registered");
        return gasLimits[msg.sender];
    }

    function _addGasUsed(address _sender, uint256 gasUsed) internal returns (uint256,uint256){
        if (gasUsed > gasLimits[_sender]){
            gasLimits[_sender] = 0;
            gasUsedLastBlocks = gasUsedLastBlocks + gasUsed;
            _penalizeNode();
        }else{
            gasLimits[_sender] = gasLimits[_sender] - gasUsed;
            gasUsedLastBlocks = gasUsedLastBlocks + gasUsed;
        }

        return (gasLimits[_sender],gasUsedLastBlocks);
    }

    function getGasUsedLastBlocks() external view returns (uint256){
        return gasUsedLastBlocks;
    }

    function setGasUsedLastBlocks(uint256 newGasUsed) onlyAdmin external {
        gasUsedLastBlocks = newGasUsed;
    }

    function recalcuteGasLimit() internal returns (bool){
        if (block.number == blockNumber){
            return false;
        }
        uint blocks = block.number - blockCalculateExecuted;
        if (blocks == blocksFrequency){
            averageLastBlocks = gasUsedLastBlocks / blocksFrequency;
            return true;
        } else if (blocks > blocksFrequency){
            if ((blocks - blocksFrequency)>=blocksFrequency){
                gasUsedLastBlocks = 0;
                averageLastBlocks = 0;
                return true;
            }else{
                gasUsedLastBlocks = (blocks - blocksFrequency) * (gasUsedLastBlocks / blocks);
                averageLastBlocks = gasUsedLastBlocks / blocksFrequency;
                return true;
            }
        }
        return false;
    }

    function exists(address node) internal view returns (bool) {
        return indexOf[node] != 0;
    }

    function addNode(address newNode, uint8 nodeType) onlyAccountContract external returns(bool){
        if(indexOf[newNode] == 0){
            writerNodes.push(newNode);
            indexOf[newNode] = writerNodes.length;
            typeNodes[newNode] = nodeType;
            if (nodeType == 3){
                premiumNodes.push(newNode);
            }
            nodeAdded = true;
            emit NodeAdded(newNode);
        }
        return true;
    }

    function deleteNode(address node) onlyAccountContract external {
        uint index = indexOf[node];
        require(index>0, "Node doesn't exist");
        
        if (index > writerNodes.length) return;

        address lastNode = writerNodes[writerNodes.length - 1];
        writerNodes[index-1] = lastNode;
        
        indexOf[lastNode] = index;
        indexOf[node] = 0;

        delete writerNodes[writerNodes.length-1];
        writerNodes.pop();
        emit NodeDeleted(node);
    }

    function getNodes() external view returns (uint){
        return writerNodes.length;
    }

    function _penalizeNode() private {
        nodeWasPenalized[msg.sender] = true;
        bytes memory payload = abi.encodeWithSignature("removeAccount(address)",msg.sender);
        (bool cResponse, bytes memory result) = trustedAccountIngress.call(payload);
        if (cResponse){
            emit NodeBlocked(msg.sender, block.number);
        }
    }

    function getMaxGasBlockLimit() public view returns (uint256){
        return maxGasBlockLimit;
    }

    function getCurrentGasLimit() public view returns (uint256){
        return currentGasLimit;
    }

    function setMaxGasBlockLimit(uint256 _maxGasBlockLimit) onlyAdmin public {
        maxGasBlockLimit = _maxGasBlockLimit;
        emit MaxGasBlockLimitChanged(block.number, msg.sender, maxGasBlockLimit);
    }

    function setGasUsedRelayHub(uint256 _gasUsedRelayHub) onlyAdmin public {
        gasUsedRelayHub = _gasUsedRelayHub;
        emit GasUsedRelayHubChanged(msg.sender, gasUsedRelayHub);
    }

    function setAccounIngress(address _accountIngress) onlyAdmin public {
        trustedAccountIngress = _accountIngress;
        emit AccountIngressChanged(msg.sender, trustedAccountIngress);
    }

    modifier onlyAdmin(){
        require(hasRole(DEFAULT_ADMIN_ROLE, msg.sender), "Caller is not Admin");
        _;
    }

    modifier onlyAccountContract(){
        require(msg.sender == trustedAccountIngress, "Caller is not Account Contract");
        _;
    }

    modifier onlyAccountPermissioned(){
        require(exists(msg.sender));
        _;   
    }

    event NodeAdded(address newNode);
    event NodeDeleted(address oldNode);
    event AccountIngressChanged(address admin, address newAddress);
    event BlockFrequencyChanged(address admin, uint16 blocksFrequency);
    event BasicGasLimitChanged(address admin, uint256 newBasicGasLimit);
    event StandardGasLimitChanged(address admin, uint256 newStandardGasLimit);
    event MaxGasBlockLimitChanged(uint256 blockNumber, address admin, uint256 maxGasBlockLimit);
    event GasUsedRelayHubChanged(address admin, uint256 gasUsedRelayHub);
    event Recalculated(bool result);
    event GasLimitSet(uint256 blockNumber, uint256 gasUsedLastBlocks, uint256 averageLastBlocks, uint256 newGasLimit);
    event GasLimitExceeded(address node, uint256 blockNumber, uint8 countExceeded);
    event NodeBlocked(address node, uint256 blockNumber);
}
