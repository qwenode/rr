package rr

import (
    "bytes"
    "crypto/md5"
    "crypto/sha1"
    "crypto/sha256"
    "fmt"
    "hash/crc32"
    "os"
    "path/filepath"
    "strings"
    "testing"
)

func TestFileSize(t *testing.T) {
    // 测试空路径
    if size := FileSize(""); size != 0 {
        t.Errorf("空路径应返回0，但得到 %d", size)
    }

    // 测试不存在的路径
    if size := FileSize("/not/exist/path"); size != 0 {
        t.Errorf("不存在路径应返回0，但得到 %d", size)
    }

    // 测试已有文件
    tmpDir := os.TempDir()
    testFile := filepath.Join(tmpDir, "test_file_size.txt")
    content := []byte("test content")
    if err := os.WriteFile(testFile, content, 0644); err != nil {
        t.Fatalf("创建测试文件失败: %v", err)
    }
    defer os.Remove(testFile)

    if size := FileSize(testFile); size != int64(len(content)) {
        t.Errorf("期望文件大小 %d，但得到 %d", len(content), size)
    }
}

func TestFileIsRegularFileName(t *testing.T) {
    tests := []struct {
        path     string
        expected bool
    }{
        {"", false},
        {"/", false},
        {"a.txt", true},
        {"path/to/file.txt", true},
    }

    for _, test := range tests {
        if got := FileIsRegularFileName(test.path); got != test.expected {
            t.Errorf("FileIsRegularFileName(%q) = %v，期望 %v", test.path, got, test.expected)
        }
    }
}

func TestFileIsDirectory(t *testing.T) {
    // 创建临时目录
    tmpDir, err := os.MkdirTemp("", "test_dir_*")
    if err != nil {
        t.Fatalf("创建临时目录失败: %v", err)
    }
    defer os.RemoveAll(tmpDir)

    // 在临时目录中创建测试文件
    testFile := filepath.Join(tmpDir, "test.txt")
    if err := os.WriteFile(testFile, []byte("test"), 0644); err != nil {
        t.Fatalf("创建测试文件失败: %v", err)
    }

    // 创建符号链接（如果系统支持）
    symLink := filepath.Join(tmpDir, "symlink")
    _ = os.Symlink(testFile, symLink) // 忽略错误，因为某些系统可能不支持符号链接

    tests := []struct {
        name string
        path string
        want bool
    }{
        {
            "空路径",
            "",
            false,
        },
        {
            "不存在的路径",
            "/path/not/exists",
            false,
        },
        {
            "普通文件",
            testFile,
            false,
        },
        {
            "目录",
            tmpDir,
            true,
        },
        {
            "符号链接",
            symLink,
            false,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := FileIsDirectory(tt.path); got != tt.want {
                t.Errorf("FileIsDirectory(%q) = %v, want %v", tt.path, got, tt.want)
            }
        })
    }
}

func TestFileAppendContentsAsByte(t *testing.T) {
    // 创建临时目录
    tmpDir, err := os.MkdirTemp("", "test_append_*")
    if err != nil {
        t.Fatalf("创建临时目录失败: %v", err)
    }
    defer os.RemoveAll(tmpDir)

    // 测试用例
    tests := []struct {
        name     string
        path     string
        content  []byte
        setup    func(string) error
        wantErr  bool
        checkErr func(error) bool
    }{
        {
            name:    "空路径",
            path:    "",
            content: []byte("test"),
            wantErr: true,
            checkErr: func(err error) bool {
                return err == ErrNotRegularFile
            },
        },
        {
            name:    "正常追加到空文件",
            path:    filepath.Join(tmpDir, "normal.txt"),
            content: []byte("first line\n"),
            wantErr: false,
        },
        {
            name:    "正常追加到已有内容的文件",
            path:    filepath.Join(tmpDir, "normal.txt"),
            content: []byte("second line\n"),
            setup: func(path string) error {
                return os.WriteFile(path, []byte("existing content\n"), 0644)
            },
            wantErr: false,
        },
        {
            name:    "追加到只读文件",
            path:    filepath.Join(tmpDir, "readonly.txt"),
            content: []byte("test"),
            setup: func(path string) error {
                if err := os.WriteFile(path, []byte(""), 0644); err != nil {
                    return err
                }
                return os.Chmod(path, 0444)
            },
            wantErr: true,
        },
        {
            name:    "追加到目录",
            path:    tmpDir,
            content: []byte("test"),
            wantErr: true,
        },
        {
            name:    "大数据量追加",
            path:    filepath.Join(tmpDir, "large.txt"),
            content: bytes.Repeat([]byte("a"), 1024*1024), // 1MB 数据
            wantErr: false,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // 如果有设置函数，执行设置
            if tt.setup != nil {
                if err := tt.setup(tt.path); err != nil {
                    t.Fatalf("设置测试环境失败: %v", err)
                }
            }

            // 执行追加操作
            err := FileAppendContentsAsByte(tt.path, tt.content)

            // 检查错误
            if (err != nil) != tt.wantErr {
                t.Errorf("FileAppendContentsAsByte() error = %v, wantErr %v", err, tt.wantErr)
                return
            }

            // 如果有特定的错误检查函数
            if tt.checkErr != nil && err != nil {
                if !tt.checkErr(err) {
                    t.Errorf("FileAppendContentsAsByte() error = %v, 不符合预期的错误类型", err)
                }
            }

            // 如果期望成功，验证文件内容
            if !tt.wantErr {
                got, err := os.ReadFile(tt.path)
                if err != nil {
                    t.Fatalf("读取文件失败: %v", err)
                }

                // 如果是第二次追加，检查是否包含之前的内容
                if tt.setup != nil {
                    if !bytes.Contains(got, []byte("existing content\n")) {
                        t.Error("文件应包含原有内容")
                    }
                }

                // 检查是否包含新追加的内容
                if !bytes.Contains(got, tt.content) {
                    t.Error("文件应包含新追加的内容")
                }
            }
        })
    }
}

func TestFileAppendContents(t *testing.T) {
    // 测试空路径
    if err := FileAppendContents("", "test"); err != ErrNotRegularFile {
        t.Errorf("空路径应返回ErrNotRegularFile，但得到 %v", err)
    }

    // 测试正常追加
    tmpDir := os.TempDir()
    testFile := filepath.Join(tmpDir, "test_append_str.txt")
    defer os.Remove(testFile)

    // 第一次追加
    content1 := "first content\n"
    if err := FileAppendContents(testFile, content1); err != nil {
        t.Fatalf("第一次追加失败: %v", err)
    }

    // 第二次追加
    content2 := "second content\n"
    if err := FileAppendContents(testFile, content2); err != nil {
        t.Fatalf("第二次追加失败: %v", err)
    }

    // 验证内容
    got, err := os.ReadFile(testFile)
    if err != nil {
        t.Fatalf("读取文件失败: %v", err)
    }

    expected := content1 + content2
    if string(got) != expected {
        t.Errorf("文件内容不匹配\n期望: %q\n实际: %q", expected, got)
    }
}

func TestFilePutContentsAsByte(t *testing.T) {
    // 测试空路径
    if err := FilePutContentsAsByte("", []byte("test")); err != ErrNotRegularFile {
        t.Errorf("空路径应返回ErrNotRegularFile，但得到 %v", err)
    }

    // 测试正常写入
    tmpDir := os.TempDir()
    testFile := filepath.Join(tmpDir, "test_put.txt")
    defer os.Remove(testFile)

    // 第一次写入
    content1 := []byte("first content\n")
    if err := FilePutContentsAsByte(testFile, content1); err != nil {
        t.Fatalf("第一次写入失败: %v", err)
    }

    // 第二次写入（应覆盖）
    content2 := []byte("second content\n")
    if err := FilePutContentsAsByte(testFile, content2); err != nil {
        t.Fatalf("第二次写入失败: %v", err)
    }

    // 验证内容（应只有第二次写入的内容）
    got, err := os.ReadFile(testFile)
    if err != nil {
        t.Fatalf("读取文件失败: %v", err)
    }

    if string(got) != string(content2) {
        t.Errorf("文件内容不匹配\n期望: %q\n实际: %q", content2, got)
    }
}

func TestFilePutContents(t *testing.T) {
    // 测试空路径
    if err := FilePutContents("", "test"); err != ErrNotRegularFile {
        t.Errorf("空路径应返回ErrNotRegularFile，但得到 %v", err)
    }

    // 测试正常写入
    tmpDir := os.TempDir()
    testFile := filepath.Join(tmpDir, "test_put_str.txt")
    defer os.Remove(testFile)

    // 第一次写入
    content1 := "first content\n"
    if err := FilePutContents(testFile, content1); err != nil {
        t.Fatalf("第一次写入失败: %v", err)
    }

    // 第二次写入（应覆盖）
    content2 := "second content\n"
    if err := FilePutContents(testFile, content2); err != nil {
        t.Fatalf("第二次写入失败: %v", err)
    }

    // 验证内容（应只有第二次写入的内容）
    got, err := os.ReadFile(testFile)
    if err != nil {
        t.Fatalf("读取文件失败: %v", err)
    }

    if string(got) != content2 {
        t.Errorf("文件内容不匹配\n期望: %q\n实际: %q", content2, got)
    }
}

func TestFileExist(t *testing.T) {
    // 测试空路径
    if FileExist("") {
        t.Error("空路径应返回false")
    }

    // 测试不存在的路径
    if FileExist("/not/exist/path") {
        t.Error("不存在路径应返回false")
    }

    // 测试存在的文件
    tmpDir := os.TempDir()
    testFile := filepath.Join(tmpDir, "test_exist.txt")
    if err := os.WriteFile(testFile, []byte("test"), 0644); err != nil {
        t.Fatalf("创建测试文件失败: %v", err)
    }
    defer os.Remove(testFile)

    if !FileExist(testFile) {
        t.Error("已存在文件应返回true")
    }
}

func TestFileGetContentsAsByte(t *testing.T) {
    // 创建临时目录
    tmpDir, err := os.MkdirTemp("", "test_get_*")
    if err != nil {
        t.Fatalf("创建临时目录失败: %v", err)
    }
    defer os.RemoveAll(tmpDir)

    // 用于存储需要在测试结束时关闭的文件
    var filesToClose []*os.File

    // 创建测试文件
    normalFile := filepath.Join(tmpDir, "normal.txt")
    normalContent := []byte("test content\n")
    if err := os.WriteFile(normalFile, normalContent, 0644); err != nil {
        t.Fatalf("创建普通测试文件失败: %v", err)
    }

    // 创建大文件
    largeFile := filepath.Join(tmpDir, "large.txt")
    largeContent := bytes.Repeat([]byte("a"), 5*1024*1024) // 5MB
    if err := os.WriteFile(largeFile, largeContent, 0644); err != nil {
        t.Fatalf("创建大文件失败: %v", err)
    }

    // 创建二进制文件
    binaryFile := filepath.Join(tmpDir, "binary.dat")
    binaryContent := []byte{0x00, 0x01, 0x02, 0x03, 0xFF, 0xFE, 0xFD}
    if err := os.WriteFile(binaryFile, binaryContent, 0644); err != nil {
        t.Fatalf("创建二进制文件失败: %v", err)
    }

    // 创建只读文件
    readonlyFile := filepath.Join(tmpDir, "readonly.txt")
    readonlyContent := []byte("readonly content")
    if err := os.WriteFile(readonlyFile, readonlyContent, 0444); err != nil {
        t.Fatalf("创建只读文件失败: %v", err)
    }

    tests := []struct {
        name     string
        path     string
        want     []byte
        wantNil  bool
        setup    func() error
        teardown func() error
    }{
        {
            name:    "空路径",
            path:    "",
            wantNil: true,
        },
        {
            name:    "不存在的路径",
            path:    "/not/exist/path",
            wantNil: true,
        },
        {
            name: "正常文件",
            path: normalFile,
            want: normalContent,
        },
        {
            name: "大文件",
            path: largeFile,
            want: largeContent,
        },
        {
            name: "二进制文件",
            path: binaryFile,
            want: binaryContent,
        },
        {
            name: "只读文件",
            path: readonlyFile,
            want: readonlyContent,
        },
        {
            name:    "目录",
            path:    tmpDir,
            wantNil: true,
        },
        {
            name: "文件被锁定",
            path: filepath.Join(tmpDir, "locked.txt"),
            setup: func() error {
                content := []byte("locked content")
                if err := os.WriteFile(filepath.Join(tmpDir, "locked.txt"), content, 0644); err != nil {
                    return err
                }
                // 打开文件并保持锁定
                f, err := os.OpenFile(filepath.Join(tmpDir, "locked.txt"), os.O_RDWR, 0644)
                if err != nil {
                    return err
                }
                filesToClose = append(filesToClose, f)
                return nil
            },
            want: []byte("locked content"),
        },
    }

    // 确保在测试结束时关闭所有打开的文件
    defer func() {
        for _, f := range filesToClose {
            f.Close()
        }
    }()

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if tt.setup != nil {
                if err := tt.setup(); err != nil {
                    t.Fatalf("设置测试环境失败: %v", err)
                }
            }

            got := FileGetContentsAsByte(tt.path)

            if tt.teardown != nil {
                if err := tt.teardown(); err != nil {
                    t.Errorf("清理测试环境失败: %v", err)
                }
            }

            if tt.wantNil {
                if got != nil {
                    t.Errorf("FileGetContentsAsByte() = %v, 期望 nil", got)
                }
                return
            }

            if !bytes.Equal(got, tt.want) {
                if len(got) > 100 || len(tt.want) > 100 {
                    t.Errorf("FileGetContentsAsByte() 返回的内容长度 = %d, 期望长度 %d", len(got), len(tt.want))
                } else {
                    t.Errorf("FileGetContentsAsByte() = %v, 期望 %v", got, tt.want)
                }
            }
        })
    }
}

func TestFileGetContents(t *testing.T) {
    // 创建临时目录
    tmpDir, err := os.MkdirTemp("", "test_get_str_*")
    if err != nil {
        t.Fatalf("创建临时目录失败: %v", err)
    }
    defer os.RemoveAll(tmpDir)

    // 创建测试文件
    normalFile := filepath.Join(tmpDir, "normal.txt")
    normalContent := "test content\n"
    if err := os.WriteFile(normalFile, []byte(normalContent), 0644); err != nil {
        t.Fatalf("创建普通测试文件失败: %v", err)
    }

    // 创建包含特殊字符的文件
    specialFile := filepath.Join(tmpDir, "special.txt")
    specialContent := "特殊字符：\n中文\t制表符\r\n换行符"
    if err := os.WriteFile(specialFile, []byte(specialContent), 0644); err != nil {
        t.Fatalf("创建特殊字符文件失败: %v", err)
    }

    // 创建大文件
    largeFile := filepath.Join(tmpDir, "large.txt")
    largeContent := strings.Repeat("a", 5*1024*1024) // 5MB
    if err := os.WriteFile(largeFile, []byte(largeContent), 0644); err != nil {
        t.Fatalf("创建大文件失败: %v", err)
    }

    tests := []struct {
        name    string
        path    string
        want    string
        setup   func() error
        cleanup func() error
    }{
        {
            name: "空路径",
            path: "",
            want: "",
        },
        {
            name: "不存在的路径",
            path: "/not/exist/path",
            want: "",
        },
        {
            name: "正常文件",
            path: normalFile,
            want: normalContent,
        },
        {
            name: "特殊字符文件",
            path: specialFile,
            want: specialContent,
        },
        {
            name: "大文件",
            path: largeFile,
            want: largeContent,
        },
        {
            name: "目录",
            path: tmpDir,
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if tt.setup != nil {
                if err := tt.setup(); err != nil {
                    t.Fatalf("设置测试环境失败: %v", err)
                }
            }

            got := FileGetContents(tt.path)

            if tt.cleanup != nil {
                if err := tt.cleanup(); err != nil {
                    t.Errorf("清理测试环境失败: %v", err)
                }
            }

            if got != tt.want {
                if len(got) > 100 || len(tt.want) > 100 {
                    t.Errorf("FileGetContents() 返回的内容长度 = %d, 期望长度 %d", len(got), len(tt.want))
                } else {
                    t.Errorf("FileGetContents() = %q, 期望 %q", got, tt.want)
                }
            }
        })
    }
}

func TestFileGetExtension(t *testing.T) {
    tests := []struct {
        path     string
        expected string
    }{
        {"", ""},
        {"README", ""},
        {"a.txt", "txt"},
        {"/path/to/file.jpg", "jpg"},
        {"/path/to/archive.tar.gz", "gz"},
        {"file.", ""},
        {".gitignore", "gitignore"},
        {"path/to/.config", "config"},
    }

    for _, test := range tests {
        if got := FileGetExtension(test.path); got != test.expected {
            t.Errorf("FileGetExtension(%q) = %q，期望 %q", test.path, got, test.expected)
        }
    }
}

func TestFileGetName(t *testing.T) {
    tests := []struct {
        path     string
        expected string
    }{
        {"", ""},
        {"a.txt", "a.txt"},
        {"/path/to/file.jpg", "file.jpg"},
        {"file.txt?v=1", "file.txt"},
        {"image.jpg?size=large#preview", "image.jpg"},
        {"/path/to/doc.pdf#page=1", "doc.pdf"},
        {"archive.zip?dl=1#download", "archive.zip"},
        {"/var/www/index.html", "index.html"},
        {"C:\\Windows\\System32\\file.dll", "file.dll"},
        {"../relative/path/file.txt", "file.txt"},
        {"./current/file.txt", "file.txt"},
        {"file", "file"},
    }

    for _, test := range tests {
        if got := FileGetName(test.path); got != test.expected {
            t.Errorf("FileGetName(%q) = %q，期望 %q", test.path, got, test.expected)
        }
    }
}

func TestFileSha1(t *testing.T) {
    // 测试空路径
    if _, err := FileSha1(""); err != ErrNotRegularFile {
        t.Errorf("空路径应返回ErrNotRegularFile，但得到 %v", err)
    }

    // 测试不存在的路径
    if _, err := FileSha1("/not/exist/path"); err != ErrNotRegularFile {
        t.Errorf("不存在路径应返回ErrNotRegularFile，但得到 %v", err)
    }

    // 测试正常计算
    tmpDir := os.TempDir()
    testFile := filepath.Join(tmpDir, "test_sha1.txt")
    content := []byte("test content for sha1")
    if err := os.WriteFile(testFile, content, 0644); err != nil {
        t.Fatalf("创建测试文件失败: %v", err)
    }
    defer os.Remove(testFile)

    got, err := FileSha1(testFile)
    if err != nil {
        t.Fatalf("计算SHA1失败: %v", err)
    }

    h := sha1.New()
    h.Write(content)
    expected := fmt.Sprintf("%x", h.Sum(nil))

    if got != expected {
        t.Errorf("SHA1不匹配\n期望: %s\n实际: %s", expected, got)
    }
}

func TestFileSha256(t *testing.T) {
    // 测试空路径
    if _, err := FileSha256(""); err != ErrNotRegularFile {
        t.Errorf("空路径应返回ErrNotRegularFile，但得到 %v", err)
    }

    // 测试不存在的路径
    if _, err := FileSha256("/not/exist/path"); err != ErrNotRegularFile {
        t.Errorf("不存在路径应返回ErrNotRegularFile，但得到 %v", err)
    }

    // 测试正常计算
    tmpDir := os.TempDir()
    testFile := filepath.Join(tmpDir, "test_sha256.txt")
    content := []byte("test content for sha256")
    if err := os.WriteFile(testFile, content, 0644); err != nil {
        t.Fatalf("创建测试文件失败: %v", err)
    }
    defer os.Remove(testFile)

    got, err := FileSha256(testFile)
    if err != nil {
        t.Fatalf("计算SHA256失败: %v", err)
    }

    h := sha256.New()
    h.Write(content)
    expected := fmt.Sprintf("%x", h.Sum(nil))

    if got != expected {
        t.Errorf("SHA256不匹配\n期望: %s\n实际: %s", expected, got)
    }
}

func TestFileMd5(t *testing.T) {
    // 测试空路径
    if _, err := FileMd5(""); err != ErrNotRegularFile {
        t.Errorf("空路径应返回ErrNotRegularFile，但得到 %v", err)
    }

    // 测试不存在的路径
    if _, err := FileMd5("/not/exist/path"); err != ErrNotRegularFile {
        t.Errorf("不存在路径应返回ErrNotRegularFile，但得到 %v", err)
    }

    // 测试正常计算
    tmpDir := os.TempDir()
    testFile := filepath.Join(tmpDir, "test_md5.txt")
    content := []byte("test content for md5")
    if err := os.WriteFile(testFile, content, 0644); err != nil {
        t.Fatalf("创建测试文件失败: %v", err)
    }
    defer os.Remove(testFile)

    got, err := FileMd5(testFile)
    if err != nil {
        t.Fatalf("计算MD5失败: %v", err)
    }

    h := md5.New()
    h.Write(content)
    expected := fmt.Sprintf("%x", h.Sum(nil))

    if got != expected {
        t.Errorf("MD5不匹配\n期望: %s\n实际: %s", expected, got)
    }
}

func TestFileCrc32(t *testing.T) {
    // 测试空路径
    if _, err := FileCrc32(""); err != ErrNotRegularFile {
        t.Errorf("空路径应返回ErrNotRegularFile，但得到 %v", err)
    }

    // 测试不存在的路径
    if _, err := FileCrc32("/not/exist/path"); err != ErrNotRegularFile {
        t.Errorf("不存在路径应返回ErrNotRegularFile，但得到 %v", err)
    }

    // 测试正常计算
    tmpDir := os.TempDir()
    testFile := filepath.Join(tmpDir, "test_crc32.txt")
    content := []byte("test content for crc32")
    if err := os.WriteFile(testFile, content, 0644); err != nil {
        t.Fatalf("创建测试文件失败: %v", err)
    }
    defer os.Remove(testFile)

    got, err := FileCrc32(testFile)
    if err != nil {
        t.Fatalf("计算CRC32失败: %v", err)
    }

    table := crc32.MakeTable(crc32.IEEE)
    expected := fmt.Sprintf("%x", crc32.Checksum(content, table))

    if got != expected {
        t.Errorf("CRC32不匹配\n期望: %s\n实际: %s", expected, got)
    }
}

func TestFileCopy(t *testing.T) {
    // 测试空路径
    if err := FileCopy("", "dst.txt"); err != ErrNotRegularFile {
        t.Errorf("空源路径应返回ErrNotRegularFile，但得到 %v", err)
    }

    // 测试不存在的源文件
    if err := FileCopy("/not/exist/path", "dst.txt"); err != ErrNotRegularFile {
        t.Errorf("不存在源路径应返回ErrNotRegularFile，但得到 %v", err)
    }

    // 测试正常复制
    tmpDir := os.TempDir()
    srcFile := filepath.Join(tmpDir, "test_copy_src.txt")
    dstFile := filepath.Join(tmpDir, "test_copy_dst.txt")
    content := []byte("test content for copy")

    // 创建源文件
    if err := os.WriteFile(srcFile, content, 0644); err != nil {
        t.Fatalf("创建源文件失败: %v", err)
    }
    defer os.Remove(srcFile)
    defer os.Remove(dstFile)

    // 执行复制
    if err := FileCopy(srcFile, dstFile); err != nil {
        t.Fatalf("复制文件失败: %v", err)
    }

    // 验证目标文件内容
    got, err := os.ReadFile(dstFile)
    if err != nil {
        t.Fatalf("读取目标文件失败: %v", err)
    }

    if string(got) != string(content) {
        t.Errorf("复制后文件内容不匹配\n期望: %q\n实际: %q", content, got)
    }
}

func TestFileMove(t *testing.T) {
    // 测试空路径
    if err := FileMove("", "dst.txt"); err != ErrNotRegularFile {
        t.Errorf("空源路径应返回ErrNotRegularFile，但得到 %v", err)
    }

    // 测试不存在的源文件
    if err := FileMove("/not/exist/path", "dst.txt"); err != ErrNotRegularFile {
        t.Errorf("不存在源路径应返回ErrNotRegularFile，但得到 %v", err)
    }

    // 测试正常移动
    tmpDir := os.TempDir()
    srcFile := filepath.Join(tmpDir, "test_move_src.txt")
    dstFile := filepath.Join(tmpDir, "test_move_dst.txt")
    content := []byte("test content for move")

    // 创建源文件
    if err := os.WriteFile(srcFile, content, 0644); err != nil {
        t.Fatalf("创建源文件失败: %v", err)
    }
    defer os.Remove(srcFile) // 如果移动失败，清理源文件
    defer os.Remove(dstFile) // 清理目标文件

    // 执行移动
    if err := FileMove(srcFile, dstFile); err != nil {
        t.Fatalf("移动文件失败: %v", err)
    }

    // 验证源文件不存在
    if FileExist(srcFile) {
        t.Error("移动后源文件仍然存在")
    }

    // 验证目标文件内容
    got, err := os.ReadFile(dstFile)
    if err != nil {
        t.Fatalf("读取目标文件失败: %v", err)
    }

    if string(got) != string(content) {
        t.Errorf("移动后文件内容不匹配\n期望: %q\n实际: %q", content, got)
    }
}

func TestFileWithWorkDirectory(t *testing.T) {
    // 保存原始工作目录
    oldWorkDirectory := WorkDirectory
    defer func() {
        WorkDirectory = oldWorkDirectory
    }()

    // 设置测试工作目录
    WorkDirectory = filepath.FromSlash("/test/work/dir")

    tests := []struct {
        path     string
        expected string
    }{
        {"file.txt", filepath.FromSlash("/test/work/dir/file.txt")},
        {"subdir/file.txt", filepath.FromSlash("/test/work/dir/subdir/file.txt")},
        {"/absolute/path.txt", filepath.FromSlash("/test/work/dir/absolute/path.txt")},
        {"", filepath.FromSlash("/test/work/dir")},
        {"./file.txt", filepath.FromSlash("/test/work/dir/file.txt")},
        {"../file.txt", filepath.FromSlash("/test/work/file.txt")},
    }

    for _, test := range tests {
        if got := FileWithWorkDirectory(test.path); got != test.expected {
            t.Errorf("FileWithWorkDirectory(%q) = %q，期望 %q", test.path, got, test.expected)
        }
    }
}
