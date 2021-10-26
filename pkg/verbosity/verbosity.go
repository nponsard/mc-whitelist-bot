package verbosity

import (
	"fmt"
	"log"
	"os"
	"path"
	"runtime"

	"github.com/TwinProduction/go-color"
)

var (
	logToFile   log.Logger
	verbose     = false
	saveLog     = false
	logFilePath string
)

// Set verbosity level and log file
func SetupLog(VerbosityActive bool, logPath string) {
	verbose = VerbosityActive
	logFilePath = logPath
}

// Enable log file saving
func SetLogging(active bool) {
	saveLog = active

	// Opens the log file in append mode

	if saveLog {

		logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

		// can’t open file, exit

		if err != nil {
			log.Fatal(err)
		}

		// create logger

		logToFile = *log.New(logFile, "", log.Ldate|log.Ltime)

		logToFile.Println("------- New execution")

		logToFile.Println(os.Args)
	}

}

// Writes to terminal an log file
func doubleLog(chosenColor string, level string, v ...interface{}) {

	_, file, line, _ := runtime.Caller(2)

	// show debug to terminal only when verbose

	if level != "Debug : " || verbose {

		fmt.Print(chosenColor)
		if level == "" {
			fmt.Print(level)
		} else {
			fmt.Printf("%-20s %8s", path.Base(file)+":"+fmt.Sprint(line), level)
		}
		fmt.Print(v...)
		fmt.Print(color.Reset)
		fmt.Println()
	}

	if saveLog {

		if level == "" {
			level = "Info : "
		}

		// Save to file

		a := append([]interface{}{fmt.Sprintf("%-25s %8s", path.Base(file)+":"+fmt.Sprint(line), level)}, v...)
		logToFile.Println(a...)
	}
}

// Debug message, shown when verbose mode is on. Always logged to file
func Debug(v ...interface{}) {
	doubleLog(color.Blue, "Debug : ", v...)
}

// Send message to user
func Info(v ...interface{}) {
	doubleLog("", "", v...)
}

// Error message
func Error(v ...interface{}) {
	doubleLog(color.Red, "Error : ", v...)
}

// Show error then exits
func Fatal(v ...interface{}) {
	doubleLog(color.Red, "Fatal : ", v...)
	os.Exit(1)
}

// Show a warning
func Warning(v ...interface{}) {
	doubleLog(color.Yellow, "Warn : ", v...)
}
