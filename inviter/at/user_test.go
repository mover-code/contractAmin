/***************************
@File        : user_test.go
@Time        : 2022/09/29 14:13:19
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : test
****************************/

package at

import (
    "fmt"
    "testing"
    "time"

    "github.com/pterm/pterm"
    "github.com/schollz/progressbar/v3"
)

var (
    bar          *progressbar.ProgressBar
    relationAddr = "0x38ec08010823E8BBC3F179fF19482B11dB33D27B"
    teamAddr     = "0x9c7176738f0cecf6605665a4f0e8ecdbc2f1b2e1"
    orderAddr    = "0x3Cd35800fbd06Bd4f3550a5b000882351B67ad4B"
    r            = NewRelation(relationAddr)
    team         = NewTeam(teamAddr)
    o            = NewOrder(orderAddr)
)

func TestAll(t *testing.T) {
    bar = progressbar.Default(7)
    fmt.Println("Start test all function:")
    TestUserSum(t)
    wait()
    TestUserAddr(t)
    wait()
    TestUserReferrer(t)
    wait()
    TestUserLevel(t)
    wait()

    TestPeroid(t)
    wait()

    TestInfo(t)
    wait()

    TestUserOrder(t)
    wait()

}

func wait() {
    time.Sleep(time.Second)
    bar.Add(1)
    fmt.Println("")
}

// `TestUserSum` is a function that takes a `testing.T` and returns nothing
//
// Args:
//   t: test case
func TestUserSum(t *testing.T) {
    pterm.NewRGB(201, 144, 30).Println(fmt.Sprintf("参与AT众筹总人数测试: 结果[%v]", r.UserSum()))
}

func TestUserAddr(t *testing.T) {
    pterm.NewRGB(201, 144, 30).Println(fmt.Sprintf("获取参与AT众筹账户地址测试: 结果[%v]", r.UserAddr(10)))
}

func TestUserReferrer(t *testing.T) {
    pterm.NewRGB(201, 144, 30).Println(fmt.Sprintf("获取参与AT众筹账户地址上级测试: 结果[%v]", r.UserReferrer("0x1d066261d7cc32d00e0ac909d4b8808ab0bdfa89")))
}

func TestUserLevel(t *testing.T) {
    pterm.NewRGB(201, 144, 30).Println(fmt.Sprintf("获取参与AT众筹账户地址等级测试: 结果[%v]", team.UserLevel("0xb98c3ac4fcf2807ac235164c5920c99abf089cf0")))
}

func TestPeroid(t *testing.T) {
    pterm.NewRGB(201, 144, 30).Println(fmt.Sprintf("获取参与AT众筹所有期数测试: 结果[%v]", o.Period()))
}

func TestInfo(t *testing.T) {
    pterm.NewRGB(201, 144, 30).Println(fmt.Sprintf("获取参与AT众筹期数详情测试: 结果[%v]", o.OrderInfo(1)))
}

func TestUserOrder(t *testing.T) {
    pterm.NewRGB(201, 144, 30).Println(fmt.Sprintf("获取参与AT众筹用户本期数信息测试: 结果[%v]", o.UserOrder("0xb98c3ac4fcf2807ac235164c5920c99abf089cf0", 2)))
}
