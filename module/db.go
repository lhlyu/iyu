package module

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/lhlyu/iyu/common"
	"log"
	"time"
)

type db struct {
}

func (db) seq() int {
	return 1 << 2
}

func (db) SetUp() {
	log.Println("init db module ->")
	c := &dbConf{}
	if err := common.Cfg.UnmarshalKey("db.db_wr", c); err != nil {
		log.Fatal("db setup is err:", err)
	}

	setDb(c)
}

// 数据库连接模块
var DbModule = db{}

type dbConf struct {
	User            string `json:"user"`
	Password        string `json:"password"`
	Host            string `json:"host"`
	Port            int    `json:"port"`
	Database        string `json:"database"`
	Charset         string `json:"charset"`
	ParseTime       bool   `json:"parseTime"`
	MaxOpenConns    int    `json:"maxOpenConns"`
	MaxIdleConns    int    `json:"maxIdleConns"`
	ConnMaxLifetime int    `json:"connMaxLifetime"`
	Init            bool   `json:"init"`
}

func setDb(c *dbConf) {

	path := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=Local", c.User,
		c.Password,
		c.Host,
		c.Port,
		c.Database,
		c.Charset,
		c.ParseTime)

	db, err := gorm.Open("mysql", path)
	if err != nil {
		log.Fatal("db connect is fail,err:", err)
	}

	db.SingularTable(true)
	db.DB().SetMaxOpenConns(c.MaxOpenConns)
	db.DB().SetMaxIdleConns(c.MaxIdleConns)
	db.DB().SetConnMaxLifetime(time.Second * time.Duration(c.ConnMaxLifetime))

	if err != nil {
		panic(err)
	}
	err = db.DB().Ping()
	if err != nil {
		panic(err)
	}
	common.DB = db

	// 初始化

	initDatabase(c.Init)
}

// 数据表初始化：创建表 添加初始数据
// 如果表存在不会执行相应操作,如果想强行初始化，参数传true
func initDatabase(force bool) {
	//sql := ""
	//common.DB.Exec(sql)
}
