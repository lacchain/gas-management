// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.6.0;

import "./IRelayHub.sol";
import "./lib/ECDSA.sol";
import "./lib/SafeMath.sol";
import "./GasLimit.sol";

contract TxRelay is GasLimit,IRelayHub{

    using ECDSA for bytes32;
    using SafeMath for uint256;

    // Note: This is a local nonce.
    // Different from the nonce defined w/in protocol.
    mapping(address => uint) nonces;

    constructor(uint8 _blocksFrequency) GasLimit(_blocksFrequency) public{}

    function relayMetaTx(
        address from,
        address to,
        bytes calldata encodedFunction,
        uint256 gasLimit,
        uint256 nonce,
        bytes calldata signature
    ) external evaluateCurrencyBlock override returns (bool success){
        require(gasLimit <= block.gaslimit, "Impossible gas limit");

        nonces[msg.sender]++; //if we are going to do tx, update nonce

        // This recreates the message hash that was signed on the client.
        bytes32 hash = keccak256(abi.encodePacked(from, to, keccak256(encodedFunction), gasLimit, nonce));

        bytes32 messageHash = hash.toEthSignedMessageHash();

        // Verify that the message's signer is the owner of the order
        address signer = messageHash.recover(signature);
        emit Relayed(msg.sender,signer);
        require(from == signer, "the sender isn't the signer");

        bytes memory _data = encodedFunction;

        uint256 gasUsed = gasleft();

        require(_hasEnoughGas(gasUsed),"Exceeded the allowed gas limit");

        _executeCall(to,0,_data);

        gasUsed = gasUsed.sub(gasleft());

        (uint256 newGasLimit,uint256 gasUsedLastBlocks) = _addGasUsed(gasUsed);

        emit GasUsedByTransaction(block.number, gasUsed, newGasLimit, gasUsedLastBlocks);

        return true;
    }

    function deployMetaTx(
        address from,
        bytes calldata _byteCode,
        uint256 gasLimit,
        uint256 nonce,
        bytes calldata signature
    ) external evaluateCurrencyBlock returns (bool success, address deployedAddress){
        require(gasLimit <= block.gaslimit, "Impossible gas limit");

        nonces[msg.sender]++; //if we are going to do tx, update nonce

        // This recreates the message hash that was signed on the client.
        bytes32 hash = keccak256(abi.encodePacked(from, keccak256(_byteCode), gasLimit, nonce));

        bytes32 messageHash = hash.toEthSignedMessageHash();

        // Verify that the message's signer is the owner of the order
        address signer = messageHash.recover(signature);
        emit Relayed(msg.sender,signer);
        require(from == signer, "the sender isn't the signer");

        bytes memory _data = _byteCode;

        uint256 gasUsed = gasleft();

        require(_hasEnoughGas(gasUsed),"Exceeded the allowed gas limit");

        deployedAddress = _doCreate(0,_data);

        gasUsed = gasUsed.sub(gasleft());

        emit ContractDeployed(deployedAddress);

        (uint256 newGasLimit,uint256 gasUsedLastBlocks) = _addGasUsed(gasUsed);

        emit GasUsedByTransaction(block.number, gasUsed, newGasLimit, gasUsedLastBlocks);

        return (true,deployedAddress);
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

    event Relayed(address indexed sender, address indexed from);
    event GasUsedByTransaction(uint blockNumber, uint gasUsed, uint gasLimit, uint gasUsedLastBlocks);
    event ContractDeployed(address contractDeployed);
}
