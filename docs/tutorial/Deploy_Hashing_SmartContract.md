# Smart contract deployment tutorial
This tutorial provides a basic deployment of a smart contract on LACChain networks which are configured with gas model feature. It is intended for anyone interested in deploy smart contracts on LACChain networks base on Ethereum. You don't need any programming or blockchain experience to complete this tutorial. This is just the first step.

## 1. Running a writer node
The first step is to run a writer node in one of the LACChain networks. We recommend that you have a node running on the testnet. To be able to provision a node in the testnet network clone this repository.

```shell
$ git clone https://github.com/lacchain/besu-pro-testnet
$ git checkout david19_network
$ cd besu-pro-testnet/
```

Please, make sure you are in *david19_network* branch

After, follow this instructions [here](https://github.com/lacchain/besu-pro-testnet/tree/david19_network) to provision a *writer node*.

Finally, [check your writer](https://github.com/lacchain/besu-pro-testnet/tree/david19_network#checking-your-connection) node is working well.

## 2. Run relaysigner component

Wait your previously writer node deployed to have synced the entire blockchain.

Enter the node's console and run the following commands

```shell
$ cd /root/lacchain/gas-relay-signer
$ systemctl import-environment WRITER_KEY
$ service relaysigner start
```

Verify relaysigner is working well, go to the log file.

```shell
$ cd /root/lacchain/gas-relay-signer/log 
$ tail -100 idbServiceLog.log
```

The first lines should be something like this: 
```
General Logger:	2021/11/23 16:19:19 main.go:60: smartContract=0x3B62E51E37d090453600395Ff1f9bdf4d7398404 AgentKey=/home/adrian/.ethereum/keystore/UTC--2020-06-26T19-00-23.241896464Z--bceda2ba9af65c18c7992849c312d1db77cf008e
General Logger:	2021/11/23 16:19:19 main.go:65: Init RelaySigner
General Logger:	2021/11/23 16:19:19 client.go:46: Connected to Ethereum Node: ws://localhost:4546
```
That means relaysigner is configured and working well.

## 3. Set node address
Next step is set your node address(by default) to model gas. Please, run these commands on your node and copy your node address:

```shell
$ cd /root/lacchain/data
$ cat nodeAddress
```
Now send an email with the subject "permissioning node in testnet" to *adrianp@iadb.org* with copy to *eduardoma@iadb.org* writing as content your *node address* and your the name of your organization.

When you receive the response, your node will be ready to send transactions or deploy contracts.

## 4. Download prerequisites
A simple contract will be deployed which records a hash associated with the end user who sent the transaction.

For this part you will need to have *nodejs* installed. Check whether node is installed on your local computer by running the following command:

```shell
$ node --version
```
If the command doesn’t return a version number, download and install node by following the instructions for the operating system you use on the [Node.js](https://nodejs.org/es/download/) website. Node version should be at least v14.

Check whether yarn is installed on your local computer by running the following command:

```shell
$ yarn --version
```
If the command doesn’t return a version number, download and install yarn by running the following command:

```shell
$ npm install -g yarn
```
Please clone this repository. It contains all the code necessary to deploy a contract and send a transaction.

```shell
$ git clone https://github.com/lacchain/samples.git
```
Now download all the necessary dependencies

```shell
$ cd samples
$ yarn install
```

## 5. Configure local variables
After having downloaded the project dependencies, we are going to configure the local variables, for which make sure you have the IP of your writer node, as well as the private key of the end user with which you will deploy the contract.

```shell
$ cd samples
```
open a text editor and change the RPC_URL value to the IP of your writer node in /xxxxx/keys.js file, keep port 80 set.
```js
module.exports = {
    RPC_URL:"http://1.1.1.1:80"
}
```
save the changes. Now change the Web3 parameter to the IP of your writer node in /xxxxx/web3.js file. Also, keep port 80 set.

```js
const Web3 = require('web3')
const web3 = new Web3('http://1.1.1.1:80')
module.exports = web3
```
Next, set your user final address and private key, these values ​​will be of the end user who will send the transaction.

```js
const addressFrom = '0x173CF75f0905338597fcd38F5cE13E6840b230e9'  //SET USER ADDRESS HERE
const privKey = Buffer.from('0E273FFD9CF5214F7D6ADE5D1ABFD6D101B648AF12BC2DE6AC4AFCB4DB805CD3', 'hex') //SET USER FINAL PRIVATE_KEY HERE
```
Finally, set your node address and a timestamp which is 1 day ahead of the current date. To know how to get the timestamp, you can go to this [link](https://www.unixtimestamp.com)

```js
const nodeAddress = "0xd00e6624a73f88b39f82ab34e8bf2b4d226fd768"  //SET YOUR NODE ADDRESS HERE
const expiration = 1736394529  //SET TIMESTAMP HERE
```

## 6. Deploy your contract
To deploy the contract, please run the following command:
```shell
$ cd samples
$ node deployPublicSmartContract.js
```
If everything goes well, you will get the address of the contract deployed.
