package mysql

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DbConf struct {
	DbMs   string
	User   string
	Pass   string
	DbName string
	Host   string
	Port   string
	DbURL  string
}

func Connect() *gorm.DB {
	dc, err := SetConfig()
	if err != nil {
		panic(err.Error())
	}
	DBMS := dc.DbMs
	// USER := dc.User
	// PASS := dc.Pass
	// PROTOCOL := fmt.Sprintf("tcp(%s:%s)", dc.Host, dc.Port)
	// DBNAME := dc.DbName
	// CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	CONNECT := dc.DbURL + "?parseTime=true" // heroku対応

	fmt.Println(CONNECT)
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func SetConfig() (dc DbConf, err error) {
	dc.DbMs = "mysql"
	dc.DbName = os.Getenv("DB_DATABASE")
	dc.User = os.Getenv("DB_USERNAME")
	dc.Pass = os.Getenv("DB_PASSWORD")
	dc.Host = os.Getenv("DB_HOST")
	dc.Port = os.Getenv("DB_PORT")
	dc.DbURL = os.Getenv("CLEARDB_DATABASE_URL")
	return dc, nil
}

func TransactAndReturnData(db *gorm.DB, txFunc func(*gorm.DB) (interface{}, error)) (data interface{}, err error) {
	tx := db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit().Error
		}
	}()
	data, err = txFunc(tx)
	return
}
