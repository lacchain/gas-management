# Interacting with the LACChain Gas Model

In this section we will review the way to interacting with the Gas Model and its differences with traditional way of interacting with the node.

It is important to mention that the RelaySigner component will be running on top of Besu, which means that all transactions or queries sent to node will go through the RelaySigner. You can review the architecture for more knowledge in the following link [Gas Distribution Model Architecture](./Architecture.md)

If you prefer to go to the practical part of how to adapt your Dapp, you can check [How adapt your Dapp](How_adapt_your_Dapp.md)

## Smart Contract

### Pre-existing contracts

For these pre-existing contracts, it is necessary that their address be registered in the transaction permissioning contract, so that transactions towards these pre-existing contracts are allowed.

To interact with smart contracts that currently are running on the network is not necessary to make any changes. The interaction will be keep directly to the Besu node(rpc o websocket) without going through any of the components of the gas model(RelaySigner).

It is worth mentioning that these pre-existing contracts must be adapted to the gas model in order to function in a short time. With which we recommend seeing [How to adapt my contract to this new model](Architecture#how-to-adapt-my contract-to-this-new-model).

### New Smart Contracts

New smart contracts that will be deployed have to previously change the way to obtain the sender according to the documentation [How to adapt my contract to this new model](Architecture#how-to-adapt-my contract-to-this-new-model)

## About Nonce

### Before Gas Model

All accounts get the last pending or executed nonce to be able to send the next transaction using getTransactionCount function.

### After Gas Model deployed

All accounts will keep getting the nonce in traditional way, using [getTransactionCount](https://besu.hyperledger.org/en/stable/Reference/API-Methods/#eth_gettransactioncount), but once the gas model is deployed, all nonce will be reset to 0. This due to the RelayHub smart contract will be charge of managing the nonce.

## Deploy Smart Contract

### Before Gas Model

To deploy contracts organizations use different tools and clients such as truffle, web3, metamask, remix, openzeppelin cli, etc. In case of LACChain, when a new contract is deployed, a signed transaction is sent with the bytecode to be displayed using the `eth_sendRawTransaction` function connected directly to Besu.

In case the deployment is successful, the adress of the new deployed contract is obtained from the `contractAddress` parameter of the transaction receipt.

### After Gas Model deployed

Entities will be able to use the tools or clients to deploy contracts as normal. Instead of pointing directly to the besu it is necessary to point to the RelaySigner. They also have to add two parameters at the end of the bytecode, which are the address of the writer node through which the transactions will be sent and the expiration time (timestamp), according to the documentation [How send a transaction](Architecture.md#how-send-a-transaction)

The RelaySigner will be in charge of adding the organization's node signature and sending the transaction for the contract deployment, after the RelaySigner will be listening for events to verify that bytecode has been deployed. In case it has been deployed correctly, it will return an address of the contract deployed in the contractAddress parameter of the transaction receipt.

Additionally, if the organization wishes to verify the address of the contract deployed, This can be obtained by listening to `event ContractDeployed(address indexed relay, address indexed from, address contractDeployed)` from the RelayHub smart contract. In this event, the relay parameter is the node's address, from parameter is the original sender address that sent the transaction and contractDeployed parameter will be the address of the successfully deployed contract.

## Send a Transaction

### Before Gas Model

To send transactions organizations use different tools and clients such as truffle, web3, metamask, remix, openzeppelin cli, etc. In case of LACChain a signed transaction is sent using the `eth_sendRawTransaction` function connected directly to Besu.

In case the transaction was successful, the receipt of the transaction would return true in the status parameter and corresponding logs.

### After Gas Model Deployed

Entities will be able to use the tools or clients to deploy contracts as normal. Instead of pointing directly to the besu it is necessary to point to the RelaySigner. They also have to add two parameters at the end of the bytecode, which are the address of the writer node through which the transactions will be sent and the expiration time (timestamp), according to the documentation [How send a transaction](Architecture.md#how-send-a-transaction)

The RelaySigner will be in charge of adding the organization's node signature and sending the transaction to execute, after the RelaySigner will be listening for events to verify that the transaction has been executed successfully.

Since all transactions will go through the RelayHub smart contract and it acts as a proxy, the following scenarios may occur:

1. The transaction has not been forwarded to the recipient contract. In this case, the following event will have been emitted `event BadTransactionSent(address node, address originalSender, ErrorCode errorCode)`, which errorCode parameter might be one of the following causes:

    * MaxBlockGasLimit. It means that the transaction sent exceeded the maximum gas limit allowed.
    * BadNonce. It means that the nonce sent is not correct.
    * NodeNotAllowed. This error means that the node which sent the transaction is not allowed to send transactions to the network.
    * NotEnoughGas. This means that the node does not have enough gas to be able to execute the transaction.
    * IsNotContract. It means the recipient contract address does not have any code stored, therefore it is not a contract.

2. The transaction was forwarded to the recipient contract, but it throw an EVM error due to some contract restriction. In this case, status parameter of the transaction will return false and the EVM error message will be returned in revertReason parameter.

3. The transaction was forwarded to the recipient contract and executed successfully. In this case, the following event will be emitted `event TransactionRelayed(address indexed relay, address indexed from, address indexed to, bool executed, bytes output)` which parameters are:

    * relay: the address node's organization
    * from: account address that sent the transaction
    * to: recipient contract 
    * executed: true in case transaction was executed successfully and false otherwise
    * output: value returned by the contract after execution

## Private Transactions

### Before Gas Model

Private transactions are sent directly to the Besu node using the `eea_sendRawTransaction` and `priv_distributeRawTransaction` functions and returns a transaction hash to verify execution of the transaction.

### After Gas Model Deployed

Private transactions will have to be sent through the RelaySigner, but will not be forwarded to the RelayHub contract. This is because transactions can not be forwaded to the precompiled privacy contract from another contract (RelayHub in this case).

It is important to mention that in the same way as public transactions, the organization's node from which private transactions are sent must be registered in the account permissioning smart contract.

In this scenario, the Relay Signer will redirect the call to the Besu node for the private transaction to be executed and will return the transaction hash in the response.



 
