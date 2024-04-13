package models

import (
	"encoding/base64"
	"fmt"
	"log"
	"strings"
	"tiny-sec-kill/utils/errMsg"
	middleware "tiny-sec-kill/middlewares"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null " json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(500);not null" json:"password" validate:"required,min=6,max=120" label:"密码"`
	Role     int    `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=2" label:"角色码"`
}

// 检查用户是否存在
func CheckUser(name string) (code int) {
	var user User
	db.Select("id").Where("username = ?", name).First(&user)
	if user.ID > 0 {
		return errMsg.ERROR //1001
	}
	return errMsg.SUCCESS

}

// 新建用户
func CreateUser(data *User) int {
	fmt.Println(data.Role)
	//data.Password = ScryptPw(data.Password)
	err := db.Create(&data).Error
	if err != nil {
		return errMsg.ERROR // 500
	}
	return errMsg.SUCCESS
}

// 查询用户列表
func GetUsers(pageSize int, pageNum int, token string) ([]User, int64) {
	user := TransToken(token)
	if user.Role == 0 {
		var users []User
		var total int64
		fmt.Println(pageSize)
		fmt.Println(pageNum)
		err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Count(&total).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			fmt.Println(err)
			return nil, 0
		}
		return users, total
	} else {
		return nil, -1
	}
}

// 编辑更新用户
func EditUser(id int, data *User, token string) int {
	var user User
	auth := TransToken(token)
	if auth.Role == 0 {
		var maps = make(map[string]interface{})
		maps["username"] = data.Username
		maps["password"] = data.Password
		err = db.Model(&user).Where("id = ?", id).Updates(maps).Error
		if err != nil {
			return errMsg.ERROR
		}
		return errMsg.SUCCESS
	} else {
		return errMsg.ERROR_USER_NO_PERMISSION
	}

}

// 更改密码
func ChangePassword(id int, data *User) int {

	err = db.Select("password").Where("id = ?", id).Updates(&data).Error
	if err != nil {
		return errMsg.ERROR
	}
	return errMsg.SUCCESS
}

// 删除用户
func DeleteUser(id int, token string) int {
	var user User
	auth := TransToken(token)
	if auth.Role == 0 {
		err = db.Where("id = ?", id).Delete(&user).Error
		if err != nil {
			return errMsg.ERROR
		}
		return errMsg.SUCCESS
	} else {
		return errMsg.ERROR_USER_NO_PERMISSION
	}
}

// GetUser 查询用户
func GetUser(id int) (User, int) {
	var user User
	err := db.Limit(1).Where("ID = ?", id).Find(&user).Error
	if err != nil {
		return user, errMsg.ERROR
	}
	return user, errMsg.SUCCESS
}

// 对于用户密码进行加密
func (u *User) BeforeSave(_ *gorm.DB) (err error) {
	u.Password = ScryptPw(u.Password)
	//u.Role = 2
	return nil
}
func ScryptPw(password string) string {
	const KeyLen = 10
	salt := make([]byte, 0, 8)
	salt = []byte{12, 10, 5, 61, 89, 44, 53, 11}
	HashPw, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}
	fpw := base64.StdEncoding.EncodeToString(HashPw)
	return fpw
}
func CheckLogin(username string, password string) int {
	var user User

	db.Where("username = ?", username).First(&user)
	if user.ID == 0 {
		return errMsg.ERROR
	}
	if ScryptPw(password) != user.Password {
		return errMsg.ERROR
	}
	return errMsg.SUCCESS
}

func TransToken(token string) User {
	var user User
	checkToken := strings.SplitN(token, " ", 2)
	claims, _ := middleware.CheckToken(checkToken[1])
	db.Where("username = ?", claims.Username).First(&user)
	fmt.Println(user)
	return user
}

