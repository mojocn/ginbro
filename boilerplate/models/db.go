package models

import (
	"fmt"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

//redis client
var redisDB *redis.Client
var mysqlDB *gorm.DB

const redisPrefix = "ginbro:"

func init() {
	//initializing redis client
	redisAddr, redisPassword := viper.GetString("redis.addr"), viper.GetString("redis.password")
	if redisAddr != "" {
		redisDB = redis.NewClient(&redis.Options{
			Addr:     redisAddr,
			Password: redisPassword,                // no password set
			DB:       viper.GetInt("redis.db_idx"), // use default DB
		})
		if pong, err := redisDB.Ping().Result(); err != nil || pong != "PONG" {
			logrus.WithError(err).Fatal("could not connect to the redis server")
		}
	}

	//init mysql
	conn := fmt.Sprintf("%s:%s@(%s)/%s?charset=%s&parseTime=True&loc=Local", viper.GetString("mysql.user"),
		viper.GetString("mysql.password"), viper.GetString("mysql.addr"), viper.GetString("mysql.database"),
		viper.GetString("mysql.charset"))
	if db, err := gorm.Open("mysql", conn); err == nil {
		mysqlDB = db
	} else {
		logrus.WithError(err).Fatalln("initialize mysql database failed")
	}
	//enable Gorm mysql log
	if flag := viper.GetBool("app.enable_sql_log"); flag {
		mysqlDB.LogMode(flag)
		//f, err := os.OpenFile("mysql_gorm.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		//if err != nil {
		//	logrus.WithError(err).Fatalln("could not create mysql gorm log file")
		//}
		//logger :=  New(f,"", Ldate)
		//mysqlDB.SetLogger(logger)
	}
	//mysqlDB.AutoMigrate()

}

//Close clear db collection
func Close() {
	if mysqlDB != nil {
		mysqlDB.Close()
	}
	if redisDB != nil {
		redisDB.Close()
	}
}
