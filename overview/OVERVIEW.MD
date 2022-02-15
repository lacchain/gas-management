# 1. Context

The throughput of blockchain networks is inherently limited . There are several reasons for this. First, blockchain networks are decentralized, which leads to time constraints for replicating transactions across the network and executing these transactions by the nodes. Second, nodes maintain a copy of the entire transactional history, which leads to space constraints because the size of the history needs to be controlled and limited in order to stay manageable. For this reason, networks limit the size and execution capacity that blocks and nodes can assume.

Throughput limitations lead to a “supply and demand” challenge. How can we manage situations in which the amount of individuals or entities that want to use a blockchain network exceed the network’s ability to support them? Permissionless networks have addressed this with dynamic transaction-fee models. The more network space and computational capacity a person or entity requires for their registries/transactions, the more they have to pay. Consequently, the more users are using the network, the more expensive it becomes for each of them to use it. It is a classic supply and demand approach: transaction fees go up until supply and demand curves meet and equilibrium is reached.

Because of that, there is a big problem with transaction-fee-based networks which is that they quickly become very expensive. The more successful they are in attracting users, the more unaffordable they become for these users.  For example, Bitcoin transaction fee average oscillated between $2 and $63 over the year 2021, and Ethereum averaged transaction fees between $17 and $63 in the fourth quarter of 2021, respectively. This is why transaction-fee-based blockchain networks can hardly be an option for government and enterprise use cases that generate large amounts of transactions. This becomes even more unfeasible when taking into account fee volatility, which makes forecasts for budget allocation very difficult.

As opposed to permissionless networks, there are also permissioned blockchain networks which among other differences have usually erased transaction fees. These networks have been seen as the way to go for high transactional use cases. In the [LACChain Framework for Permissioned Public Networks](https://publications.iadb.org/en/lacchain-framework-permissioned-public-blockchain-networks-blockchain-technology-blockchain) we discussed how a permissioned public network can meet the benefits of permissionless networks and be at the same better for high transactional applications. 

Erasing transaction fees is definitely good for users, as they can send transactions for free or with a fixed membership cost. However, two main trade-offs must be addressed. One is what are the incentives for the entities participating and looking after the network’s welness if nobody gets rewarded with transaction fees (assuming that there is neither a native token or cryptocurrency serving this purpose). This is addressed in the LACChain Framework mentioned above.  The second and probably most critical issue is how is the access and use of the network (i.e., the distribution of resources) managed, in a way that the network does not collapse due to the fact that everyone could in principle send as many transactions as they want for free. This paper describes the open-source solution developed by LACNet to address this issue in Ethereum networks.

# 2. Solution

In order to manage the use of GAS in Ethereum-based permissioned public blockchain networks that have erased transaction fees, LACNet has developed the first GAS distribution model of its kind. LACNet’s GAS distribution model presented in this document is aligned with the LACChain Framework for Permissioned Public Networks.
Often, gas distribution models for blockchain networks consist of a faucet from which account owners can request gas. This has notable limitations. One limitation is that it might be difficult or impossible to prevent permissioned entities from transferring gas to non-permissioned entities, which, in a permissioned network, might enable non-permissioned users to broadcast transactions. Another limitation is that it might be difficult or impossible to prevent a permissioned or non-permissioned user to accumulate gas and use it to provoke a DoS attack and collapse the network by using all that gas in a short period of time.
LACNet’s gas distribution model for Ethereum-based Networks consists of a set of smart contracts that, among other things:
*	Assign GAS per block to the accounts associated with permissioned writer nodes in a dynamic way based on how stressed the network is at each time point (the more stressed it is, the less gas is distributed). GAS is not distributed or made available directly to end-user accounts.
*	Serve as a proxy to evaluate every transaction sent to the network and check that (i) each transaction has been signed by a permissioned writer node and (ii) the writer node has enough gas left for registering that transaction in the current block.
This model gives the power and responsibility to writer nodes rather than end users, as writer nodes will have the GAS that broadcasts transactions. Therefore, writer nodes decide which end-users they allow to use their GAS.
Writer nodes are required to intercept each transaction they receive from end-users or client applications and wrap them into meta-transactions they then broadcast to the network. These meta-transactions contain the original transaction as well as the signature of the permissioned writer node that is broadcasting the transaction to the network. 

This solution intentionally requires that writer nodes sign the transactions they send to the network in order to establish an accountability framework that makes them legally accountable for those transactions. Currently, Ethereum-based network transactions are signed only by anonymous end-users and it is not possible to track which nodes introduced or replicated a specific transaction. Thus, it is uncertain who is responsible for the information that is registered in a blockchain network. Therefore, it is possible that all nodes in the network are blamed for some data they did not introduce nor validate. However, it is immutably recorded in the copy that they maintain locally. For a more detailed discussion on this topic, read the LACChain Framework for Permissioned Public Blockchain Networks.

Instead of sending the transactions to the recipient contract indicated by the original sender, the writer node is required to send the transactions to the a of proxy smart contracts. In general terms, these smart contracts verify that the writer node has enough GAS and the signature is valid. 

## 2.1. Architecture

A high-level architecture diagram is presented in Figure 1. The solution has two main components: 
1) The smart contracts that contain the logic described in the previous section 
2) The writer node back-end components necessary to generate, sign, and send the meta-transactions to the proxy smart contracts. 

[FIGURE1]

### 2.1.1 Smart Contracts

This solution has six smart contracts: the Transaction Permissioning Smart Contract, the Node Permission Smart Contract, the RelayHub Smart Contract, the Gas Management Smart Contract, the Recipient Smart Contract, and the Local Account Permissioning Smart Contract. These smart contracts play two roles: 1) realize the gas distribution model perse by carrying out gas accounting and distribution, and 2) verify the transaction permissioning (both at the network and writer node levels).

*	Transaction Permissioning Smart Contract: The role of this contract is to only allow transactions that meet the network requirements (see Section 2.2 for more information):
    
    -Transactions must be directed to the RelayHub smart contract.
    
    -Transaction gas price must be 0.
    
    -The gas limit of the transaction must be sufficient enough to execute the RelayHub smart contract.
    
    -The writer node that broadcasts the meta-transaction must be the one indicated by the end-user or application in the original transaction (, as explained in Section 2.3) and the expiration time for the transaction set by the end-user must be respected.

Transactions that do not satisfy these rules will be rejected by validator nodes.

*	Node Permissioning Smart Contract: This contract contains the list of permissioned nodes and defines the rules for TCP connections between nodes according to the topology and the routing rules explained in the LACChain Blockchain Framework for Permissioned Public Networks. Nodes that are not registered in this contract cannot connect to the network. Before two nodes establish a connection, they must verify against this contract whether the connection is allowed. 
*	Relay Hub Smart Contract: This contract is based on the EIP 1077. The contract receives a meta-transaction (the transaction generated and signed by the writer node containing the original transaction sent by the sender) whose data field is decoded by RLP to obtain the parameters of the original transaction, which are the address of the original sender, gas limit, nonce, original data sent, and destination address. The RelayHub Smart Contract validates the signer's signature against the original message to guarantee that the message came from the signer. Then, it verifies that the node has not reached the gas limit for that block. If everything is correct, the transaction is forwarded to the recipient contract indicated by the original sender (if there is one) or creates a new smart contract (if that was the intention). 
*	Gas Management Smart Contract: Keeps track of how much gas each permissioned node (the address or addresses associated to it) has per block, and registers how much gas is being consumed in the last N blocks. It also sets the gas limit that each node can consume depending on the network’s degree of stress in the last N blocks.
*	Recipient Smart Contract: This contract is the final destination of the transaction sent by a user or client application (unless the transaction was not intended to call an existing smart contract, for instance, because it was intended to deploy a new contract). This contract executes the function chosen with the parameters sent in the original transaction.
*	Local Account Permissioning Contract: Writer nodes can deploy, maintain, and customize a Local Account Permissioning Contract, which is a permissioning layer that writer node operators can use to filter (by whitelisting) the reliable addresses (senders) they allow to send transactions to the network through their writer nodes under rules customized by the writer node operator such that transactions that do not meet the requirements (presented in Section 2.2) are broadcasted to the network. This allows for each writer node operator to define its own rules for transactions to be broadcasted to the network. For writer nodes exposed to external users, services, and applications, this is extremely important because writer node operators are accountable for transactions that their node broadcasts, despite who the original sender is.
When a new node is added to the network, the Permissioning Committee registers it into the Node Permissioning Smart Contract. Next, the organization senders’ addresses, whose node was previously added to the Node Permissioning Smart Contract, are added to the Transaction Permissioning Smart Contract. Automatically, the Transaction Permissioning Smart Contract invokes the Gas Management Contract which sets a gas limit for the new node addresses and recalculates the gas limit for the existing nodes. 

[FIGURE 2]

Each time a transaction is sent to the network, nodes check it against the Transaction Permissioning Smart Contract. First, writers check the transaction coming from the off-chain service or application, then boots check the transaction coming from the writer, and finally validators check the transaction coming from the boots.  If the transaction does not satisfy the requirements (introduced above and covered in detail in Section 2.2), the transaction will be rejected (i.e., not executed and erased from the transaction pool).
The validator nodes of the network, which are responsible for generating blocks, have the Transaction Permissioning configuration set, which makes them check every transaction they receive against the Transaction Permissioning Contract before they are executed in the network. It is worth mentioning that if a validator node becomes malicious by modifying the configuration of its node to accept invalid transactions into the network, this validator node will be banned by the other validators. For an invalid transaction to be introduced into a block, ⅔ +1 validator nodes need to be corrupted. If validator node operators tried such an attack, they would incur a legal penalty because validator nodes are committed contractually to not attack the network.

[FIGURE 3]

When the transaction satisfies the requirements, validator nodes execute the transaction against the RelayHub contract, where the nonce is verified to be greater than the previous one in order to guarantee that the transaction is unique and not duplicated. Before the transaction is forwarded to the Recipient Smart Contract (or creates a new smart contract itself), it is verified that the node has not consumed all the gas assigned to it for that block Gas Management Contract. If the transaction gas limit is lower than the amount of gas remaining for the writer node in the current block, the transaction is executed and the amount of gas used is discounted in the Gas Management Contract. If the transaction execution cannot be completed, the amount of gas consumed is discounted and a “not enough gas” error event is emitted. 

[FIGURE 4]

The RelayHub emits an event when a bad or doubtful transaction is sent, including:

•	The nonce sent is repeated or lower than the one registered in the contract for that user's address.
•	The gas limit sent in the transaction exceeds the block logical gas limit set in the contract (it is not the same as the genesis gas limit).
•	The transaction is sent to a contract that has an empty code.
•	The transaction tries to deploy an empty code contract.
•	The user transaction has an invalid signature.
•	The gas limit of the transaction exceeds the amount of gas left available for the writer node that is broadcasting the transaction. 

[FIGURE 5]

When an organization that owns a writer node exceeds the gas limit assigned to them in a block, it will be marked as an “exceeded gas limit” error, and the organization responsible for the writer node is automatically removed from the Transaction Permissioning Contract, which prevents it from sending transactions. When a node is banned, the Permissioning Committee reaches out to the node operator to clarify the situation and add the node back into the Transaction Permissioning Contract. 

[FIGURE 6]

When a private transaction is sent by the end-user or application to the writer node, it goes through the RelaySigner, which calls the Gas Management Smart Contract to discount a fixed amount of 25000. Then, it redirects the transaction to the Besu process, and this is communicated with the Orion service, which shares the transaction with the private Orion or Tessera nodes indicated trough the Orion or Tessera private channel.

[FIGURE 7]

### 2.1.2. Writer Node

The LACNet writer nodes consist of a Besu node configured as a writer (according to the LACChain topology) and is provided with a nginx and a RelaySigner. We have designed the gas management solution to be as transparent as possible to end-users, meaning that services and applications on top can send transactions almost as if the proxy smart contracts do not exist. It is the writer node that ties the transaction to compliance with the gas and node-signature requirements. Writer nodes can also deploy a Local Account Permissioning Contract as a filter to avoid unknown accounts from using their node, explained in Section 2.1.1.

•	Nginx: Acts as a reverse proxy and SSL termination (terminates HTTPS traffic from clients). It accepts all transactions to the node and forwards them to the RelaySigner. Direct queries to the node do not go through the RelaySigner; nginx redirects them to the Besu RPC port.
•	RelaySigner: Evaluates the type of message from an application (service or end-user) that arrives to the node. If it is a raw transaction (RLP unsigned from final user or application), then the RelaySigner decomposes it, builds a meta transaction, signs it, and sends it to the network. If it is another type of message, such as contract calls, the RelaySigner gets the recipient, number of connected nodes, block information or transaction, and forwards the message directly to the Besu component. 
•	Besu: A client of the LACNet network, with which the p2p connection is maintained. It broadcasts transactions to the other nodes of the network. This component receives meta-transactions from the RelaySigner to broadcast to network.

## 2.2. Requirements for Transactions to be Valid

The node back-end components described in Section 2.1 allow writer nodes to send only valid transactions to the network, and they will be banned if they decide to manipulate these components such that invalid transactions are sent. Transactions that do not meet the requirements below are rejected by all the non-corrupted nodes after they check them against the Transaction Permissioning Contract. The Transaction Permissioning Contract is a security layer of rules to determine whether a transaction should be executed and included into a block (see Section 2.1.1). As explained in Section 2.2.1 every node that listens to a transaction check that it a valid one before replicating it. Therefore, it is first the writer’s check followed by the boot's check, and afterwards by the validator’s check. 
•	The destination of the transaction must be the RelayHub (relay contract), through which all transactions must go through.
•	The gas limit set for the transaction must cover the RelayHub execution (fixed and variable parts). The variable part of the gas limit depends on the amount of bytes sent in the transaction, which guarantees the execution of the RelayHub contract and avoids exceptions for lack of gas (an out of gas exception).
•	The gas price sent in the transaction must be 0, because in networks orchestrated by LACNet the gas price is always 0 (see Section 2).
•	In order to guarantee that the transactions sent by end-users or services and applications on top are not altered by writer nodes or malicious parties, the original sender is required to add two additional parameters to the function parameters of the destination contract: the “nodeAddress” and the “expiration”. Doing this guarantees that a transaction sent by an end-user cannot be re-sent by any node other than the one selected by the end-user and the transaction is not executed after the expiration time indicated by the end user. 
•	The transaction sender must be an address registered by the organization running a writer node in the Transaction Permissioning Contract, which prevents anyone from sending transactions to the network. Additionally, transactions must be signed with the private key of the organization node’s registered address, which must also match the nodeAdress selected by the end user to broadcast its transaction. Only permissioned writers can send transactions to the network.

## 2.3. Gas Distribution Algorithm

The frequency with which the new gas limit is set for each participating writer node of the network is every N blocks. This parameter is configurable by the Permissioning Committee. In order to determine the gas limit for each node per block, the Permissioning Committee establishes the maximum amount of gas that the network can process per block without suffering delays in the block generation, G_T (see Section 3). 
The type of membership the node has purchased determines the amount of gas per block the node has access to. There are three types of memberships:
	
* Basic membership: The node gas limit is set to 500K per block.
* Standard membership: The node gas limit is set to 1.5M per block.
* Premium membership: There is not a pre-defined gas limit. The gas available per block R_T minus the gas reserved for basic and standard nodes is distributed equally among all the premium nodes. 

Table 1 gives the ideal amount of gas consumed on average by different typical transactions in the networks orchestrated by LACNet.

[TABLE 1]

The idea of “reserving gas” can lead to scenarios where some nodes might not be using all their gas per block whereas others might need more. For this reason, this gas distribution model is very flexible and attempts to maximize the amount of gas per block made available to each premium writer node. This is achieved by updating the gas limit every Y blocks for premium nodes according to how much the network is being used by all the permissioned nodes.  

The amount of gas available for each premium writer node per block changes dynamically depending on how much gas all the writer nodes have used in the last N blocks (this parameter is adjustable). The higher the amount of gas used by the writer nodes, the lower the gas limit will be. If all the nodes (basic, standard, and premium) have reached the gas limit available to them per block, then each premium node will have a gas limit per block equal to the total amount of gas minus the gas reserved for basic and standard, divided equally by the number of premium nodes. If writer nodes have not used all the gas assign to them, premium nodes are allowed to spend up to 5 times the previous amount up to 50% of the total gas available for all nodes-.

The formula to calculate the gas limit per block is as follows:

N_B:= Number of writer nodes with Basic membership package.

N_S:= Number of writer nodes with Standard membership package.

N_P:= Number of writer nodes with Premium membership package.

G_T:= Amount of gas per block that the network can process without operating under stress.

G_B:= Gas limit per block for writer nodes with Basic membership package.

G_S:= Gas limit per block for writer nodes with Standard membership package.

G_P:= Gas limit per block for writer nodes with Premium membership package.

ϑ≔ Fraction of G_T  consumed on average in the last 10 minutes by all the writer nodes.

[Formula 1]

This formula leads to the following scenarios:

[TABLE 2]

Let’s explore the formula with an example. Let us suppose that the amount of gas per block that the network can process without operating under stress 〖(G〗_T) is 150M. In a network with 20 basic, 20 standard, and 20 premium nodes, basic nodes would have access to a maximum of (0.5M gas limit x 20 nodes) 10M of gas, and standard nodes would have access to a maximum of (1.5M gas limit x 20 nodes) 30M of gas. The 20 premium nodes would then have access to a maximum of (150M gas total – 10M gas for basic nodes – 30M gas for standard nodes) 110M of gas to be equally used among them, which is equal to a gas limit per premium node of 5.5M gas per block. But this is the case only if these 60 nodes are using all the gas they have available (thus the parameter ϑ is equal to 1). If, conversely, the under 20% of the network network is used (the 150M of gas per block), then basic and standard nodes would keep their 0.5M and 1.5M gas limit per block, respectively, but premium nodes have up to 5 times higher gas limit, which in this example, would mean (5.5M gas limit x 5) 22.5M gas limit per block. Appendix A presents results of stress tests performed in the LACNet Networks using the GAS distribution mechanism. Appendix B analyzes the monthly cost in Ether and USD of this amount of gas in the Ethereum Mainnet.

# 3. Potential Attacks to the Network and Mitigations

As explained in Section 1, the GAS distribution model is a layer designed and developed by LACNet to prevent DoS attacks and require writer nodes to sign the transactions they broadcast to the network so they can be made accountable for them. As in every GAS distribution model, there are always potential attacks to be carried out by malicious node operators or end-users. It is essential to be able to identify these attacks in order to mitigate them. With the help of Coinspect, the LACNet team has developed mechanisms to mitigate all potential following attacks from an end-user/application/end-service with access to a writer node or from a malicious writer node to the network. 

By attacks to the network we are referring to any behavior than can impact the network negatively, including sending invalid transactions, attempting to modify Admin Smart Contracts (such as the Node Permissioning Contract, the RelyHub, or the Transaction Permissioning Contract) without permission, allowing non-permissioned nodes to broadcast transactions, attempting to use more gas than the gas available, attempting to attack other writer nodes or end-users by forwarding or overwriting pending transactions, or trying to overload the transaction pool of other nodes with transactions that will not be executed for different reasons. LACNet, as an orchestration vehicle, is responsible for ensuring that all these potential attacks are prevented or mitigated before they can impact the network.

## 3.1. DoS Attack by Exceeding GAS Limit

The first and most basic attack that any malicious writer node operator or end-user/application/service with access to a writer node can attempt is sending transactions that exceed the GAS limit assigned for that writer node in that block, which could lead to a DoS attack. Writer nodes need to modify their signer to perform this attack because the signer is set by default to refuse sending transactions that overpass the GAS remaining in each block (see Section 2.1.2).

The GAS distribution model prevents the attack by banning the writer node as soon as the first transaction that exceeds the GAS limit is broadcasted. This ban consists of removing the writer node from the Transaction Permissioning Contract so new transactions will automatically be rejected by validator nodes. To be added back into the Transaction Permissioning Contract, the writer node operator needs to clarify the potential DoS attempt.

## 3.2. Transactions with Lower GAS Limit than Required

As the base technology of the LACNet Besu network is Ethereum, every transaction executed by the nodes involves a computational cost, which is measured in units of GAS. When a transaction is sent with insufficient GAS, the transaction will not be fully executed by the validators. This releases an out of GAS (OOG) exception indicating that there was not enough GAS to complete the full execution of the transaction. 

A malicious end-user could send multiple transactions with a gas limit lower than what the transaction actually requires for execution, thereby attempting to generate the OOG exception. When the exception is launched, it does not execute the step of GAS reduction in the GAS Management Smart Contract. This makes it such that a malicious end-user or node operator could evade the accounting of how much GAS the node has consumed in that block. This situation could incur a ban for the writer node operator if it leads to the writer node attempting to exceed its gas limit. 

This potential attack is prevented with a prior check by validator nodes against the Transaction Permissioning Contract that discards these types of transactions before they are executed. Prior to the execution of the transaction by validator nodes, validator nodes check against the Transaction Permissioning Contract to confirm that the GAS limit of the transaction , based on the amount of bytes sent, is higher than what will be required to execute the RelayHub contract. By doing this, we guarantee that the amount of GAS is sufficient enough to complete the execution of the RelayHub contract and ensure we account for the GAS used by the node that sent the transaction, as well as the OOG exception. In case the GAS limit sent in the transaction does not cover the RelayHub execution, the transaction is rejected prior to the execution, and thus, no computation is involved. 

## 3.3. Setting the RelayHub as the Final Destionation for a Transaction

All transactions must go through the RelayHub contract before they reach their destination contract. A malicious user may want to re-enter the RelayHub contract with the intention of consuming GAS in a block without going through the accounting of the GAS model. The malicious user sends a transaction in which the target contract, after going through the RelayHub, becomes the RelayHub itself. This potential attack is prevented by the Rely Hub by not allowing re-entry. The Rely Hub would return an error and  the gas consumed by the transaction is discounted to the writer node

## 3.4. Trying to Modify Admin Contracts

A malicious user might want to modify admin contracts such as the Node Permissioning Contract, the Transaction Permissioning Contract, or the GAS Management Contract. For example, the malicious user would send a transaction in which the final destination contract after going through the RelayHub would be the method of removing nodes from the GAS Management Contract.

[FIGURE 8]

This potential attack is prevented by the RelayHub. When the validator nodes execute transactions against the RelayHub, the RelayHub confirms that the target contract is not any of the management contracts. This prevents end-users from reaching the contracts through the RelayHub, which has permission, through these contracts, to enable authorized users to interact with these contracts. Only authorized users specified in the RelayHub are allowed to interact with the Admin contract. Through this process, the malicious transaction would not be executed against the Admin Contract and the GAS would be discounted to the writer node that broadcasted it.

## 3.5. Transactions Broadcasted by Non-Permissioned Nodes

LACNet networks are permissioned networks, which means that permission is required for the node to join the network and connect to other nodes in the network. This permission is based on a permissioning smart contract. The Permissioning Committee updates the smart contract every time a node is added o removed from the permissioning, and all the permissioned nodes accept or reject connections based on the list of nodes reflected in the permissioning smart contracts.

A malicious writer node operator could remove its default node's permissioning settings (so it stops following the connections indicated in the Node Permissioning Smart Contracts), thereby accepting connections from new nodes. The malicious writer node operator would deploy a new malicious node or allow a third-party unauthorized node to connect to its node, allowing the unauthorized node or party to broadcast transactions that are replicated to the network by the permissioned node with the corrupted permissioning configuration. 

[FIGURE 9]

This potential attack is prevented by the Transaction Permissioning Smart Contract which rejects the transaction before it is execute because it is an invalid transaction as it is not signed by a permissioned node. The malicious node serving as a proxy will also be banned.

# 4. Guidelines and Recommendations for Writer Node Operators

In this section we analyze the GAS distribution mechanism from the perspective of writer node operators. First, we review the only two modifications they need to introduce into solutions developed for Ethereum and Hyperledger Besu Networks in order to be adapted for LACNet GAS distribution mechanism and therefore to work in LACNet Networks. Second, we provide recommendations for writer node operators to protect their nodes from potential attacks.

## 4.1.	Modifications to Solutions Developed for Ethereum Networks

As mentioned before, this GAS distribution mechanism has been designed to be as transparent as possible for end-users. Thus, the writer node back-end wraps the original transaction into a meta-transaction and signs it as requested by the proxy smart contracts. 

This GAS distribution mechanism brings only two small changes for end-users. The first is a difference in how to retrieve the original sender of a transaction and the second is indicating the writer node and the expiration time selected to broadcast the transaction.

### 4.1.1. Retrieving the Original Sender

The original transactions are not sent directly to the Recipient Smart Contract but by the RelyHub Contract (once it is checked that they meet the requirements detailed in Section 2.2). This makes it such that the sender of the transactions reaching the Recipient Smart Contract is the RelayHub contract.

Because of this, it is necessary to have a mechanism to obtain the address of the original sender (i.e., the client or user who sent the transaction to the writer node). To achieve this, we take advantage of the atomicity of the transaction execution in the EVM. That is, every time a transaction is sent to the RelayHub, the address of the original sender is stored, which is then retrieved by making a call to the RelayHub from the recipient contract.

This function to obtain the original sender is located in an abstract contract, which has to be inherited by all the contracts that will be deployed in the network.

[FIGURE 10]

For more detailed information, including a comparison between how transactions work with and without the GAS distribution model, please see https://github.com/lacchain/gas-relay-signer/blob/master/docs/Interaction%20with%20RelaySigner.md. 

### 4.1.2. Specifying Writer Node Address and Expiration Time

To avoid various attacks, such as the one described in Section XXX, end-users, applications and services on top that send transactions must add parameters to their transaction to indicate the address of the node they want to broadcast their transaction to as well as the expiration time for the transaction to be executed. These two parameters are checked by validator nodes against the Transaction Permissioning Smart Contract (see Section 2.1.1). Parameters are as follows:

* nodeAddress(type:address): This parameter is the address of the private key that signs the transactions in the RelaySigner. Otherwise, by default, it is the address of the writer node through which the transactions will be sent.
*	expiration(type:uint256): This parameter is the timestamp (Unix timestamp) that determines the expiration time for transaction to be executed. After this time, the transaction cannot be executed (added in a block).

## 4.2. Security Recommendations for Writer Node Operators

Writer node operators are contractually committed to LACNet to not send transactions with a gas limit higher than what they have left per block, not violating any other rule in the network-under penalty of being banned. Therefore, the entity operating the writer node must ensure that neither its use of the node nor the use by authorized third parties (e.g., applications and end users) violates these clauses. 

It is not the responsibility of LACNet to mitigate attacks from writer nodes by corrupt end-users or node administrators, because the operation of these nodes is the responsibility of the writer node operators. Further, even if security layers were developed, the operators of these nodes could always modify them, generating vulnerabilities that would fall under their responsibility. This is due to the decentralized architecture of the network, which allows each node to be self-managed. 

However, LACNet is not intended to let these writer node operators down with regard to potential attacks by end-users, services and applications on top, or others. On the contrary, LACNet has identified a set of threats posed to writer node operators and has provided tools, protocols, and suggestions for the writer node operators to mitigate any attack either to them or to the network through them. 

It is important to note that if the writer node operator does not want or cannot protect the writer node from the attack, the attack will not affect the network and other node operators because the malicious node will be immediately banned as soon as any attempt of misbehavior in the network is identified (such as exceeding GAS limit).

### 4.2.1. Restricting and Protecting Access to the Wrter Node

With prior knowledge that a writer node is exposed, attacks could be carried out by malicious users from both outside and inside the network. It is essential that writer nodes’ RPC and WebSocket ports are never open publicly without some layer of security. In the case of boots and validators, both types of nodes are committed via SLA to not have RPC nor WebSocket open, and to not have APIs enabled for sending transactions.

### 4.2.2. Checking and Filtering Transactions

Each network writer node has the ability to deploy and customize a Local Account Permissioning Contract, which is a permissioning layer that writer node operators can use to filter (by whitelisting) the reliable addresses (senders) they allow to send transactions to the network through their writer node under rules customizable by the writer node operator. Transactions that do not meet the requirements presented in Section 2.2 are broadcasted to the network.

This allows for each writer node operator to define its own rules for broadcasting transactions to the network. For writer nodes exposed to external users, services, and applications, this is extremely important, because writer node operators will be accountable for the transactions their node broadcasts, no matter who the original sender is.

Examples of rules include verifying the trail, the gas limit, the destination of the transaction or the gas price. They can also be more granular rules, such as verifying the data being sent in the transaction. 

[FIGURE 11]

### 4.2.3. Resending and Overwriting Transactions

Transactions sent by an innocent node to the network are replicated to all bootnodes. The bootnodes then replicate the transaction to the other writers and validators, which then execute the transactions. A malicious node could see a transaction in its pool and attempt to extract and forward it through its malicious node. It would do so with the intention that the second transaction is executed by the validators and the original transaction is executed with errors with its gas consumed anyway. Even if the original transaction is executed at the attacker’s expense, the end-user might be cheated because the end-user’s attempt would be unsuccessful.

This potential attack is mitigated by the requirement that transactions must specify the writer node that is selected to broadcast them (otherwise they are rejected by the Transaction Permissioning Contract). Therefore, transactions will only be accepted if they have the signature of the node address. Additionally, end-users must indicate an expiration time. For more information, see Annex I.

An alternative version of this attack could be to attempt to overwrite a transaction by increasing the gas price of the duplicated transaction so validators are more likely to accept it first. This is impossible because the network has a GAS price of zero, so any transaction with a GAS price higher than zero is also rejected by the Transaction Permissioning Smart Contract.


