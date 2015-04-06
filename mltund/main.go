package main

import (
    "log"
    "github.com/xowap/mltun/tunlog"
    "github.com/xowap/mltun/proto"
    "net"
    "encoding/json"
    "io/ioutil"
    "os"
)


type configData struct {
    Users []proto.User
    Bind string
}


var logger *log.Logger
var config configData

func init() {
    logger = tunlog.New("mltund")
    config = configData{
        Bind: ":1842",
    }
}

func handle(conn net.Conn) {
    logger.Printf("Incoming connection from %s\n", conn.RemoteAddr())

    success, err := proto.EnsureMltun(conn)

    if err != nil {
        logger.Printf("Aborting %s because of error: %s\n", conn.RemoteAddr(), err)
        conn.Close()
        return
    }

    if !success {
        logger.Printf("%s doesn't speak the same language, aborting\n", conn.RemoteAddr())
        conn.Close()
        return
    }

    conn.Close()
}

func readConf(file_name string) {
    logger.Printf("Loading conf file \"%s\"\n", file_name)
    data, err := ioutil.ReadFile(file_name)

    if err != nil {
        logger.Fatalf("Could not load configuration: %s\n", err)
    }

    err = json.Unmarshal(data, &config)

    if err != nil {
        logger.Fatalf("Could not decode configuration file: %s\n", err)
    }
}

func main() {
    logger.Println("Starting...")

    if len(os.Args) != 2 {
        logger.Fatalln("Wrong arguments syntax. You must provide a configuration file as sole argument.")
    }

    readConf(os.Args[1])

    listener, err := net.Listen("tcp", config.Bind)

    if err != nil {
        logger.Fatalf("Could not listen to %s. Reason: %s\n", config.Bind, err)
    }

    logger.Println("Started on " + config.Bind)

    for {
        conn, err := listener.Accept()

        if err != nil {
            logger.Printf("Impossible to accept incoming connection. Your system must be under serious stress. " +
                          "Reason: %s\n", err)
        } else {
            go handle(conn)
        }
    }
}