//binary to deploy the smart contract
 const simpleStorageByteCode = "0x608060405234801561001057600080fd5b5060c78061001f6000396000f3fe6080604052348015600f57600080fd5b506004361060325760003560e01c80636057361d146037578063b05784b8146062575b600080fd5b606060048036036020811015604b57600080fd5b8101908080359060200190929190505050607e565b005b60686088565b6040518082815260200191505060405180910390f35b8060008190555050565b6000805490509056fea26469706673582212208595760c8d4272f1711ffb94441dddc78e22630630efc2bc60b984de1caeb06c64736f6c63430006030033";

const addressFrom = '0x173CF75f0905338597fcd38F5cE13E6840b230e9'
const privKey = Buffer.from('0E273FFD9CF5214F7D6ADE5D1ABFD6D101B648AF12BC2DE6AC4AFCB4DB805CD3', 'hex')

let contractData

const nodeAddress = "0x971bb94d235a4ba42d53ab6fb0a86b12c73ba460"  //additional nodeAddress parameter
const expiration = 1636394529  //timestamp unix parameter
let value = web3.eth.abi.encodeParameters(
  ["address","uint256"],
  [nodeAddress,expiration])

contractData = simpleStorageByteCode + value.substr(2)

console.log("txData:"+contractData)

const deploy = async () => {
  console.log("#######################Deploying smart contract#######################")
  return await deploySmartContract(contractData,addressFrom,privKey)
}

deploy()
