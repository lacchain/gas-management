// SPDX-License-Identifier: UNLICENSED

pragma solidity >=0.6.0 <0.7.0;

import "./IRelayHub.sol";
import "./lib/ECDSA.sol";
import "./lib/SafeMath.sol";
import "./lib/RLPEncode.sol";
import "./GasLimit.sol";

contract TxRelay is GasLimit,IRelayHub{

    using ECDSA for bytes32;
    using SafeMath for uint256;

    // Note: This is a local nonce.
    // Different from the nonce defined w/in protocol.
    mapping(address => uint) nonces;

    address msgSender;

    constructor(uint8 _blocksFrequency, address _accountIngress) GasLimit(_blocksFrequency,_accountIngress) public{
    //    _setupRole(DEFAULT_ADMIN_ROLE, msg.sender);
    }

    function relayMetaTx(
        address to,
        bytes calldata encodedFunction,
        uint256 gasLimit,
        uint256 nonce,
        bytes calldata senderSignature
    ) external evaluateCurrencyBlock override returns (bool success){
        address from = _getOriginalSender(to,encodedFunction,gasLimit,nonce,senderSignature,false);
        nonces[from]++; //if we are going to do tx, update nonce

        if (nonce != nonces[from] -1){
            _decreaseGasUsed(gasUsedRelayHub);
            emit BadTransactionSent(msg.sender, from, ErrorCode.BadNonce);
            return false;
        }

        if (gasLimit > maxGasBlockLimit){
            _decreaseGasUsed(gasUsedRelayHub);
            emit BadTransactionSent(msg.sender, from, ErrorCode.MaxBlockGasLimit);
            return false;
        }

        if (!exists(msg.sender)){
            _decreaseGasUsed(gasUsedRelayHub);
            emit BadTransactionSent(msg.sender, from, ErrorCode.NodeNotAllowed);
            return false;
        }

        emit Relayed(from,msg.sender);

        uint256 gasUsed = gasleft();

        if(!_hasEnoughGas(gasUsed)){
            _decreaseGasUsed(gasUsedRelayHub);
            emit BadTransactionSent(msg.sender, from, ErrorCode.NotEnoughGas);
            return false;
        }

        if (_isContract(to)){
            msgSender = from;
            (bool executed, bytes memory output) = _executeCall(to,0,encodedFunction);
            emit TransactionRelayed(msg.sender, from, to, executed, output);
            gasUsed = gasUsed.sub(gasleft());
            _decreaseGasUsed(gasUsed);
        }else{
            _decreaseGasUsed(gasUsedRelayHub);
            emit BadTransactionSent(msg.sender, from, ErrorCode.IsNotContract);
            return false;
        }

        return true;
    }

    function deployMetaTx(
        bytes calldata _byteCode,
        uint256 gasLimit,
        uint256 nonce,
        bytes calldata senderSignature
    ) external evaluateCurrencyBlock returns (bool success, address deployedAddress){
        address from = _getOriginalSender(address(0),_byteCode,gasLimit,nonce,senderSignature,true);
        nonces[from]++; //if we are going to do tx, update nonce

        if (nonce != nonces[from] -1){
            _decreaseGasUsed(gasUsedRelayHub);
            emit BadTransactionSent(msg.sender, from, ErrorCode.BadNonce);
            return (false, address(0));
        }

        if (gasLimit > maxGasBlockLimit){
            _decreaseGasUsed(gasUsedRelayHub);
             emit BadTransactionSent(msg.sender, from, ErrorCode.MaxBlockGasLimit);
            return (false, address(0));
        }

        if (!exists(msg.sender)){
            _decreaseGasUsed(gasUsedRelayHub);
            emit BadTransactionSent(msg.sender, from, ErrorCode.NodeNotAllowed);
            return (false, address(0));
        }

        emit Relayed(from,msg.sender);

        uint256 gasUsed = gasleft();

        if(!_hasEnoughGas(gasUsed)){
            _decreaseGasUsed(gasUsedRelayHub);
            emit BadTransactionSent(msg.sender, from, ErrorCode.NotEnoughGas);
            return (false,address(0));
        }

        if (_byteCode.length > 0){
            msgSender = from;
            deployedAddress = _doCreate(0,_byteCode);
            gasUsed = gasUsed.sub(gasleft());
            emit ContractDeployed(msg.sender, from, deployedAddress);
            _decreaseGasUsed(gasUsed);
            return (true,deployedAddress);
        }else{
            _decreaseGasUsed(gasUsedRelayHub);
            emit BadTransactionSent(msg.sender, from, ErrorCode.EmptyCode);
            return (false,address(0));
        }
    }

    /**
     * @param _value amount ether to send
     * @param _code contract code
     */
    function _doCreate(uint _value, bytes memory _code) internal returns (address createdContract){
        uint contractDeployed;
        // solium-disable-next-line security/no-inline-assembly
        assembly {
            createdContract := create(_value, add(_code, 0x20), mload(_code))
            contractDeployed := gt(extcodesize(createdContract),0)
        }
        require(contractDeployed>0,"Failed");
        return createdContract;
    }

    /**
     * @param _to contract destination
     * @param _value amount ether to send
     * @param _data transaction to send
     */
    function _executeCall(address _to, uint256 _value, bytes memory _data) private returns (bool success, bytes memory output){
        // solium-disable-next-line security/no-inline-assembly
        /*assembly {
            success:= call(gas(), _to, _value, add(_data, 0x20), mload(_data), 0, 0)
            let size := returndatasize()
            output:= returndatacopy(add(_data, 0x20), 0, size)
        }*/
        (success, output) = _to.call{gas:gasleft(), value:_value}(_data);
    }

    function getNonce(address from) external override view returns (uint256){
         return nonces[from];
    }

    function getMsgSender() external view returns (address){
        return msgSender;
    }

    event Relayed(address indexed sender, address indexed from);
    event GasUsedByTransaction(address node, uint blockNumber, uint gasUsed, uint gasLimit, uint gasUsedLastBlocks);
    event ContractDeployed(address indexed relay, address indexed from, address contractDeployed);
    event BadTransactionSent(address node, address originalSender, ErrorCode errorCode);

    function _isContract(address _addr) private view returns (bool isContract){
        uint32 size;
        assembly {
            size := extcodesize(_addr)
        }
        return (size > 0);
    }

    function decreaseGasUsed(uint256 gasUsed) evaluateCurrencyBlock onlyAccountPermissioned external returns (bool){
        _decreaseGasUsed(gasUsed);
    }

    function _decreaseGasUsed(uint256 gasUsed) private {
        (uint256 newGasLimit,uint256 gasUsedLastBlocks) = _addGasUsed(msg.sender, gasUsed);
        emit GasUsedByTransaction(msg.sender, block.number, gasUsed, newGasLimit, gasUsedLastBlocks);
    }

    function _getOriginalSender(address to, bytes memory encodedFunction, uint256 gasLimit, uint256 nonce, bytes memory signature, bool deployContract)private pure returns(address){
        bytes[] memory rawTx =  new bytes[](6);
        rawTx[0]=RLPEncode.encodeUint(nonce);
        rawTx[1]=RLPEncode.encodeUint(0); //gasPrice
        rawTx[2]=RLPEncode.encodeUint(gasLimit);
        if (deployContract){
            rawTx[3]=RLPEncode.encodeUint(0);
        }else{
            rawTx[3]=RLPEncode.encodeAddress(to);
        }
        rawTx[4]=RLPEncode.encodeUint(0); //value
        rawTx[5]=RLPEncode.encodeBytes(encodedFunction);

        bytes32 digest = keccak256(RLPEncode.encodeList(rawTx));
        address sender = digest.recover(signature);
        return sender;
    }
}
