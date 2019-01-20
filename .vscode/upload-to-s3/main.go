package main

import (
	"os"
	"os/exec"
	"path"
)

type LProject struct {
	Name   string
	Bucket string
	Role   string
	path   string
}

func (lp LProject) UploadLambda(name string) error {
	fpath := path.Join(lp.path, name)
	os.Setenv("GOOS", "linux")
	os.Setenv("GOARCH", "amd64")

	_, err := run("go", "build", "-o", fpath, fpath+".go")
	if err != nil {
		return err
	}

	_, err = run("zip", "-j", fpath+".zip", fpath)
	if err != nil {
		return err
	}

	lamname := lp.Name + "_" + name
	upcmd := exec.Command("aws", "s3", "cp", fpath+".zip", "s3://"+lp.Bucket+"/"+lamname+".zip")

	upOut, err := upcmd.StdoutPipe()
}
