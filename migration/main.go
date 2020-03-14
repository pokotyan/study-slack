// シンプルな例: https://github.com/golang-migrate/migrate/tree/master/database/mysql
// パクってきた例: https://github.com/k-yomo/go_echo_api_boilerplate/blob/57ea72eb4e/config/db.go
package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

var migrationFilePath = "file://./migration/migrations/"

func main() {
	fmt.Println("start migration")
	// flag.Parse()
	// if flag.Arg(0) == "" {
	// 	showUsage()
	// 	os.Exit(1)
	// }

	m := newMigrate()
	version, dirty, _ := m.Version()
	force := flag.Bool("f", false, "force execute fixed sql")
	if dirty && *force {
		fmt.Println("force=true: force execute current version sql")
		m.Force(int(version))
	}

	switch flag.Arg(0) {
	case "new":
		newMigration(flag.Arg(1))
	case "up":
		up(m)
	case "down":
		down(m)
	case "drop":
		drop(m)
	case "version":
		showVersionInfo(m.Version())
	default:
		fmt.Println("\nerror: invalid command '", flag.Arg(0), "'")
		showUsage()
		os.Exit(0)
	}
}

func newMigrate() *migrate.Migrate {
	err := godotenv.Load()
	if err != nil {
		fmt.Errorf("err %s", "load error .env")
		os.Exit(1)
	}

	// user := os.Getenv("DB_USERNAME")
	// pass := os.Getenv("DB_PASSWORD")
	// host := os.Getenv("DB_HOST")
	// port := os.Getenv("DB_PORT")
	// dbName := os.Getenv("DB_DATABASE")

	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, pass, host, port, dbName)

	dbURL := os.Getenv("CLEARDB_DATABASE_URL") // heroku対応
	dsn := dbURL + "?parseTime=true"

	fmt.Println(dsn)

	db, _ := sql.Open("mysql", dsn)
	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		migrationFilePath,
		"mysql",
		driver,
	)

	if err != nil {
		fmt.Println(errors.Wrap(err, "initialize Migrate instance failed"))
		os.Exit(1)
	}
	return m
}

func showUsage() {
	fmt.Println(`
-------------------------------------
Usage:
  go run migrate.go <command>
Commands:
  new NAME	Create new up & down migration files
  up		Apply up migrations
  down		Apply down migrations
  drop		Drop everything
  version	Check current migrate version
-------------------------------------`)
}

func newMigration(name string) {
	if name == "" {
		fmt.Println("\nerror: migration file name must be supplied as an argument")
		os.Exit(1)
	}
	base := fmt.Sprintf("./migration/migrations/%s_%s", time.Now().Format("20060102030405"), name)
	ext := ".sql"
	createFile(base + ".up" + ext)
	createFile(base + ".down" + ext)
}

func createFile(fname string) {
	if _, err := os.Create(fname); err != nil {
		panic(err)
	}
}

func up(m *migrate.Migrate) {
	fmt.Println("Before:")
	showVersionInfo(m.Version())
	err := m.Up()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("\nUpdated:")
		version, dirty, err := m.Version()
		showVersionInfo(version, dirty, err)
	}
}

func down(m *migrate.Migrate) {
	fmt.Println("Before:")
	showVersionInfo(m.Version())
	err := m.Steps(-1)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("\nUpdated:")
		showVersionInfo(m.Version())
	}
}

func drop(m *migrate.Migrate) {
	err := m.Drop()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Dropped all migrations")
		return
	}
}

func showVersionInfo(version uint, dirty bool, err error) {
	fmt.Println("-------------------")
	fmt.Println("version : ", version)
	fmt.Println("dirty   : ", dirty)
	fmt.Println("error   : ", err)
	fmt.Println("-------------------")
}
