package rr

import (
    "encoding/base64"
    "encoding/json"
    "net/url"
    "regexp"
    "strconv"
    "strings"
)

var (
    JsonMarshalAdapter   = json.Marshal
    JsonUnmarshalAdapter = json.Unmarshal
)

type RString string

func (r RString) String() string {
    return string(r)
}
func (r RString) GetFirst(sep string) RString {
    v := r.String()
    if !strings.Contains(v, sep) {
        return r
    }
    return RString(strings.Split(v, sep)[0])
}
func (r RString) GetLast(sep string) RString {
    v := r.String()
    if !strings.Contains(v, sep) {
        return r
    }
    split := strings.Split(v, sep)
    return RString(split[len(split)-1])
}
func (r RString) RemoveLast(sep string) RString {
    v := r.String()
    if !strings.Contains(v, sep) {
        return r
    }
    split := strings.Split(v, sep)
    return RString(strings.Join(split[:len(split)-1], sep))
}
func (r RString) RemoveFirst(sep string) RString {
    v := r.String()
    if !strings.Contains(v, sep) {
        return r
    }
    split := strings.Split(v, sep)
    return RString(strings.Join(split[1:], sep))
}
func (r RString) GetSecond(sep string) RString {
    v := r.String()
    if !strings.Contains(v, sep) {
        return r
    }
    split := strings.Split(v, sep)
    return RString(split[1])
}
func (r RString) AsUrl(https ...bool) string {
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
func (r RString) IsUrl() bool {
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
func (r RString) UrlDecode() RString {
    if r == "" {
        return r
    }
    unescape, _ := url.QueryUnescape(r.String())
    return RString(unescape)
}
func (r RString) AsString() string {
    return string(r)
}
func (r RString) AsFloat() float64 {
    parsed, _ := strconv.ParseFloat(r.AsString(), 64)
    return parsed
}
func (r RString) AsInt() int {
    atoi, _ := strconv.Atoi(string(r))
    return atoi
}
func (r RString) AsInt64() int64 {
    atoi, _ := strconv.ParseInt(string(r), 10, 64)
    return atoi
}
func (r RString) TrimSpace() RString {
    return RString(strings.TrimSpace(r.String()))
}
func (r RString) TrimLeft(cut string) RString {
    return RString(strings.TrimLeft(r.String(), cut))
}
func (r RString) Lower() RString {
    return RString(strings.ToLower(r.String()))
}
func (r RString) Upper() RString {
    return RString(strings.ToUpper(r.String()))
}
func (r RString) AsBool() bool {
    v := r.TrimSpace().Lower().String()
    switch v {
    case "1", "t", "true", "ok", "yes", "sure":
        return true
    }
    return false
}
func (r RString) JsonUnmarshal(v interface{}) error {
    return JsonUnmarshalAdapter([]byte(r), v)
}
func (r RString) JsonUnSerialize(v interface{}) error {
    return r.JsonUnmarshal(v)
}
func (r RString) Base64Decode() RString {
    return RString(r.Base64DecodeAsBytes())
}
func (r RString) Base64DecodeAsBytes() []byte {
    decodeString, _ := base64.StdEncoding.DecodeString(r.String())
    return decodeString
}

func JsonSerialize(v interface{}) string {
    return JsonMarshal(v)
}
func JsonMarshal(v interface{}) string {
    return string(JsonMarshalAsBytes(v))
}
func JsonMarshalAsBytes(v interface{}) []byte {
    adapter, _ := JsonMarshalAdapter(v)
    return adapter
}
func (r RString) TrimRight(cut string) RString {
    return RString(strings.TrimRight(r.String(), cut))
}

func (r RString) AsUint64() uint64 {
    if r == "" {
        return 0
    }
    parseint, _ := strconv.ParseUint(string(r), 10, 64)
    return parseint
}

var sanitizeIntCompile = regexp.MustCompile("[0-9]+")

func (r RString) SanitizeAsInt() int {
    if r == "" {
        return 0
    }
    v := r.String()
    match := sanitizeIntCompile.FindAllString(v, -1)
    i, _ := strconv.Atoi(strings.Join(match, ""))
    if strings.Index(v, "-") == 0 && i > 0 {
        i *= -1
    }
    return i
}
func (r RString) SanitizeAsInt64() int64 {
    if r == "" {
        return 0
    }
    v := r.String()
    match := sanitizeIntCompile.FindAllString(v, -1)
    i, _ := strconv.ParseInt(strings.Join(match, ""), 10, 64)
    if strings.Index(v, "-") == 0 && i > 0 {
        i *= -1
    }
    return i
}
func StringJoin(strs ...string) RString {
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
        return RString(b.String())
    }
    return RString(strs[0] + strs[1])
}
func (r RString) Prepend(s string) RString {
    return StringJoin(s, r.String())
}
func (r RString) Append(s string) RString {
    return StringJoin(r.String(), s)
}
func (r RString) SanitizeAsHostname() string {
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

var sanitizeAlphabetCompile = regexp.MustCompile("([a-zA-Z ]+)")

func (r RString) SanitizeAsAlphabet() RString {
    if r == "" {
        return r
    }
    match := sanitizeAlphabetCompile.FindAllString(r.String(), -1)
    s := strings.Join(match, "")
    // s = strings.Join(strings.Fields(s), " ")
    return RString(s)
}

var sanitizeAlphabetSpaceCompile = regexp.MustCompile("([a-zA-Z]+)")

func (r RString) SanitizeAsAlphabetWithoutSpace() RString {
    if r == "" {
        return r
    }
    match := sanitizeAlphabetSpaceCompile.FindAllString(r.String(), -1)
    s := strings.Join(match, "")
    return RString(s)
}

var sanitizeAlphabetNumberCompile = regexp.MustCompile("([a-zA-Z0-9 ]+)")

func (r RString) SanitizeAsAlphabetNumber() RString {
    if r == "" {
        return r
    }
    match := sanitizeAlphabetNumberCompile.FindAllString(r.String(), -1)
    s := strings.Join(match, "")
    return RString(s)
}

var sanitizeAlphabetNumberSpaceCompile = regexp.MustCompile("([a-zA-Z0-9]+)")

func (r RString) SanitizeAsAlphabetNumberWithoutSpace() RString {
    if r == "" {
        return r
    }
    match := sanitizeAlphabetNumberSpaceCompile.FindAllString(r.String(), -1)
    s := strings.Join(match, "")
    return RString(s)
}
func (r RString) StripHtml() RString {
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
    return RString(data)
}
