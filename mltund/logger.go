package main

import (
    "github.com/xowap/mltun/tunlog"
    "log"
)

var logger *log.Logger

func init() {
    logger = tunlog.New("mltund")
}