/***************************
@File        : run.go
@Time        : 2022/09/30 09:53:34
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : 计划任务
****************************/

package job

import (
    "fmt"
    "math/big"

    "github.com/robfig/cron"
    "github.com/shopspring/decimal"
)

var ()

func Run() {
    c := cron.New()
    c.AddFunc("@every 1h", CheckNewUser)
    c.Start()
}

func CheckNewUser() {
    fmt.Println("CheckNewUser")
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
