package logic

import (
	"fmt"
	"log"
	"os"
	"time"
)

func getFileName(DBName string) (fileName string, err error) {
	path := "pg_backup"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Println("make dir pg_backup")
		os.Mkdir(path, os.ModePerm)
	}
	if err != nil {
		return
	}
	t := time.Now()
	now := t.Format("20060102150405")
	fileName = fmt.Sprintf("%s/%s_%s", path, DBName, now)
	return
}
