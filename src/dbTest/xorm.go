package main

import (
	"github.com/go-xorm/xorm"
	"fmt"
	"time"
 _ "github.com/go-sql-driver/mysql"
	"math/rand"
	"strconv"
)

func mysqlEngine() (*xorm.Engine, error) {
	return xorm.NewEngine("mysql", "root:@192.168.99.64：3306/webcron?charset=utf8")
}
type t_task struct {
	Id int64
	User_id int64
	Group_id int64
	Task_name string `xorm:"varchar(200)"`
	Task_type int
	Description string `xorm:"varchar(200)"`
	Cron_spec string `xorm:"created"`
	Concurrent int64
	Status int
	Notify int
	Notify_email int
	Timeout int
	Execute_times int
	prev_time time.Time `xorm:"created"`
	create_time time.Time `xorm:"updated"`

}
type User struct {
	Id int
	Name string
	Salt string
	Age int
	Passwd string `xorm:"varchar(200)"`
	Created int `xorm:"created"`
	Updated int `xorm:"updated"`
}

// 批量生产数据
func GenBatchData() []*User {
	users := make([]*User, 0)
	for i := 1; i <= 10 ; i++  {
		user := &User{
			//Id:rand.Intn(100),
			Name:strconv.Itoa(rand.Int()),
			Salt:strconv.Itoa(rand.Int()),
			Age:rand.Intn(100),
			Passwd:strconv.Itoa(rand.Int()),
			Created:rand.Int(),
			Updated:rand.Int(),
		}
		users = append(users,user)
	}
	return users
}
// 生产单条数据
func GenOneData() *User {
	user := &User{
		Name:strconv.Itoa(rand.Intn(100)),
		Salt:strconv.Itoa(rand.Int()),
		Age:rand.Intn(100),
		Passwd:strconv.Itoa(rand.Intn(1000)),
		Created:rand.Int(),
		Updated:rand.Int(),
	}
	return user
}
func main()  {
	// 初始化数据库
	engine, _ := xorm.NewEngine("mysql", "root:123456@(192.168.99.64:3306)/webcron?charset=utf8")

	//engine.Sync2(&User{})
	//db, _ :=engine.DBMetas()
	//users := GenOneData()
	//tt := engine.TableInfo(users)
	// 测试数据
	//users := GenBatchData()

	//insert 批量插入数据  返回插入成功的数据个数，如果失败，返回0， err有错误信息，当然这样可以插入单条数据
	//affected, err := engine.Insert(users)

	//insert 单条数据插入 仅支持单条插入
	//affected,err := engine.InsertOne(GenOneData())
	//fmt.Println(affected, err)

	//select 查询数据
	//user  := &User{}
	// 获取单条数据的全部全部字段
	//has, err := engine.Get(User)
	// SELECT * FROM user LIMIT 1

	//has, err := engine.Where("Name = ?", "12").Desc("id").Get(user)
	// SELECT `id`, `name`, `salt`, `age`, `passwd`, `created`, `updated` FROM `user` WHERE (Name=12) ORDER BY `id` DESC LIMIT 1
	//fmt.Println(has, user, "6666")
	//has, err := engine.Where("id =115").Cols("name").Get(User)
	// SELECT `name` FROM `user` WHERE (id =115) LIMIT 1

	//has, err := engine.Where("Age > 50").Cols("name, id").GroupBy("name").Get(User)
	// SELECT `name`, `id` FROM `user` WHERE (Age > 50) GROUP BY name LIMIT 1

	//has, err := engine.Where("name=12").Cols("id").Get(User)
	// SELECT `id` FROM `user` WHERE (name=12) LIMIT 1

	// 获取多条数据
	//var users =[]User{}
	//err := engine.Where("name!=12").And("age >85").Limit(10,0).Find(&users)

	type UserDetail struct {
		User `xorm:"extends"`
		t_task `xorm:"extends"`
	}

	var uu []UserDetail

	//复杂查询多表查询
	err := engine.Table("user").Select("user.name,user.id,t_task.task_name").
		Join("left","t_task", "user.id=t_task.id").
		Where("user.name=12").Limit(1,0).
		Find(&uu)

	fmt.Println(err,  uu)
	return
	//var user User
	//user := new(User)
	//has, err := engine.Where("Name = ?", 11).Desc("id").Get(user)
	//var tt []Task
	results, err := engine.Query("select * from t_task")
	//has, err := engine.Where("name = ?", 123).Desc("id").Get(&tt)
	fmt.Println(results, err)
	//for k,v := range results {
	//	fmt.Println(k, v)
	//
	//}
	//fmt.Println(has,err)
}









