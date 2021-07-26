// test/TxRelay.test.js
// Load dependencies
const { BN, expectEvent, time } = require('@openzeppelin/test-helpers');
const { expect } = require('chai');

// Load compiled artifacts
const TxRelay = artifacts.require('TxRelay');

// Start test block
contract('TxRelay', function (accounts) {
  const [ originalSender, admin, permissioningContract ] = accounts;

  beforeEach(async function () {
    // Deploy a new Box contract for each test
    this.txRelay = await TxRelay.new(10,permissioningContract, {from: admin});
  });

  // Test case
  it('exceded gas limit', async function () {
    // Relay a Transaction
    const encodedFunction = '0xf84180808316e36094eae291f161ba5a4d4e2c883bf5f6d9bc2a6a4cf980a46057361d000000000000000000000000000000000000000000000000000000000000002d';
    const v = 28;
    const r = '0x00568422bc039183ba811f2c94cbfa8c168e2304af0861d8a63429709fa9fc57';
    const s = '0x5362336809363bdd1d3f891ba7913586e11ea79ae3a96409e17fc8362624689f';
    
    const receipt = await this.txRelay.addNode(originalSender, {from: permissioningContract});
    expectEvent(receipt, 'NodeAdded', { newNode: originalSender});
    
    expect((await this.txRelay.getNodes()).toString()).to.equal('1');

    const receipt2 = await this.txRelay.setMaxGasBlockLimit(1000000, {from: admin});
    
    expectEvent(receipt2, 'MaxGasBlockLimitChanged', { admin: admin, maxGasBlockLimit: new BN(1000000)});

    const receipt3 = await this.txRelay.relayMetaTx(encodedFunction,v,r,s, {from: originalSender, gasPrice: 0, gas:2000000});
    expectEvent(receipt3, 'BadTransactionSent', { node: originalSender, errorCode: new BN(0) });
  });

  // Test case
  it('get node gasLimit ', async function () {
    const encodedFunction = '0xf861808082ea6094fd32cfc2e71611626d6368a41f915d0077a306a180b8446057361d000000000000000000000000000000000000000000000000000000000000003c000000000000000000000000173cf75f0905338597fcd38f5ce13e6840b230e9';
    const v = 27;
    const r = '0xc799efa03b5704c80758759aa0844544a5afba312f1d6f4a49b27dcbd75d14fa';
    const s = '0x0f1ab4f7308759e83db3820377fc595f10bdf103273a54fc061c33746a296cd0';

    const receipt = await this.txRelay.addNode(originalSender, {from: permissioningContract});
    expectEvent(receipt, 'NodeAdded', { newNode: originalSender});
    
    expect((await this.txRelay.getNodes()).toString()).to.equal('1');
    
    expect((await this.txRelay.getGasLimit()).toString()).to.equal('0');

    await this.txRelay.relayMetaTx(encodedFunction,v,r,s);

    expect(await this.txRelay.getGasLimit({from: originalSender })).to.be.bignumber.gt('150000000');
    

  });

  // Test case
  it('recipient contract doesnt exist' , async function () {
    // Relay a Transaction
    const encodedFunction = '0xf861808082ea6094fd32cfc2e71611626d6368a41f915d0077a306a180b8446057361d000000000000000000000000000000000000000000000000000000000000003c000000000000000000000000173cf75f0905338597fcd38f5ce13e6840b230e9';
    const v = 27;
    const r = '0xc799efa03b5704c80758759aa0844544a5afba312f1d6f4a49b27dcbd75d14fa';
    const s = '0x0f1ab4f7308759e83db3820377fc595f10bdf103273a54fc061c33746a296cd0';

    const receipt = await this.txRelay.addNode(originalSender, {from: permissioningContract});
    expectEvent(receipt, 'NodeAdded', { newNode: originalSender});
    
    expect((await this.txRelay.getNodes()).toString()).to.equal('1');
    
    expect((await this.txRelay.getGasLimit({ from: originalSender, gasPrice: 0})).toString()).to.equal('0');

    const receipt2 = await this.txRelay.relayMetaTx(encodedFunction,v,r,s,{ from: originalSender, gasPrice: 0});
    expectEvent(receipt2, 'BadTransactionSent', { node: originalSender, errorCode: new BN(4) });
  });

  // Test case
  it('invalid recipient smart contract' , async function () {
    // Relay a Transaction
    const encodedFunction = '0xf861808082ea6094'+permissioningContract.substring(2, permissioningContract.length)+'80b8446057361d000000000000000000000000000000000000000000000000000000000000003c000000000000000000000000173cf75f0905338597fcd38f5ce13e6840b230e9';
    const v = 27;
    const r = '0xc799efa03b5704c80758759aa0844544a5afba312f1d6f4a49b27dcbd75d14fa';
    const s = '0x0f1ab4f7308759e83db3820377fc595f10bdf103273a54fc061c33746a296cd0';

    const receipt = await this.txRelay.addNode(originalSender, {from: permissioningContract});
    expectEvent(receipt, 'NodeAdded', { newNode: originalSender});
    
    expect((await this.txRelay.getNodes()).toString()).to.equal('1');
    
    expect((await this.txRelay.getGasLimit({ from: originalSender, gasPrice: 0})).toString()).to.equal('0');

    const receipt2 = await this.txRelay.relayMetaTx(encodedFunction,v,r,s,{ from: originalSender, gasPrice: 0});
    expectEvent(receipt2, 'BadTransactionSent', { node: originalSender, errorCode: new BN(7) });
  });

  //Test case
  it('decrease Gas used by a Node' , async function () {
    // Relay a Transaction
    const encodedFunction = '0xf861808082ea6094fd32cfc2e71611626d6368a41f915d0077a306a180b8446057361d000000000000000000000000000000000000000000000000000000000000003c000000000000000000000000173cf75f0905338597fcd38f5ce13e6840b230e9';
    const v = 27;
    const r = '0xc799efa03b5704c80758759aa0844544a5afba312f1d6f4a49b27dcbd75d14fa';
    const s = '0x0f1ab4f7308759e83db3820377fc595f10bdf103273a54fc061c33746a296cd0';

    const receipt = await this.txRelay.addNode(originalSender, {from: permissioningContract});
    expectEvent(receipt, 'NodeAdded', { newNode: originalSender});
    
    expect((await this.txRelay.getNodes()).toString()).to.equal('1');
    
    expect((await this.txRelay.getGasLimit({ from: originalSender, gasPrice: 0})).toString()).to.equal('0');

    await this.txRelay.relayMetaTx(encodedFunction,v,r,s, {from: originalSender, gasPrice: 0});

    expect(await this.txRelay.getGasLimit({ from: originalSender, gasPrice: 0})).to.be.bignumber.gt('150000000');

    const receipt2 = await this.txRelay.increaseGasUsed(400000, {from: originalSender});
    expectEvent(receipt2, 'GasUsedByTransaction', { node: originalSender, gasUsed: new BN(400000), gasLimit: new BN(159600000) });
  });

  // Test case
  it('sent an incorrect nonce' , async function () {
    // Relay a Transaction
    const encodedFunction = '0xf8620180832dc6c094a3efcbced03c37600f8506de61835b4ca6f0573880b8446057361d000000000000000000000000000000000000000000000000000000000000003c000000000000000000000000173cf75f0905338597fcd38f5ce13e6840b230e9';
    const v = 27;
    const r = '0x00eb4d4c7eb32408c1ed44f915c44d817288b1d09a9fc2419d00679702bf346c';
    const s = '0x4990ec46cb11c8fe6fdd67b1157f267bc0d111a5934fd1ffabf8d20e6cd31ead';

    const receipt = await this.txRelay.addNode(originalSender, {from: permissioningContract});
    expectEvent(receipt, 'NodeAdded', { newNode: originalSender});
    
    expect((await this.txRelay.getNodes()).toString()).to.equal('1');
    
    expect((await this.txRelay.getGasLimit({ from: originalSender, gasPrice: 0})).toString()).to.equal('0');

    await this.txRelay.relayMetaTx(encodedFunction,v,r,s, {from: originalSender, gasPrice: 0});

    expect(await this.txRelay.getGasLimit({ from: originalSender, gasPrice: 0})).to.be.bignumber.gt('150000000');

    await this.txRelay.relayMetaTx(encodedFunction,v,r,s,{ from: originalSender, gasPrice: 0});

    const receipt3 = await this.txRelay.relayMetaTx(encodedFunction,v,r,s,{ from: originalSender, gasPrice: 0});
    expectEvent(receipt3, 'BadTransactionSent', { node: originalSender, errorCode: new BN(2) });
  });

  // Test case
  it('sent an invalid signature' , async function () {
    // Relay a Transaction
    const encodedFunction = '0xf861808082ea6094fd32cfc2e71611626d6368a41f915d0077a306a180b8446057361d000000000000000000000000000000000000000000000000000000000000003c000000000000000000000000173cf75f0905338597fcd38f5ce13e6840b230e9';
    const v = 29;
    const r = '0xc799efa03b5704c80758759aa0844544a5afba312f1d6f4a49b27dcbd75d14fa';
    const s = '0x0f1ab4f7308759e83db3820377fc595f10bdf103273a54fc061c33746a296cd0';

    const receipt = await this.txRelay.addNode(originalSender, {from: permissioningContract});
    expectEvent(receipt, 'NodeAdded', { newNode: originalSender});
    
    expect((await this.txRelay.getNodes()).toString()).to.equal('1');
    
    expect((await this.txRelay.getGasLimit({ from: originalSender, gasPrice: 0})).toString()).to.equal('0');

    const receipt2 = await this.txRelay.relayMetaTx(encodedFunction,v,r,s,{ from: originalSender, gasPrice: 0});
    expectEvent(receipt2, 'BadTransactionSent', { node: originalSender, errorCode: new BN(6) });
  });

  // Test case
  it('penalize excess gas used by a node' , async function () {
    // Relay a Transaction
    const encodedFunction = '0xf8628080833877fb94a3efcbced03c37600f8506de61835b4ca6f0573880b8446057361d000000000000000000000000000000000000000000000000000000000000003c000000000000000000000000173cf75f0905338597fcd38f5ce13e6840b230e9';
    const v = 27;
    const r = '0x3dee366368e92130b7c7f9fa93af9f4bbc7958d6d94f090f99395f3644362dca';
    const s = '0x27e5be796f1f162976abf24b0e46755554a41147f67379a558844c75de714a9b';

    const encodedFunction1 = '0xf8620180833877fb94a3efcbced03c37600f8506de61835b4ca6f0573880b8446057361d000000000000000000000000000000000000000000000000000000000000003c000000000000000000000000173cf75f0905338597fcd38f5ce13e6840b230e9';
    const v1 = 27;
    const r1 = '0x5434f994aba8ab3e88dbf62adb4508fbbb1e4c6d09cb7c687e0208dab3d9c25e';
    const s1 = '0x67cb28ed12b84c43ce3404ae3ac8b311f4e017294bfba610ce266851c6e75c61';

    const encodedFunction2 = '0xf8620280833877fb94a3efcbced03c37600f8506de61835b4ca6f0573880b8446057361d000000000000000000000000000000000000000000000000000000000000003c000000000000000000000000173cf75f0905338597fcd38f5ce13e6840b230e9';
    const v2 = 27;
    const r2 = '0x5d74c3c58fd92e9891ca0aa31b3d7c1c79068e9f70b7bedffdfe60e17d2bd202';
    const s2 = '0x6b2f6c4c5009baf2e11766a58dbd39de9caf0e3bca5f0641f8c60366402c7ec8';

    const encodedFunction3 = '0xf8620380833877fb94a3efcbced03c37600f8506de61835b4ca6f0573880b8446057361d000000000000000000000000000000000000000000000000000000000000003c000000000000000000000000173cf75f0905338597fcd38f5ce13e6840b230e9';
    const v3 = 27;
    const r3 = '0x55ff90b768c4a814707bc46bb8259926d3ba817392421c464c7b93258e9d10d8';
    const s3 = '0x57021df83bf7932c72dbd404bd31dc0aee23e5aa9c321b584897fc7873bee863';

    const receipt = await this.txRelay.addNode(originalSender, {from: permissioningContract});
    expectEvent(receipt, 'NodeAdded', { newNode: originalSender});
    
    expect((await this.txRelay.getNodes()).toString()).to.equal('1');
    
    expect((await this.txRelay.getGasLimit({ from: originalSender, gasPrice: 0})).toString()).to.equal('0');

    await this.txRelay.relayMetaTx(encodedFunction,v,r,s, {from: originalSender, gasPrice: 0});

    expect(await this.txRelay.getGasLimit({ from: originalSender, gasPrice: 0})).to.be.bignumber.gt('150000000');

    const receipt2 = await this.txRelay.setMaxGasBlockLimit(10000, {from: admin});
    
    expectEvent(receipt2, 'MaxGasBlockLimitChanged', { admin: admin, maxGasBlockLimit: new BN(10000)});

    for (i = 0; i < 9; i++) {
      time.advanceBlock();
    }

    const receipt3 = await this.txRelay.relayMetaTx(encodedFunction1,v1,r1,s1,{ from: originalSender, gasPrice: 0});
    expectEvent(receipt3, 'NodeBlocked', { node: originalSender });
  });

  //Test
  it('get nonce' , async function () {
    // Relay a Transaction
    const encodedFunction = '0xf8628080833877fb94a3efcbced03c37600f8506de61835b4ca6f0573880b8446057361d000000000000000000000000000000000000000000000000000000000000003c000000000000000000000000173cf75f0905338597fcd38f5ce13e6840b230e9';
    const v = 27;
    const r = '0x3dee366368e92130b7c7f9fa93af9f4bbc7958d6d94f090f99395f3644362dca';
    const s = '0x27e5be796f1f162976abf24b0e46755554a41147f67379a558844c75de714a9b';

    const encodedFunction1 = '0xf8620180833877fb94a3efcbced03c37600f8506de61835b4ca6f0573880b8446057361d000000000000000000000000000000000000000000000000000000000000003c000000000000000000000000173cf75f0905338597fcd38f5ce13e6840b230e9';
    const v1 = 27;
    const r1 = '0x5434f994aba8ab3e88dbf62adb4508fbbb1e4c6d09cb7c687e0208dab3d9c25e';
    const s1 = '0x67cb28ed12b84c43ce3404ae3ac8b311f4e017294bfba610ce266851c6e75c61';

    const encodedFunction2 = '0xf8620280833877fb94a3efcbced03c37600f8506de61835b4ca6f0573880b8446057361d000000000000000000000000000000000000000000000000000000000000003c000000000000000000000000173cf75f0905338597fcd38f5ce13e6840b230e9';
    const v2 = 27;
    const r2 = '0x5d74c3c58fd92e9891ca0aa31b3d7c1c79068e9f70b7bedffdfe60e17d2bd202';
    const s2 = '0x6b2f6c4c5009baf2e11766a58dbd39de9caf0e3bca5f0641f8c60366402c7ec8';

    const encodedFunction3 = '0xf8620380833877fb94a3efcbced03c37600f8506de61835b4ca6f0573880b8446057361d000000000000000000000000000000000000000000000000000000000000003c000000000000000000000000173cf75f0905338597fcd38f5ce13e6840b230e9';
    const v3 = 27;
    const r3 = '0x55ff90b768c4a814707bc46bb8259926d3ba817392421c464c7b93258e9d10d8';
    const s3 = '0x57021df83bf7932c72dbd404bd31dc0aee23e5aa9c321b584897fc7873bee863';

    const receipt = await this.txRelay.addNode(originalSender, {from: permissioningContract});
    expectEvent(receipt, 'NodeAdded', { newNode: originalSender});
    
    expect((await this.txRelay.getNodes()).toString()).to.equal('1');
    
    expect((await this.txRelay.getGasLimit({ from: originalSender, gasPrice: 0})).toString()).to.equal('0');

    await this.txRelay.relayMetaTx(encodedFunction,v,r,s, {from: originalSender, gasPrice: 0});

    expect(await this.txRelay.getGasLimit({ from: originalSender, gasPrice: 0})).to.be.bignumber.gt('150000000');

    for (i = 0; i < 9; i++) {
      time.advanceBlock();
    }

    await this.txRelay.relayMetaTx(encodedFunction1,v1,r1,s1,{ from: originalSender, gasPrice: 0});
    await this.txRelay.relayMetaTx(encodedFunction2,v2,r2,s2,{ from: originalSender, gasPrice: 0});
    
    expect((await this.txRelay.getNonce('0x173CF75f0905338597fcd38F5cE13E6840b230e9',{from: originalSender})).toString()).to.equal('3');
  });
  
});