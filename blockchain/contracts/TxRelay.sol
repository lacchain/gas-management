// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.6.0;

import "./IRelayHub.sol";
import "./lib/ECDSA.sol";
import "./lib/SafeMath.sol";

contract TxRelay is IRelayHub {
    
    using ECDSA for bytes32;
    using SafeMath for uint256;

    // Note: This is a local nonce.
    // Different from the nonce defined w/in protocol.
    mapping(address => uint) nonces;
    
    mapping(address => uint) gasLimits;

    function relayMetaTx(
        address from,
        address to,
        bytes calldata encodedFunction,
        uint256 gasLimit,
        uint256 nonce,
        bytes calldata signature
    ) override external returns (bool success){
        
        require(gasLimit <= block.gaslimit, "Impossible gas limit");
        
        // This recreates the message hash that was signed on the client.
        bytes32 hash = keccak256(abi.encodePacked(from, to, encodedFunction, gasLimit, nonce));
        
        bytes32 messageHash = hash.toEthSignedMessageHash();

        // Verify that the message's signer is the owner of the order
        address signer = messageHash.recover(signature);
        emit Relayed(msg.sender,signer);
        require(msg.sender == signer, "El sender no es el mismo");

        nonces[msg.sender]++; //if we are going to do tx, update nonce 
        
        bytes memory _data = encodedFunction;
        
        uint256 gasUsed = gasleft();
  
        require(gasUsed<=gasLimits[msg.sender],"Exceeded the allowed gas limit");
  
        _executeCall(to,0,_data);

        gasUsed = gasUsed.sub(gasleft());
        
        gasLimits[msg.sender] = gasLimits[msg.sender].sub(gasUsed);
     
        emit GasLimit(gasUsed,  gasLimits[msg.sender]);
        
        return true;
    }
    
    function _executeCall(address _to, uint256 _value, bytes memory _data) internal returns (bool success){
        // solium-disable-next-line security/no-inline-assembly
        assembly {
            success:= call(gas(), _to, _value, add(_data, 0x20), mload(_data), 0, 0)
        }
    }
    
    function getNonce(address from) override external view returns (uint256){
         return nonces[from];
    }
    
    function setGasLimit(address to, uint256 gasLimit) external returns (bool){
         gasLimits[to] = gasLimit;
         return true;
    }
    
    function getGasLimit() external view returns (uint256){
        return gasLimits[msg.sender];
    }
    
    event Relayed(address indexed sender, address indexed from);
    event Hashed(bytes32 hash);
    event GasUsed(uint256 gasUsed);
    event GasLimit(uint gasUsed, uint gasLimit);
    
}