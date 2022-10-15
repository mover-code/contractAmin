/***************************
@File        : db.go
@Time        : 2022/09/30 10:02:04
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : db
****************************/

package job

import (
    "github.com/GoAdminGroup/example/models"
)

var orm = models.Orm

func userSum() (sum int64) {
    orm.Model(User{}).Count(&sum)
    return
}

// insert new user and user referee
func insertUser(addr, referee string, level int64) (id int64) {
    u := User{Addr: addr, Level: level}
    orm.Create(&u)
    var refereeUser User
    orm.Select("id").Where("addr = ?", referee).Take(&refereeUser)
    var sum int64
    orm.Model(Refer{}).Where("user_id = ? and refer_id = ?", u.ID, refereeUser.ID).Count(&sum)
    if sum == 0 {
        re := Refer{UserId: int64(u.ID), ReferId: int64(refereeUser.ID)}
        orm.Create(&re)
        id = int64(re.ID)
        return
    }
    return
}

func exceptSum() (sum int64) {
    orm.Model(Expect{}).Count(&sum)
    return
}

func insertExpect(e Expect) (id int64) {
    orm.Create(&e)
    id = int64(e.ID)
    return
}

func insertOrder(o Order) (id int64) {
    orm.Create(&o)
    id = int64(o.ID)
    return
}

func allUser() (users []User) {
    orm.Select("id", "addr").Find(&users)
    return
}

func overExpect() (expects []Expect) {
    orm.Model(Expect{}).Where("is_over=?", 0).Find(&expects)
    return
}

func updateUser(u User) {
    orm.Model(&u).Updates(&u)
}

func upsUser(id int64) (ups []User) {
    for {
        var up int64
        orm.Model(Refer{}).Where("user_id=?", id).Select("refer_id").Scan(&up)
        if up > 0 {
            var upUser User
            orm.Model(User{}).Where("id=?", up).First(&upUser)
            ups = append(ups, upUser)
            id = up
        } else {
            break
        }
    }
    // fmt.Println("获取此账号所有上级",ups)
    return
}

func ordAmount(index int64) (amount float64) {
    orm.Model(Order{}).Where("expect_id=?", index).Select("sum(amount)").Scan(&amount)
    return
}

func expectAmount(index int64) (amount float64) {
    orm.Model(Expect{}).Where("expect=?", index).Select("amount").Scan(&amount)
    return
}

func ordOrder(index int64) (ords []Order) {
    orm.Model(Order{}).Where("expect_id=?", index).Find(&ords)
    return
}

func ordUser(id int64) (u User) {
    orm.Model(User{}).Where("id=?", id).First(&u)
    return
}
