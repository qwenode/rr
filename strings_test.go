package ff

import (
    "bytes"
    "errors"
    "testing"
)

func TestRString_JsonUnmarshal(t *testing.T) {
    type args struct {
        v interface{}
    }
    tests := []struct {
        name    string
        r       RString
        args    args
        wantErr bool
    }{
        {
            "正常情况有效 JSON 字符串和正确类型接收对象",
            RString(`{"key": "value"}`),
            args{&struct{ Key string }{Key: ""}},
            false,
        },
        {
            "边界情况空字符串",
            RString(""),
            args{&struct{}{}},
            true,
        },
        {
            "错误输入无效 JSON 字符串",
            RString("invalid_json"),
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
        r    RString
        args args
        want RString
    }{
        {
            "正常情况非空字符串",
            RString("encoded_value"),
            args{},
            RString("encoded_value"), // 假设实际解码后的结果
        },
        {
            "边界情况空字符串",
            RString(""),
            args{},
            RString(""),
        },
        {
            "逻辑检查特殊编码字符串",
            RString("%2B"),
            args{},
            RString("+"),
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
        r     RString
        https []bool
    }
    tests := []struct {
        name string
        args args
        want string
    }{
        {
            "有内容且不指定 https 为 http 格式",
            args{RString("example.com"), []bool{}},
            "http://example.com",
        },
        {
            "有内容且指定 https 为 https 格式",
            args{RString("example.com"), []bool{true}},
            "https://example.com",
        },
        {
            "已有协议的内容且不指定 https 保留原协议",
            args{RString("http://example.org"), []bool{}},
            "http://example.org",
        },
        {
            "空字符串",
            args{RString(""), []bool{}},
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
        r RString
    }
    tests := []struct {
        name string
        args args
        want bool
    }{
        {
            "有效 URL",
            args{RString("https://example.com")},
            true,
        },
        {
            "另一个有效 URL",
            args{RString("http://example.org")},
            true,
        },
        {
            "无效 URL",
            args{RString("invalid-url")},
            false,
        },
        {
            "空字符串",
            args{RString("")},
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
    r := RString("apple,banana,cherry")
    expected := RString("apple")
    result := r.GetFirst(",")
    if result != expected {
        t.Errorf("Expected %s, got %s", expected, result)
    }
    
    // 测试不包含分隔符的情况
    r = RString("orange")
    expected = RString("orange")
    result = r.GetFirst(",")
    if result != expected {
        t.Errorf("Expected %s, got %s", expected, result)
    }
    r = RString(",orange")
    expected = RString("")
    result = r.GetFirst(",")
    if result != expected {
        t.Errorf("Expected %s, got %s", expected, result)
    }
}
func TestRemoveLast(t *testing.T) {
    // 正常情况
    input1 := RString("Hello, World!")
    sep1 := ", "
    expected1 := RString("Hello")
    output1 := input1.RemoveLast(sep1)
    if output1 != expected1 {
        t.Errorf("Expected %v but got %v", expected1, output1)
    }
    
    input2 := RString("Hello, World!")
    sep2 := "!"
    expected2 := RString("Hello, World")
    output2 := input2.RemoveLast(sep2)
    if output2 != expected2 {
        t.Errorf("Expected %v but got %v", expected2, output2)
    }
    
    // 边界情况
    input3 := RString("Hello, World!")
    sep3 := ":"
    expected3 := input3
    output3 := input3.RemoveLast(sep3)
    if output3 != expected3 {
        t.Errorf("Expected %v but got %v", expected3, output3)
    }
    
    input4 := RString("")
    sep4 := ","
    expected4 := input4
    output4 := input4.RemoveLast(sep4)
    if output4 != expected4 {
        t.Errorf("Expected %v but got %v", expected4, output4)
    }
    
    input5 := RString("Hello, World!")
    sep5 := ","
    expected5 := RString("Hello")
    output5 := input5.RemoveLast(sep5)
    if output5 != expected5 {
        t.Errorf("Expected %v but got %v", expected5, output5)
    }
    
    // 错误情况
    input6 := RString("Hello, World!")
    sep6 := "W"
    expected6 := RString("Hello, ")
    output6 := input6.RemoveLast(sep6)
    if output6 != expected6 {
        t.Errorf("Expected %v but got %v", expected6, output6)
    }
}
func TestGetLast(t *testing.T) {
    // 测试包含分隔符的情况
    r := RString("apple,banana,cherry")
    expected := RString("cherry")
    result := r.GetLast(",")
    if result != expected {
        t.Errorf("Expected %s, got %s", expected, result)
    }
    
    // 测试不包含分隔符的情况
    r = RString("lemon")
    expected = RString("lemon")
    result = r.GetLast(",")
    if result != expected {
        t.Errorf("Expected %s, got %s", expected, result)
    }
    // 
    r = RString("lemon,")
    expected = RString("")
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
        r    RString
        args args
        want RString
    }{
        {
            "正常情况",
            RString("abcdefg"),
            args{"cd"},
            RString("efg"),
        },
        {
            "边界情况，开头就是分隔符",
            RString("cdabc"),
            args{"cd"},
            RString("abc"),
        },
        {
            "不存在分隔符的情况",
            RString("abc"),
            args{"cd"},
            RString("abc"),
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
        r    RString
        args args
        want RString
    }{
        {
            "正常情况有分隔符",
            RString("abcdefg"),
            args{"cd"},
            RString("efg"),
        },
        {
            "边界情况开头就是分隔符",
            RString("cdabc"),
            args{"cd"},
            RString("abc"),
        },
        {
            "不存在分隔符",
            RString("abc"),
            args{"cd"},
            RString("abc"),
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
        r    RString
        args args
        want []byte
    }{
        {
            "正常情况",
            RString("SGVsbG8="), // 对应 "Hello" 的 base64 编码
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
        r    RString
        args args
        want float64
    }{
        {
            "正常转换",
            RString("123.45"),
            args{},
            123.45,
        },
        {
            "非数字字符串",
            RString("abc"),
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
        r    RString
        args args
        want int
    }{
        {
            "正常转换数字",
            RString("123"),
            args{},
            123,
        },
        {
            "非数字字符串",
            RString("abc"),
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
        r    RString
        args args
        want int64
    }{
        {
            "正常数字转换",
            RString("123456"),
            args{},
            123456,
        },
        {
            "非数字字符串",
            RString("abc"),
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
        r    RString
        args args
        want RString
    }{
        {
            "正常去空格",
            RString("  abc   "),
            args{},
            RString("abc"),
        },
        {
            "本身无空格",
            RString("abc"),
            args{},
            RString("abc"),
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
        r    RString
        args args
        want error
    }{
        {
            "正常解析结构体",
            RString(`{"name":"Alice","age":25}`),
            args{&struct {
                Name string
                Age  int
            }{}},
            nil,
        },
        {
            "无效 JSON 字符串",
            RString("invalid json"),
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
        r    RString
        args args
        want bool
    }{
        {
            "正常转换为 true - '1'",
            RString("1"),
            args{},
            true,
        },
        {
            "正常转换为 true - 't'",
            RString("t"),
            args{},
            true,
        },
        {
            "正常转换为 true - 'true'",
            RString("true"),
            args{},
            true,
        },
        {
            "正常转换为 true - 'ok'",
            RString("ok"),
            args{},
            true,
        },
        {
            "正常转换为 true - 'yes'",
            RString("yes"),
            args{},
            true,
        },
        {
            "正常转换为 true -'sure'",
            RString("sure"),
            args{},
            true,
        },
        {
            "不转换为 true - 'abc'",
            RString("abc"),
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
        r    RString
        args args
        want RString
    }{
        {
            "正常转换为大写",
            RString("abc"),
            args{},
            RString("ABC"),
        },
        {
            "本身就是大写",
            RString("ABC"),
            args{},
            RString("ABC"),
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
        r    RString
        args args
        want RString
    }{
        {
            "正常转换为小写",
            RString("ABC"),
            args{},
            RString("abc"),
        },
        {
            "本身就是小写",
            RString("abc"),
            args{},
            RString("abc"),
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
        r    RString
        args args
        want RString
    }{
        {
            "正常裁剪左边",
            RString("abcdefg"),
            args{"abc"},
            RString("defg"),
        },
        {
            "裁剪部分不匹配",
            RString("abcdefg"),
            args{"xyz"},
            RString("abcdefg"),
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
        r    RString
        args args
        want RString
    }{
        {
            "正常裁剪右边",
            RString("abcdefg"),
            args{"efg"},
            RString("abcd"),
        },
        {
            "裁剪部分不匹配",
            RString("abcdefg"),
            args{"xyz"},
            RString("abcdefg"),
        },
        {
            "裁剪部分不匹配",
            RString("abcdefg"),
            args{"cde"},
            RString("abcdefg"),
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
        r    RString
        args args
        want uint64
    }{
        {
            "正常转换",
            RString("123"),
            args{},
            123,
        },
        {
            "转换失败",
            RString("abc"),
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
        r    RString
        args args
        want int
    }{
        {
            "空字符串返回 0",
            RString(""),
            args{},
            0,
        },
        {
            "正常数字转换",
            RString("1x2#3"),
            args{},
            123,
        },
        {
            "有非数字字符但能提取出数字",
            RString("abc12_-3"),
            args{},
            123,
        },
        {
            "带负号且转换正确",
            RString("-123"),
            args{},
            -123,
        },
    }
    for _, tt := range tests {
        t.Run(
            tt.name, func(t *testing.T) {
                if got := tt.r.SanitizeAsInt(); got != tt.want {
                    t.Errorf("SanitizeAsInt() = %v, want %v", got, tt.want)
                }
            },
        )
    }
}
func TestRString_SanitizeAsInt64(t *testing.T) {
    type args struct{}
    tests := []struct {
        name string
        r    RString
        args args
        want int64
    }{
        {
            "空字符串返回 0",
            RString(""),
            args{},
            0,
        },
        {
            "正常数字转换",
            RString("1)23"),
            args{},
            123,
        },
        {
            "有非数字字符但能提取出数字",
            RString("abc12#3"),
            args{},
            123,
        },
        {
            "带负号且转换正确",
            RString("-12#3"),
            args{},
            -123,
        },
    }
    for _, tt := range tests {
        t.Run(
            tt.name, func(t *testing.T) {
                if got := tt.r.SanitizeAsInt64(); got != tt.want {
                    t.Errorf("SanitizeAsInt64() = %v, want %v", got, tt.want)
                }
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
        r    RString
        args args
        want RString
    }{
        {
            "正常添加前缀",
            RString("abc"),
            args{"xyz"},
            RString("xyzabc"),
        },
        {
            "空字符串添加前缀",
            RString("def"),
            args{""},
            RString("def"),
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
        r    RString
        args args
        want RString
    }{
        {
            "正常添加后缀",
            RString("abc"),
            args{"xyz"},
            RString("abcxyz"),
        },
        {
            "空字符串添加后缀",
            RString("def"),
            args{""},
            RString("def"),
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
        r    RString
        args args
        want string
    }{
        {
            "正常有协议的 URL",
            RString("https://example.com"),
            args{},
            "example.com",
        },
        {
            "无协议添加 http:// 后提取",
            RString("example.com"),
            args{},
            "example.com",
        },
        {
            "无协议添加 http:// 后提取",
            RString("127.0.0.1"),
            args{},
            "127.0.0.1",
        },
        {
            "无协议添加 http:// 后提取",
            RString("127.0.0.1:222/wowig"),
            args{},
            "127.0.0.1",
        },
        {
            "无效 URL",
            RString("invalid url"),
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
        r    RString
        args args
        want RString
    }{
        {
            "正常包含字母",
            RString("abc123xyz"),
            args{},
            RString("abcxyz"),
        }, {
            "正常包含字母",
            RString("ab c123x yz"),
            args{},
            RString("ab cx yz"),
        },
        {
            "全数字",
            RString("123"),
            args{},
            RString(""),
        },
        {
            "空字符串",
            RString(""),
            args{},
            RString(""),
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
        r    RString
        args args
        want RString
    }{
        {
            "正常有字母和其他字符",
            RString(" abc 123 xyz "),
            args{},
            RString("abcxyz"),
        },
        {
            "全数字",
            RString("123"),
            args{},
            RString(""),
        },
        {
            "空字符串",
            RString(""),
            args{},
            RString(""),
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
        r    RString
        args args
        want RString
    }{
        {
            "包含字母和数字",
            RString("a#bc1#23xyz_"),
            args{},
            RString("abc123xyz"),
        },
        {
            "只有字母",
            RString("abc#xyz#"),
            args{},
            RString("abcxyz"),
        },
        {
            "只有数字",
            RString("12#3"),
            args{},
            RString("123"),
        },
        {
            "空字符串",
            RString(""),
            args{},
            RString(""),
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
        r    RString
        args args
        want RString
    }{
        {
            "有字母、数字和空格",
            RString("a b c# 1 2 3"),
            args{},
            RString("abc123"),
        },
        {
            "只有字母和数字",
            RString("#abc#123 "),
            args{},
            RString("abc123"),
        },
        {
            "空字符串",
            RString(""),
            args{},
            RString(""),
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
        r    RString
        args args
        want RString
    }{
        {
            "正常包含 html 标签",
            RString("<<h1>Hello</h1> World"),
            args{},
            RString("Hello World"),
        },
        {
            "全是 html 标签",
            RString("<html<body></body></html><script ss>"),
            args{},
            RString(""),
        },
        {
            "没有 html 标签",
            RString("Just text"),
            args{},
            RString("Just text"),
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
