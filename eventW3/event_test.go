/*
 * @Author: small_ant xms.chnb@gmail.com
 * @Time: 2023-04-13 17:35:03
 * @LastAuthor: small_ant xms.chnb@gmail.com
 * @lastTime: 2023-04-13 18:09:39
 * @FileName: event_test
 * @Desc: test event 
 *
 * Copyright (c) 2023 by small_ant, All Rights Reserved.
 */
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
    c.GetHistoryLogs(22158180, Deposit{})
}
