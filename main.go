package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-sql-driver/mysql"
)

func main() {
	conf := &mysql.Config{
		Net:                  "tcp",
		ParseTime:            true,
		AllowNativePasswords: true,
	}
	flag.StringVar(&conf.User, "user", "", "DB user.")
	flag.StringVar(&conf.Passwd, "pass", "", "DB password.")
	flag.StringVar(&conf.Addr, "addr", "0.0.0.0:3306", "DB address.")
	flag.StringVar(&conf.DBName, "dbname", "", "DB name.")

	var seedDir string
	flag.StringVar(&seedDir, "dir", "_seeds", "seeds directory.")

	flag.Parse()

	// initialize DB
	db, err := sql.Open("mysql", conf.FormatDSN())
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	defer db.Close()

	// get Seed directory
	seedsDir, err := filepath.Abs(seedDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	seed, err := NewSeed(db, seedsDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	fmt.Println("Start seeding")
	if err := seed.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	fmt.Println("Success seeding!!")
	os.Exit(0)
}
