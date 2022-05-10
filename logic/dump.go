package logic

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/bigzhu/gobz/confbz"
)

const tool string = "pg_dump"

func checkExists() error {
	path, err := exec.LookPath(tool)
	if err != nil {
		log.Printf("didn't find 'pg_dump' executable, you need install postgresql first\n")
	} else {
		log.Printf("'pg_dump' executable is in '%s'\n", path)
	}
	return err
}
func prepareCMD() (cmd *exec.Cmd, err error) {
	dbConf := confbz.GetDBConf()
	if dbConf.Port == "" {
		dbConf.Port = "5432"
	}
	password := fmt.Sprintf("PGPASSWORD=%s", dbConf.Password)
	host := fmt.Sprintf("--host=%s", dbConf.Host)
	port := fmt.Sprintf("--port=%s", dbConf.Port)
	user := fmt.Sprintf(`--username=%s`, dbConf.User)
	params := "--format=c"
	fileName, err := getFileName(dbConf.DBName)
	if err != nil {
		return
	}

	file := fmt.Sprintf("--file=%s", fileName)

	cmd = exec.Command(tool, host, port, user, params, "-v", file, dbConf.DBName)
	//cmd := exec.Command("pg_dump -h 123.176.102.187 -p 5432 -U learn_ngsl -Fc -f learn_ngsl.dump learn_ngsl")
	cmd.Env = append(os.Environ(),
		password,
	)

	err = cmd.Start()

	return
}

// Dump main dump
func Dump() (err error) {
	err = checkExists()
	if err != nil {
		return
	}
	cmd, err := prepareCMD()
	if err != nil {
		return
	}
	log.Printf("Starting dump DB from remote server...")
	err = cmd.Wait()
	if err != nil {
		return
	}
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	outStr := string(stdout.Bytes())
	fmt.Println(outStr)
	log.Printf("DB dum is Done!")
	return
}
