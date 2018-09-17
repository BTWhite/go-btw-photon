// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3

package logger

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/fatih/color"
)

const (
	logDebug = 0x00
	logInfo  = 0x01
	logErr   = 0x02
)

var (
	logLevel = logDebug
	f, _     = os.OpenFile("debug.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	c        = color.New(color.FgHiBlack).Add(color.BgHiYellow)
	mu       = new(sync.Mutex)
)

// Init initializes a logger with the required type.
func Init(level string) {
	switch level {
	case "debug":
		logLevel = logDebug
	case "info":
		logLevel = logInfo
	case "err":
		logLevel = logErr
	default:
		{
			panic("Invalid logLevel type")
		}

	}
}

// Info responsible for the output of logs with general information.
func Info(text ...interface{}) {
	if logLevel == logErr {
		return
	}
	print("info", text...)
}

// Debug responsible for the output of logs with information for debugging.
func Debug(text ...interface{}) {
	if logLevel != logDebug {
		return
	}
	print("debug", text...)
}

// Err responsible for the output of logs with error/warnign.
func Err(text ...interface{}) {
	print("error", text...)
}

// Fatal responsible for the output of logs with the fatal error (process exit).
func Fatal(text ...interface{}) {
	print("fatal", text...)
	panic(text[0])

}

func print(prefix interface{}, text ...interface{}) {
	tm := time.Now().Format("2006-01-02 15:04:05 ")

	mu.Lock()
	f.WriteString(tm)
	f.WriteString(fmt.Sprintln(text...))

	c.Print(prefix)
	c.Add(color.FgBlack).Print(" ", tm)
	c.Add(color.FgHiBlack)
	fmt.Print(" ")
	fmt.Println(text...)
	mu.Unlock()
}
