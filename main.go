package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/mafuyuk/goseed/db"
)

func main() {
	var dbConf = &db.Config{}
	flag.StringVar(&dbConf.Host, "h", "0.0.0.0", "DB host.")
	flag.StringVar(&dbConf.Port, "P", ":3306", "DB port.")
	flag.StringVar(&dbConf.User, "u", "user", "DB user.")
	flag.StringVar(&dbConf.Password, "p", "password", "DB password.")
	flag.StringVar(&dbConf.DBName, "n", "dbname", "DB name.")

	var seedDir string
	flag.StringVar(&seedDir, "d", "_seeds", "seeds directory.")

	flag.Parse()

	// initialize DB
	db, err := db.NewMySQL(dbConf)
	if err != nil {
		os.Exit(1)
	}
	defer db.Close()

	// Seedディレクトリの取得
	seedsDir, err := filepath.Abs(seedDir)
	if err != nil {
		os.Exit(2)
	}

	seed, err := NewSeed(db, seedsDir)
	if err != nil {
		os.Exit(3)
	}

	fmt.Println("Start seeding")
	if err := seed.Execute(); err != nil {
		fmt.Printf("Failed seeding. err: %#v\n", err)
		os.Exit(4)
	}
	fmt.Println("Success seeding!!")
}
