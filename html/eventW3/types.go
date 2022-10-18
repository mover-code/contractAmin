/***************************
@File        : types.go
@Time        : 2022/09/29 13:41:59
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        :
****************************/

package event

import (
	"fmt"
	"math/big"
	"reflect"
	"time"

	"github.com/GoAdminGroup/example/models"
	"github.com/mitchellh/mapstructure"
	web3 "github.com/mover-code/golang-web3"
	"github.com/mover-code/golang-web3/abi"
	"github.com/mover-code/golang-web3/contract"
	"github.com/mover-code/golang-web3/jsonrpc"
)

// `Deposit` is a struct that contains four fields, `FaNum`, `Inviter`, `UsdtNum`, and `User`.
//
// The first field, `FaNum`, is a pointer to a `big.Int`. The second field, `Inviter`, is a
// `web3.Address`. The third field, `UsdtNum`, is a pointer to a `big.Int`. The fourth field, `User`,
// is a `web3.Address`.
//
// The `Deposit` struct is used to represent a single deposit event.
// @property {string} Addr - The address of the contract
// @property Contract - The contract object that is used to call the contract.
type (
	MyContract struct {
		Addr     string
		Contract *contract.Contract
	}

	Deposit struct {
		FaNum   *big.Int     `json:"faNum"`
		Inviter web3.Address `json:"inviter"`
		UsdtNum *big.Int     `json:"usdtNum"`
		User    web3.Address `json:"User"`
	}
)

func (d Deposit) Parse() {
	d.UsdtNum = web3.FWei(d.UsdtNum)
	d.FaNum = web3.FWei(d.FaNum)
	amount := new(big.Int).Mul(d.UsdtNum, big.NewInt(2))
	u := models.FirstFaUser()
	u.User = d.User.String()
	u.Amount = amount.String()
	u.UId = d.Inviter.String()
	u.Insert()
	uId := u.FindByAddr(u.UId)
	var upUser models.FaUser
Loop:
	for {
		upUser.FindById(uId)
		// fmt.Println("上级", upUser)
		if upUser.Id > 0 {
			old, b := new(big.Int).SetString(upUser.TeamAmount, 0)
			if b {
				upUser.TeamAmount = new(big.Int).Add(amount, old).String()
			} else {
				upUser.TeamAmount = amount.String()
			}
			upUser.Update()
			uId = upUser.FindByAddr(upUser.UId)
			if uId == 0 {
				break Loop
			}
		} else {
			break Loop
		}
	}
}

// It creates a new client for the given url.
//
// Args:
//   url (string): The address of the RPC service.
func NewCli(url string) *jsonrpc.Client {
	cli, err := jsonrpc.NewClient(url)
	if err != nil {
		panic(fmt.Sprintf("error:%s", url))
	}
	return cli
}

// If the ABI is invalid, panic.
//
// Args:
//   s (string): The ABI string of the contract
func NewAbi(s string) *abi.ABI {
	a, err := abi.NewABI(s)
	if err != nil {
		panic("error")
	}
	return a
}

// A function that takes in a map of string to interface and a struct. It then decodes the map into the
// struct.
//
// Args:
//   d: the struct you want to load the data into
//   v: the map[string]interface{} that you want to decode
func Load(d interface{}, v map[string]interface{}) {
	if err := mapstructure.Decode(v, &d); err == nil {
		switch d.(type) {
		case Deposit:
			d.(Deposit).Parse()
		default:
			fmt.Println("sorry we did not support this type!")
		}
	}
}

// `NewContract` takes an address, an ABI string, and a data struct, and returns a `MyContract` struct
//
// Args:
//   addr (string): The address of the contract
//   abiStr (string): The ABI of the contract.
func NewContract(addr, abiStr string) *MyContract {
	return &MyContract{
		Addr:     addr,
		Contract: contract.NewContract(web3.HexToAddress(addr), NewAbi(abiStr), w3),
	}
}

// Parsing the log and loading the data into the struct.
func (d *MyContract) ParseLog(l *web3.Log, name ...interface{}) {
	for _, n := range name {
		if e, b := d.Contract.Event(reflect.TypeOf(n).Name()); b {
			data, err := e.ParseLog(l)
			if err == nil {
				Load(n, data)
			}
		}
	}
}

// Creating a filter for the contract events.
func (d *MyContract) NewFilter(name ...interface{}) *web3.LogFilter {
	topics := []*web3.Hash{}
	for _, n := range name {
		e, b := d.Contract.Event(reflect.TypeOf(n).Name())
		if b {
			topic := e.Encode()
			topics = append(topics, &topic)
		}
	}
	return &web3.LogFilter{
		Address: []web3.Address{web3.HexToAddress(d.Addr)},
		Topics:  topics,
	}
}

// It returns the current block number
func NowBlock() web3.BlockNumber {
	blockNumber, _ := w3.Eth().BlockNumber()
	return web3.BlockNumber(blockNumber)
}

// The above code is a Go function that is used to get the logs of a contract.
func (d *MyContract) GetLogs(name ...interface{}) {
	f := d.NewFilter(name...)

	logsInfo := make(chan *web3.Log)
	go func() {
		d := NowBlock()
		f.SetToUint64(uint64(d - 5))
		old := *f.To
		f.SetFromUint64(uint64(old))
		// start := 0
		for {
			time.Sleep(time.Second * 3)
			now := NowBlock()
			if now > old {
				logs, err := w3.Eth().GetLogs(f)
				if err == nil && len(logs) > 0 {
					for _, l := range logs {
						logsInfo <- l
					}
				}
				f.SetFromUint64(uint64(now))
				f.SetToUint64(uint64(now))
				old = now
			}
		}
	}()

	for {
		select {
		case l := <-logsInfo:
			d.ParseLog(l, name...)
		}
	}
}

// The above code is a function that is used to get the history logs of a contract.
func (d *MyContract) GetHistoryLogs(start int64, name ...interface{}) {
	f := d.NewFilter(name...)

	stop := false
	block := NowBlock()
	for {
		time.Sleep(time.Second)
		if start < int64(block) {
			newBlock := start + 5000

			// new := start
			if newBlock > int64(block) {
				newBlock = int64(block)
				stop = true
			}

			f.SetFromUint64(uint64(start))
			f.SetToUint64(uint64(newBlock))

			logs, err := w3.Eth().GetLogs(f)
			// fmt.Println(fmt.Sprintf("catch logs length %v;blockNum:from[%v]-to[%v] ----", len(logs), start, newBlock), err)
			if err == nil && len(logs) > 0 {
				// fmt.Println(fmt.Sprintf("catch logs length %v;blockNum:from[%v]-to[%v] ----", len(logs), start, newBlock))
				for _, l := range logs {
					// fmt.Println("block: ", l.BlockNumber)
					d.ParseLog(l, name...)
				}
			}
			start = newBlock
		}
		if stop {
			break
		}
	}

}

var (
	w3         = NewCli("https://bsc-dataseed3.defibit.io")
	DepositAbi = `[
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "_start",
				"type": "uint256"
			},
			{
				"internalType": "address",
				"name": "_treasury",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "_topInviter",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "_nftReward",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "_pointReward",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "_gnosis",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "_usdt",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "_fa",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "_fac",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "_faPair",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "_relation",
				"type": "address"
			}
		],
		"stateMutability": "nonpayable",
		"type": "constructor"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"internalType": "address",
				"name": "user",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "address",
				"name": "inviter",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "usdtNum",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "faNum",
				"type": "uint256"
			}
		],
		"name": "Deposit",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "address",
				"name": "previousOwner",
				"type": "address"
			},
			{
				"indexed": true,
				"internalType": "address",
				"name": "newOwner",
				"type": "address"
			}
		],
		"name": "OwnershipTransferred",
		"type": "event"
	},
	{
		"inputs": [],
		"name": "accTokenPerShare",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "blackhole",
		"outputs": [
			{
				"internalType": "address",
				"name": "",
				"type": "address"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "usdtAmount",
				"type": "uint256"
			}
		],
		"name": "calculateFaAmount",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "user",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "index",
				"type": "uint256"
			}
		],
		"name": "checkInviter",
		"outputs": [
			{
				"internalType": "bool",
				"name": "",
				"type": "bool"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "usdtAmount",
				"type": "uint256"
			},
			{
				"internalType": "address",
				"name": "inviter",
				"type": "address"
			}
		],
		"name": "deposit",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "faToken",
		"outputs": [
			{
				"internalType": "contract IERC20",
				"name": "",
				"type": "address"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "facToken",
		"outputs": [
			{
				"internalType": "contract IERC20",
				"name": "",
				"type": "address"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "getCurrentPrice",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "user",
				"type": "address"
			}
		],
		"name": "getInviteNum",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "user",
				"type": "address"
			}
		],
		"name": "getInviter",
		"outputs": [
			{
				"internalType": "address",
				"name": "",
				"type": "address"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "_from",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "_to",
				"type": "uint256"
			}
		],
		"name": "getMultiplier",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "pure",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "gnosisAddress",
		"outputs": [
			{
				"internalType": "address",
				"name": "",
				"type": "address"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "lastRewardBlock",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "minAmount",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "nftReward",
		"outputs": [
			{
				"internalType": "address",
				"name": "",
				"type": "address"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "owner",
		"outputs": [
			{
				"internalType": "address",
				"name": "",
				"type": "address"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "pairAddress",
		"outputs": [
			{
				"internalType": "address",
				"name": "",
				"type": "address"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "_user",
				"type": "address"
			}
		],
		"name": "pendingFAC",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "pointReward",
		"outputs": [
			{
				"internalType": "address",
				"name": "",
				"type": "address"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "relation",
		"outputs": [
			{
				"internalType": "contract IRelation",
				"name": "",
				"type": "address"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "renounceOwnership",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "_treasury",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "_topInviter",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "_nftReward",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "_pointReward",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "_gnosis",
				"type": "address"
			}
		],
		"name": "setAddress",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "num",
				"type": "uint256"
			}
		],
		"name": "setMinAmount",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "num",
				"type": "uint256"
			}
		],
		"name": "setPrice",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "num",
				"type": "uint256"
			}
		],
		"name": "setTokenPerBlock",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "settlePrice",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "startBlock",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "_token",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "user",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "amount",
				"type": "uint256"
			}
		],
		"name": "takeOut",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "tokenPerBlock",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "topInviter",
		"outputs": [
			{
				"internalType": "address",
				"name": "",
				"type": "address"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "totalPower",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "newOwner",
				"type": "address"
			}
		],
		"name": "transferOwnership",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "treasury",
		"outputs": [
			{
				"internalType": "address",
				"name": "",
				"type": "address"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "update",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "usdt",
		"outputs": [
			{
				"internalType": "contract IERC20",
				"name": "",
				"type": "address"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "",
				"type": "address"
			}
		],
		"name": "userInfo",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "faAmount",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "usdtPower",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "totalPower",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "hasPower",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "rewardDebt",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "withdraw",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	}
]`
)
