package proto

import "encoding/binary"

type User struct {
    Username, Password string
}

type Remote struct {
    Address, LocalAddress string
}

var BYTE_ORDER = binary.LittleEndian
const PROTO_MAGIC_KEY = 42