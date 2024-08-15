package rr

import "sync"

type CC interface {
    Add()
    Done()
    Wait()
}
type cc struct {
    rate  chan int
    group *sync.WaitGroup
}

// 并行任务控制
func NewCC(concurrency int) CC {
    r := new(cc)
    r.rate = make(chan int, concurrency)
    r.group = new(sync.WaitGroup)
    return r
}

//  添加一项
func (r *cc) Add() {
    r.rate <- 1
    r.group.Add(1)
}

//  完成一项
func (r *cc) Done() {
    r.group.Done()
    <-r.rate
}

//  等待所有项目完成
func (r *cc) Wait() {
    r.group.Wait()
}
