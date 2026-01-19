package main

import (
	"sync"
	"time"
)

var ready bool = false
var userList []int = []int{}

func main() {

	cond := sync.NewCond(&sync.Mutex{})
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {

		go func() {
			defer wg.Done()
			NewRequesst(i, cond)
		}()
	}

	wg.Wait()

	println("End ......")
}

func NewRequesst(userId int, cond *sync.Cond) {
	Checking(userId, cond)
	cond.L.Lock()
	defer cond.L.Unlock()
	for !ready {
		cond.Wait()
	}

	println("user", userId, "Start Stream")
}

func Checking(userid int, cond *sync.Cond) {

	println(userid, "Wating To Start Process")
	time.Sleep(50 * time.Millisecond)
	cond.L.Lock()
	defer cond.L.Unlock()
	userList = append(userList, userid)

	if len(userList) == 5 {
		ready = true
		cond.Broadcast()
		println("userid ", userid, "end wait")
	}

}
