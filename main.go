package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	"github.com/bigzhu/gobz/confbz"
)

const tool string = "pg_dump"

func checkLsExists() error {
	path, err := exec.LookPath(tool)
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
	password := fmt.Sprintf(`export PGPASSWORD="%s";`, dbConf.Password)
	host := fmt.Sprintf("--host=%s", dbConf.Host)
	port := fmt.Sprintf("--port=%s", dbConf.Port)
	user := fmt.Sprintf("--username=%s", dbConf.User)
	params := "--format=c"
	file := fmt.Sprintf("--file=%s", "test.dump")

	dbName := fmt.Sprintf("-U %s", dbConf.DBName)
	cmd := exec.Command(tool, host, port, user, params, file, dbName)
	cmd.Env = append(os.Environ(),
		password,
	)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err = cmd.Run()

	if err != nil {
		// log.Fatal(err)
	}
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	fmt.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)
}
