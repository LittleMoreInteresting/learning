package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)
type User struct {
	Id int
	Name string
	Age int
}

func connect() (*sql.DB,error) {
	user := "root"
	pwd := "root"
	host := "127.0.0.1"
	dbname := "gorm_demo"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",user,pwd,host,dbname)
	return sql.Open("mysql", dsn)
}

func QueryUser(user *User,id int) error {
	db, err := connect()
	if err != nil {
		return  err
	}
	querysql := "select id,name,age from users where id = ?;"
	err = db.QueryRow(querysql, id).Scan(&user.Id,&user.Name,&user.Age)
	if err == sql.ErrNoRows {
		return  nil
	}
	return err
}
func main() {
	/**
	1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。
	为什么，应该怎么做请写出代码？
	 */
	// Answer:  dao 层遇到 sql.ErrNoRows 应该在dao层处理这个错误，不抛给上层，
	// 业务层可以根据查询结课判断是否存在数据。也可以区分数据库层面的错误和业务数据的异常
	user := &User{}
	err := QueryUser(user,1)
	if err != nil {
		fmt.Printf("Err:%v\n",err)
		return
	}
	if user.Id > 0 {
		fmt.Printf("User:%v\n", user)
		return
	}
	fmt.Printf("User Not Found\n")
}

