package main

import (
	"encoding/base64"
	"io"
	"log"
	"os"
)

var ENCODING = base64.StdEncoding

func LockFile(f os.FileInfo) {
	iF, oF, ok := OpenFiles(f.Name(), f.Name()+".locked")
	if !ok {
		return
	}
	defer oF.Close()
	defer iF.Close()
	encoder := base64.NewEncoder(ENCODING, oF)
	defer encoder.Close()
	_, err := io.Copy(encoder, iF)
	if err != nil {
		log.Println("Encode write error: ", err)
	} else {
		log.Println("Encoding successful!")
	}
}
func UnLockFile(f os.FileInfo) {
	iF, oF, ok := OpenFiles(f.Name(), f.Name()+".unlocked")
	if !ok {
		return
	}
	defer oF.Close()
	defer iF.Close()
	decoder := base64.NewDecoder(ENCODING, iF)
	_, err := io.Copy(oF, decoder)
	if err != nil {
		log.Println("Decode copy error: ", err)
	} else {
		log.Println("Decoding successful!")
	}
}

func OpenFiles(inName, outName string) (iF, oF *os.File, ok bool) {
	_, err := os.Stat(outName)
	if err == nil || !os.IsNotExist(err) {
		log.Println("Cannot create file: ", outName, " already exists!")
		return nil, nil, false
	}
	oF, err = os.Create(outName)
	if err != nil {
		log.Println("Cannot create file ", outName, ": ", err)
		return nil, nil, false
	}
	iF, err = os.Open(inName)
	if err != nil {
		oF.Close()
		log.Println("Cannot read file ", inName, ": ", err)
		return nil, nil, false
	}
	return iF, oF, true
}
