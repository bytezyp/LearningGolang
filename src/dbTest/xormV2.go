package main

import (
	"github.com/go-xorm/xorm"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"os"
)

type Config struct {
	Name   string `xorm:"not null index varchar(20)" form:"name" json:"name"`
	Value  string `xorm:"varchar(1024)" form:"value" json:"value"`
	Label  string `xorm:"varchar(40)" form:"label" json:"label"`
	Format string `xorm:"varchar(10)" form:"format" json:"format"`
}

func main()  {
	engine, _ := xorm.NewEngine("mysql", "root:123456@(192.168.99.64:3306)/webcron?charset=utf8")
	engine.ShowSQL(true) // 在控制台，打印sql
	engine.Logger().SetLevel(core.LOG_INFO) // 一般日志输出级别,输出到控制台


	f,err := os.Create("sql.log")
	if err != nil {
		panic(err)
	}
	engine.SetLogger(xorm.NewSimpleLogger(f))
	//logger := xorm.NewSimpleLogger()
	//con := &Config{
	//	Name:"demo4",
	//	Value:"zhang4",
	//	Label:"demo4",
	//	Format:"ceshi4",
	//}
	//inser, err := engine.Insert(con)
	//var conn Config
	//cc := make(map[int]Config, 100)
	res,err := engine.Query("select * from config")
	for _, v:=range res{
		//fmt.Println(k)
		for _,vv := range v {
			fmt.Println(string(vv), vv)
		}
	}

	fmt.Println(res,err)
}
