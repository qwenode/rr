package tz

import "time"

func UTC() time.Time {
    return time.Now().UTC()
}
func Unix() int64 {
    return time.Now().Unix()
}
