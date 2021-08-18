package main

import "fmt"

func main() {
	stateManager := NewStateManager()
	ps := &PendingState{
		name: "pending",
	}
	rs := &RunningState{
		name: "running",
	}
	ss := &SuccessState{
		name: "success",
	}
	fs := &FailedState{
		name: "failed",
	}
	stateManager.Add(ps)
	stateManager.Add(rs)
	stateManager.Add(ss)
	stateManager.Add(fs)

	err := stateManager.Transit("running")
	if err != nil {
		fmt.Println("切换running失败： ", err)
	} else {
		fmt.Println("切换running成功： ")
	}
	errS := stateManager.Transit("success")
	if errS != nil {
		fmt.Println("切换success失败： ", errS)
	} else {
		fmt.Println("切换success成功： ")
	}
	errF := stateManager.Transit("failed")
	if errF != nil {
		fmt.Println("切换failed失败： ", errF)
	} else {
		fmt.Println("切换failed成功： ")
	}
}

