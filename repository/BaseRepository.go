package repository

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"gopkg.in/ini.v1"
	"log"
)

var db *gorm.DB
var Cfg *ini.File

type BaseRepository struct {
}

func init() {

	//读取配置
	var err error
	Cfg, err = ini.Load("config.ini")
	if err != nil {
		log.Fatalln("配置文件读取失败", err)
	}
	host := Cfg.Section("mysql").Key("host").String()
	port := Cfg.Section("mysql").Key("port").String()
	username := Cfg.Section("mysql").Key("username").String()
	password := Cfg.Section("mysql").Key("password").String()
	database := Cfg.Section("mysql").Key("database").String()

	//连接数据库
	connectionArgs := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, database)
	db, err = gorm.Open("mysql", connectionArgs)
	if err != nil {
		log.Fatal("数据库连接失败", err)
	}

	//设置最大空闲和连接数
	db.DB().SetMaxIdleConns(5)
	db.DB().SetMaxOpenConns(20)
}
