package gologs

import (
	//"io"
	"log"
	"os"
	"time"
)

func Init(args ...string) {

	fileName := ""

	if len(args) > 0 {
		fileName = args[0]
	}

	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		os.Mkdir("logs", 0755)
	}

	go func() {
		for {

			var logFile *os.File

			fullName := "logs/" + fileName + time.Now().Format("2006-01-02 15") +
				"-00-" + "00" + ".log"

			if _, err := os.Stat(fullName); os.IsNotExist(err) {
				logFile, err = os.Create(fullName)
				if err != nil {
					log.Println(err)
					os.Exit(0)
				}
			} else {
				logFile, err = os.OpenFile(
					fullName,
					os.O_CREATE|os.O_WRONLY|os.O_APPEND,
					0666,
				)
				if err != nil {
					log.Fatalln("Failed to open log file:", err)
				}
			}

			//multi := io.MultiWriter(logFile, os.Stdout)

			Debug = log.New(logFile,
				"DEBUG: ",
				log.Ldate|log.Ltime|log.Lshortfile)

			Info = log.New(logFile,
				"INFO: ",
				log.Ldate|log.Ltime|log.Lshortfile)

			Warning = log.New(logFile,
				"WARNING: ",
				log.Ldate|log.Ltime|log.Lshortfile)

			Error = log.New(logFile,
				"ERROR: ",
				log.Ldate|log.Ltime|log.Lshortfile)

			for time.Now().Format("04-05") != "00-00" {
				time.Sleep(time.Second * 1)
			}

			logFile.Close()
		}
	}()
}
