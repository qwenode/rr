package rr

import (
    "crypto/md5"
    "crypto/sha1"
    "crypto/sha256"
    "fmt"
    "hash/crc32"
    "io"
    "os"
    "path/filepath"
)

func FileSize(path string) int64 {
    if path == "" {
        return 0
    }
    f, err := os.Stat(path)
    if err != nil {
        return 0
    }
    return f.Size()
}

func FileIsRegularFileName(path string) bool {
    if path == "" || path == "/" {
        return false
    }
    return true
}

func FileIsDirectory(path string) bool {
    if path == "" {
        return false
    }
    stat, err := os.Stat(path)
    if err != nil {
        return false
    }
    return stat.IsDir()
}

func FileAppendContentsAsByte(path string, v []byte) error {
    if !FileIsRegularFileName(path) {
        return ErrNotRegularFile
    }
    f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
    if err != nil {
        return err
    }
    _, err = f.Write(v)
    if err1 := f.Close(); err1 == nil {
        err = err1
    }
    return err
}

func FileAppendContents(path string, v string) error {
    return FileAppendContentsAsByte(path, []byte(v))
}

func FilePutContentsAsByte(path string, v []byte) error {
    if !FileIsRegularFileName(path) {
        return ErrNotRegularFile
    }
    f, err := os.OpenFile(path, os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0644)
    if err != nil {
        return err
    }
    _, err = f.Write(v)
    if err1 := f.Close(); err1 == nil {
        err = err1
    }
    return err
}

func FilePutContents(path string, v string) error {
    return FilePutContentsAsByte(path, []byte(v))
}

func FileExist(path string) bool {
    _, err := os.Stat(path)
    if err != nil {
        if os.IsNotExist(err) {
            return false
        }
    }
    return true
}

func FileGetContents(path string) string {
    asByte := FileGetContentsAsByte(path)
    return string(asByte)
}

func FileGetContentsAsByte(path string) []byte {
    bytes, err := os.ReadFile(path)
    if err != nil {
        return nil
    }
    return bytes
}

func FileGetExtension(path string) string {
    return S(path).GetExtension().String()
}

func FileGetName(path string) string {
    if path == "" {
        return ""
    }
    base := filepath.Base(path)
    return S(base).GetFirst("?").GetFirst("/").GetFirst("#").String()
}

func FileSha1(path string) (string, error) {
    if !FileIsRegularFileName(path) || !FileExist(path) {
        return "", ErrNotRegularFile
    }
    f, err := os.Open(path)
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

func FileSha256(path string) (string, error) {
    if !FileIsRegularFileName(path) || !FileExist(path) {
        return "", ErrNotRegularFile
    }
    f, err := os.Open(path)
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

func FileMd5(path string) (string, error) {
    if !FileIsRegularFileName(path) || !FileExist(path) {
        return "", ErrNotRegularFile
    }
    f, err := os.Open(path)
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

func FileCrc32(path string) (string, error) {
    if !FileIsRegularFileName(path) || !FileExist(path) {
        return "", ErrNotRegularFile
    }
    f, err := os.Open(path)
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

func FileCopy(src string, dst string) error {
    if !FileIsRegularFileName(src) || !FileExist(src) {
        return ErrNotRegularFile
    }
    in, err := os.Open(src)
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

func FileMove(src string, dst string) error {
    err := FileCopy(src, dst)
    if err != nil {
        return err
    }
    return os.Remove(src)
}

func FileWithWorkDirectory(path string) string {
    return filepath.Join(WorkDirectory, path)
}
