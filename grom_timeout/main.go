package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	// 初始化 MySql
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


type Product struct {
	gorm.Model
	Code sql.NullString
	Price uint
}

func main() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel: logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:false,
		},
	)

	dsn := "root:root@tcp(127.0.0.1:3306)/gorm_demo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(),5*time.Second)
	begin := db.WithContext(ctx).Begin(&sql.TxOptions{Isolation: sql.LevelSerializable})
	go func(){
		select {
		case <-ctx.Done():
			begin.Rollback()
			cancel()
			fmt.Println("Time out")
		}
	}()
	//db.Create(&Product{Code: "Meta40", Price: 5000})
	var product Product
	//db.First(&product, 1) // 根据整形主键查找
	begin.First(&product, "code = ?", "Meta40") // 查找 code 字段值为 D42 的记录
	fmt.Println(product)
	len := begin.Model(&product).Updates(Product{Price: 1200, Code: sql.NullString{"", true}}).RowsAffected
	fmt.Println(len)
	time.Sleep(6*time.Second)
	begin.Commit()

}
