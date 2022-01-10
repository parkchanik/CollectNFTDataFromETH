package logger

import (
	"fmt"
	"io"
	"log"

	"os"

	"runtime"
	"time"
)

var datalogging *log.Logger
var errorlogging *log.Logger
var infologging *log.Logger

func LoggerInit() {

	var path string

	if runtime.GOOS == "windows" {
		path = "..\\TokenData\\"
		fmt.Println("runtime.GOOS =", runtime.GOOS)
	} else if runtime.GOOS == "linux" {
		path = "../data/TokenData/"
		fmt.Println("runtime.GOOS =", runtime.GOOS)
	}
	t := time.Now()
	//logfile := path + "transactiondata" + t.Format("20060102_15") + ".log"
	filenameprefix := "transactiondata"
	logfile := fmt.Sprintf("%s%s_%s.log", path, filenameprefix, t.Format("20060102_15"))
	datafile, err := os.OpenFile(logfile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}

	datalogging = log.New(io.MultiWriter(datafile, os.Stdout), "", 0)

	errorfilenameprefix := "transactiondata_error"
	errorlogfile := fmt.Sprintf("%s%s_%s.log", path, errorfilenameprefix, t.Format("20060102_15"))
	errorfile, err := os.OpenFile(errorlogfile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}

	errorlogging = log.New(io.MultiWriter(errorfile, os.Stdout), "", 0)

	infofilenameprefix := "transactiondata_info"
	infologfile := fmt.Sprintf("%s%s_%s.log", path, infofilenameprefix, t.Format("20060102_15"))
	infofile, err := os.OpenFile(infologfile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}

	infologging = log.New(io.MultiWriter(infofile, os.Stdout), "", 0)

}

func DataLog(data string) {
	datalogging.Println(data)
}
func ErrorLog(format string, v ...interface{}) {
	errorlogging.Printf(format, v...)
}

func InfoLog(format string, v ...interface{}) {
	infologging.Printf(format, v...)
}
