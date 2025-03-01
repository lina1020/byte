package database

//package main
import (
	"byteSmart/src/constant/config"
	"byteSmart/src/utils/logging"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var Client *gorm.DB

func main() {
	var err error
	gormLogrus := logging.GetGormLogger()
	cfg := gorm.Config{
		PrepareStmt: true,
		Logger:      gormLogrus,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: config.EnvCfg.PostgreSQLSchema + "." + config.EnvCfg.MySQLPrefix,
		},
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.EnvCfg.MySQLUser,
		config.EnvCfg.MySQLPassword,
		config.EnvCfg.MySQLHost,
		config.EnvCfg.MySQLPort,
		config.EnvCfg.MySQLDataBase,
	)
	if Client, err = gorm.Open(mysql.Open(dsn), &cfg); err != nil {
		fmt.Println(dsn)
		panic(err)
	}

	testConnection()
	//读写分离、连接池、链路追踪配置
	/* 		// 读写分离配置
	   		if config.EnvCfg.MySQLReplicaState == "enable" {
	   			var replicas []gorm.Dialector
	   			for _, addr := range strings.Split(config.EnvCfg.MySQLReplicaAddress, ",") {
	   				pair := strings.Split(addr, ":")
	   				if len(pair) != 2 {
	   					continue
	   				}

	   				replicaDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	   					config.EnvCfg.MySQLReplicaUsername,
	   					config.EnvCfg.MySQLReplicaPassword,
	   					pair[0],
	   					pair[1],
	   					config.EnvCfg.MySQLDataBase,
	   				)
	   				replicas = append(replicas, mysql.Open(replicaDSN))
	   			}

	   			err := Client.Use(dbresolver.Register(dbresolver.Config{
	   				Replicas: replicas,
	   				Policy:   dbresolver.RandomPolicy{},
	   			}))
	   			if err != nil {
	   				panic(err)
	   			}
	   		}

	   		// 连接池配置
	   		sqlDB, err := Client.DB()
	   		if err != nil {
	   			panic(err)
	   		}

	   		sqlDB.SetMaxIdleConns(100)
	   		sqlDB.SetMaxOpenConns(200)
	   		sqlDB.SetConnMaxLifetime(24 * time.Hour)
	   		sqlDB.SetConnMaxIdleTime(time.Hour)

	   		// 分布式追踪
	   		if err := Client.Use(tracing.NewPlugin()); err != nil {
	   			panic(err)
	   			} */

}

// 测试是否成功连接
func testConnection() {
	sqlDB, err := Client.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	log.Println("Successfully connected to the database!")
}
