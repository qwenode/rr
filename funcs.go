package rr

import "time"

func Catch(callable func() error) error {
    return callable()
}
func Retry(count int, callable func() error) error {
    var err error
    for i := 0; i < count; i++ {
        err = callable()
        if err == nil {
            return nil
        }
    }
    return err
}
func RetryInterval(count int, sleepFor time.Duration, callable func() error) error {
    var err error
    for i := 0; i < count; i++ {
        err = callable()
        if err == nil {
            return nil
        }
        time.Sleep(sleepFor)
    }
    return err
}
