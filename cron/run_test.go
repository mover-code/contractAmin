/***************************
@File        : run_test.go
@Time        : 2022/09/30 11:55:36
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : test
****************************/

package job

import (
	"testing"
	"time"
)

func TestCheckUser(t *testing.T) {
    CheckNewUser()
}

func TestOrderInfo(t *testing.T) {
    CheckNewPeriod()
}

func TestUserInfo(t *testing.T) {
    for i := 1; i < 2; i++ {
        UserForExpect(int64(i))
        time.Sleep(time.Second * 10)
        // upsUser(51)
    }
}
