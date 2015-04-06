package tunlog

import (
    "log"
    "os"
)

func New(name string) *log.Logger {
    return log.New(os.Stderr, name + ": ", log.Ldate | log.Ltime | log.Lshortfile)
}