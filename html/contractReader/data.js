const ETHERSCAN_API_KEY = 'H9Q2BDA55J85PISTNG8BDM5IKD4MGTUVW4';
const ALCHEMY_API_KEY = 'FY6BwiO9_hzVN4N2Fx8Ti-BeukyI2XiM';

const _web3Instance = {};
let _contractInstance;

async function dataFetchAbi(contractAddress, network) {
  if (!contractAddress) {
    throw new Error('Contract address is empty.');
  }
  if (!contractAddress.startsWith('0x')) {
    throw new Error('Contract address does not start with 0x.');
  }
  if (contractAddress.length != 42) {
    throw new Error('Contract address should be 20 bytes.');
  }
  return JSON.parse(_abi)
}

function dataEncodeFunctionSignature(abiField, network) {
  return _getWeb3(network).eth.abi.encodeFunctionSignature(abiField);
}

function dataInitializeContractInstance(contractAddress, network, abi) {
  console.log(network)
  _contractInstance = new (_getWeb3(network).eth.Contract)(abi, contractAddress);
  _contractInstance.address = contractAddress
}

async function dataQueryFunction(abiField, inputs, blockNumber, from) {
  return _contractInstance.methods[abiField.name](...inputs).call();
}

async function dataWriteFunction(abiField, inputs, blockNumber, from) {
  // console.log(abiField.name, _contractInstance.address)
  const transactionParameters = {
    from: $("#addr").text(),
    to: _contractInstance.address,
    data: _contractInstance.methods[abiField.name](...inputs).encodeABI(),
    // custom gas price
  };

  return ethereum.request({
    method: 'eth_sendTransaction',
    params: [transactionParameters],
  });
}

function dataValidateType(type, value) {
  if (value === '') {
    throw new Error('Value is empty.');
  }
  if (type === 'address') {
    if (!value.startsWith('0x')) {
      throw new Error('Value does not start with 0x.');
    }
    if (value.length != 42) {
      throw new Error('Value should be 20 bytes.');
    }
  }
}

function _getWeb3(network) {
  console.log(network)
  return new Web3(network);
}

