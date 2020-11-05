// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.6.0;

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

    constructor(uint8 _blocksFrequency, address _accountIngress) GasLimit(_blocksFrequency,_accountIngress) public{}

    function relayMetaTx(
        address from,
        address to,
        bytes memory encodedFunction,
        uint256 gasLimit,
        uint256 nonce,
        bytes calldata signature,
        bytes calldata senderSignature
    ) external evaluateCurrencyBlock override returns (bool success){
        if (gasLimit > block.gaslimit){
            _decreaseGasUsed(GASUSED_RELAYHUB);
            emit BadTransactionSent(msg.sender, from, ErrorCode.MaxBlockGasLimit);
            return false;
        }

        if (from != _getOriginalSender(to,encodedFunction,gasLimit,nonce,senderSignature)){
            _decreaseGasUsed(GASUSED_RELAYHUB);
            emit BadTransactionSent(msg.sender, from, ErrorCode.BadOriginalSender);
            return false;
        }

        nonces[msg.sender]++; //if we are going to do tx, update nonce

        // This recreates the message hash that was signed on the client.
        bytes32 hash = keccak256(abi.encodePacked(from, to, keccak256(encodedFunction), gasLimit, nonce));

        bytes32 messageHash = hash.toEthSignedMessageHash();

        // Verify that the message's signer is the owner of the order
        address signer = messageHash.recover(signature);
        emit Relayed(from,signer);
        if (msg.sender != signer){
            _decreaseGasUsed(GASUSED_RELAYHUB);
            emit BadTransactionSent(msg.sender, from, ErrorCode.BadNodeSigner);
            return false;
        }

        //bytes memory _data = encodedFunction;

        uint256 gasUsed = gasleft();

        if(!_hasEnoughGas(gasUsed)){
            _decreaseGasUsed(GASUSED_RELAYHUB);
            emit BadTransactionSent(msg.sender, from, ErrorCode.NotEnoughGas);
            return false;
        }

        if (_isContract(to)){
            msgSender = from;
            bool executed = _executeCall(to,0,encodedFunction);
            emit TransactionExecuted(executed);
            gasUsed = gasUsed.sub(gasleft());
            _decreaseGasUsed(gasUsed);
        }else{
            _decreaseGasUsed(GASUSED_RELAYHUB);
            emit BadTransactionSent(msg.sender, from, ErrorCode.IsNotContract);
            return false;
        }

        return true;
    }

    function deployMetaTx(
        address from,
        bytes calldata _byteCode,
        uint256 gasLimit,
        uint256 nonce,
        bytes calldata signature,
        bytes calldata senderSignature
    ) external evaluateCurrencyBlock returns (bool success, address deployedAddress){
        if (gasLimit > block.gaslimit){
            _decreaseGasUsed(GASUSED_RELAYHUB);
            return (false, address(0));
        }

        if (from != _getOriginalSender(address(0),_byteCode,gasLimit,nonce,senderSignature)){
            _decreaseGasUsed(GASUSED_RELAYHUB);
            emit BadTransactionSent(msg.sender, from, ErrorCode.BadOriginalSender);
            return (false, address(0));
        }

        nonces[msg.sender]++; //if we are going to do tx, update nonce

        // This recreates the message hash that was signed on the client.
        bytes32 hash = keccak256(abi.encodePacked(from, keccak256(_byteCode), gasLimit, nonce));

        bytes32 messageHash = hash.toEthSignedMessageHash();

        // Verify that the message's signer is the owner of the order
        address signer = messageHash.recover(signature);
        emit Relayed(msg.sender,signer);
        if (msg.sender != signer){
            _decreaseGasUsed(GASUSED_RELAYHUB);
            emit BadSigner(true);
            return (false,address(0));
        }

        //bytes memory _data = _byteCode;

        uint256 gasUsed = gasleft();

        if(!_hasEnoughGas(gasUsed)){
            _decreaseGasUsed(GASUSED_RELAYHUB);
            return (false,address(0));
        }

        if (_byteCode.length > 0){
            msgSender = from;
            deployedAddress = _doCreate(0,_byteCode);
            gasUsed = gasUsed.sub(gasleft());
            emit ContractDeployed(deployedAddress);
            _decreaseGasUsed(gasUsed);
            return (true,deployedAddress);
        }else{
            _decreaseGasUsed(GASUSED_RELAYHUB);
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
    function _executeCall(address _to, uint256 _value, bytes memory _data) private returns (bool success){
        // solium-disable-next-line security/no-inline-assembly
        assembly {
            success:= call(gas(), _to, _value, add(_data, 0x20), mload(_data), 0, 0)
        }
    }

    function getNonce(address from) external override view returns (uint256){
         return nonces[from];
    }

    function getMsgSender() external view returns (address){
        return msgSender;
    }

    event Relayed(address indexed sender, address indexed from);
    event GasUsedByTransaction(address node, uint blockNumber, uint gasUsed, uint gasLimit, uint gasUsedLastBlocks);
    event ContractDeployed(address contractDeployed);

    function _isContract(address _addr) private view returns (bool isContract){
        uint32 size;
        assembly {
            size := extcodesize(_addr)
        }
        return (size > 0);
    }

    function _decreaseGasUsed(uint256 gasUsed) private {
        (uint256 newGasLimit,uint256 gasUsedLastBlocks) = _addGasUsed(gasUsed);
        emit GasUsedByTransaction(msg.sender, block.number, gasUsed, newGasLimit, gasUsedLastBlocks);
    }

    function _getOriginalSender(address to, bytes memory encodedFunction, uint256 gasLimit, uint256 nonce, bytes memory signature)private pure returns(address){
        bytes[] memory rawTx =  new bytes[](6);
        rawTx[0]=RLPEncode.encodeUint(nonce);
        rawTx[1]=RLPEncode.encodeUint(0); //gasPrice
        rawTx[2]=RLPEncode.encodeUint(gasLimit);
        rawTx[3]=RLPEncode.encodeAddress(to);
        rawTx[4]=RLPEncode.encodeUint(0); //value
        rawTx[5]=RLPEncode.encodeBytes(encodedFunction);

        bytes32 digest = keccak256(RLPEncode.encodeList(rawTx));
        address sender = digest.recover(signature);
        return sender;
    }

    event BadTransactionSent(address node, address originalSender, ErrorCode errorCode);

    //TODO delete
    event TransactionExecuted(bool executed);

    //TODO delete
    event BadSigner(bool isSigner);
}
