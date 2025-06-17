package rr

import (
    "crypto/md5"
    "crypto/sha1"
    "crypto/sha256"
    "errors"
    "fmt"
    "hash/crc32"
    "io"
    "os"
    "path/filepath"
)

var WorkDirectory, _ = os.Getwd()
var (
    ErrNotRegularFile = errors.New("not a regular file")
)

type F string

func NewF(v string) F {
    return F(v)
}
func (r F) String() string {
    return string(r)
}
func (r F) AsS() S {
    return S(r)
}
func (r F) Size() int64 {
    if r == "" {
        return 0
    }
    f, err := os.Stat(r.String())
    if err != nil {
        return 0
    }
    return f.Size()
}
func (r F) IsRegularFileName() bool {
    if r == "" || r == "/" {
        return false
    }
    return true
}
func (r F) IsDirectory() bool {
    if r == "" {
        return false
    }
    stat, err := os.Stat(r.String())
    if err != nil {
        return false
    }
    return stat.IsDir()
}
func (r F) AppendContentsAsByte(v []byte) error {
    if !r.IsRegularFileName() {
        return ErrNotRegularFile
    }
    f, err := os.OpenFile(r.String(), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
    if err != nil {
        return err
    }
    _, err = f.Write(v)
    if err1 := f.Close(); err1 == nil {
        err = err1
    }

    return err
}
func (r F) AppendContents(v string) error {
    return r.AppendContentsAsByte([]byte(v))
}
func (r F) PutContentsAsByte(v []byte) error {
    if !r.IsRegularFileName() {
        return ErrNotRegularFile
    }
    f, err := os.OpenFile(r.String(), os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0644)
    if err != nil {
        return err
    }
    _, err = f.Write(v)
    if err1 := f.Close(); err1 == nil {
        err = err1
    }

    return err
}
func (r F) PutContents(v string) error {
    return r.PutContentsAsByte([]byte(v))
}
func (r F) Exist() bool {
    _, err := os.Stat(r.String())

    if err != nil {
        if os.IsNotExist(err) {
            return false
        }
    }
    return true
}
func (r F) GetContents() S {
    asByte := r.GetContentsAsByte()
    return S(asByte)
}
func (r F) GetContentsAsByte() []byte {
    bytes, err := os.ReadFile(r.String())
    if err != nil {
        return nil
    }
    return bytes
}

func (r F) GetExtension() S {
    return r.AsS().GetExtension()
}
func (r F) GetName() S {
    if r == "" {
        return ""
    }
    base := filepath.Base(r.String())
    return S(base).GetFirst("?").GetFirst("/").GetFirst("#")
}

// Sha1 get file sha1 hash
func (r F) Sha1() (string, error) {
    if !r.IsRegularFileName() || !r.Exist() {
        return "", ErrNotRegularFile
    }
    f, err := os.Open(r.String())
    if err != nil {
        return "", err
    }
    defer f.Close()
    h := sha1.New()
    if _, err = io.Copy(h, f); err != nil {
        return "", err
    }
    return fmt.Sprintf("%x", h.Sum(nil)), nil
}

// Sha256 get file sha256 hash
func (r F) Sha256() (string, error) {
    if !r.IsRegularFileName() || !r.Exist() {
        return "", ErrNotRegularFile
    }
    f, err := os.Open(r.String())
    if err != nil {
        return "", err
    }
    defer f.Close()
    h := sha256.New()
    if _, err = io.Copy(h, f); err != nil {
        return "", err
    }
    return fmt.Sprintf("%x", h.Sum(nil)), nil
}

// Md5 get file md5 hash
func (r F) Md5() (string, error) {
    if !r.IsRegularFileName() || !r.Exist() {
        return "", ErrNotRegularFile
    }
    f, err := os.Open(r.String())
    if err != nil {
        return "", err
    }
    defer f.Close()
    h := md5.New()
    if _, err = io.Copy(h, f); err != nil {
        return "", err
    }
    return fmt.Sprintf("%x", h.Sum(nil)), nil
}

// Crc32 get file crc32 hash
func (r F) Crc32() (string, error) {
    if !r.IsRegularFileName() || !r.Exist() {
        return "", ErrNotRegularFile
    }
    f, err := os.Open(r.String())
    if err != nil {
        return "", err
    }
    defer f.Close()
    table := crc32.MakeTable(crc32.IEEE)
    hash32 := crc32.New(table)
    if _, err := io.Copy(hash32, f); err != nil {
        return "", err
    }
    return fmt.Sprintf("%x", hash32.Sum(nil)), nil
}
func (r F) CopyFile(dst string) error {
    if !r.IsRegularFileName() || !r.Exist() {
        return ErrNotRegularFile
    }
    in, err := os.Open(r.String())
    if err != nil {
        return err
    }
    defer in.Close()
    out, err := os.Create(dst)
    if err != nil {
        return err
    }
    defer out.Close()
    _, err = io.Copy(out, in)
    if err != nil {
        return err
    }
    return nil
}

// MoveFile move file
func (r F) MoveFile(dst string) error {
    err := r.CopyFile(dst)
    if err != nil {
        return err
    }
    return os.Remove(r.String())
}

func (r F) WithWorkDirectory() F {
    return F(filepath.Join(WorkDirectory, r.String()))
}
