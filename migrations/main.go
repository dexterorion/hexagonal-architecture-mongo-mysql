package main

import (
	"database/sql"
	"embed"
	"fmt"

	"github.com/Boostport/migration"
	"github.com/Boostport/migration/driver/mysql"
	"github.com/dexterorion/hexagonal-architecture-mongo-mysql/helpers"
)

var (
	//go:embed sql/*.sql
	migFs embed.FS
)

func main() {
	db, err := sql.Open("mysql", helpers.BuildMysqlConnUrl())
	if err != nil {
		panic(err)
	}

	driver, err := mysql.NewFromDB(db)
	if err != nil {
		panic(err)
	}

	dirs, err := migFs.ReadDir("sql")
	if err != nil {
		panic(fmt.Sprintf("error reading migrations files from embed: %v", err))
	} else {
		println("List of migrations found:")
		for _, d := range dirs {
			println(fmt.Sprintf(" - %s", d.Name()))
		}
		println("End of list")
	}

	embedSource := &migration.EmbedMigrationSource{
		EmbedFS: migFs,
		Dir:     "sql",
	}

	// Run all up migrations
	applied, err := migration.Migrate(driver, embedSource, migration.Up, 0)
	if err != nil {
		panic(fmt.Sprintf("Error applying migrations: %s", err.Error()))
	} else {
		println(fmt.Sprintf("last applied %d", applied))
	}
}
