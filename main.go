package main

import (
	"fmt"
	"os/exec"

	"github.com/bigzhu/gobz/confbz"
)

func checkLsExists() error {
	path, err := exec.LookPath("pg_dump")
	if err != nil {
		fmt.Printf("didn't find 'pg_dump' executable, you need install postgresql first\n")
	} else {
		fmt.Printf("'pg_dump' executable is in '%s'\n", path)
	}
	return err
}
func main() {
	// fmt.Println("vim-go")
	err := checkLsExists()
	if err != nil {
		return
	}
	dbConf := confbz.GetDBConf()
	fmt.Println(dbConf)
}
