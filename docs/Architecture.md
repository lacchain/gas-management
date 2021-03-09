# Solution Architecture

In this section we will review the different components of the architecture, function and relationship with the other components.

![Architecture](images/architecture.png)

## Backend Components

### Node
Is a writer node part of the LACChain network. The node is composed of Nginx, relaySigner layer and Hyperledger Besu.

### Nginx
Its role is act as reverse proxy and SSL termination (terminate HTTPS traffic from clients), it accepts all transactions to the node and forwards them to the RelaySigner.

### RelaySigner Layer
Its role is to evaluate the type of message that arrives. In case it is a raw transaction, then it is decomposed and builds a meta transaction to send it to the network. If it is another type of message such as contract calls, get the recipient, number of connected nodes, block information or transaction then the message is directly forwarded to the Besu component.

### Besu
Their role is to be a client of the LACChain network, with which the p2p connection is maintained, broadcast transactions with the other nodes of the network. This component receive meta-transactions from RelaySigner to broadcast to network.

## Smart Contracts

### Account Permission Smart Contract
This smart contract is deployed on all nodes in the LACChain network. The role of this contract is to allow only transactions that are directed towards the RelayHub smart contract. Transactions that do not go through the RelayHub will be rejected.

### Node Permission Smart Contract
This smart contract is deployed on all nodes in the LACChain network. The contract allows connections between the nodes that are registered in the contract. 

Its role in the model is to invoke the gas management contract when nodes are added or removed to the network to perform the new gas limit assignment on the nodes.

### Relay Hub Smart Contract
This contract is based on EIP 1077. The contract receives the meta-transaction and validates the signer's signature against the sent message to guarantee that the message came from the signer. Then verifies that the node has not reached the gas limit for that block. In case everything is correct, the transaction is forwarded to the recipient contract or create a new smart contract.

### Gas Managenment Smart Contract
This contract keeps track of how much gas is consumed by each node of the network. It also assigns a new maximum gas limit that each node can consume depending on the use of the network in the last blocks.

### Recipient Smart Contract
This contract will be the final destination of the transaction sent by an user or client application. It is the contract that executes the function choosen with parameters sent in the original transaction.

## Behavior

### Add New Node

![add_new_node](images/add_new_nodo.png)

When a new node is added to the network, it is added in the Node Permissioning Contract. Then, the permissioning contract invokes the gas management contract which will set a GAS limit that the new node can use.

### New Gas Node Limit

The frequency with which the new gas limit is set to each participating writer node of the network is every 10 blocks, this is configurable by the consensus group.

![calculate_variables](images/calculate_variables.png)

The formula as the gas is distributed depends on how much GAS was used in the last blocks frequency (10 blocks). While higher the amount of Gas used among all the organizations, then lower the maximum GAS limit will be. It is very important to mention that the GAS is distributed in the same way among all the members of the network.

![calculate_continous](images/calculate_continous.png)

![calculate_discrete](images/calculate_discrete.png)

### Send Transaction

![send_transaction](images/send_transaction.png)

Each time a transaction is sent to the network, it is first verified against the account permissioning contract if the transaction is directed to the Relayhub. In case the transaction does not go to the RelayHub the transaction is canceled and will not be executed. 

After the transaction is verified, this transaction will go directly to the RelayHub contract, which will validate if the transaction was signed by a permissioned node and the transaction has not been altered.

As additional validation, it is verified that the nonce sent is greater than the previous one to guarantee that the transaction is unique and does not repeat itself. Before the transaction is forwarded to a recipient contract, it is verified that the node has not consumed all the assigned GAS, in case the node has enough gas then the transaction is forwarded and after execution in the recipient contract the amount of GAS used is reduced for that node.

### Private Transactions

![private_transactions](images/private_transactions.png)

When a private transaction is sent, it goes through the RelaySigner which call to Management Gas contrat to decrease the gas used, which is 25000. Then it redirects the transaction without modifying to the Besu process and this communicates with the Orion service which share the transaction with the participating nodes of the private transaction.

### Bad Transactions

![send_bad_transaction](images/send_bad_transaction.png)

The RelayHub emits a event when a bad or doubtful transaction is sended in following cases:
- The nonce sent is repeated or lower than registered in contract for that user's address.
- The gas limit sent in the transaction exceeds the logical gas limit set in the contract(it is not the same as the genesis gas limit).
- A transaction sent to a contract that has an empty code.
- A transaction that tries to display an empty code contract.
- In case that execution of the sent transaction exceeds the total GAS assigned to the writer node through which the transaction is sent.

### Node Ban and DoS

![gas_exceeded](images/gas_exceeded.png)

When an organization that owns a writer node exceeds the GAS limit assigned to be used in a block, then it will be counted as an exceeded limit and informed by mail to the institution of such action. 

The excesses are reset in the following block frequency (every 10), that means that when the 10 blocks generated are completed, the count of excesses of all the writer nodes will be reset to 0 again.

![penalize](images/penalize.png)

If action of exceeding the gas limit assigned occurs more than 3 times then the node will be banned and transactions will not be accepted from that writer node. This is because it could be a denial of service attack.

The node banned can be reestablished after the organization has contacted the consensus group.

## How to adapt my contract to this new model

As the sender with which the transactions arrive at the receipient contract is the address of the RelayHub contract, a mechanism is necessary to obtain the original sender of the client or user who sent the transaction. 

To make this possible, we take advantage of the atomicity of the execution of the transactions in the EVM. That is, every time a transaction is sent to the RelayHub, the address of the original sender is stored, which is then retrieved by making a call to the RelayHub from the recipient contract.

This function to obtain the original sender is located in an abstract contract, which has to be inherited by all the contracts that will be deployed in the network.

![recepeint](images/recepeint.png)


