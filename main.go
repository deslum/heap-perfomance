package main

import (
	"fmt"
	"math/rand"

	"test-heap/heap"
	"test-heap/types"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	actions := &types.Actions{}
	heap.Init(actions)
	for i := 0; i < 10000000; i++ {
		heap.Push(actions, &types.Action{
			EventTimestamp: time.Now().Unix() - rand.Int63n(10000000),
			ConditionA:     rand.Intn(2) == 0,
			ConditionB:     rand.Intn(2) == 1,
		})
	}

	for i:=0;i<100;i++ {
		fmt.Println(heap.Pop(actions))
	}
}
