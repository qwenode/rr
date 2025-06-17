package rr

import (
    "crypto/md5"
    "crypto/sha1"
    "crypto/sha256"
    "crypto/sha512"
    "fmt"
    "hash/crc32"
)

type B []byte

func NewB(v []byte) B {
    return B(v)
}
func (r B) AsS() S {
    return S(r)
}
func (r B) String() string {
    return string(r)
}
func (r B) Sha1() [20]byte {
    return sha1.Sum(r)
}
func (r B) Sha256() [32]byte {
    return sha256.Sum256(r)
}

func (r B) Sha512() [64]byte {
    return sha512.Sum512(r)
}
func (r B) Crc32() string {
    table := crc32.MakeTable(crc32.IEEE)
    byteString := crc32.Checksum(r, table)
    return fmt.Sprintf("%x", byteString)
}

func (r B) Md5() [16]byte {
    return md5.Sum(r)
}

func (r B) Sha1String() string {
    return fmt.Sprintf("%x", r.Sha1())
}

func (r B) Md5String() string {
    return fmt.Sprintf("%x", r.Md5())
}
