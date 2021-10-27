// SPDX-License-Identifier: UNLICENSED

pragma solidity >=0.8.0 <0.9.0;

import "./IRelayHub.sol";
import "./lib/ECDSA.sol";
import "./GasLimit.sol";
import "solidity-rlp/contracts/RLPReader.sol";
import "./security/ReentrancyGuard.sol";

contract TxRelay is ReentrancyGuard,GasLimit,IRelayHub{

    using ECDSA for bytes32;
    using RLPReader for RLPReader.RLPItem;
    using RLPReader for RLPReader.Iterator;
    using RLPReader for bytes;

    // Note: This is a local nonce.
    // Different from the nonce defined w/in protocol.
    //    writerNode => userAddress => nonce
    mapping(address => mapping (address => uint)) nonces;

    address msgSender;

    constructor(uint16 _blocksFrequency, address _accountIngress) GasLimit(_blocksFrequency,_accountIngress){
        _setupRole(DEFAULT_ADMIN_ROLE, msg.sender);
    }

    function relayMetaTx(
        uint256 gasLimit,
        bytes memory signingData, ////0xf8..RLP final user + nodeAddress + expiration
        uint8 v,
        bytes32 r,
        bytes32 s
    ) external nonReentrant evaluateCurrentBlock override returns (ErrorCode success){
        (ErrorCode errorCode, TransactionData memory data) = _verifyTransaction(signingData,v,r,s,false);

        if (errorCode!=ErrorCode.OK){
            if (!_increaseGasUsed(gasLimit - gasleft() + 6494)){  //6494 cost to set gasUsed
                return ErrorCode.NotEnoughGas;    
            }
            emit BadTransactionSent(msg.sender, data.from, errorCode);
            return errorCode;
        }else{
            if (_isContract(data.to)){
                msgSender = data.from;
                if (gasleft()>(data.gasLimit + 35000)){  //could be done in transaction permissioning, but decode payload to get gaslimit of user is expensive and we already have here
                    (bool executed, bytes memory output) = _executeCall(data.to,0,data.gasLimit+600,data.encodedFunction); //600 because 64 bytes are added as parameters (nodeAddress, expiration)

                    emit TransactionRelayed(msg.sender, data.from, data.to, executed, output);
                }else{
                    emit BadTransactionSent(msg.sender, data.from, ErrorCode.NotEnoughGas);
                    if (!_increaseGasUsed(gasLimit - gasleft() + 6494)){    //6494 cost to set gasUsed
                        return ErrorCode.NotEnoughGas;    
                    }
                    return ErrorCode.NotEnoughGas;
                }
            }else{
                emit BadTransactionSent(msg.sender, data.from, ErrorCode.IsNotContract);
            }
            if (!_increaseGasUsed(gasLimit - gasleft() + 6494)){   //6494 cost to set gasUsed
                return ErrorCode.NotEnoughGas;
            }
            return ErrorCode.OK;
        }
    }

    function deployMetaTx(
        uint256 gasLimit,
        bytes memory signingData,
        uint8 v,
        bytes32 r,
        bytes32 s
    ) external nonReentrant override evaluateCurrentBlock returns (ErrorCode success, address deployedAddress){
        (ErrorCode errorCode, TransactionData memory data) = _verifyTransaction(signingData,v,r,s,true);

        if (errorCode!=ErrorCode.OK){
            emit BadTransactionSent(msg.sender, data.from, errorCode);
            if (!_increaseGasUsed(gasLimit - gasleft() + 6494)){  //6494 cost to set gasUsed
                return (ErrorCode.NotEnoughGas,address(0));   
            }
            return (errorCode,address(0));
        }else{
            if (data.encodedFunction.length > 0){
                msgSender = data.from;
                if (gasleft()>(data.gasLimit + 70000)){ 
                    (uint8 isContractDeployed, address createdContract) = _doCreate(0,data.encodedFunction);
                    deployedAddress = createdContract;
                    if (isContractDeployed>0){
                        emit ContractDeployed(msg.sender, data.from, deployedAddress);
                    }
                } else{
                    emit BadTransactionSent(msg.sender, data.from, ErrorCode.NotEnoughGas);
                    if (!_increaseGasUsed(gasLimit - gasleft() + 6494)){   //6494 cost to set gasUsed
                        return (ErrorCode.NotEnoughGas, address(0));
                    }
                    return (ErrorCode.NotEnoughGas,address(0));
                }
            }else{
                emit BadTransactionSent(msg.sender, data.from, ErrorCode.EmptyCode);
            }
            if (!_increaseGasUsed(gasLimit - gasleft() + 6494)){   //6494 cost to set gasUsed
                return (ErrorCode.NotEnoughGas,address(0));
            }
            
            return (ErrorCode.OK,deployedAddress);
        }
    }

    function _verifyTransaction(
        bytes memory signingData,
        uint8 v,
        bytes32 r,
        bytes32 s,
        bool isContractDeploy) private returns (ErrorCode, TransactionData memory){    
        TransactionData memory data;
        
        (bool success, address signer) = _getOriginalSender(signingData,v,r,s);

        if (!success){
            data.from = address(0);
            return (ErrorCode.InvalidSignature,data);
        }
        data.from = signer;
        (data.nonce,data.gasLimit, data.to, data.encodedFunction) = _getParametersTx(signingData,isContractDeploy);
        
        //emit Parameters(data.nonce, data.gasLimit, data.to, data.encodedFunction);
        if (data.nonce != nonces[msg.sender][data.from]){
            return (ErrorCode.BadNonce,data);
        }

        if (!_destinationPermitted(data.to)){
            return (ErrorCode.InvalidDestination,data);
        }

        if (data.gasLimit > maxGasBlockLimit){
            return (ErrorCode.MaxBlockGasLimit,data);
        }

        if(!_hasEnoughGas(data.gasLimit)){
            return (ErrorCode.NotEnoughGas, data);
        }

        emit Relayed(data.from,msg.sender);
        nonces[msg.sender][data.from]++; //if we are going to do tx, update nonce
        
        return (ErrorCode.OK, data);
    }

    /**
     * @param _value amount ether to send
     * @param _code contract code
     */
    function _doCreate(uint _value, bytes memory _code) internal returns (uint8 success, address createdContract){
        uint8 contractDeployed;
        // solium-disable-next-line security/no-inline-assembly
        assembly {
            createdContract := create(_value, add(_code, 0x20), mload(_code))
            contractDeployed := gt(extcodesize(createdContract),0)
        }
        return (contractDeployed, createdContract);
    }

    /**
     * @param _untrustedRecipient contract destination
     * @param _value amount ether to send
     * @param _data transaction to send
     */
    function _executeCall(address _untrustedRecipient, uint256 _value, uint256 _gasLimit, bytes memory _data) private returns (bool success, bytes memory output){
        (success, output) = _untrustedRecipient.call{gas:_gasLimit, value:_value}(_data);
    }

    function getNonce(address from) external override view returns (uint256){
         return nonces[msg.sender][from];
    }

    function getMsgSender() external view returns (address){
        return msgSender;
    }

    function _isContract(address _addr) private view returns (bool isContract){
        uint32 size;
        assembly {
            size := extcodesize(_addr)
        }
        return (size > 0);
    }

    function increaseGasUsed(uint256 gasUsed) nonReentrant evaluateCurrentBlock onlyAccountPermissioned external returns (bool){
        return _increaseGasUsed(gasUsed);
    }

    function _increaseGasUsed(uint256 gasUsed) private returns (bool) {
        (uint256 newGasLimit,uint256 gasUsedLastBlocks) = _addGasUsed(msg.sender, gasUsed);
        emit GasUsedByTransaction(msg.sender, block.number, gasUsed, newGasLimit, gasUsedLastBlocks);
        if (newGasLimit == 0){
            return false;
        }
        return true;
    }

    function _getOriginalSender(bytes memory signingData, uint8 v, bytes32 r, bytes32 s)private pure returns(bool, address){
        bytes32 digest = keccak256(signingData);
        return digest.recover(v,r,s);
    }

    function _destinationPermitted(address _to)private view returns(bool){
        if (_to == trustedAccountIngress){
            return false;
        }
        return true;
    }

    function _getParametersTx(bytes memory signingData, bool isDeployContract) private pure returns (uint256, uint256, address, bytes memory){    //    bytes memory ddd = "0x6057361d0000000000000000000000000000000000000000000000000000000000000057";
        RLPReader.RLPItem[] memory ls = signingData.toRlpItem().toList();
        if (!isDeployContract){
            return (ls[0].toUint(),ls[2].toUint(),ls[3].toAddress(),ls[5].toBytes());
        }else{
            return (ls[0].toUint(),ls[2].toUint(),address(0),ls[5].toBytes());
        }
    //    return (0,500000,address(0xf9F76f30dcA57501e132d43D26FDab10A313CcE8),ddd);
    }

    event Relayed(address indexed sender, address indexed from);
    event GasUsedByTransaction(address node, uint blockNumber, uint gasUsed, uint gasLimit, uint gasUsedLastBlocks);
    event ContractDeployed(address indexed relay, address indexed from, address contractDeployed);
    event BadTransactionSent(address node, address originalSender, ErrorCode errorCode);
    //TODO REMOVE
    event Parameters(uint256 nonce, uint256 gasLimit, address to, bytes decodedFunction);
//    event GasSigner(uint256 gasSig);
//    event GasPara(uint256 gasPar);
//    event GasTotal(uint256 gasTotal);
}
