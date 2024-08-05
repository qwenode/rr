package rr

import "strings"

const (
    ExceptionMaxStack = 10
)

var (
    ErrExceptionNetwork         = NewExceptionT("Network error")
    ErrExceptionNotFound        = NewExceptionT("Not found")
    ErrExceptionServer          = NewExceptionT("Server error")
    ErrExceptionTimeout         = NewExceptionT("Timeout")
    ErrExceptionBadRequest      = NewExceptionT("Bad request")
    ErrExceptionUnauthorized    = NewExceptionT("Unauthorized")
    ErrExceptionForbidden       = NewExceptionT("Forbidden")
    ErrExceptionTooManyRequests = NewExceptionT("Too many requests")
)

type Exception interface {
    Error() string
    Is(v error) bool
    IsT(v string) bool
    With(v error) Exception
    WithT(v string) Exception
    WithException(v Exception) Exception
    StackMessages() string
}
type exceptionStack struct {
    text   string
    with   []string
    index  int
    length int
}

func (r *exceptionStack) Error() string {
    return r.text
}

func (r *exceptionStack) IsT(v string) bool {
    if v == "" {
        return false
    }
    if v == r.text {
        return true
    }
    for _, err := range r.with {
        if err == r.text {
            return true
        }
    }
    return false
}
func (r *exceptionStack) Is(v error) bool {
    if v == nil {
        return false
    }
    s := v.Error()
    if s == r.text {
        return true
    }
    for _, err := range r.with {
        if err == s {
            return true
        }
    }
    return false
}
func (r *exceptionStack) StackMessages() string {
    var sb strings.Builder
    sb.Grow(r.length)
    sb.WriteString(r.text)
    sb.WriteString("\n")
    for _, s := range r.with {
        sb.WriteString(s)
        sb.WriteString("\n")
    }
    return sb.String()
}
func (r *exceptionStack) WithException(v Exception) Exception {
    return r.WithT(v.Error())
}
func (r *exceptionStack) WithT(v string) Exception {
    if r.index >= ExceptionMaxStack {
        return r
    }
    r.with[r.index] = v
    r.index++
    return r
}
func (r *exceptionStack) With(v error) Exception {
    if v == nil {
        return r
    }
    return r.WithT(v.Error())
}

func NewException(e error) Exception {
    var v string
    if e != nil {
        v = e.Error()
    }
    return NewExceptionT(v)
}
func NewExceptionT(v string) Exception {
    return &exceptionStack{text: v, with: make([]string, 0, ExceptionMaxStack), index: 0, length: len(v)}
}
