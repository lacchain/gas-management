var TxRelay = artifacts.require("TxRelay");

//const Rules = artifacts.require("./AccountRules.sol");

const blocksFrequency = 60;
const account_Ingress = '0xbcEda2Ba9aF65c18C7992849C312d1Db77cF008E';

module.exports = async(deployer, network) => {
//    const rules = await Rules.deployed();

    await deployer.deploy(TxRelay, blocksFrequency, account_Ingress);
    console.log("   > Deployed TxRelay contract to address = " + TxRelay.address);

    const relayHub = await TxRelay.deployed();

//    await rules.setRelay(rules.address);
//    await relayHub.addNode("0xd00e6624a73f88b39f82ab34e8bf2b4d226fd768")

//    console.log("node Added")

    await relayHub.addNode("0xfaae4e8e9dabf9859db1601024191f3c97302230")

    console.log("node Added")

};
