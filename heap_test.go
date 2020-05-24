package main

import (
	"math/rand"
	"sort"
	"testing"
	"time"

	"test-heap/heap"
	"test-heap/types"
)

func BenchmarkMinElementWithTuneHeap(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for n := 0; n < 10; n++ {
		b.StopTimer()
		rand.Seed(time.Now().Unix())
		actions := &types.Actions{}
		heap.Init(actions)
		for i := 0; i < 1000000; i++ {
			heap.Push(actions, &types.Action{
				EventTimestamp: time.Now().Unix() - rand.Int63n(10000000),
				ConditionA:     rand.Intn(1) == 0,
				ConditionB:     rand.Intn(1) == 1,
			})
		}

		b.StartTimer()
		_ = heap.Pop(actions)
	}
}

func BenchmarkMinElementWithSort(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for n := 0; n < 10; n++ {
		b.StopTimer()
		actions := types.Actions{}
		for i := 0; i < 1000000; i++ {
			actions = append(actions, &types.Action{
				EventTimestamp: time.Now().Unix() - rand.Int63n(10000000),
				ConditionA:     rand.Intn(1) == 0,
				ConditionB:     rand.Intn(1) == 1,
			})
		}

		b.StartTimer()
		sort.Sort(actions)
		_ = actions[0]
	}
}

func BenchmarkMinElementWithLoop(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for n := 0; n < 10; n++ {
		b.StopTimer()
		actions := types.Actions{}
		for i := 0; i < 1000000; i++ {
			actions = append(actions, &types.Action{
				EventTimestamp: time.Now().Unix() - rand.Int63n(10000000),
				ConditionA:     rand.Intn(1) == 0,
				ConditionB:     rand.Intn(1) == 1,
			})
		}

		b.StartTimer()
		min := actions[0]
		for _, x := range actions {
			if x.EventTimestamp < min.EventTimestamp && x.ConditionA && x.ConditionB {
				min = x
			}
		}

		_ = min
	}

}
