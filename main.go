package main

import (
	"bytes"
	"fmt"
	"log"
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
	password := fmt.Sprintf("PGPASSWORD=%s", dbConf.Password)
	host := fmt.Sprintf("--host=%s", dbConf.Host)
	port := fmt.Sprintf("--port=%s", dbConf.Port)
	user := fmt.Sprintf(`--username=%s`, dbConf.User)
	params := "--format=c"
	file := fmt.Sprintf("--file=%s", "test.dump")

	cmd := exec.Command(tool, host, port, user, params, "-v", file, dbConf.DBName)
	//cmd := exec.Command("pg_dump -h 123.176.102.187 -p 5432 -U learn_ngsl -Fc -f learn_ngsl.dump learn_ngsl")
	cmd.Env = append(os.Environ(),
		password,
	)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Starting dump DB from remote server...")
	err = cmd.Wait()
	if err != nil {
		log.Fatal(err)
	}
	outStr := string(stdout.Bytes())
	fmt.Println(outStr)
	log.Printf("DB dum is Done!")
}
