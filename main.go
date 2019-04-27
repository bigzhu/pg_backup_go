package main

import "log"
import "github.com/bigzhu/pg_backup_go/logic"

func main() {
	err := logic.Dump()
	if err != nil {
		log.Fatal(err)
	}
}
