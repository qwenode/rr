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

type S string

func (r S) String() string {
    return string(r)
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

var sanitizeIntCompile = regexp.MustCompile("[0-9]+")

func (r S) SanitizeAsInt() int {
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
func (r S) SanitizeAsInt64() int64 {
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
func StringJoin(strs ...string) S {
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
        return S(b.String())
    }
    return S(strs[0] + strs[1])
}
func (r S) Prepend(s string) S {
    return StringJoin(s, r.String())
}
func (r S) Append(s string) S {
    return StringJoin(r.String(), s)
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

var sanitizeAlphabetCompile = regexp.MustCompile("([a-zA-Z ]+)")

func (r S) SanitizeAsAlphabet() S {
    if r == "" {
        return r
    }
    match := sanitizeAlphabetCompile.FindAllString(r.String(), -1)
    s := strings.Join(match, "")
    // s = strings.Join(strings.Fields(s), " ")
    return S(s)
}

var sanitizeAlphabetSpaceCompile = regexp.MustCompile("([a-zA-Z]+)")

func (r S) SanitizeAsAlphabetWithoutSpace() S {
    if r == "" {
        return r
    }
    match := sanitizeAlphabetSpaceCompile.FindAllString(r.String(), -1)
    s := strings.Join(match, "")
    return S(s)
}

var sanitizeAlphabetNumberCompile = regexp.MustCompile("([a-zA-Z0-9 ]+)")

func (r S) SanitizeAsAlphabetNumber() S {
    if r == "" {
        return r
    }
    match := sanitizeAlphabetNumberCompile.FindAllString(r.String(), -1)
    s := strings.Join(match, "")
    return S(s)
}

var sanitizeAlphabetNumberSpaceCompile = regexp.MustCompile("([a-zA-Z0-9]+)")

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
