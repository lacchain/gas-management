# Solution Architecture

In this section we will review the different components of the architecture, function and relationship with the other components.

![Architecture](images/architecture.png)

## Backend Components

### Node
Is a writer node part of the LACChain network. The node is composed by Nginx, relaySigner layer and Hyperledger Besu.

### Nginx
Its role is act as reverse proxy and SSL termination (terminate HTTPS traffic from clients), it accepts all transactions to the node and forwards them to the RelaySigner. Direct queries to the node do not go through the RelaySigner, in this case nginx redirects them to the RPC port of Besu.

### RelaySigner Layer
Its role is to evaluate the type of message that arrives. In case it is a raw transaction, then the node generates a new transaction encapsulating the original transaction(RLP unsigned) in the data field of the new transaction and sends this new transaction to the network. If it is another type of message such as private transactions, get the recipient, block information or get transaction count then the message is directly forwarded to the Besu component.

### Besu
Their role is to be a client of the LACChain network, with which the p2p connection is maintained, broadcast transactions with the other nodes of the network. This component receive transactions from RelaySigner to broadcast to network.

## Smart Contracts

### Node Permission Smart Contract
This smart contract is deployed on the LACChain network. The contract allows connections between the nodes that are registered in the contract. 

Its role in the model is to allow the connection through the TCP protocol of the nodes in the network. Nodes that are not registered in this contract cannot connect to the permissioned network. It means that before two nodes make the connection, they previously verify against this contract if both are allowed to make the connection.

### Transaction Permission Smart Contract
This smart contract is deployed on the LACChain network. The role of this contract is to allow only transactions that satisfy the follow rules:
* Transactions have to be directed towards the RelayHub smart contract.
* Gas price of transaction must be 0.
* Gas limit of transaction must be sufficient to run the entire RelayHub.
* To transactions protected the sender of the transaction must be the same as indicated by the end user or application that generated the transaction.    

Transactions that do not satisfy these rules will be rejected.

### Relay Hub Smart Contract
This contract is based on EIP 1077. The contract receives a transaction whose data field is decoded by RLP to obtain the parameters of the original transaction, which are the address of the original sender, the gas limit, the nonce, the original data sent and the destination address of this data. 

Then verifies that address which the node of the organization has chosen to send transactions has not reached the gas limit for that block. In case everything is correct, the transaction is forwarded to the recipient contract or create a new smart contract.

### Gas Management Smart Contract
This contract keeps track of how much gas is consumed by each organization's address of the network. It also assigns a new maximum gas limit that each node's address can consume depending on the use of the network in the last blocks.

### Recipient Smart Contract
This contract will be the final destination of the transaction sent by an user or client application. It is the contract that executes the function chosen with parameters sent in the original transaction.

## Behavior

### Add New Node and sender addresses

![add_new_node](images/add_node_and_address.png)

When a new node is added to the network, it is added in the Node Permissioning Contract. Then the addresses of the senders of the organization, whose node was previously added to the permissioning smart contract, are added to the Transaction permissioning contract.

Then, this last transaction permissioning contract invokes the gas management contract which will set a GAS limit that new node's addresses can use.

### New Gas Address Limit

The frequency with which the new gas limit is set to each participating address of a writer node of the network is every 300 blocks (10 min), this is configurable by the permissioning committee.

![calculate_variables](images/calculate_variables.png)

The formula as the gas is distributed depends on how much GAS was used in the last blocks frequency (300 blocks, which is configurable). While higher the amount of Gas used among all the organizations, then lower the maximum GAS limit will be. It is very important to mention that the GAS is distributed in the same way among all the members of the network.

![calculate_continous](images/calculate_continous.png)

![calculate_discrete](images/calculate_discrete.png)

### Send Transaction

![send_transaction](images/send_transaction.png)

Each time a transaction is sent to the network, it is first verified against the transaction permissioning contract. The transaction is verified to satisfy the follow rules:

* Transactions have to be directed towards the RelayHub smart contract.
* Gas price of transaction must be 0.
* Gas limit of transaction must be sufficient to run the entire RelayHub.
* To transactions protected the sender of the transaction must be the same as indicated by the end user or application that generated the transaction.  

In case the transaction does not comply these rules then is canceled and will not be executed. 

After the transaction is verified and accepted, this transaction will go directly to the RelayHub contract where it is verified that the nonce sent is greater than the previous one to guarantee that the transaction is unique and does not repeat itself. Before the transaction is forwarded to a recipient contract, it is verified that the node has not consumed all the assigned GAS, in case the node has enough gas then the transaction is forwarded, after execution in the recipient contract the amount of GAS used is reduced for that address which sent the original transaction.

### Protected vs Simple Transaction

The GAS model has two types of transactions, protected and simple. 

Protected transactions are those that ensure that a transaction will only be executed if the writer designated by the end user or application co-signs the transaction and sends it to the network. In addition, the protected transaction has an expiration time in which it must be executed.

In another side, simple transactions do not have execution protection, they may be violated by malicious writer nodes or malicious users could use these types of transactions to ban innocent nodes that do not have protected transactions.

Writer node operators are advised to use protected transactions to avoid any type of attack.

### Private Transactions

![private_transactions](images/private_transactions.png)

When a private transaction is sent, it goes through the RelaySigner which call to Management Gas contrat to decrease the gas used, which is 25000. Then it redirects the transaction without modifying to the Besu process and this communicates with the Orion service which share the transaction with the participating nodes of the private transaction.

### Bad Transactions

![send_bad_transaction](images/send_bad_transaction.png)

The RelayHub emits a event when a bad or doubtful transaction is sent in following cases:
* The nonce sent is repeated or lower than registered in contract for that user's address.
* The gas limit sent in the transaction exceeds the logical gas limit set in the contract(it is not the same as the genesis gas limit).
* A transaction sent to a contract that has an empty code.
* A transaction that tries to display an empty code contract.
* In case that execution of the sent transaction exceeds the total GAS assigned to the writer node through which the transaction is sent.

### Node Ban and DoS

![gas_exceeded](images/gas_exceeded.png)

When an organization that owns a address which exceeds the GAS limit assigned to be used in a block, then it will be banned.

![penalize](images/penalize.png)

If action of exceeding the gas limit assigned occurs in one block address will be banned and transactions will not be accepted from that address. This is because it could be a denial of service attack.

The address banned can be reestablished after the organization has contacted the consensus group.

## How to adapt my contract to this new model

As the sender with which the transactions arrive at the receipient contract is the address of the RelayHub contract, a mechanism is necessary to obtain the original sender of the client or user who sent the transaction. 

To make this possible, we take advantage of the atomicity of the execution of the transactions in the EVM. That is, every time a transaction is sent to the RelayHub, the address of the original sender is stored, which is then retrieved by making a call to the RelayHub from the recipient contract.

This function to obtain the original sender is located in an abstract contract, which has to be inherited by all the contracts that will be deployed in the network.

![recipient](images/recipient.png)

## How send a protected transaction

In order to send a protected transaction, an end user or application that generates and signs the transaction must add two additional parameters to the function parameters of the destination contract. These parameters are as follows:

* nodeAddress(type:address): This parameter is the address of the private key that signs the transactions in the RelaySigner or by default it will be the address of the writer node through which the transactions will be sent.

* expiration(type:uint256): This parameter is the timestamp (Unix timestamp) that determines until when the transaction can be executed, after this time the transaction cannot be executed (added in a block).
