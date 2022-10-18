/***************************
@File        : event_test.go
@Time        : 2022/10/17 14:43:48
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : test Event
****************************/

package event

import (
    "testing"
)

func TestEvent(t *testing.T) {

    c := NewContract("0x8028905608F120FaC0c90F8f467e6807A3D4e5f9", DepositAbi)
    c.GetHistoryLogs(22134550, Deposit{})
}
