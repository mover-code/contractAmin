/***************************
@File        : searchInviter.go
@Time        : 2022/05/12 15:17:16
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : search inviter user
****************************/

package at

import (
	"math/big"

	web3 "github.com/mover-code/golang-web3"
)

// Getting the number of users in the contract.
func (r *Relation) UserSum() int64 {
    if data, err := r.Contract.Call("getUsersNum", NowBlock()); err == nil {
        // fmt.Println(data)
        return data["0"].(*big.Int).Int64()
    }
    return 0
}

// Getting the address of the user.
func (r *Relation) UserAddr(index int64) string {
    if data, err := r.Contract.Call("users", NowBlock(), big.NewInt(index)); err == nil {
        return data["0"].(web3.Address).String()
    }
    return ""
}

// Getting the referrer of the user.
func (r *Relation) UserReferrer(addr string) string {
    if data, err := r.Contract.Call("getReferrer", NowBlock(), web3.HexToAddress(addr)); err == nil {
        return data["0"].(web3.Address).String()
    }
    return ""
}

// Getting the level of the user.
func (t *Team) UserLevel(addr string) int64 {
    if data, err := t.Contract.Call("getUserLevel", NowBlock(), web3.HexToAddress(addr)); err == nil {
        return data["0"].(*big.Int).Int64()
    }
    return 0
}

// Getting the total period of the contract.
func (o *Order) Period() int64 {
    if data, err := o.Contract.Call("getTotalPeriod", NowBlock()); err == nil {
        return data["0"].(*big.Int).Int64()
    }
    return 0
}

// Getting the information of the order.
func (o *Order) OrderInfo(index int64) interface{} {
    if data, err := o.Contract.Call("getFundInfo", NowBlock(), big.NewInt(index)); err == nil {
        return data["0"]
    }
    return nil
}

// Getting the information of the order.
func (o *Order) UserOrder(addr string, index int64) string {
    if data, err := o.Contract.Call("getUserInfo", NowBlock(), web3.HexToAddress(addr), big.NewInt(index)); err == nil {
        return data["0"].(map[string]interface{})["crowdfunding"].(*big.Int).String()
    }
    return "0"
}
