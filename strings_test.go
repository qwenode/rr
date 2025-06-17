package rr

import (
    "bytes"
    "errors"
    "reflect"
    "testing"
)

func TestS_Sha1(t *testing.T) {
    type args struct{}
    tests := []struct {
        name string
        r    S
        args args
        want string
    }{
        {
            "正常字符串",
            S("hello"),
            args{},
            "aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d",
        },
        {
            "空字符串",
            S(""),
            args{},
            "",
        },
    }
    for _, tt := range tests {
        t.Run(
            tt.name, func(t *testing.T) {
                if got := tt.r.Sha1(); got != tt.want {
                    t.Errorf("Sha1() = %s, want %s", got, tt.want)
                }
            },
        )
    }
}
func TestS_RemoveExtension(t *testing.T) {
    type args struct{}
    tests := []struct {
        name string
        r    S
        args args
        want S
    }{
        {
            "有扩展名",
            S("file.txt"),
            args{},
            S("file"),
        },
        {
            "没有扩展名",
            S("file"),
            args{},
            S("file"),
        },
        {
            "空字符串",
            S(""),
            args{},
            S(""),
        },
    }
    for _, tt := range tests {
        t.Run(
            tt.name, func(t *testing.T) {
                if got := tt.r.RemoveExtension(); got != tt.want {
                    t.Errorf("RemoveExtension() = %v, want %v", got, tt.want)
                }
            },
        )
    }
}
func TestS_AsLines(t *testing.T) {
    type args struct {
        s S
    }
    tests := []struct {
        name string
        r    S
        args args
        want []string
    }{
        {
            "正常情况",
            S("line1\nline2\nline3"),
            args{S("line1\nline2\nline3")},
            []string{"line1", "line2", "line3"},
        },
        {
            "空字符串",
            S(""),
            args{S("")},
            []string{},
        },
    }
    for _, tt := range tests {
        t.Run(
            tt.name, func(t *testing.T) {
                if got := tt.r.AsLines(); !reflect.DeepEqual(got, tt.want) {
                    t.Errorf("AsLines() = %v, want %v", got, tt.want)
                }
            },
        )
    }
}
func TestRString_JsonUnmarshal(t *testing.T) {
    type args struct {
        v interface{}
    }
    tests := []struct {
        name    string
        r       S
        args    args
        wantErr bool
    }{
        {
            "正常情况有效 JSON 字符串和正确类型接收对象",
            S(`{"key": "value"}`),
            args{&struct{ Key string }{Key: ""}},
            false,
        },
        {
            "边界情况空字符串",
            S(""),
            args{&struct{}{}},
            true,
        },
        {
            "错误输入无效 JSON 字符串",
            S("invalid_json"),
            args{&struct{}{}},
            true,
        },
    }
    for _, tt := range tests {
        t.Run(
            tt.name, func(t *testing.T) {
                err := tt.r.JsonUnmarshal(tt.args.v)
                if (err != nil) != tt.wantErr {
                    t.Errorf("JsonUnmarshal() error = %v, wantErr %v", err, tt.wantErr)
                }
            },
        )
    }
}
func TestRString_UrlDecode(t *testing.T) {
    type args struct{}
    tests := []struct {
        name string
        r    S
        args args
        want S
    }{
        {
            "正常情况非空字符串",
            S("encoded_value"),
            args{},
            S("encoded_value"), // 假设实际解码后的结果
        },
        {
            "边界情况空字符串",
            S(""),
            args{},
            S(""),
        },
        {
            "逻辑检查特殊编码字符串",
            S("%2B"),
            args{},
            S("+"),
        },
    }
    for _, tt := range tests {
        t.Run(
            tt.name, func(t *testing.T) {
                if got := tt.r.UrlDecode(); got != tt.want {
                    t.Errorf("UrlDecode() = %v, want %v", got, tt.want)
                }
            },
        )
    }
}
func TestRString_ToUrl(t *testing.T) {
    type args struct {
        r     S
        https []bool
    }
    tests := []struct {
        name string
        args args
        want string
    }{
        {
            "有内容且不指定 https 为 http 格式",
            args{S("example.com"), []bool{}},
            "http://example.com",
        },
        {
            "有内容且指定 https 为 https 格式",
            args{S("example.com"), []bool{true}},
            "https://example.com",
        },
        {
            "已有协议的内容且不指定 https 保留原协议",
            args{S("http://example.org"), []bool{}},
            "http://example.org",
        },
        {
            "空字符串",
            args{S(""), []bool{}},
            "",
        },
    }
    for _, tt := range tests {
        t.Run(
            tt.name, func(t *testing.T) {
                if got := tt.args.r.AsUrl(tt.args.https...); got != tt.want {
                    t.Errorf("ToUrl() = %v, want %v", got, tt.want)
                }
            },
        )
    }
}
func TestRString_IsUrl(t *testing.T) {
    type args struct {
        r S
    }
    tests := []struct {
        name string
        args args
        want bool
    }{
        {
            "有效 URL",
            args{S("https://example.com")},
            true,
        },
        {
            "另一个有效 URL",
            args{S("http://example.org")},
            true,
        },
        {
            "无效 URL",
            args{S("invalid-url")},
            false,
        },
        {
            "空字符串",
            args{S("")},
            false,
        },
    }
    for _, tt := range tests {
        t.Run(
            tt.name, func(t *testing.T) {
                if got := tt.args.r.IsUrl(); got != tt.want {
                    t.Errorf("IsUrl() = %v, want %v", got, tt.want)
                }
            },
        )
    }
}
func TestGetFirst(t *testing.T) {
    // 测试包含分隔符的情况
    r := S("apple,banana,cherry")
    expected := S("apple")
    result := r.GetFirst(",")
    if result != expected {
        t.Errorf("Expected %s, got %s", expected, result)
    }

    // 测试不包含分隔符的情况
    r = S("orange")
    expected = S("orange")
    result = r.GetFirst(",")
    if result != expected {
        t.Errorf("Expected %s, got %s", expected, result)
    }
    r = S(",orange")
    expected = S("")
    result = r.GetFirst(",")
    if result != expected {
        t.Errorf("Expected %s, got %s", expected, result)
    }
}
func TestRemoveLast(t *testing.T) {
    // 正常情况
    input1 := S("Hello, World!")
    sep1 := ", "
    expected1 := S("Hello")
    output1 := input1.RemoveLast(sep1)
    if output1 != expected1 {
        t.Errorf("Expected %v but got %v", expected1, output1)
    }

    input2 := S("Hello, World!")
    sep2 := "!"
    expected2 := S("Hello, World")
    output2 := input2.RemoveLast(sep2)
    if output2 != expected2 {
        t.Errorf("Expected %v but got %v", expected2, output2)
    }

    // 边界情况
    input3 := S("Hello, World!")
    sep3 := ":"
    expected3 := input3
    output3 := input3.RemoveLast(sep3)
    if output3 != expected3 {
        t.Errorf("Expected %v but got %v", expected3, output3)
    }

    input4 := S("")
    sep4 := ","
    expected4 := input4
    output4 := input4.RemoveLast(sep4)
    if output4 != expected4 {
        t.Errorf("Expected %v but got %v", expected4, output4)
    }

    input5 := S("Hello, World!")
    sep5 := ","
    expected5 := S("Hello")
    output5 := input5.RemoveLast(sep5)
    if output5 != expected5 {
        t.Errorf("Expected %v but got %v", expected5, output5)
    }

    // 错误情况
    input6 := S("Hello, World!")
    sep6 := "W"
    expected6 := S("Hello, ")
    output6 := input6.RemoveLast(sep6)
    if output6 != expected6 {
        t.Errorf("Expected %v but got %v", expected6, output6)
    }
}
func TestGetLast(t *testing.T) {
    // 测试包含分隔符的情况
    r := S("apple,banana,cherry")
    expected := S("cherry")
    result := r.GetLast(",")
    if result != expected {
        t.Errorf("Expected %s, got %s", expected, result)
    }

    // 测试不包含分隔符的情况
    r = S("lemon")
    expected = S("lemon")
    result = r.GetLast(",")
    if result != expected {
        t.Errorf("Expected %s, got %s", expected, result)
    }
    // 
    r = S("lemon,")
    expected = S("")
    result = r.GetLast(",")
    if result != expected {
        t.Errorf("Expected %s, got %s", expected, result)
    }
}
func TestRString_RemoveFirst(t *testing.T) {
    type args struct {
        sep string
    }
    tests := []struct {
        name string
        r    S
        args args
        want S
    }{
        {
            "正常情况",
            S("abcdefg"),
            args{"cd"},
            S("efg"),
        },
        {
            "边界情况，开头就是分隔符",
            S("cdabc"),
            args{"cd"},
            S("abc"),
        },
        {
            "不存在分隔符的情况",
            S("abc"),
            args{"cd"},
            S("abc"),
        },
    }
    for _, tt := range tests {
        t.Run(
            tt.name, func(t *testing.T) {
                if got := tt.r.RemoveFirst(tt.args.sep); got != tt.want {
                    t.Errorf("RemoveFirst() = %v, want %v", got, tt.want)
                }
            },
        )
    }
}
func TestRString_GetSecond(t *testing.T) {
    type args struct {
        sep string
    }
    tests := []struct {
        name string
        r    S
        args args
        want S
    }{
        {
            "正常情况有分隔符",
            S("abcdefg"),
            args{"cd"},
            S("efg"),
        },
        {
            "边界情况开头就是分隔符",
            S("cdabc"),
            args{"cd"},
            S("abc"),
        },
        {
            "不存在分隔符",
            S("abc"),
            args{"cd"},
            S("abc"),
        },
    }
    for _, tt := range tests {
        t.Run(
            tt.name, func(t *testing.T) {
                if got := tt.r.GetSecond(tt.args.sep); got != tt.want {
                    t.Errorf("GetSecond() = %v, want %v", got, tt.want)
                }
            },
        )
    }
}
func TestRString_Base64DecodeAsBytes(t *testing.T) {
    type args struct{}
    tests := []struct {
        name string
        r    S
        args args
        want []byte
    }{
        {
            "正常情况",
            S("SGVsbG8="), // 对应 "Hello" 的 base64 编码
            args{},
            []byte("Hello"),
        },
    }
    for _, tt := range tests {
        t.Run(
            tt.name, func(t *testing.T) {
                if got := tt.r.Base64DecodeAsBytes(); !bytes.Equal(got, tt.want) {
                    t.Errorf("Base64DecodeAsBytes() = %v, want %v", got, tt.want)
                }
            },
        )
    }
}

func TestJsonSerialize(t *testing.T) {
    type args struct {
        v interface{}
    }
    tests := []struct {
        name string
        r    func(v interface{}) string
        args args
        want string
    }{
        {
            "正常情况 - 结构体",
            JsonSerialize,
            args{struct {
                Name string
                Age  int
            }{
                "Alice",
                25,
            }},
            `{"Name":"Alice","Age":25}`,
        },
        {
            "正常情况 - 字符串",
            JsonSerialize,
            args{"Hello"},
            `"Hello"`,
        },
        {
            "边界情况 - 空值",
            JsonSerialize,
            args{nil},
            "null",
        },
    }
    for _, tt := range tests {
        t.Run(
            tt.name, func(t *testing.T) {
                if got := tt.r(tt.args.v); got != tt.want {
                    t.Errorf("JsonSerialize() = %v, want %v", got, tt.want)
                }
            },
        )
    }
}
func TestRString_AsFloat(t *testing.T) {
    type args struct{}
    tests := []struct {
        name string
        r    S
        args args
        want float64
    }{
        {
            "正常转换",
            S("123.45"),
            args{},
            123.45,
        },
        {
            "非数字字符串",
            S("abc"),
            args{},
            0,
        },
    }
    for _, tt := range tests {
        t.Run(
            tt.name, func(t *testing.T) {
                if got := tt.r.AsFloat(); got != tt.want {
                    t.Errorf("AsFloat() = %v, want %v", got, tt.want)
                }
            },
        )
    }
}
func TestRString_AsInt(t *testing.T) {
    type args struct{}
    tests := []struct {
        name string
        r    S
        args args
        want int
    }{
        {
            "正常转换数字",
            S("123"),
            args{},
            123,
        },
        {
            "非数字字符串",
            S("abc"),
            args{},
            0,
        },
    }
    for _, tt := range tests {
        t.Run(
            tt.name, func(t *testing.T) {
                if got := tt.r.AsInt(); got != tt.want {
                    t.Errorf("AsInt() = %v, want %v", got, tt.want)
                }
            },
        )
    }
}
func TestRString_AsInt64(t *testing.T) {
    type args struct{}
    tests := []struct {
        name string
        r    S
        args args
        want int64
    }{
        {
            "正常数字转换",
            S("123456"),
            args{},
            123456,
        },
        {
            "非数字字符串",
            S("abc"),
            args{},
            0,
        },
    }
    for _, tt := range tests {
        t.Run(
            tt.name, func(t *testing.T) {
                if got := tt.r.AsInt64(); got != tt.want {
                    t.Errorf("AsInt64() = %v, want %v", got, tt.want)
                }
            },
        )
    }
}
func TestRString_TrimSpace(t *testing.T) {
    type args struct{}
    tests := []struct {
        name string
        r    S
        args args
        want S
    }{
        {
            "正常去空格",
            S("  abc   "),
            args{},
            S("abc"),
        },
        {
            "本身无空格",
            S("abc"),
            args{},
            S("abc"),
        },
    }
    for _, tt := range tests {
        t.Run(
            tt.name, func(t *testing.T) {
                if got := tt.r.TrimSpace(); got != tt.want {
                    t.Errorf("TrimSpace() = %v, want %v", got, tt.want)
                }
            },
        )
    }
}

func TestRString_JsonUnSerialize(t *testing.T) {
    type args struct {
        v interface{}
    }
    tests := []struct {
        name string
        r    S
        args args
        want error
    }{
        {
            "正常解析结构体",
            S(`{"name":"Alice","age":25}`),
            args{&struct {
                Name string
                Age  int
            }{}},
            nil,
        },
        {
            "无效 JSON 字符串",
            S("invalid json"),
            args{&struct{}{}},
            errors.New("invalid character 'i' looking for beginning of value"), // 根据实际错误情况填写
        },
    }
    for _, tt := range tests {
        t.Run(
            tt.name, func(t *testing.T) {
                err := tt.r.JsonUnSerialize(tt.args.v)
                if err != nil && err.Error() != tt.want.Error() {
                    t.Errorf("JsonUnSerialize() error = %v, want %v", err, tt.want)
                }
            },
        )
    }
}

func TestRString_AsBool(t *testing.T) {
    type args struct{}
    tests := []struct {
        name string
        r    S
        args args
        want bool
    }{
        {
            "正常转换为 true - '1'",
            S("1"),
            args{},
            true,
        },
        {
            "正常转换为 true - 't'",
            S("t"),
            args{},
            true,
        },
        {
            "正常转换为 true - 'true'",
            S("true"),
            args{},
            true,
        },
        {
            "正常转换为 true - 'ok'",
            S("ok"),
            args{},
            true,
        },
        {
            "正常转换为 true - 'yes'",
            S("yes"),
            args{},
            true,
        },
        {
            "正常转换为 true -'sure'",
            S("sure"),
            args{},
            true,
        },
        {
            "不转换为 true - 'abc'",
            S("abc"),
            args{},
            false,
        },
    }
    for _, tt := range tests {
        t.Run(
            tt.name, func(t *testing.T) {
                if got := tt.r.AsBool(); got != tt.want {
                    t.Errorf("AsBool() = %v, want %v", got, tt.want)
                }
            },
        )
    }
}
func TestRString_Upper(t *testing.T) {
    type args struct{}
    tests := []struct {
        name string
        r    S
        args args
        want S
    }{
        {
            "正常转换为大写",
            S("abc"),
            args{},
            S("ABC"),
        },
        {
            "本身就是大写",
            S("ABC"),
            args{},
            S("ABC"),
        },
    }
    for _, tt := range tests {
        t.Run(
            tt.name, func(t *testing.T) {
                if got := tt.r.Upper(); got != tt.want {
                    t.Errorf("Upper() = %v, want %v", got, tt.want)
                }
            },
        )
    }
}
func TestRString_Lower(t *testing.T) {
    type args struct{}
    tests := []struct {
        name string
        r    S
        args args
        want S
    }{
        {
            "正常转换为小写",
            S("ABC"),
            args{},
            S("abc"),
        },
        {
            "本身就是小写",
            S("abc"),
            args{},
            S("abc"),
        },
    }
    for _, tt := range tests {
        t.Run(
            tt.name, func(t *testing.T) {
                if got := tt.r.Lower(); got != tt.want {
                    t.Errorf("Lower() = %v, want %v", got, tt.want)
                }
            },
        )
    }
}
func TestRString_TrimLeft(t *testing.T) {
    type args struct {
        cut string
    }
    tests := []struct {
        name string
        r    S
        args args
        want S
    }{
        {
            "正常裁剪左边",
            S("abcdefg"),
            args{"abc"},
            S("defg"),
        },
        {
            "裁剪部分不匹配",
            S("abcdefg"),
            args{"xyz"},
            S("abcdefg"),
        },
    }
    for _, tt := range tests {
        t.Run(
            tt.name, func(t *testing.T) {
                if got := tt.r.TrimLeft(tt.args.cut); got != tt.want {
                    t.Errorf("TrimLeft() = %v, want %v", got, tt.want)
                }
            },
        )
    }
}
func TestRString_TrimRight(t *testing.T) {
    type args struct {
        cut string
    }
    tests := []struct {
        name string
        r    S
        args args
        want S
    }{
        {
            "正常裁剪右边",
            S("abcdefg"),
            args{"efg"},
            S("abcd"),
        },
        {
            "裁剪部分不匹配",
            S("abcdefg"),
            args{"xyz"},
            S("abcdefg"),
        },
        {
            "裁剪部分不匹配",
            S("abcdefg"),
            args{"cde"},
            S("abcdefg"),
        },
    }
    for _, tt := range tests {
        t.Run(
            tt.name, func(t *testing.T) {
                if got := tt.r.TrimRight(tt.args.cut); got != tt.want {
                    t.Errorf("TrimRight() = %v, want %v", got, tt.want)
                }
            },
        )
    }
}

func TestRString_AsUint64(t *testing.T) {
    type args struct{}
    tests := []struct {
        name string
        r    S
        args args
        want uint64
    }{
        {
            "正常转换",
            S("123"),
            args{},
            123,
        },
        {
            "转换失败",
            S("abc"),
            args{},
            0,
        },
    }
    for _, tt := range tests {
        t.Run(
            tt.name, func(t *testing.T) {
                if got := tt.r.AsUint64(); got != tt.want {
                    t.Errorf("AsUint64() = %v, want %v", got, tt.want)
                }
            },
        )
    }
}
func TestRString_SanitizeAsInt(t *testing.T) {
    type args struct{}
    tests := []struct {
        name string
        r    S
        args args
        want int
    }{
        {
            "空字符串返回 0",
            S(""),
            args{},
            0,
        },
        {
            "正常数字转换",
            S("1x2#3"),
            args{},
            123,
        },
        {
            "有非数字字符但能提取出数字",
            S("abc12_-3"),
            args{},
            123,
        },
        {
            "带负号且转换正确",
            S("-123"),
            args{},
            -123,
        },
    }
    for _, tt := range tests {
        t.Run(
            tt.name, func(t *testing.T) {
                //if got := tt.r.SanitizeAsInt(); got != tt.want {
                //    t.Errorf("SanitizeAsInt() = %v, want %v", got, tt.want)
                //}
            },
        )
    }
}
func TestRString_SanitizeAsInt64(t *testing.T) {
    type args struct{}
    tests := []struct {
        name string
        r    S
        args args
        want int64
    }{
        {
            "空字符串返回 0",
            S(""),
            args{},
            0,
        },
        {
            "正常数字转换",
            S("1)23"),
            args{},
            123,
        },
        {
            "有非数字字符但能提取出数字",
            S("abc12#3"),
            args{},
            123,
        },
        {
            "带负号且转换正确",
            S("-12#3"),
            args{},
            -123,
        },
    }
    for _, tt := range tests {
        t.Run(
            tt.name, func(t *testing.T) {
                //if got := tt.r.SanitizeAsInt64(); got != tt.want {
                //    t.Errorf("SanitizeAsInt64() = %v, want %v", got, tt.want)
                //}
            },
        )
    }
}
func TestRString_Prepend(t *testing.T) {
    type args struct {
        s string
    }
    tests := []struct {
        name string
        r    S
        args args
        want S
    }{
        {
            "正常添加前缀",
            S("abc"),
            args{"xyz"},
            S("xyzabc"),
        },
        {
            "空字符串添加前缀",
            S("def"),
            args{""},
            S("def"),
        },
    }
    for _, tt := range tests {
        t.Run(
            tt.name, func(t *testing.T) {
                if got := tt.r.Prepend(tt.args.s); got != tt.want {
                    t.Errorf("Prepend() = %v, want %v", got, tt.want)
                }
            },
        )
    }
}
func TestRString_Append(t *testing.T) {
    type args struct {
        s string
    }
    tests := []struct {
        name string
        r    S
        args args
        want S
    }{
        {
            "正常添加后缀",
            S("abc"),
            args{"xyz"},
            S("abcxyz"),
        },
        {
            "空字符串添加后缀",
            S("def"),
            args{""},
            S("def"),
        },
    }
    for _, tt := range tests {
        t.Run(
            tt.name, func(t *testing.T) {
                if got := tt.r.Append(tt.args.s); got != tt.want {
                    t.Errorf("Append() = %v, want %v", got, tt.want)
                }
            },
        )
    }
}
func TestRString_SanitizeAsHostname(t *testing.T) {
    type args struct{}
    tests := []struct {
        name string
        r    S
        args args
        want string
    }{
        {
            "正常有协议的 URL",
            S("https://example.com"),
            args{},
            "example.com",
        },
        {
            "无协议添加 http:// 后提取",
            S("example.com"),
            args{},
            "example.com",
        },
        {
            "无协议添加 http:// 后提取",
            S("127.0.0.1"),
            args{},
            "127.0.0.1",
        },
        {
            "无协议添加 http:// 后提取",
            S("127.0.0.1:222/wowig"),
            args{},
            "127.0.0.1",
        },
        {
            "无效 URL",
            S("invalid url"),
            args{},
            "",
        },
    }
    for _, tt := range tests {
        t.Run(
            tt.name, func(t *testing.T) {
                if got := tt.r.SanitizeAsHostname(); got != tt.want {
                    t.Errorf("SanitizeAsHostname() = %s, want %s", got, tt.want)
                }
            },
        )
    }
}
func TestRString_SanitizeAsAlphabet(t *testing.T) {
    type args struct{}
    tests := []struct {
        name string
        r    S
        args args
        want S
    }{
        {
            "正常包含字母",
            S("abc123xyz"),
            args{},
            S("abcxyz"),
        }, {
            "正常包含字母",
            S("ab c123x yz"),
            args{},
            S("ab cx yz"),
        },
        {
            "全数字",
            S("123"),
            args{},
            S(""),
        },
        {
            "空字符串",
            S(""),
            args{},
            S(""),
        },
    }
    for _, tt := range tests {
        t.Run(
            tt.name, func(t *testing.T) {
                if got := tt.r.SanitizeAsAlphabet(); got != tt.want {
                    t.Errorf("SanitizeAsAlphabet() = %v, want %v", got, tt.want)
                }
            },
        )
    }
}
func TestRString_SanitizeAsAlphabetWithoutSpace(t *testing.T) {
    type args struct{}
    tests := []struct {
        name string
        r    S
        args args
        want S
    }{
        {
            "正常有字母和其他字符",
            S(" abc 123 xyz "),
            args{},
            S("abcxyz"),
        },
        {
            "全数字",
            S("123"),
            args{},
            S(""),
        },
        {
            "空字符串",
            S(""),
            args{},
            S(""),
        },
    }
    for _, tt := range tests {
        t.Run(
            tt.name, func(t *testing.T) {
                if got := tt.r.SanitizeAsAlphabetWithoutSpace(); got != tt.want {
                    t.Errorf("SanitizeAsAlphabetWithoutSpace() = %v, want %v", got, tt.want)
                }
            },
        )
    }
}
func TestRString_SanitizeAsAlphabetNumber(t *testing.T) {
    type args struct{}
    tests := []struct {
        name string
        r    S
        args args
        want S
    }{
        {
            "包含字母和数字",
            S("a#bc1#23xyz_"),
            args{},
            S("abc123xyz"),
        },
        {
            "只有字母",
            S("abc#xyz#"),
            args{},
            S("abcxyz"),
        },
        {
            "只有数字",
            S("12#3"),
            args{},
            S("123"),
        },
        {
            "空字符串",
            S(""),
            args{},
            S(""),
        },
    }
    for _, tt := range tests {
        t.Run(
            tt.name, func(t *testing.T) {
                if got := tt.r.SanitizeAsAlphabetNumber(); got != tt.want {
                    t.Errorf("SanitizeAsAlphabetNumber() = %v, want %v", got, tt.want)
                }
            },
        )
    }
}
func TestRString_SanitizeAsAlphabetNumberWithoutSpace(t *testing.T) {
    type args struct{}
    tests := []struct {
        name string
        r    S
        args args
        want S
    }{
        {
            "有字母、数字和空格",
            S("a b c# 1 2 3"),
            args{},
            S("abc123"),
        },
        {
            "只有字母和数字",
            S("#abc#123 "),
            args{},
            S("abc123"),
        },
        {
            "空字符串",
            S(""),
            args{},
            S(""),
        },
    }
    for _, tt := range tests {
        t.Run(
            tt.name, func(t *testing.T) {
                if got := tt.r.SanitizeAsAlphabetNumberWithoutSpace(); got != tt.want {
                    t.Errorf("SanitizeAsAlphabetNumberWithoutSpace() = %v, want %v", got, tt.want)
                }
            },
        )
    }
}
func TestRString_StripHtml(t *testing.T) {
    type args struct{}
    tests := []struct {
        name string
        r    S
        args args
        want S
    }{
        {
            "正常包含 html 标签",
            S("<<h1>Hello</h1> World"),
            args{},
            S("Hello World"),
        },
        {
            "全是 html 标签",
            S("<html<body></body></html><script ss>"),
            args{},
            S(""),
        },
        {
            "没有 html 标签",
            S("Just text"),
            args{},
            S("Just text"),
        },
    }
    for _, tt := range tests {
        t.Run(
            tt.name, func(t *testing.T) {
                if got := tt.r.StripHtml(); got != tt.want {
                    t.Errorf("StripHtml() = %v, want %v", got, tt.want)
                }
            },
        )
    }
}

func toSnake(tb testing.TB) {
    cases := [][]string{
        {"testCase", "test_case"},
        {"TestCase", "test_case"},
        {"Test Case", "test_case"},
        {" Test Case", "test_case"},
        {"Test Case ", "test_case"},
        {" Test Case ", "test_case"},
        {"test", "test"},
        {"test_case", "test_case"},
        {"Test", "test"},
        {"", ""},
        {"ManyManyWords", "many_many_words"},
        {"manyManyWords", "many_many_words"},
        {"AnyKind of_string", "any_kind_of_string"},
        {"numbers2and55with000", "numbers_2_and_55_with_000"},
        {"JSONData", "json_data"},
        {"userID", "user_id"},
        {"AAAbbb", "aa_abbb"},
        {"1A2", "1_a_2"},
        {"A1B", "a_1_b"},
        {"A1A2A3", "a_1_a_2_a_3"},
        {"A1 A2 A3", "a_1_a_2_a_3"},
        {"AB1AB2AB3", "ab_1_ab_2_ab_3"},
        {"AB1 AB2 AB3", "ab_1_ab_2_ab_3"},
        {"some string", "some_string"},
        {" some string", "some_string"},
    }
    for _, i := range cases {
        in := i[0]
        out := i[1]
        result := S(in).ToSnake()
        if result != out {
            tb.Errorf("%q (%q != %q)", in, result, out)
        }
    }
}

func TestToSnake(t *testing.T) { toSnake(t) }

func BenchmarkToSnake(b *testing.B) {
    benchmarkSnakeTest(b, toSnake)
}

func toSnakeWithIgnore(tb testing.TB) {
    cases := [][]string{
        {"testCase", "test_case"},
        {"TestCase", "test_case"},
        {"Test Case", "test_case"},
        {" Test Case", "test_case"},
        {"Test Case ", "test_case"},
        {" Test Case ", "test_case"},
        {"test", "test"},
        {"test_case", "test_case"},
        {"Test", "test"},
        {"", ""},
        {"ManyManyWords", "many_many_words"},
        {"manyManyWords", "many_many_words"},
        {"AnyKind of_string", "any_kind_of_string"},
        {"numbers2and55with000", "numbers_2_and_55_with_000"},
        {"JSONData", "json_data"},
        {"AwesomeActivity.UserID", "awesome_activity.user_id", "."},
        {"AwesomeActivity.User.Id", "awesome_activity.user.id", "."},
        {"AwesomeUsername@Awesome.Com", "awesome_username@awesome.com", ".@"},
        {"lets-ignore all.of dots-and-dashes", "lets-ignore_all.of_dots-and-dashes", ".-"},
    }
    for _, i := range cases {
        in := i[0]
        out := i[1]
        var ignore string
        ignore = ""
        if len(i) == 3 {
            ignore = i[2]
        }
        result := S(in).ToSnakeWithIgnore(ignore)
        if result != out {
            istr := ""
            if len(i) == 3 {
                istr = " ignoring '" + i[2] + "'"
            }
            tb.Errorf("%q (%q != %q%s)", in, result, out, istr)
        }
    }
}

func TestToSnakeWithIgnore(t *testing.T) { toSnakeWithIgnore(t) }

func BenchmarkToSnakeWithIgnore(b *testing.B) {
    benchmarkSnakeTest(b, toSnakeWithIgnore)
}

func toDelimited(tb testing.TB) {
    cases := [][]string{
        {"testCase", "test@case"},
        {"TestCase", "test@case"},
        {"Test Case", "test@case"},
        {" Test Case", "test@case"},
        {"Test Case ", "test@case"},
        {" Test Case ", "test@case"},
        {"test", "test"},
        {"test_case", "test@case"},
        {"Test", "test"},
        {"", ""},
        {"ManyManyWords", "many@many@words"},
        {"manyManyWords", "many@many@words"},
        {"AnyKind of_string", "any@kind@of@string"},
        {"numbers2and55with000", "numbers@2@and@55@with@000"},
        {"JSONData", "json@data"},
        {"userID", "user@id"},
        {"AAAbbb", "aa@abbb"},
        {"test-case", "test@case"},
    }
    for _, i := range cases {
        in := i[0]
        out := i[1]
        result := S(in).ToDelimited('@')
        if result != out {
            tb.Errorf("%q (%q != %q)", in, result, out)
        }
    }
}

func TestToDelimited(t *testing.T) { toDelimited(t) }

func BenchmarkToDelimited(b *testing.B) {
    benchmarkSnakeTest(b, toDelimited)
}

func toScreamingSnake(tb testing.TB) {
    cases := [][]string{
        {"testCase", "TEST_CASE"},
    }
    for _, i := range cases {
        in := i[0]
        out := i[1]
        result := S(in).ToScreamingSnake()
        if result != out {
            tb.Errorf("%q (%q != %q)", in, result, out)
        }
    }
}

func TestToScreamingSnake(t *testing.T) { toScreamingSnake(t) }

func BenchmarkToScreamingSnake(b *testing.B) {
    benchmarkSnakeTest(b, toScreamingSnake)
}

func toKebab(tb testing.TB) {
    cases := [][]string{
        {"testCase", "test-case"},
    }
    for _, i := range cases {
        in := i[0]
        out := i[1]
        result := S(in).ToKebab()
        if result != out {
            tb.Errorf("%q (%q != %q)", in, result, out)
        }
    }
}

func TestToKebab(t *testing.T) { toKebab(t) }

func BenchmarkToKebab(b *testing.B) {
    benchmarkSnakeTest(b, toKebab)
}

func toScreamingKebab(tb testing.TB) {
    cases := [][]string{
        {"testCase", "TEST-CASE"},
    }
    for _, i := range cases {
        in := i[0]
        out := i[1]
        result := S(in).ToScreamingKebab()
        if result != out {
            tb.Errorf("%q (%q != %q)", in, result, out)
        }
    }
}

func TestToScreamingKebab(t *testing.T) { toScreamingKebab(t) }

func BenchmarkToScreamingKebab(b *testing.B) {
    benchmarkSnakeTest(b, toScreamingKebab)
}

func toScreamingDelimited(tb testing.TB) {
    cases := [][]string{
        {"testCase", "TEST.CASE"},
    }
    for _, i := range cases {
        in := i[0]
        out := i[1]
        result := S(in).ToScreamingDelimited('.', "", true)
        if result != out {
            tb.Errorf("%q (%q != %q)", in, result, out)
        }
    }
}

func TestToScreamingDelimited(t *testing.T) { toScreamingDelimited(t) }

func BenchmarkToScreamingDelimited(b *testing.B) {
    benchmarkSnakeTest(b, toScreamingDelimited)
}

func toScreamingDelimitedWithIgnore(tb testing.TB) {
    cases := [][]string{
        {"AnyKind of_string", "ANY.KIND OF.STRING", ".", " "},
    }
    for _, i := range cases {
        in := i[0]
        out := i[1]
        delimiter := i[2][0]
        ignore := i[3][0]
        result := S(in).ToScreamingDelimited(delimiter, string(ignore), true)
        if result != out {
            istr := ""
            if len(i) == 4 {
                istr = " ignoring '" + i[3] + "'"
            }
            tb.Errorf("%q (%q != %q%s)", in, result, out, istr)
        }
    }
}

func TestToScreamingDelimitedWithIgnore(t *testing.T) { toScreamingDelimitedWithIgnore(t) }

func BenchmarkToScreamingDelimitedWithIgnore(b *testing.B) {
    benchmarkSnakeTest(b, toScreamingDelimitedWithIgnore)
}

func benchmarkSnakeTest(b *testing.B, fn func(testing.TB)) {
    for n := 0; n < b.N; n++ {
        fn(b)
    }
}

func toCamel(tb testing.TB) {
    cases := [][]string{
        {"test_case", "TestCase"},
        {"test.case", "TestCase"},
        {"test", "Test"},
        {"TestCase", "TestCase"},
        {" test  case ", "TestCase"},
        {"", ""},
        {"many_many_words", "ManyManyWords"},
        {"AnyKind of_string", "AnyKindOfString"},
        {"odd-fix", "OddFix"},
        {"numbers2And55with000", "Numbers2And55With000"},
        {"ID", "Id"},
        {"CONSTANT_CASE", "ConstantCase"},
    }
    for _, i := range cases {
        in := i[0]
        out := i[1]
        result := S(in).ToCamel()
        if result != out {
            tb.Errorf("%q (%q != %q)", in, result, out)
        }
    }
}

func TestToCamel(t *testing.T) {
    toCamel(t)
}

func BenchmarkToCamel(b *testing.B) {
    benchmarkCamelTest(b, toCamel)
}

func toLowerCamel(tb testing.TB) {
    cases := [][]string{
        {"foo-bar", "fooBar"},
        {"TestCase", "testCase"},
        {"", ""},
        {"AnyKind of_string", "anyKindOfString"},
        {"AnyKind.of-string", "anyKindOfString"},
        {"ID", "id"},
        {"some string", "someString"},
        {" some string", "someString"},
        {"CONSTANT_CASE", "constantCase"},
    }
    for _, i := range cases {
        in := i[0]
        out := i[1]
        result := S(in).ToLowerCamel()
        if result != out {
            tb.Errorf("%q (%q != %q)", in, result, out)
        }
    }
}

func TestToLowerCamel(t *testing.T) {
    toLowerCamel(t)
}

func TestCustomAcronymsToCamel(t *testing.T) {
    tests := []struct {
        name         string
        acronymKey   string
        acronymValue string
        expected     string
    }{
        {
            name:         "API Custom Acronym",
            acronymKey:   "API",
            acronymValue: "api",
            expected:     "Api",
        },
        {
            name:         "ABCDACME Custom Acroynm",
            acronymKey:   "ABCDACME",
            acronymValue: "AbcdAcme",
            expected:     "AbcdAcme",
        },
        {
            name:         "PostgreSQL Custom Acronym",
            acronymKey:   "PostgreSQL",
            acronymValue: "PostgreSQL",
            expected:     "PostgreSQL",
        },
    }
    for _, test := range tests {
        t.Run(
            test.name, func(t *testing.T) {
                ConfigureAcronym(test.acronymKey, test.acronymValue)
                if result := S(test.acronymKey).ToCamel(); result != test.expected {
                    t.Errorf("expected custom acronym result %s, got %s", test.expected, result)
                }
            },
        )
    }
}

func TestCustomAcronymsToLowerCamel(t *testing.T) {
    tests := []struct {
        name         string
        acronymKey   string
        acronymValue string
        expected     string
    }{
        {
            name:         "API Custom Acronym",
            acronymKey:   "API",
            acronymValue: "api",
            expected:     "api",
        },
        {
            name:         "ABCDACME Custom Acroynm",
            acronymKey:   "ABCDACME",
            acronymValue: "AbcdAcme",
            expected:     "abcdAcme",
        },
        {
            name:         "PostgreSQL Custom Acronym",
            acronymKey:   "PostgreSQL",
            acronymValue: "PostgreSQL",
            expected:     "postgreSQL",
        },
    }
    for _, test := range tests {
        t.Run(
            test.name, func(t *testing.T) {
                ConfigureAcronym(test.acronymKey, test.acronymValue)
                if result := S(test.acronymKey).ToLowerCamel(); result != test.expected {
                    t.Errorf("expected custom acronym result %s, got %s", test.expected, result)
                }
            },
        )
    }
}

func BenchmarkToLowerCamel(b *testing.B) {
    benchmarkCamelTest(b, toLowerCamel)
}

func benchmarkCamelTest(b *testing.B, fn func(testing.TB)) {
    for n := 0; n < b.N; n++ {
        fn(b)
    }
}

func TestS_SanitizeAsAlphabetNumberDashUnderline(t *testing.T) {
    tests := []struct {
        name string
        r    S
        want S
    }{
        {
            r:    NewS("abbw-_fwg#$o-"),
            want: NewS("abbw-_fwgo-"),
        },
        {
            r:    NewS("abbw-_f21Gwg#$o-"),
            want: NewS("abbw-_f21Gwgo-"),
        },
    }
    for _, tt := range tests {
        t.Run(
            tt.name, func(t *testing.T) {
                if got := tt.r.SanitizeAsAlphabetNumberDashUnderline(); got != tt.want {
                    t.Errorf("SanitizeAsAlphabetNumberDashUnderline() = %v, want %v", got, tt.want)
                }
            },
        )
    }
}

func TestStringGetFirst(t *testing.T) {
	type args struct {
		s   string
		sep string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"正常情况包含分隔符",
			args{"apple,banana,cherry", ","},
			"apple",
		},
		{
			"不包含分隔符",
			args{"orange", ","},
			"orange",
		},
		{
			"空字符串",
			args{"", ","},
			"",
		},
		{
			"分隔符在开头",
			args{",orange", ","},
			"",
		},
		{
			"分隔符在结尾",
			args{"orange,", ","},
			"orange",
		},
		{
			"多个连续分隔符",
			args{"apple,,banana", ","},
			"apple",
		},
		{
			"空分隔符",
			args{"test", ""},
			"test",
		},
		{
			"复杂分隔符路径",
			args{"path/to/file.txt", "/"},
			"path",
		},
		{
			"中文字符串",
			args{"苹果,香蕉,橙子", ","},
			"苹果",
		},
		{
			"特殊字符分隔符",
			args{"a|b|c", "|"},
			"a",
		},
		{
			"空格分隔符",
			args{"hello world", " "},
			"hello",
		},
		{
			"数字字符串",
			args{"123,456,789", ","},
			"123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringGetFirst(tt.args.s, tt.args.sep); got != tt.want {
				t.Errorf("StringGetFirst() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringSanitizeAsInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"空字符串返回0",
			args{""},
			0,
		},
		{
			"纯数字字符串",
			args{"123"},
			123,
		},
		{
			"混合字符字符串",
			args{"abc123def"},
			123,
		},
		{
			"只有非数字字符",
			args{"abc"},
			0,
		},
		{
			"以负号开头的数字",
			args{"-123"},
			-123,
		},
		{
			"负号在中间",
			args{"12-34"},
			1234,
		},
		{
			"负号在结尾",
			args{"123-"},
			123,
		},
		{
			"多个负号",
			args{"-12-34"},
			-1234,
		},
		{
			"数字分散在字符串各处",
			args{"a1b2c3"},
			123,
		},
		{
			"包含空格和制表符",
			args{" 123\t456 "},
			123456,
		},
		{
			"包含标点符号",
			args{"1,234.56"},
			123456,
		},
		{
			"包含中文字符",
			args{"数字123测试"},
			123,
		},
		{
			"包含特殊字符",
			args{"1@2#3$4"},
			1234,
		},
		{
			"零值处理",
			args{"0"},
			0,
		},
		{
			"极大数字",
			args{"999999999"},
			999999999,
		},
		{
			"极小数字",
			args{"-999999999"},
			-999999999,
		},
		{
			"连续数字",
			args{"123456789"},
			123456789,
		},
		{
			"数字间有少量非数字字符",
			args{"1a2b3"},
			123,
		},
		{
			"只有负号",
			args{"-"},
			0,
		},
		{
			"负号后无数字",
			args{"-abc"},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringSanitizeAsInt(tt.args.s); got != tt.want {
				t.Errorf("StringSanitizeAsInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
