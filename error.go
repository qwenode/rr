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
    ErrExceptionUnauthorized    = NewExceptionT("Unauthorized")
    ErrExceptionForbidden       = NewExceptionT("Forbidden")
    ErrExceptionTooManyRequests = NewExceptionT("Too many requests")
    ErrExceptionInvalidArgs = NewExceptionT("Invalid arguments")
)

//  异常错误处理
type Exception interface {
    Error() string
    Is(v error) bool
    IsT(v string) bool
    IsE(exception Exception) bool
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
func (r exceptionStack) IsE(exception Exception) bool {
    if exception == nil {
        return false
    }
    return r.IsT(exception.Error())
}
func (r *exceptionStack) IsT(v string) bool {
    if v == "" {
        return false
    }
    if v == r.text {
        return true
    }
    for _, s := range r.with {
        if s == v {
            return true
        }
    }
    return false
}
func (r *exceptionStack) Is(v error) bool {
    if v == nil {
        return false
    }
    return r.IsT(v.Error())
}
func (r *exceptionStack) StackMessages() string {
    var sb strings.Builder
    sb.Grow(r.length)
    sb.WriteString(r.text)
    
    for _, s := range r.with {
        if s == "" {
            continue
        }
        sb.WriteString("\n")
        sb.WriteString(s)
        
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

func NewException(e error, with ...Exception) Exception {
    var v string
    if e != nil {
        v = e.Error()
    }
    
    t := NewExceptionT(v)
    if len(with) > 0 {
        for _, exception := range with {
            t.WithException(exception)
        }
    }
    return t
}
func NewExceptionT(v string, with ...Exception) Exception {
    t := &exceptionStack{text: v, with: make([]string, ExceptionMaxStack), index: 0, length: len(v)}
    if len(with) > 0 {
        for _, exception := range with {
            t.WithException(exception)
        }
    }
    return t
}
