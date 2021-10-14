package mysql

import (
    "fmt"
    "github.com/itp-backend/backend-a-co-create/model/domain"
    log "github.com/sirupsen/logrus"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

type Client interface {
	Ping() error
}

type client struct {
	dbConnection *gorm.DB
}

func (c *client) Ping() error {
	var result int64
	tx := c.dbConnection.Raw("select 1").Scan(&result)
	if tx.Error != nil {
		return fmt.Errorf("mysql unable to serve basic query. %v", tx.Error)
	}
	return nil
}

func (c *client) OpenDB() *gorm.DB {
    return c.dbConnection
}

func NewMysqlClient(config ClientConfig) *client {
	//connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=UTC",
	//	config.Username,
	//	config.Password,
	//	config.Host,
	//	config.Port,
	//	config.DBName,
	//)
    connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=require TimeZone=Asia/Jakarta",
        config.Host,
        config.Username,
        config.Password,
        config.DBName,
        config.Port,
    )
    //host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai
	dbConn, err := gorm.Open(postgres.Open(connStr), &gorm.Config{
		SkipDefaultTransaction:                   true,
		PrepareStmt:                              true,
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		log.Fatalf("unable to initiate mysql connection. %v", err)
	}

    err = dbConn.AutoMigrate(&domain.User{}, &domain.Enrollment{}, &domain.Article{}, domain.Project{})
    if err != nil {
        log.Fatalf("unable to migrate db. %v", err)
    }

	return &client{
		dbConnection: dbConn,
	}
}
