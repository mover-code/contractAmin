/***************************
@File        : type.go
@Time        : 2022/09/30 10:15:02
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        :
****************************/

package job

type (
    User struct {
        ID         uint `gorm:"primarykey"`
        Addr       string
        Team       int64
        Level      int64
        Amount     float64
        TeamAmount float64
    }

    Refer struct {
        ID      uint `gorm:"primarykey"`
        UserId  int64
        ReferId int64
    }

    Expect struct {
        ID     uint `gorm:"primarykey"`
        IsOver int64
        IsWin  int64
        Expect int64
        Count  int64
        Amount float64
    }

    Order struct {
        ID       uint `gorm:"primarykey"`
        UserId   int64
        ExpectId int64
        Amount   float64
    }
)
