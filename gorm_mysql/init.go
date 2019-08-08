package gorm_mysql

import (
	"awesomePet/api/debug"
	. "awesomePet/models"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB // 全局变量用 =

func Init(args *string) {
	var err error
	db, err = gorm.Open("mysql", *args)
	debug.PanicErr(err)
	db.DB().SetMaxIdleConns(100)
	db.DB().SetMaxOpenConns(1000)
	db.LogMode(true)       // 启用Logger，显示详细日志
	db.SingularTable(true) // 全局禁用表名复数
	fmt.Println("mysql数据库已连接，检查表结构中...")
	// 暂时手动创建
	if !db.HasTable(&User{}) {
		fmt.Println("表:user不存在，正在创建中")
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&User{}).Error; err != nil {
			panic(err)
		}
	}
	if !db.HasTable(&UserInfo{}) {
		fmt.Println("表:userInfo不存在，正在创建中")
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&UserInfo{}).Error; err != nil {
			panic(err)
		}
	}
	if !db.HasTable(&Pet{}) {
		fmt.Println("表:pet不存在，正在创建中")
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&Pet{}).Error; err != nil {
			panic(err)
		}
	}
	if !db.HasTable(&Pic{}) {
		fmt.Println("表:pic不存在，正在创建中")
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&Pic{}).Error; err != nil {
			panic(err)
		}
	}
	//defer db.Close()
}
