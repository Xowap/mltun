package main

import (
    "log"
    "github.com/xowap/mltun/tunlog"
    "net"
)


var logger *log.Logger

func init() {
    logger = tunlog.New("mltund")
}

func handle(conn net.Conn) {
    logger.Println("Incoming connection from " + conn.RemoteAddr().String())

    conn.Write([]byte("coucou\n"))
    conn.Close()
}

func main() {
    logger.Println("Starting")

    bind_addr := ":1842"
    listener, err := net.Listen("tcp", bind_addr)

    if err != nil {
        logger.Fatalln("Could not listen to " + bind_addr)
    }

    for {
        conn, err := listener.Accept()

        if err != nil {
            logger.Println("Impossible to accept incoming connection. Your system must be under serious stress")
            continue
        }

        go handle(conn)
    }
}