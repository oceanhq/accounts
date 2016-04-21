package uuid

import (
	"crypto/rand"
	"fmt"
	"io"
)

type UUID interface {
	String() string
}

func Generate() UUID {
	uuid := make([]byte, 16)
	io.ReadFull(rand.Reader, uuid)

	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40

	return &uuidImpl{val: uuid}
}

type uuidImpl struct {
	val []byte
}

func (u *uuidImpl) String() string {
	return fmt.Sprintf("%x-%x-%x-%x-%x", u.val[0:4], u.val[4:6], u.val[6:8], u.val[8:10], u.val[10:])
}
