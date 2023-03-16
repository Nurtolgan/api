package debugger

import (
	"log"
	"os"
)

func CheckError(errName string, e error) {
	if e != nil {
		f, err := os.OpenFile("logs.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			log.Panicf("[Error: %s]: %s\n", errName, e.Error())
		}
		defer f.Close()
		log.SetOutput(f)
		log.Panicf("[Error: %s]: %s\n", errName, e.Error())
	}
}
