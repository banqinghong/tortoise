package main

import "fmt"

type State interface {
	Name()                        string
	CanTransitTo(to string)       bool
	OnBegin()
	OnEnd()
}

func StateName(s State) string {
	if s == nil {
		return "none"
	}
	return s.Name()
}

type PendingState struct {
	name  string
}

func (ps *PendingState) Name () string {
	return ps.name
}

func (ps *PendingState) CanTransitTo (to string) bool {
	if to == "running" {
		return true
	}
	return false
}

func (ps *PendingState) OnBegin () {
	fmt.Println("任务等待中")
}

func (ps *PendingState) OnEnd () {
	fmt.Println("任务停止等待")
}

type RunningState struct {
	name  string
}

func (rs *RunningState) Name () string {
	return rs.name
}

func (rs *RunningState) CanTransitTo (to string) bool {
	if to == "success" || to == "failed" {
		return true
	}
	return false
}

func (rs *RunningState) OnBegin () {
	fmt.Println("任务执行中")
}

func (rs *RunningState) OnEnd () {
	fmt.Println("执行任务退出")
}

type SuccessState struct {
	name  string
}

func (ss *SuccessState) Name () string {
	return ss.name
}

func (ss *SuccessState) CanTransitTo (to string) bool {
	return false
}

func (ss *SuccessState) OnBegin () {
	fmt.Println("任务执行成功")
}

func (ss *SuccessState) OnEnd () {
	fmt.Println("success 退出")
}

type FailedState struct {
	name  string
}

func (fs *FailedState) Name () string {
	return fs.name
}

func (fs *FailedState) CanTransitTo (to string) bool {
	return false
}

func (fs *FailedState) OnBegin () {
	fmt.Println("任务执行失败")
}

func (fs *FailedState) OnEnd () {
	fmt.Println("failed 退出")
}
