package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os/exec"
	"sync"
)

func main() {
	rm, err := RoleMap()
	if err != nil {
		log.Fatal(err)
	}
	for k, v := range rm {
		fmt.Println(k + "--" + v)
	}
}

func run(prog string, args ...string) ([]byte, error) {
	cmd := exec.Command(prog, args...)
	outPipe, err := cmd.StdoutPipe() //listen to stdout
	if err != nil {
		return []byte{}, err
	}

	errPipe, err := cmd.StderrPipe() //listen to stderror
	if err != nil {
		return []byte{}, err
	}
	err = cmd.Start()
	if err != nil {
		return []byte{}, err
	}

	var outBuf bytes.Buffer
	var errBuf bytes.Buffer

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		io.Copy(&outBuf, outPipe)
		wg.Done()
	}()

	io.Copy(&errBuf, errPipe)
	wg.Wait()

	if len(errBuf.Bytes()) != 0 {
		return outBuf.Bytes(), fmt.Errorf("%s", errBuf.Bytes())
	}
	return outBuf.Bytes(), nil
}

type RoleList struct {
	Roles []Role
}

type Role struct {
	RoleName string
	Arn      string
}

func RoleMap() (map[string]string, error) {
	res := make(map[string]string)
	data, err := run("aws", "iam", "list-roles")
	if err != nil {
		return res, err
	}

	var rlist RoleList
	err = json.Unmarshal(data, &rlist)
	if err != nil {
		return res, err
	}
	//fmt.Println("Rlist is", rlist)
	for _, v := range rlist.Roles {
		res[v.RoleName] = v.Arn
	}
	return res, nil

}
