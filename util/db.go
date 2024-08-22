package util

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

const DB_USERNAME = "root"

const DB_PASSWORD = "Baidu01)!"

// const DB_PASSWORD = "Admin01)!"
const DB_NAME = "apptools"

const DB_HOST = "10.138.170.14"

//const DB_HOST = "sh-cynosdbmysql-grp-iyzml20m.sql.tencentcdb.com"

// const DB_HOST = "127.0.0.1"

const DB_PORT = "8306"

//const DB_PORT = "23591"

var Db *gorm.DB

func InitDb() *gorm.DB {
	if Db != nil {
		return Db
	}
	Db = connectDB()
	return Db
}

func connectDB() *gorm.DB {
	var err error
	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" +
		"charset=utf8mb4&parseTime=true&loc=Local"
	fmt.Println("dsn : ", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名，启用该选项，此时，`Article` 的表名应该是 `it_article`
		},
	})
	shutdown := make(chan os.Signal)
	//监听指定信号 ctrl+c kill
	signal.Notify(shutdown, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM,
		syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
	go func() {
		for s := range shutdown {
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				if(Db != nil) {
					dbInstance, _ := Db.DB()
					_ = dbInstance.Close()
				}
				fmt.Println("Program Exit...", s)
				os.Exit(0)
				//GracefullExit()
		//case syscall.SIGUSR1:
		//		fmt.Println("usr1 signal", s)
		//	case syscall.SIGUSR2:
		//		fmt.Println("usr2 signal", s)
			default:
				fmt.Println("other signal", s)
			}
		}
	}()
	if err != nil {
		fmt.Println("Error connecting to database : error=%v", err)
		return nil
	}

	return db
}

func InitRedisDB() *redis.Client{
	rdb := redis.NewClient(&redis.Options{
		Addr:	  "10.138.170.14:8379",
		Password: "Baidu01)!", // 没有密码，默认值
		DB:		  0,  // 默认DB 0
	})
	return rdb
}
