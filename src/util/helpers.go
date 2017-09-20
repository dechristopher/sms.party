package util

import (
	"fmt"
	"log"
	"os"
	str "strings"
	"time"

	s "github.com/dechristopher/sms.party/src/strings"
)

// Log logs information to console
func Log(message string) {
	fmt.Println(s.LogPrefix + message)
}

// LogErr logs information using log.Fatal
func LogErr(message string) {
	log.Fatal(s.LogPrefix + message)
}

// LogToFile logs info to a flat file
func LogToFile(m string) {
	now := getCurrentTime()
	if f, ferr := os.OpenFile("./log", os.O_APPEND|os.O_WRONLY, os.ModeAppend); ferr != nil {
		if str.Contains(ferr.Error(), "no such") {
			os.Create("log")
			LogToFile(m)
		} else {
			fmt.Println("[logging] Failed to open log file")
		}
	} else {
		defer f.Close()

		if _, werr := f.WriteString("[" + now + "] " + m + "\n"); werr != nil {
			fmt.Println("[logging] Failed to log: " + m)
			fmt.Println("[logging] ERROR: " + werr.Error())
			return
		}
		fmt.Printf("[logging] " + m + "\n")
		f.Sync()
		return
	}
}

// LogToRedis logs SMS messages sent into redis
func LogToRedis(ip string, uri string, number string, message string) {
	// Get the current time for log
	now := getCurrentTime()
	// Log the events
	if logged, err := R.LPush("sms-log", "["+now+"] "+ip+" "+uri+" ["+number+" - "+message+"]").Result(); err != nil {
		LogErr("[" + now + "] Failed to connect to Redis!")
	} else {
		if logged == 1 {
			return
		}

		Log("[" + now + "] Failed to log to Redis!")
	}
}

// getCurrentTime returns current time in RFC1123 format as string
func getCurrentTime() string {
	return time.Now().Format(time.RFC1123)
}
