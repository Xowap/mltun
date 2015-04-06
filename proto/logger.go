package proto

import (
    "log"
    "github.com/xowap/mltun/tunlog"
)

var logger *log.Logger

func init() {
    logger = tunlog.New("proto")
}