package proto

import (
    "errors"
    "net"
    "math/rand"
    "bytes"
    "encoding/binary"
)

func EnsureMltun(conn net.Conn) (bool, error) {
    fail := func (err error, reason string) (bool, error) {
        logger.Println(err)
        return false, errors.New(reason)
    }

    num := rand.Uint32()
    sign := num ^ uint32(PROTO_MAGIC_KEY)
    buf := new(bytes.Buffer)

    err := binary.Write(buf, BYTE_ORDER, num)

    if err != nil {
        return fail(err, "memory allocation failed")
    }

    err = binary.Write(buf, BYTE_ORDER, sign)

    if err != nil {
        return fail(err, "memory allocation failed")
    }

    _, err = conn.Write(buf.Bytes())

    if err != nil {
        return fail(err, "network error")
    }

    var r_num, r_sign uint32
    err = binary.Read(conn, BYTE_ORDER, &r_num);
    if err != nil {
        return fail(err, "network error")
    }

    err = binary.Read(conn, BYTE_ORDER, &r_sign);
    if err != nil {
        return fail(err, "network error")
    }

    if (r_num ^ r_sign) == PROTO_MAGIC_KEY {
        logger.Println("Protocol match")
        return true, nil
    } else {
        logger.Println("Protocol mismatch")
        return false, nil
    }
}