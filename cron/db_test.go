/***************************
@File        : db_test.go
@Time        : 2022/09/30 10:37:07
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : test
****************************/

package job

import (
    "fmt"
    "testing"
    "time"

    "github.com/pterm/pterm"
    "github.com/schollz/progressbar/v3"
)

var bar *progressbar.ProgressBar

func wait() {
    time.Sleep(time.Second)
    bar.Add(1)
    fmt.Println("")
}

func TestAll(t *testing.T) {
    bar = progressbar.Default(5)
    fmt.Println("Start test all function:")

    TestSum(t)
    wait()

    TestInsert(t)
    wait()

    TestExceptSum(t)
    wait()

    TestAddExcept(t)
    wait()

    TestOrder(t)
    wait()
}

func TestSum(t *testing.T) {
    pterm.NewRGB(201, 144, 30).Println(fmt.Sprintf("AT数据库总人数测试: 结果[%v]", userSum()))
}

func TestInsert(t *testing.T) {
    pterm.NewRGB(201, 144, 30).Println(fmt.Sprintf("AT数据库插入用户测试: 结果[%v]", insertUser("bbbb", "aaaa", 2)))
}

func TestExceptSum(t *testing.T) {
    pterm.NewRGB(201, 144, 30).Println(fmt.Sprintf("AT数据库总期数测试: 结果[%v]", exceptSum()))
}

func TestAddExcept(t *testing.T) {
    e := Expect{IsOver: 1, IsWin: 1, Expect: 10, Count: 1}
    pterm.NewRGB(201, 144, 30).Println(fmt.Sprintf("AT数据库插入用户测试: 结果[%v]", insertExpect(e)))
}

func TestOrder(t *testing.T) {
    o := Order{UserId: 2, ExpectId: 1, Amount: 0.02}
    pterm.NewRGB(201, 144, 30).Println(fmt.Sprintf("AT数据库插入用户测试: 结果[%v]", insertOrder(o)))
}
