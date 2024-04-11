package models

import (
	"fmt"
	"strconv"
	"tiny-sec-kill/utils/errMsg"

	"gorm.io/gorm"
)

type Good struct {
	gorm.Model
	Gid      int    `gorm:"type:int;not null" json:"gid"`
	GoodName string `gorm:"type:varchar(20);not null " json:"goodName"`
	StockNum int    `gorm:"type:int;not null" json:"stockNum"`
}

// 新建商品
func CreateGood(data *Good, token string) int {
	user := TransToken(token)
	if user.Role < 2 {
		err := db.Create(&data).Error
		if err != nil {
			return errMsg.ERROR
		}
		return errMsg.SUCCESS
	} else {
		return errMsg.ERROR_USER_NO_PERMISSION
	}
}

// 获取商品信息
func GetGood(id int) (Good, int) {
	var good Good
	err := db.Where("gid = ?", id).First(&good).Error
	if err != nil {
		return good, errMsg.ERROR
	}
	return good, errMsg.SUCCESS
}

// 编辑并更改商品
func EditGood(id int, data *Good, token string) int {
	user := TransToken(token)
	if user.Role < 2 {
		var good Good
		var maps = make(map[string]interface{})
		maps["gid"] = data.Gid
		maps["good_name"] = data.GoodName
		maps["stock_num"] = data.StockNum
		//dd, _ := GetGood(id)
		//fmt.Println(dd)
		err := db.Debug().Model(&good).Where("Gid = ?", id).First(&good).Updates(maps).Error
		fmt.Println(maps)
		if err != nil {
			return errMsg.ERROR
		}
		return errMsg.SUCCESS
	} else {
		return errMsg.ERROR_USER_NO_PERMISSION
	}
}

// 商品预加载
func PreSecKill(id int, token string) (int, Good) {
	var good Good
	user := TransToken(token)
	//fmt.Println(user)
	if user.Role < 2 {
		db.Where("gid= ?", id).First(&good)
		fmt.Println(good)
		err := rdb.Set(strconv.Itoa(good.Gid), good.StockNum, 0).Err()
		if err != nil {
			fmt.Printf("set score failed, err:%v\n", err)
			return errMsg.ERROR, good
		}
		return errMsg.SUCCESS, good
	} else {
		return errMsg.ERROR_USER_NO_PERMISSION, good
	}
}


