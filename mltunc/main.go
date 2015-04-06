package main

import (
    "github.com/xowap/mltun/proto"
    "github.com/xowap/mltun/tunlog"
    "log"
    "os"
    "io/ioutil"
    "encoding/json"
    "net"
)


type configData struct {
    User proto.User
    Remotes []proto.Remote
}

var logger *log.Logger
var config configData


func init() {
    logger = tunlog.New("mltunc")
}

func readConf(file_name string) {
    logger.Printf("Loading conf file \"%s\"\n", file_name)
    data, err := ioutil.ReadFile(file_name)

    if err != nil {
        logger.Fatalf("Could not load configuration file: %s\n", err)
    }

    err = json.Unmarshal(data, &config)

    if err != nil {
        logger.Fatalf("Could not decode configuration file: %s\n", err)
    }

    if config.User.Username == "" || config.User.Password == "" {
        logger.Fatalf("No username/password defined, can't do anything without")
    }
}

func run(remote proto.Remote) {
    logger.Printf("Connecting remote \"%s\" with address \"%s\"\n", remote.Address, remote.LocalAddress)

    remote_addr, err := net.ResolveTCPAddr("tcp", remote.Address)

    if err != nil {
        logger.Printf("Could not parse given remote address: %s\n", err)
        return
    }

    var local_addr *net.TCPAddr

    if remote.LocalAddress != "" {
        local_addr, err = net.ResolveTCPAddr("tcp", remote.LocalAddress)

        if err != nil {
            logger.Printf("Could not parse given local address: %s\n", err)
            return
        }
    } else {
        local_addr = nil
    }

    conn, err := net.DialTCP("tcp", local_addr, remote_addr)

    if err != nil {
        logger.Printf("Failed to connect \"%s\" with address \"%s\"\n", remote.Address, remote.LocalAddress)
        return
    }

    proto.EnsureMltun(conn)
}

func main() {
    logger.Println("Starting...")
    done := make(chan bool)

    if len(os.Args) != 2 {
        logger.Fatalln("Wrong arguments syntax. You must provide a configuration file a sole argument")
    }

    readConf(os.Args[1])

    for i := 0; i < len(config.Remotes); i += 1 {
        go run(config.Remotes[i])
    }

    <- done
}