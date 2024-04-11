package models

import (
	"fmt"
	"strconv"
	"sync"
	"tiny-sec-kill/utils/errMsg"
	"tiny-sec-kill/middlewares/lockers"

	"gorm.io/gorm"
)

var mutex sync.RWMutex

type Order struct {
	User User `gorm:"foreignkey:Uid"`
	gorm.Model
	Ip   string `gorm:"type:varchar(20);not null " json:"ip"`
	Gid  int    `gorm:"type:int;not null" json:"gid"`
	Uid  int    `gorm:"type:int;not null" json:"uid"`
	Msg  string `gorm:"type:varchar(20);not null " json:"msg"`
	Left int    `gorm:"type:int;not null" json:"left"`
}

func sec(id int) (int, int64) {
	stock := rdb.Decr(strconv.Itoa(id)).Val()
	if stock >= 0 {
		return errMsg.SUCCESS, stock
	} else {
		return errMsg.ERROR, -1
	}
}

func SecKill(id int, order *Order, name string) int {
	code, left := sec(id)
	if code == errMsg.SUCCESS {
		var user User
		order.Msg = "秒杀成功"
		order.Left = int(left)
		err := db.Where("username = ?", name).First(&user).Error
		if err != nil {
			return errMsg.ERROR
		}
		order.Uid = int(user.ID)

		err = db.Create(&order).Error
		if err != nil {
			return errMsg.ERROR
		}
		var good Good
		err = db.Where("gid=?", id).First(&good).Error
		if err != nil {
			return errMsg.ERROR
		}

		//db.Model(&Good{}).Where("gid=?", id).Update("stock_num", gorm.Expr("stock_num - ?", 1))
		lockers.PremissionLock(&mutex)
		db.Model(&Good{}).Where("gid=?", id).Update("stock_num", gorm.Expr("stock_num - ?", 1))
		lockers.PremissionUnLock(&mutex)

		fmt.Println(good.StockNum - 1)
		return errMsg.SUCCESS
	} else {
		order.Msg = "库存不足，秒杀失败"
		order.Left = -1
		//err := db.Create(&data).Error
		if err != nil {
			return errMsg.ERROR
		}
		return code
	}
}
