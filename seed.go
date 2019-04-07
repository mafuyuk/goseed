package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/go-sql-driver/mysql"
)

// Seed is  set initial data for structure.
type Seed struct {
	db      *sql.DB
	dirPath string
	files   []os.FileInfo
}

func NewSeed(db *sql.DB, dirPath string) (*Seed, error) {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}
	return &Seed{
		db:      db,
		dirPath: dirPath,
		files:   files,
	}, nil
}

// Execute method can insert initialize data to database
func (s *Seed) Execute() error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	for _, file := range s.files {
		ext := filepath.Ext(file.Name())
		if ext != ".csv" {
			continue
		}

		table := file.Name()[:len(file.Name())-len(ext)]
		csvFilePath := filepath.Join(s.dirPath, file.Name())

		if _, err := loadDataFromCSV(tx, table, csvFilePath); err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func loadDataFromCSV(tx *sql.Tx, table, filePath string) (sql.Result, error) {
	query := `
        LOAD DATA
            LOCAL INFILE '%s'
        INTO TABLE %s
        FIELDS
            TERMINATED BY ','
        LINES
            TERMINATED BY '\n'
            IGNORE 1 LINES
    `
	mysql.RegisterLocalFile(filePath)
	return tx.Exec(fmt.Sprintf(query, filePath, table))
}
