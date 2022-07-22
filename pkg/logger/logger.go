package logger

import (
	"io"
	"log"
)

type Loggers struct {
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
}

var (
	Warning *log.Logger
	Info    *log.Logger
	Error   *log.Logger
)

func InitLogger(handle io.Writer) Loggers {
	return Loggers{
		log.New(handle, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		log.New(handle, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile),
		log.New(handle, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}
