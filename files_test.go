package rr

import "testing"

func TestF_GetName(t *testing.T) {
    type args struct {
    }
    tests := []struct {
        name string
        r    string
        args args
        want S
    }{
        {
            "正常有内容有分隔符",
            "file.txt?v=ggw/",
            args{},
            S("file.txt"),
        },
        {
            "正常有内容有分隔符",
            "file.txt/aa",
            args{},
            S("aa"),
        },
        {
            "正常有内容有分隔符",
            "file.txt?v=ggw",
            args{},
            S("file.txt"),
        },
        {
            "正常有内容有分隔符",
            "file.txt",
            args{},
            S("file.txt"),
        },
        {
            "空字符串",
            "",
            args{},
            S(""),
        },
        {
            "无分隔符",
            "fileonly",
            args{},
            S("fileonly"),
        },
    }
    for _, tt := range tests {
        t.Run(
            tt.name, func(t *testing.T) {
                if got := F(tt.r).GetName(); got != tt.want {
                    t.Errorf("GetName() = %v, want %v", got, tt.want)
                }
            },
        )
    }
}
func TestF_GetExtension(t *testing.T) {
    type args struct {
    }
    tests := []struct {
        name string
        r    F
        args args
        want S
    }{
        {
            "正常有扩展名",
            "a.js?v",
            args{},
            S("js"),
        },
        {
            "正常有扩展名",
            "a.js/v",
            args{},
            S("js"),
        },
    }
    for _, tt := range tests {
        t.Run(
            tt.name, func(t *testing.T) {
                if got := tt.r.GetExtension(); got != tt.want {
                    t.Errorf("GetExtension() = %v, want %v", got, tt.want)
                }
            },
        )
    }
}
func TestF_GetContents(t *testing.T) {
    type args struct {
    }
    tests := []struct {
        name string
        r    F
        args args
        want S
    }{
        {
            "正常情况",
            "testdata/f.txt",
            args{},
            S("test contents"),
        },
        {
            "正常情况",
            "files1.go",
            args{},
            S(""),
        },
    }
    for _, tt := range tests {
        t.Run(
            tt.name, func(t *testing.T) {
                if got := tt.r.GetContents(); got != tt.want {
                    t.Errorf("GetContents() = %v, want %v", got, tt.want)
                }
            },
        )
    }
}
