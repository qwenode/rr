package rr

import (
    "encoding/base64"
    "fmt"
    "net/url"
    "regexp"
    "strconv"
    "strings"
    "sync"
)

// 字符串转bytes 20250617
func StringToBytes(s string) []byte {
    return []byte(s)
}

// 按分割符切割字符串并获取第一个元素,如果分隔符不存在,则原样返回 20250617
func StringGetFirst(s, sep string) string {
    if sep == "" || !strings.Contains(s, sep) {
        return s
    }
    return strings.Split(s, sep)[0]
}

// 按分割符切割字符串并获取最后一个元素,如果分隔符不存在,则原样返回 20250617
func StringGetLast(s, sep string) string {
    if sep == "" || !strings.Contains(s, sep) {
        return s
    }
    split := strings.Split(s, sep)
    return split[len(split)-1]
}

// 按分割符切割字符串并移除最后一个元素,如果分隔符不存在,则原样返回 20250617
func StringRemoveLast(s, sep string) string {
    if sep == "" || !strings.Contains(s, sep) {
        return s
    }
    split := strings.Split(s, sep)
    return strings.Join(split[:len(split)-1], sep)
}

// 按分割符切割字符串并移除第一个元素,如果分隔符不存在,则原样返回 20250617
func StringRemoveFirst(s, sep string) string {
    if sep == "" || !strings.Contains(s, sep) {
        return s
    }
    split := strings.Split(s, sep)
    return strings.Join(split[1:], sep)
}

// 按分割符切割字符串并获取第二个元素,如果分隔符不存在,则原样返回 20250617
func StringGetSecond(s, sep string) string {
    if sep == "" || !strings.Contains(s, sep) {
        return s
    }
    split := strings.Split(s, sep)
    return split[1]
}

// 检查字符串是否为空 20250617
func StringIsEmpty(s string) bool {
    return s == ""
}

// 获取字符串长度 20250617
func StringLen(s string) int {
    return len(s)
}

// 将字符串转换为URL格式 20250617
func StringAsUrl(s string, https ...bool) string {
    if s == "" {
        return ""
    }
    v := s
    if strings.Contains(v, "://") {
        v = StringGetLast(v, "://")
    }
    if len(https) > 0 && https[0] {
        return "https://" + v
    }
    return "http://" + v
}

// 检查字符串是否为有效URL 20250617
func StringIsUrl(s string) bool {
    if s == "" {
        return false
    }
    x, err := url.Parse(s)
    if err != nil {
        return false
    }
    if x.Host == "" || x.Scheme == "" {
        return false
    }
    return true
}

// URL解码 20250617
func StringUrlDecode(s string) string {
    if s == "" {
        return s
    }
    unescape, _ := url.QueryUnescape(s)
    return unescape
}

// 字符串转float64 20250617
func StringAsFloat(s string) float64 {
    parsed, _ := strconv.ParseFloat(s, 64)
    return parsed
}

// 字符串转int 20250617
func StringAsInt(s string) int {
    atoi, _ := strconv.Atoi(s)
    return atoi
}

// 字符串转int64 20250617
func StringAsInt64(s string) int64 {
    atoi, _ := strconv.ParseInt(s, 10, 64)
    return atoi
}

// 去除字符串首尾空格 20250617
func StringTrimSpace(s string) string {
    return strings.TrimSpace(s)
}

// 去除字符串左侧指定字符 20250617
func StringTrimLeft(s, cut string) string {
    return strings.TrimLeft(s, cut)
}

// 字符串转小写 20250617
func StringLower(s string) string {
    return strings.ToLower(s)
}

// 字符串转大写 20250617
func StringUpper(s string) string {
    return strings.ToUpper(s)
}

// 字符串转bool 20250617
func StringAsBool(s string) bool {
    v := StringTrimSpace(StringLower(s))
    switch v {
    case "1", "t", "true", "ok", "yes", "sure":
        return true
    }
    return false
}

// 字符串安全转换为bool 20250617
func StringSanitizeAsBool(s string) bool {
    return StringAsBool(s)
}

// JSON反序列化 20250617
func StringJsonUnmarshal(s string, v interface{}) error {
    return JsonUnmarshalAdapter([]byte(s), v)
}

// JSON反序列化别名 20250617
func StringJsonUnSerialize(s string, v interface{}) error {
    return StringJsonUnmarshal(s, v)
}

// Base64解码为字符串 20250617
func StringBase64Decode(s string) string {
    return string(StringBase64DecodeAsBytes(s))
}

// Base64解码为字节数组 20250617
func StringBase64DecodeAsBytes(s string) []byte {
    decodeString, _ := base64.StdEncoding.DecodeString(s)
    return decodeString
}

// 去除字符串右侧指定字符 20250617
func StringTrimRight(s, cut string) string {
    return strings.TrimRight(s, cut)
}

// 字符串转uint64 20250617
func StringAsUint64(s string) uint64 {
    if s == "" {
        return 0
    }
    parseint, _ := strconv.ParseUint(s, 10, 64)
    return parseint
}

var sanitizeIntCompile = regexp.MustCompile("[0-9]+")

// 字符串安全转换为int 20250617
func StringSanitizeAsInt(s string) int {
    if s == "" {
        return 0
    }
    v := s
    match := sanitizeIntCompile.FindAllString(v, -1)
    i, _ := strconv.Atoi(strings.Join(match, ""))
    if strings.Index(v, "-") == 0 && i > 0 {
        i *= -1
    }
    return i
}

// 字符串安全转换为int64 20250617
func StringSanitizeAsInt64(s string) int64 {
    if s == "" {
        return 0
    }
    v := s
    match := sanitizeIntCompile.FindAllString(v, -1)
    i, _ := strconv.ParseInt(strings.Join(match, ""), 10, 64)
    if strings.Index(v, "-") == 0 && i > 0 {
        i *= -1
    }
    return i
}

// 字符串连接 20250617
func StringJoin(strs ...string) string {
    if len(strs) > 2 {
        var b strings.Builder
        var l int
        for _, str := range strs {
            l += len(str)
        }
        b.Grow(l)
        for _, str := range strs {
            b.WriteString(str)
        }
        return b.String()
    }
    return strs[0] + strs[1]
}

// 字符串前置拼接 20250617
func StringPrepend(s, prefix string) string {
    return StringJoin(prefix, s)
}

// 字符串后置拼接 20250617
func StringAppend(s, suffix string) string {
    return StringJoin(s, suffix)
}

// 字符串安全转换为主机名 20250617
func StringSanitizeAsHostname(s string) string {
    if s == "" {
        return ""
    }
    v := StringGetLast(s, "://")
    v = StringPrepend(v, "http://")
    parse, err := url.Parse(v)
    if err != nil {
        return ""
    }
    return parse.Hostname()
}

var sanitizeANDS = regexp.MustCompile("[a-zA-Z0-9_-]+")

// 字符串安全转换为字母数字下划线 20250617
func StringSanitizeAsAlphabetNumberDashUnderline(s string) string {
    if s == "" {
        return s
    }
    match := sanitizeANDS.FindAllString(s, -1)
    return strings.Join(match, "")
}

var sanitizeAlphabetCompile = regexp.MustCompile("([a-zA-Z ]+)")

// 字符串安全转换为字母 20250617
func StringSanitizeAsAlphabet(s string) string {
    if s == "" {
        return s
    }
    match := sanitizeAlphabetCompile.FindAllString(s, -1)
    return strings.Join(match, "")
}

var sanitizeAlphabetSpaceCompile = regexp.MustCompile("([a-zA-Z]+)")

// 字符串安全转换为字母(不含空格) 20250617
func StringSanitizeAsAlphabetWithoutSpace(s string) string {
    if s == "" {
        return s
    }
    match := sanitizeAlphabetSpaceCompile.FindAllString(s, -1)
    return strings.Join(match, "")
}

var sanitizeAlphabetNumberCompile = regexp.MustCompile("([a-zA-Z0-9 ]+)")

// 字符串安全转换为字母数字 20250617
func StringSanitizeAsAlphabetNumber(s string) string {
    if s == "" {
        return s
    }
    match := sanitizeAlphabetNumberCompile.FindAllString(s, -1)
    return strings.Join(match, "")
}

var sanitizeAlphabetNumberSpaceCompile = regexp.MustCompile("([a-zA-Z0-9]+)")

// 字符串安全转换为字母数字(不含空格) 20250617
func StringSanitizeAsAlphabetNumberWithoutSpace(s string) string {
    if s == "" {
        return s
    }
    match := sanitizeAlphabetNumberSpaceCompile.FindAllString(s, -1)
    return strings.Join(match, "")
}

// 去除HTML标签 20250617
func StringStripHtml(s string) string {
    data := make([]rune, 0, len(s))
    inside := false
    for _, c := range s {
        if c == '<' {
            inside = true
            continue
        }
        if c == '>' {
            inside = false
            continue
        }
        if !inside {
            data = append(data, c)
        }
    }
    return string(data)
}

// 字符串按行分割 20250617
func StringAsLines(s string) []string {
    if s == "" {
        return make([]string, 0)
    }
    return strings.Split(s, "\n")
}

// 获取文件扩展名 20250617
func StringGetExtension(s string) string {
    if !strings.Contains(s, ".") {
        return ""
    }
    return StringGetFirst(StringGetFirst(StringGetFirst(StringGetLast(s, "."), "/"), "#"), "?")
}

// 移除文件扩展名 20250617
func StringRemoveExtension(s string) string {
    return StringGetFirst(s, ".")
}

// 计算字符串SHA1哈希 20250617
func StringSha1(s string) string {
    if s == "" {
        return ""
    }
    return fmt.Sprintf("%x", B([]byte(s)).Sha1())
}

// 计算字符串SHA256哈希 20250617
func StringSha256(s string) string {
    if s == "" {
        return ""
    }
    return fmt.Sprintf("%x", B([]byte(s)).Sha256())
}

// 计算字符串SHA512哈希 20250617
func StringSha512(s string) string {
    if s == "" {
        return ""
    }
    return fmt.Sprintf("%x", B([]byte(s)).Sha512())
}

// 计算字符串MD5哈希 20250617
func StringMd5(s string) string {
    if s == "" {
        return ""
    }
    return fmt.Sprintf("%x", B([]byte(s)).Md5())
}

// 计算字符串CRC32校验 20250617
func StringCrc32(s string) string {
    if s == "" {
        return ""
    }
    return B([]byte(s)).Crc32()
}

// 限制字符串长度 20250617
func StringLimit(s string, length int) string {
    return StringSubstr(s, 0, length)
}

// 字符串截取 20250617
func StringSubstr(s string, start, length int) string {
    if s == "" {
        return s
    }
    if length < 0 {
        length = 0
    }
    runes := []rune(s)
    sLen := len(runes)
    begin := start
    if start < 0 {
        begin = sLen + start
    }
    if begin < 0 {
        begin = 0
    }

    end := begin + length
    if end > sLen {
        end = sLen
    }
    if begin == 0 && end == sLen {
        return s
    }
    return string(runes[begin:end])
}

// 检查字符串长度是否在指定范围内 20250617
func StringLenBetween(s string, start, end int) bool {
    i := len(s)
    return i >= start && i < end
}

// 转换为snake_case 20250617
func StringToSnake(s string) string {
    return StringToDelimited(s, '_')
}

// 转换为snake_case(忽略指定字符) 20250617
func StringToSnakeWithIgnore(s, ignore string) string {
    return StringToScreamingDelimited(s, '_', ignore, false)
}

// 转换为SCREAMING_SNAKE_CASE 20250617
func StringToScreamingSnake(s string) string {
    return StringToScreamingDelimited(s, '_', "", true)
}

// 转换为kebab-case 20250617
func StringToKebab(s string) string {
    return StringToDelimited(s, '-')
}

// 转换为SCREAMING-KEBAB-CASE 20250617
func StringToScreamingKebab(s string) string {
    return StringToScreamingDelimited(s, '-', "", true)
}

// 转换为分隔符格式 20250617
func StringToDelimited(s string, delimiter uint8) string {
    return StringToScreamingDelimited(s, delimiter, "", false)
}

// 转换为分隔符格式(支持大小写控制) 20250617
func StringToScreamingDelimited(s string, delimiter uint8, ignore string, screaming bool) string {
    s = StringTrimSpace(s)
    n := strings.Builder{}
    n.Grow(len(s) + 2) // nominal 2 bytes of extra space for inserted delimiters
    for i, v := range []byte(s) {
        vIsCap := v >= 'A' && v <= 'Z'
        vIsLow := v >= 'a' && v <= 'z'
        if vIsLow && screaming {
            v += 'A'
            v -= 'a'
        } else if vIsCap && !screaming {
            v += 'a'
            v -= 'A'
        }

        // treat acronyms as words, eg for JSONData -> JSON is a whole word
        if i+1 < len(s) {
            next := s[i+1]
            vIsNum := v >= '0' && v <= '9'
            nextIsCap := next >= 'A' && next <= 'Z'
            nextIsLow := next >= 'a' && next <= 'z'
            nextIsNum := next >= '0' && next <= '9'
            // add underscore if next letter case type is changed
            if (vIsCap && (nextIsLow || nextIsNum)) || (vIsLow && (nextIsCap || nextIsNum)) || (vIsNum && (nextIsCap || nextIsLow)) {
                prevIgnore := ignore != "" && i > 0 && strings.ContainsAny(string(s[i-1]), ignore)
                if !prevIgnore {
                    if vIsCap && nextIsLow {
                        if prevIsCap := i > 0 && s[i-1] >= 'A' && s[i-1] <= 'Z'; prevIsCap {
                            n.WriteByte(delimiter)
                        }
                    }
                    n.WriteByte(v)
                    if vIsLow || vIsNum || nextIsNum {
                        n.WriteByte(delimiter)
                    }
                    continue
                }
            }
        }

        if (v == ' ' || v == '_' || v == '-' || v == '.') && !strings.ContainsAny(string(v), ignore) {
            // replace space/underscore/hyphen/dot with delimiter
            n.WriteByte(delimiter)
        } else {
            n.WriteByte(v)
        }
    }

    return n.String()
}

// 转换为CamelCase 20250617
func StringToCamel(s string) string {
    return stringToCamelInitCase(s, true)
}

// 转换为lowerCamelCase 20250617
func StringToLowerCamel(s string) string {
    return stringToCamelInitCase(s, false)
}

// 检查字符串是否包含指定子串 20250617
func StringContains(s, v string) bool {
    return strings.Contains(s, v)
}

// 转换为CamelCase的内部函数 20250617
func stringToCamelInitCase(s string, initCase bool) string {
    s = StringTrimSpace(s)
    if s == "" {
        return s
    }
    a, hasAcronym := uppercaseAcronym.Load(s)
    if hasAcronym {
        s = a.(string)
    }

    n := strings.Builder{}
    n.Grow(len(s))
    capNext := initCase
    prevIsCap := false
    for i, v := range []byte(s) {
        vIsCap := v >= 'A' && v <= 'Z'
        vIsLow := v >= 'a' && v <= 'z'
        if capNext {
            if vIsLow {
                v += 'A'
                v -= 'a'
            }
        } else if i == 0 {
            if vIsCap {
                v += 'a'
                v -= 'A'
            }
        } else if prevIsCap && vIsCap && !hasAcronym {
            v += 'a'
            v -= 'A'
        }
        prevIsCap = vIsCap

        if vIsCap || vIsLow {
            n.WriteByte(v)
            capNext = false
        } else if vIsNum := v >= '0' && v <= '9'; vIsNum {
            n.WriteByte(v)
            capNext = true
        } else {
            capNext = v == '_' || v == ' ' || v == '-' || v == '.'
        }
    }
    return n.String()
}

type S string

func NewS(v string) S {
    return S(v)
}
func (r S) String() string {
    return string(r)
}
func (r S) Bytes() []byte {
    return []byte(r)
}
func (r S) GetFirst(sep string) S {
    v := r.String()
    if !strings.Contains(v, sep) {
        return r
    }
    return S(strings.Split(v, sep)[0])
}
func (r S) GetLast(sep string) S {
    v := r.String()
    if !strings.Contains(v, sep) {
        return r
    }
    split := strings.Split(v, sep)
    return S(split[len(split)-1])
}
func (r S) RemoveLast(sep string) S {
    v := r.String()
    if !strings.Contains(v, sep) {
        return r
    }
    split := strings.Split(v, sep)
    return S(strings.Join(split[:len(split)-1], sep))
}
func (r S) RemoveFirst(sep string) S {
    v := r.String()
    if !strings.Contains(v, sep) {
        return r
    }
    split := strings.Split(v, sep)
    return S(strings.Join(split[1:], sep))
}
func (r S) GetSecond(sep string) S {
    v := r.String()
    if !strings.Contains(v, sep) {
        return r
    }
    split := strings.Split(v, sep)
    return S(split[1])
}
func (r S) IsEmpty() bool {
    return r == ""
}
func (r S) Len() int {
    return len(r)
}
func (r S) AsUrl(https ...bool) string {
    if r == "" {
        return ""
    }
    v := r.String()
    if strings.Contains(v, "://") {
        v = r.GetLast("://").String()
    }
    if len(https) > 0 && https[0] {
        return "https://" + v
    }
    return "http://" + v
}
func (r S) IsUrl() bool {
    if r == "" {
        return false
    }
    x, err := url.Parse(r.String())
    if err != nil {
        return false
    }
    if x.Host == "" || x.Scheme == "" {
        return false
    }
    return true
}
func (r S) UrlDecode() S {
    if r == "" {
        return r
    }
    unescape, _ := url.QueryUnescape(r.String())
    return S(unescape)
}
func (r S) AsString() string {
    return string(r)
}
func (r S) AsFloat() float64 {
    parsed, _ := strconv.ParseFloat(r.AsString(), 64)
    return parsed
}
func (r S) AsInt() int {
    atoi, _ := strconv.Atoi(string(r))
    return atoi
}
func (r S) AsInt64() int64 {
    atoi, _ := strconv.ParseInt(string(r), 10, 64)
    return atoi
}
func (r S) TrimSpace() S {
    return S(strings.TrimSpace(r.String()))
}
func (r S) TrimLeft(cut string) S {
    return S(strings.TrimLeft(r.String(), cut))
}
func (r S) Lower() S {
    return S(strings.ToLower(r.String()))
}
func (r S) Upper() S {
    return S(strings.ToUpper(r.String()))
}
func (r S) AsBool() bool {
    v := r.TrimSpace().Lower().String()
    switch v {
    case "1", "t", "true", "ok", "yes", "sure":
        return true
    }
    return false
}
func (r S) SanitizeAsBool() bool {
    return r.AsBool()
}
func (r S) JsonUnmarshal(v interface{}) error {
    return JsonUnmarshalAdapter([]byte(r), v)
}
func (r S) JsonUnSerialize(v interface{}) error {
    return r.JsonUnmarshal(v)
}
func (r S) Base64Decode() S {
    return S(r.Base64DecodeAsBytes())
}
func (r S) Base64DecodeAsBytes() []byte {
    decodeString, _ := base64.StdEncoding.DecodeString(r.String())
    return decodeString
}

func (r S) TrimRight(cut string) S {
    return S(strings.TrimRight(r.String(), cut))
}

func (r S) AsUint64() uint64 {
    if r == "" {
        return 0
    }
    parseint, _ := strconv.ParseUint(string(r), 10, 64)
    return parseint
}

func (r S) Prepend(s string) S {
    return S(StringJoin(s, r.String()))
}
func (r S) Append(s string) S {
    return S(StringJoin(r.String(), s))
}

func (r S) SanitizeAsHostname() string {
    if r == "" {
        return ""
    }
    v := r.GetLast("://").Prepend("http://").String()
    parse, err := url.Parse(v)
    if err != nil {
        return ""
    }
    return parse.Hostname()
}

func (r S) SanitizeAsAlphabetNumberDashUnderline() S {
    if r == "" {
        return r
    }
    match := sanitizeANDS.FindAllString(r.String(), -1)
    s := strings.Join(match, "")
    // s = strings.Join(strings.Fields(s), " ")
    return S(s)
}

func (r S) SanitizeAsAlphabet() S {
    if r == "" {
        return r
    }
    match := sanitizeAlphabetCompile.FindAllString(r.String(), -1)
    s := strings.Join(match, "")
    // s = strings.Join(strings.Fields(s), " ")
    return S(s)
}

func (r S) SanitizeAsAlphabetWithoutSpace() S {
    if r == "" {
        return r
    }
    match := sanitizeAlphabetSpaceCompile.FindAllString(r.String(), -1)
    s := strings.Join(match, "")
    return S(s)
}

func (r S) SanitizeAsAlphabetNumber() S {
    if r == "" {
        return r
    }
    match := sanitizeAlphabetNumberCompile.FindAllString(r.String(), -1)
    s := strings.Join(match, "")
    return S(s)
}

func (r S) SanitizeAsAlphabetNumberWithoutSpace() S {
    if r == "" {
        return r
    }
    match := sanitizeAlphabetNumberSpaceCompile.FindAllString(r.String(), -1)
    s := strings.Join(match, "")
    return S(s)
}
func (r S) StripHtml() S {
    data := make([]rune, 0, len(r))
    inside := false
    for _, c := range r {
        if c == '<' {
            inside = true
            continue
        }
        if c == '>' {
            inside = false
            continue
        }
        if !inside {
            data = append(data, c)
        }
    }
    return S(data)
}
func (r S) AsLines() []string {
    if r == "" {
        return make([]string, 0)
    }
    return strings.Split(r.String(), "\n")
}

func (r S) GetExtension() S {
    s := r.String()
    if !strings.Contains(s, ".") {
        return ""
    }

    return r.GetLast(".").GetFirst("/").GetFirst("#").GetFirst("?")
}
func (r S) AsF() F {
    return F(r)
}
func (r S) RemoveExtension() S {
    return r.GetFirst(".")
}

func (r S) Sha1() string {
    if r == "" {
        return ""
    }
    return fmt.Sprintf("%x", B(r).Sha1())
}
func (r S) Sha256() string {
    if r == "" {
        return ""
    }
    return fmt.Sprintf("%x", B(r).Sha256())
}
func (r S) Sha512() string {
    if r == "" {
        return ""
    }
    return fmt.Sprintf("%x", B(r).Sha512())
}
func (r S) Md5() string {
    if r == "" {
        return ""
    }
    return fmt.Sprintf("%x", B(r).Md5())
}
func (r S) Crc32() string {
    if r == "" {
        return ""
    }
    return B(r).Crc32()
}
func (r S) Limit(length int) S {
    return r.Substr(0, length)
}
func (r S) Substr(start, length int) S {
    if r == "" {
        return r
    }
    if length < 0 {
        length = 0
    }
    s := []rune(r.String())
    sLen := len(s)
    begin := start
    if start < 0 {
        begin = sLen + start
    }
    if begin < 0 {
        begin = 0
    }

    end := begin + length
    if end > sLen {
        end = sLen
    }
    if begin == 0 && end == sLen {
        return r
    }
    return S(s[begin:end])
}
func (r S) LenBetween(start, end int) bool {
    i := len(r)
    return i >= start && i < end
}

// ToSnake converts a string to snake_case
func (r S) ToSnake() string {
    return r.ToDelimited('_')
}

func (r S) ToSnakeWithIgnore(ignore string) string {
    return r.ToScreamingDelimited('_', ignore, false)
}

// ToScreamingSnake converts a string to SCREAMING_SNAKE_CASE
func (r S) ToScreamingSnake() string {
    return r.ToScreamingDelimited('_', "", true)
}

// ToKebab converts a string to kebab-case
func (r S) ToKebab() string {
    return r.ToDelimited('-')
}

// ToScreamingKebab converts a string to SCREAMING-KEBAB-CASE
func (r S) ToScreamingKebab() string {
    return r.ToScreamingDelimited('-', "", true)
}

// ToDelimited converts a string to delimited.snake.case
// (in this case `delimiter = '.'`)
func (r S) ToDelimited(delimiter uint8) string {
    return r.ToScreamingDelimited(delimiter, "", false)
}

// ToScreamingDelimited converts a string to SCREAMING.DELIMITED.SNAKE.CASE
// (in this case `delimiter = '.'; screaming = true`)
// or delimited.snake.case
// (in this case `delimiter = '.'; screaming = false`)
func (r S) ToScreamingDelimited(delimiter uint8, ignore string, screaming bool) string {
    s := r.TrimSpace().String()
    n := strings.Builder{}
    n.Grow(len(s) + 2) // nominal 2 bytes of extra space for inserted delimiters
    for i, v := range []byte(s) {
        vIsCap := v >= 'A' && v <= 'Z'
        vIsLow := v >= 'a' && v <= 'z'
        if vIsLow && screaming {
            v += 'A'
            v -= 'a'
        } else if vIsCap && !screaming {
            v += 'a'
            v -= 'A'
        }

        // treat acronyms as words, eg for JSONData -> JSON is a whole word
        if i+1 < len(s) {
            next := s[i+1]
            vIsNum := v >= '0' && v <= '9'
            nextIsCap := next >= 'A' && next <= 'Z'
            nextIsLow := next >= 'a' && next <= 'z'
            nextIsNum := next >= '0' && next <= '9'
            // add underscore if next letter case type is changed
            if (vIsCap && (nextIsLow || nextIsNum)) || (vIsLow && (nextIsCap || nextIsNum)) || (vIsNum && (nextIsCap || nextIsLow)) {
                prevIgnore := ignore != "" && i > 0 && strings.ContainsAny(string(s[i-1]), ignore)
                if !prevIgnore {
                    if vIsCap && nextIsLow {
                        if prevIsCap := i > 0 && s[i-1] >= 'A' && s[i-1] <= 'Z'; prevIsCap {
                            n.WriteByte(delimiter)
                        }
                    }
                    n.WriteByte(v)
                    if vIsLow || vIsNum || nextIsNum {
                        n.WriteByte(delimiter)
                    }
                    continue
                }
            }
        }

        if (v == ' ' || v == '_' || v == '-' || v == '.') && !strings.ContainsAny(string(v), ignore) {
            // replace space/underscore/hyphen/dot with delimiter
            n.WriteByte(delimiter)
        } else {
            n.WriteByte(v)
        }
    }

    return n.String()
}

var uppercaseAcronym = sync.Map{}

// "ID": "id",

// ConfigureAcronym allows you to add additional words which will be considered acronyms
func ConfigureAcronym(key, val string) {
    uppercaseAcronym.Store(key, val)
}

// Converts a string to CamelCase
func (r S) toCamelInitCase(initCase bool) string {
    s := r.TrimSpace().String()
    if s == "" {
        return s
    }
    a, hasAcronym := uppercaseAcronym.Load(s)
    if hasAcronym {
        s = a.(string)
    }

    n := strings.Builder{}
    n.Grow(len(s))
    capNext := initCase
    prevIsCap := false
    for i, v := range []byte(s) {
        vIsCap := v >= 'A' && v <= 'Z'
        vIsLow := v >= 'a' && v <= 'z'
        if capNext {
            if vIsLow {
                v += 'A'
                v -= 'a'
            }
        } else if i == 0 {
            if vIsCap {
                v += 'a'
                v -= 'A'
            }
        } else if prevIsCap && vIsCap && !hasAcronym {
            v += 'a'
            v -= 'A'
        }
        prevIsCap = vIsCap

        if vIsCap || vIsLow {
            n.WriteByte(v)
            capNext = false
        } else if vIsNum := v >= '0' && v <= '9'; vIsNum {
            n.WriteByte(v)
            capNext = true
        } else {
            capNext = v == '_' || v == ' ' || v == '-' || v == '.'
        }
    }
    return n.String()
}

// ToCamel converts a string to CamelCase
func (r S) ToCamel() string {
    return r.toCamelInitCase(true)
}

// ToLowerCamel converts a string to lowerCamelCase
func (r S) ToLowerCamel() string {
    return r.toCamelInitCase(false)
}

func (r S) Contains(v string) bool {
    return strings.Contains(r.String(), v)
}

var sanitizeNumberSpaceCompile = regexp.MustCompile("([0-9]+)")

func (r S) SanitizeAsNumberWithoutSpace() S {
    if r == "" {
        return r
    }
    match := sanitizeNumberSpaceCompile.FindAllString(r.String(), -1)
    s := strings.Join(match, "")
    return S(s)
}
