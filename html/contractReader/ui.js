let _currentNetwork = 'https://chain.chymeta.com';
let _addr = $('#addr').val();
let _chainId = '11927';
let _scanAddr = 'https://scan.chymeta.com/';
let _abi = '';

// inits the UI, should be called in window.onload
async function uiInit() {
  _initUiComponents();
  $('#main-menu > a').click(_mainMenuOnClick);
  // $('#load-contract').click(_loadContractOnClick);
  // _loadContractOnClick();
  const params = $.deparam();
  if (params.contract) {
    const d = await _fetchJson('/admin/getContract?addr=' + params.contract);
    // console.log(d)
    // let d = JSON.stringify(abiReq)
    if (d.data.Abi.length > 0) {
      _abi = JSON.parse(JSON.stringify(d.data.Abi))
      _chainId = d.data.ChainId
      _chainRpc = d.data.ChainRpc
      _scanAddr = d.data.Scan
      $('#contractName').html('当前合约信息 - ' + d.data.Name + ':')
      $('#contract').html(d.data.Addr)
      _loadContract(params.contract);
      login();
    } else {
      alert('没有abi数据!')
    }
    // $('#contract-address').val(params.contract);
  } else {
    alert("缺少合约地址!")
  }
  $(`#main-menu > a[target="blank" href='${window.location.hash}']`).trigger('click');
}


async function login() {
  if (typeof window.ethereum !== 'undefined') {
    let addr = await ethereum.request({ method: 'eth_requestAccounts' });//授权连接钱包
    // console.log('用户钱包地址:', addr[0]);
    $('#addr').html(addr[0]);
    getChainId();
  } else {
    alert('没有可用钱包!');
  }
}


async function getChainId() {
  let chainId = await ethereum.request({ method: 'eth_chainId' });
  //！链id不是马上拿到的，如果通过链id来判断是不是主网的方式，请注意异步
  let id = eval(chainId).toString(10);
  $('#chain').html(id)
  if (id != _chainId) {
    alert('请更换网络,合约网络id为' + _chainId)
  }
}

ethereum.on("accountsChanged", function (accounts) {
  window.location.reload();
});
//监听链网络改变
ethereum.on("chainChanged", () => {
  window.location.reload();
});

// helper to init special SemanticUI components that have JS behavior
function _initUiComponents() {
  $('.ui.dropdown').dropdown();
  $('.ui.accordion').accordion();
  // custom dropdown component for changing decimals of numbers
  $('.ui.dropdown[data-tpl=decimals]').dropdown({
    onChange: (decimals, _, $choice) => {
      const inputEl = $choice.closest('.ui.input').find('[data-tpl=input]');
      const prevDecimals = inputEl.data('prevDecimals') || 0;
      inputEl.data('prevDecimals', decimals);
      if (inputEl.val() === '') return;
      const rawNum = new BigNumber(`${inputEl.val().replaceAll(',', '')}e${prevDecimals}`);
      if (decimals != 0) {
        inputEl.val(rawNum.dividedBy(new BigNumber(`1e${decimals}`)).toFormat());
      } else {
        inputEl.val(rawNum.dividedBy(new BigNumber(`1e${decimals}`)).toFixed());
      }
      inputEl.data('rawVal', rawNum.toFixed());
    }
  });
}

// loads a new contract into the page, downloads its ABI and generates all UI
async function _loadContract(contractAddress) {
  try {
    const abi = await dataFetchAbi(contractAddress, _currentNetwork);
    dataInitializeContractInstance(contractAddress, _currentNetwork, abi);
    _mainError();
    _generateUiForAbi(abi);
  } catch (e) {
    _mainError(e.message);
    _generateUiForAbi();
  }
}

// event handler when user clicks the main "Load" button to load a new contract
function _loadContractOnClick() {
  const contract = $.param['contract'];
  const network = "bsc";
  // window.location.search = $.param({
  //   contract,
  //   network
  // });
}

// event handler when user clicks the main sections menu (read, write, events, etc)
function _mainMenuOnClick(event) {
  $(event.delegateTarget).parent().find('.active').removeClass('active');
  $(event.delegateTarget).addClass('active');
  $('.main-menu-content').hide();
  const contentSelector = `#${$(event.delegateTarget).data('target')}`;
  $(contentSelector).show();
}

// if error is non-empty, shows the main error dialog
// if error is empty, hides the main error dialog
function _mainError(errorMessage) {
  if (errorMessage) {
    $('#contract-loaded').hide();
    $('#main-error > p').text(errorMessage);
    $('#main-error').fadeIn(100);
  } else {
    $('#main-error').hide();
    $('#contract-loaded').fadeIn(100);
  }
}

// populate all the UI for a newly loaded contract (all functions, events, etc)
function _generateUiForAbi(abi) {
  if (abi) {
    for (let i = 0; i < abi.length; i++) {
      const field = abi[i];
      if (field.type === 'function') {
        if (field.stateMutability === 'view') {
          _addUiForFunctionField(field, '#abi-read-functions');
        } else {
          _addUiForFunctionField(field, '#abi-write-functions');
        }
      }
      if (field.type === 'event') {
        // todo
      }
      _initUiComponents();
    }
  } else {
    $('#abi-fields > div').empty();
  }
}

// populate the UI for a single contract function
function _addUiForFunctionField(field, toSelector) {
  console.log(field)
  const functionEl = $('#templates > #abi-function').clone();
  $(toSelector).append(functionEl);
  $(functionEl).find('[data-tpl=name]').text(field.name + ' / ' + field.cnName);
  $(functionEl).find('[data-tpl=signature]').text(dataEncodeFunctionSignature(field, _currentNetwork));
  $(functionEl).find('[data-tpl=query]').click(field, _queryFunctionOnClick);
  for (let i = 0; i < field.inputs.length; i++) {
    const functionArgEl = _getElementForFunctionArg(field.inputs[i]);
    _appendInputFunctionArg(field.inputs[i], functionEl, functionArgEl);
  }
}

// event handler when user clicks the "Query" button of a function to execute a function call
async function _queryFunctionOnClick(event) {
  const field = event.data;
  const functionEl = $(event.delegateTarget).parent();
  $(functionEl).find('[data-tpl=outputs]').empty();
  const inputArray = _validateQueryInputArray(field, functionEl);
  if (!Array.isArray(inputArray)) return;
  try {
    // read call contract
    if (event.data.stateMutability == "view") {
      const result = await dataQueryFunction(field, inputArray);
      for (let i = 0; i < field.outputs.length; i++) {
        const outputValue = result; // todo fix for multiple
        const functionArgEl = _getElementForFunctionArg(field.outputs[i]);
        _appendOutputFunctionArg(field.outputs[i], outputValue, functionEl, functionArgEl);
      }

    } else {
      const result = await dataWriteFunction(field, inputArray);
      for (let i = 0; i < field.outputs.length; i++) {
        const outputValue = result; // todo fix for multiple
        const functionArgEl = _getElementForFunctionArg(field.outputs[i]);
        _appendOutputFunctionArg(field.outputs[i], outputValue, functionEl, functionArgEl);
      }
    }

  } catch (e) {
    _appendOutputFunctionError(e.message, functionEl);
  }
  _initUiComponents();
}

// prepare the HTML element for a single function argument
function _getElementForFunctionArg(fieldArg) {
  let templateSelector = '#function-arg';
  if (fieldArg.type === 'uint256') {
    templateSelector = '#function-arg-decimals';
  }
  const functionArgEl = $(`#templates > ${templateSelector}`).clone();
  $(functionArgEl).find('[data-tpl=name]').text(fieldArg.name);
  $(functionArgEl).find('[data-tpl=type]').text(fieldArg.type);
  if (templateSelector === '#function-arg-decimals') {
    $(functionArgEl).find('[data-tpl=menu]').find('.item').text((_, text) => {
      return text.replace('type', fieldArg.type);
    });
  }
  return functionArgEl;
}

// populate the UI for a single function argument for input
function _appendInputFunctionArg(fieldArg, functionEl, functionArgEl) {
  $(functionEl).find('[data-tpl=inputs]').append(functionArgEl);
  let placeholder = '';
  if (fieldArg.type === 'address') {
    placeholder = '20 byte hex string starting with 0x';
  }
  if (fieldArg.type === 'uint256') {
    placeholder = 'unsigned decimal';
  }
  $(functionArgEl).find('[data-tpl=input]').attr('placeholder', placeholder);
}

// populate the UI for a single function argument for output
function _appendOutputFunctionArg(fieldArg, outputValue, functionEl, functionArgEl) {
  $(functionArgEl).hide();
  $(functionEl).find('[data-tpl=outputs]').append(functionArgEl);
  $(functionArgEl).find('[data-tpl=input]').val(outputValue);
  $(functionArgEl).find('[data-tpl=input]').attr('readonly', '1');
  if (fieldArg.type === 'address') {
    $(functionArgEl).find('[data-tpl=input]').parent().wrap(`<a target="blank" style="color:#F00" href="${_getLinkForAddress(outputValue)}"></a>`);
  } else if (fieldArg.type === 'bool') {
    $(functionArgEl).find('[data-tpl=input]').parent().wrap(`<a target="blank" style="color:#F00" href="${_getTxLinkForAddress(outputValue)}"></a>`);
  }
  console.log(fieldArg.type, "-----")
  $(functionArgEl).slideDown(100);
}

// validate all input arguments and show error UI before executing a function call
// returns array if the inputs are valid, false if not
function _validateQueryInputArray(field, functionEl) {
  const inputArray = [];
  let error = false;
  $(functionEl).find('[data-tpl=inputs]').find('[data-tpl=input]').each((i, inputEl) => {
    const fieldArg = field.inputs[i];
    const inputValue = $(inputEl).data('rawVal') || $(inputEl).val();
    try {
      dataValidateType(fieldArg.type, inputValue);
    } catch (e) {
      const errorMessage = `Input '${fieldArg.name}' is invalid: ${e.message}`;
      alert(errorMessage);
      _appendOutputFunctionError(errorMessage);
      errorMessage = true;
    }
    inputArray.push(inputValue);
  });
  if (error) return false;
  return inputArray;
}

// populate the UI showing an error in the function output
function _appendOutputFunctionError(errorMessage, functionEl) {
  const errorEl = $('#templates > #function-arg-error').clone();
  $(errorEl).hide();
  $(errorEl).find('[data-tpl=text]').text(errorMessage);
  $(functionEl).find('[data-tpl=outputs]').append(errorEl);
  $(errorEl).slideDown(100);
}

// get the URL for etherscan for exploring an Ethereum/BSC address further
function _getLinkForAddress(address) {
  return _scanAddr + `${address}`;
}

function _getTxLinkForAddress(address) {
  return _scanAddr + `tx/${address}`;
}

(function ($) {
  $.deparam = $.deparam || function (uri) {
    if (uri === undefined) {
      uri = window.location.search;
    }
    var queryString = {};
    uri.replace(
      new RegExp(
        "([^?=&]+)(=([^&#]*))?", "g"),
      function ($0, $1, $2, $3) { queryString[$1] = $3; }
    );
    return queryString;
  };
})(jQuery);


async function _fetchJson(url) {
  const res = await fetch(url);
  return await res.json();
}