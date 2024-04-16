package models

import (
	"fmt"
	"time"
	"tiny-sec-kill/utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var db *gorm.DB
var err error

func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DbUser, // 使用具有创建数据库权限的用户（如root用户）
		utils.DbPassword,
		utils.DbIp,
		utils.DbPort,
	)
	rootDb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		fmt.Println(dsn)
		fmt.Println("连接数据库失败，请检查参数：", err)
		return
	}

	// 检查数据库是否存在，不存在则创建
	var count int64
	err = rootDb.Raw("SELECT COUNT(*) FROM INFORMATION_SCHEMA.SCHEMATA WHERE SCHEMA_NAME = ?", utils.DbName).Scan(&count).Error
	if err != nil {
		fmt.Println("检查数据库是否存在时出错：", err)
		return
	}

	if count == 0 {
		err = rootDb.Exec(fmt.Sprintf("CREATE DATABASE `%s` CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;", utils.DbName)).Error
		if err != nil {
			fmt.Println("创建数据库失败：", err)
			return
		}
		fmt.Println("成功创建数据库:", utils.DbName)
	}

	// 使用新创建的或已存在的数据库进行连接
	dsnWithDb := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassword,
		utils.DbIp,
		utils.DbPort,
		utils.DbName,
	)
	db, err = gorm.Open(mysql.Open(dsnWithDb), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
		// 外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
		// 禁用默认事务（提高运行速度）
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			// 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Println("连接数据库失败，请检查参数：", err)
		return
	}
	fmt.Println("成功链接mysql")

	db.Set("gorm:table_option", "ENGINE=InnoDB")
	_ = db.AutoMigrate(&User{}, &Good{}, &Order{})

	sqlDB, _ := db.DB()
	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenCons 设置数据库的最大连接数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetiment 设置连接的最大可复用时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)







	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	// 	utils.DbUser,
	// 	utils.DbPassword,
	// 	utils.DbIp,
	// 	utils.DbPort,
	// 	utils.DbName,
	// )

	// db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
	// 	// gorm日志模式：silent
	// 	Logger: logger.Default.LogMode(logger.Silent),
	// 	// 外键约束
	// 	DisableForeignKeyConstraintWhenMigrating: true,
	// 	// 禁用默认事务（提高运行速度）
	// 	SkipDefaultTransaction: true,
	// 	NamingStrategy: schema.NamingStrategy{
	// 		// 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
	// 		SingularTable: true,
	// 	},
	// })
	// if err != nil {
	// 	fmt.Println("连接数据库失败，请检查参数：", err)
	// } else {
	// 	fmt.Println("成功链接mysql")
	// }
	// db.Set("gorm:table_option", "ENGINE=InnoDB")
	// _ = db.AutoMigrate(&User{},&Good{},&Order{})
	// sqlDB, _ := db.DB()
	// // SetMaxIdleCons 设置连接池中的最大闲置连接数。
	// sqlDB.SetMaxIdleConns(10)

	// // SetMaxOpenCons 设置数据库的最大连接数量。
	// sqlDB.SetMaxOpenConns(100)

	// // SetConnMaxLifetiment 设置连接的最大可复用时间。
	// sqlDB.SetConnMaxLifetime(10 * time.Second)
}
