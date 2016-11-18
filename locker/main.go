package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"time"
)

var UNLOCK_DATE = time.Date(2017, time.May, 12, 0, 0, 0, 0, time.Now().Location())

func main() {
	if len(os.Args) != 3 || (os.Args[1] != "lock" && os.Args[1] != "unlock") {
		PrintUsage()
		return
	}
	f, err := os.Stat(os.Args[2])
	if err != nil {
		log.Println("File access error: ", err)
		return
	}
	if f.IsDir() {
		PrintUsage()
		return
	}
	if os.Args[1] == "lock" {
		LockFile(f)
		return
	}
	if TimeToUnlock() > 0 {
		PrintUsage()
		return
	}
	UnLockFile(f)
}

func PrintUsage() {
	fmt.Println("Usage: locker [lock|unlock] FILENAME")
	t := TimeToUnlock()
	if t > 0 {
		if t > 24*time.Hour {
			fmt.Printf("\tTime left till unlock available: %d days\n", int(math.Floor(t.Hours()))/24)
		} else {
			fmt.Printf("\tTime left till unlock available: %s\n", t)
		}
	} else {
		fmt.Println("\tUnlock available!")
	}
}

func TimeToUnlock() time.Duration {
	return UNLOCK_DATE.Sub(time.Now())
}
