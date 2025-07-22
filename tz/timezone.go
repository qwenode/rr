package tz

import "time"

// Deprecated: 用新的 20250722 by Node
func UTC() time.Time {
    return time.Now().UTC()
}

// Deprecated: 用新的 20250722 by Node
func Unix() int64 {
    return time.Now().Unix()
}
