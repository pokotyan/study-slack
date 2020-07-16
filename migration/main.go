// シンプルな例:
// https://github.com/golang-migrate/migrate/tree/master/database/mysql
// パクってきた例:
// https://github.com/k-yomo/go_echo_api_boilerplate/blob/57ea72eb4e/config/db.go
// https://github.com/k-yomo/go_echo_api_boilerplate/blob/57ea72eb4ef8ca490b1b801aac9b279a391b3597/cmd/db/migrate.go
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
	flag.Parse()
	command := flag.Arg(0)
	migrationFileName := flag.Arg(1)

	if command == "" {
		showUsage()
		os.Exit(1)
	}

	err := godotenv.Load()
	if err != nil {
		fmt.Println(errors.Wrap(err, "load error .env"))
	}

	m := newMigrate()
	version, dirty, _ := m.Version()
	force := flag.Bool("f", false, "force execute fixed sql")
	if dirty && *force {
		fmt.Println("force=true: force execute current version sql")
		m.Force(int(version))
	}

	switch command {
	case "new":
		newMigration(migrationFileName)
	case "up":
		up(m)
	case "down":
		down(m)
	case "drop":
		drop(m)
	case "version":
		showVersionInfo(m.Version())
	default:
		fmt.Println("\nerror: invalid command '", command, "'")
		showUsage()
		os.Exit(0)
	}
}

func generateDsn() string {
	apiRevision := os.Getenv("API_REVISION")
	var dsn string

	if apiRevision == "release" {
		dsn = os.Getenv("DATABASE_URL") + "?multiStatements=true" // heroku対応
	} else {
		user := os.Getenv("DB_USERNAME")
		pass := os.Getenv("DB_PASSWORD")
		host := os.Getenv("DB_HOST")
		port := os.Getenv("DB_PORT")
		dbName := os.Getenv("DB_DATABASE")

		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&multiStatements=true", user, pass, host, port, dbName)
	}

	return dsn
}

func newMigrate() *migrate.Migrate {
	dsn := generateDsn()

	db, openErr := sql.Open("mysql", dsn)
	if openErr != nil {
		fmt.Println(errors.Wrap(openErr, "error occurred. sql.Open()"))
		os.Exit(1)
	}

	driver, instanceErr := mysql.WithInstance(db, &mysql.Config{})
	if instanceErr != nil {
		fmt.Println(errors.Wrap(instanceErr, "error occurred. mysql.WithInstance()"))
		os.Exit(1)
	}

	m, err := migrate.NewWithDatabaseInstance(
		migrationFilePath,
		"mysql",
		driver,
	)

	if err != nil {
		fmt.Println(errors.Wrap(err, "error occurred. migrate.NewWithDatabaseInstance()"))
		os.Exit(1)
	}
	return m
}

func showUsage() {
	fmt.Println(`
-------------------------------------
Usage:
  go run migration/main.go <command>
Commands:
  new FILENAME	Create new up & down migration files
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
		if err.Error() != "no change" {
			panic(err)
		}
		fmt.Println("\nno change")
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
