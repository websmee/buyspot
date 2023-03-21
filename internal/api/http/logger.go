package http

import (
	"log"
	"os"
)

var logger = log.New(os.Stdout, "[HTTP HANDLER] ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
