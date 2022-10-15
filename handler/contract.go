/***************************
@File        : contract.go
@Time        : 2022/10/14 16:02:21
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        :
****************************/

package handler

import (
    "net/http"

    "github.com/GoAdminGroup/example/models"
    "github.com/gin-gonic/gin"
)

type Response struct {
    Msg  string      `json:"msg"`
    Data interface{} `json:"data"`
    Err  string      `json:"err"`
}

func Contract(c *gin.Context) {
    var (
        msg, err string
        data     interface{}
    )
    addr, _ := c.GetQuery("addr")
    if addr == "" {
        msg = "查寻合约地址不能为空"
    } else {
        contract := models.FirstContract()
        data = contract.ContractInfo(addr)
    }
    c.JSON(http.StatusOK, &Response{Data: &data, Msg: msg, Err: err})
}
