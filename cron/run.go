/***************************
@File        : run.go
@Time        : 2022/09/30 09:53:34
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : 计划任务
****************************/

package job

import (
	"AT/at"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/robfig/cron"
	"github.com/shopspring/decimal"
)

var (
    relationAddr = "0x38ec08010823E8BBC3F179fF19482B11dB33D27B"
    teamAddr     = "0x9c7176738f0cecf6605665a4f0e8ecdbc2f1b2e1"
    orderAddr    = "0x3Cd35800fbd06Bd4f3550a5b000882351B67ad4B"
    r            = at.NewRelation(relationAddr)
    team         = at.NewTeam(teamAddr)
    o            = at.NewOrder(orderAddr)
    wait         sync.WaitGroup
)

func Run() {
    c := cron.New()
    c.AddFunc("@every 1h", CheckNewUser)
    c.AddFunc("@every 4h", CheckNewPeriod)
    c.AddFunc("@every 4h", CheckExpect)
    c.Start()
}

func CheckNewUser() {
    dbSum := userSum()
    cSum := r.UserSum()
    fmt.Println("new user ", cSum-dbSum)

    if dbSum != cSum {
        // bar := progressbar.Default(cSum - dbSum)

        for i := dbSum; i < cSum; i++ {
            addr := r.UserAddr(i)
            refer := r.UserReferrer(addr)
            level := team.UserLevel(addr)
            insertUser(addr, refer, level)
            // fmt.Println(i)
            time.Sleep(10 * time.Millisecond)
            // bar.Add(1)
            // bar.Describe(fmt.Sprintf("获取用户数进行...[%v]", i))
        }
    }
}

func CheckNewPeriod() {
    // fmt.Println("check except")
    dbSum := exceptSum()
    cSum := o.Period()
    // if dbSum == 0 {
    dbSum++

    // dbSum = 15
    // fmt.Println(dbSum, cSum)
    if dbSum != cSum {
        for i := dbSum; i < cSum; i++ {
            time.Sleep(time.Second)
            order := o.OrderInfo(i)
            amount, _ := ToDecimal(order.(map[string]interface{})["gole"], 18).Float64()
            // fmt.Println(order)
            e := Expect{
                IsOver: order.(map[string]interface{})["rewardType"].(*big.Int).Int64(),
                Expect: order.(map[string]interface{})["id"].(*big.Int).Int64(),
                Count:  order.(map[string]interface{})["roundId"].(*big.Int).Int64(),
                IsWin:  order.(map[string]interface{})["state"].(*big.Int).Int64(),
                Amount: amount,
            }
            insertExpect(e)

        }
        // TODO 更新状态
    }
}

func CheckExpect() {
    NotOver := overExpect()
    for _, ord := range NotOver {
        order := o.OrderInfo(ord.Expect)
        // amount, _ := ToDecimal(order.(map[string]interface{})["gole"], 18).Float64()
        over := order.(map[string]interface{})["rewardType"].(*big.Int).Int64()
        if over == 1 {
            ord.IsOver = 1
            orm.Updates(&ord)
            UserForExpect(ord.Expect)
            fmt.Println(fmt.Sprintf("更新第%v期业绩供查询", ord.Expect))
        }
        time.Sleep(time.Second * 10)
    }
}

func UserForExpect(index int64) {
    if ordAmount(index) == expectAmount(index) {
        ords := ordOrder(index)
        for _, ord := range ords {
            u := ordUser(ord.UserId)
            u.Amount += ord.Amount
            updateUser(u)
            ups := upsUser(int64(u.ID))
            // if len(ups) == 0{
            //     fmt.Println(u.ID)
            // }
            // fmt.Println("添加业绩==>", i, u.ID, u.Amount)
            for _, upU := range ups {
                upU.TeamAmount += ord.Amount
                // fmt.Println("向上级添加业绩==>", upU.ID, upU.TeamAmount)
                // time.Sleep(time.Second)
                updateUser(upU)
            }
        }

    } else {
        user := allUser()
        // bar := progressbar.Default(int64(len(user)), fmt.Sprintf("获取用户众筹第%v期", index))
        // 使用协程导致rpc挂掉
        // ch := make(chan struct{}, 5)

        for _, u := range user {
            // wait.Add(1)
            time.Sleep(time.Millisecond * 100)
            // ch <- struct{}{}
            // go func(u User, index int64) {
            order := o.UserOrder(u.Addr, index)
            amount, _ := ToDecimal(order, 18).Float64()
            if amount > 0 {
                // fmt.Println(order, u.ID, amount)
                u.Amount += amount
                updateUser(u)
                insertOrder(Order{
                    UserId:   int64(u.ID),
                    ExpectId: index,
                    Amount:   amount,
                })
                ups := upsUser(int64(u.ID))
                for _, upU := range ups {
                    upU.TeamAmount += amount
                    updateUser(upU)
                }
                // bar.Describe(fmt.Sprintf("获取用户众筹投入进行...[%v:%v]", u.ID, amount))
            }
            // bar.Add(1)
            //     <-ch
            //     wait.Done()
            // }(u, index)
            // wait.Wait()
        }
    }
    fmt.Println("第", index, "期业绩统计")
}

func ToDecimal(ivalue interface{}, decimals int) decimal.Decimal {
    value := new(big.Int)
    switch v := ivalue.(type) {
    case string:
        value.SetString(v, 10)
    case *big.Int:
        value = v
    }

    mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
    num, _ := decimal.NewFromString(value.String())
    result := num.Div(mul)

    return result
}
