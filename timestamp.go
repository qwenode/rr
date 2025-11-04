package rr

import "time"

var zeroTime = time.Unix(0, 0)

// 检查时间是否为零值
// 返回 true 如果时间是零值或等于 Unix 时间戳 0 (1970-01-01 00:00:00 UTC)
func TimeIsZero(t time.Time) bool {
    return t.IsZero() || t.Equal(zeroTime)
}

// 检查时间指针是否为零值
// 返回 true 如果指针为 nil、时间是零值或等于 Unix 时间戳 0 (1970-01-01 00:00:00 UTC)
func TimeIsZeroPtr(t *time.Time) bool {
    return t == nil || t.IsZero() || t.Equal(zeroTime)
}
func Utc() time.Time {
    return time.Now().UTC()
}

func Unix() int64 {
    return time.Now().Unix()
}

// 获取今天的日期，时间部分设为零
func TimeGetToday() time.Time {
    now := Utc()
    return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
}
