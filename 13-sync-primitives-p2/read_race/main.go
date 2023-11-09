package main

// Данный пример показывает, почему важно использовать примитивы синхронизации при конкурентном чтении и записи
// в одну и туже область памяти.
// В выводе можно увидеть как отставание от актуального состояния при чтении, так и частично обновленную область
// или даже обновленную в "неправильном" порядке.
//
// Пример возможного состояния массива при чтении без блока для чтения (RLock):
// [99 99 99 102 102 102 101 101 101]
// Подобный результат попадается нечасто, но иногда можно получить даже такое.

import (
	"fmt"
	"sync"
	"time"
)

type Object struct {
	IsOdd bool
	Nums  []int
}

func (o *Object) Copy() *Object {
	cp := &Object{IsOdd: o.IsOdd, Nums: make([]int, len(o.Nums))}
	copy(cp.Nums, o.Nums)
	return cp
}

func main() {

	const (
		routines   = 8
		iterations = 100
		sliceSize  = 100

		totalIterations = routines * iterations
	)

	mu := sync.RWMutex{}
	num := 0
	obj := &Object{
		Nums: make([]int, sliceSize),
	}

	for i := 0; i < routines; i++ {
		go func(i int) {
			for j := 0; j < iterations; j++ {
				mu.Lock()
				num++
				obj.IsOdd = num%2 == 1
				for numIdx := range obj.Nums {
					obj.Nums[numIdx] = num
					//fmt.Printf("routine - %d, numIdx - %d, mum: %d", i, numIdx, num)
				}
				fmt.Println(i, " increment ", num)
				mu.Unlock()
			}
		}(i)
	}

	slicesWithRaces := [][]int{}
	for num < totalIterations {

		//mu.RLock()
		cp := obj.Copy()
		//mu.RUnlock()
		fmt.Printf("main read %+v \n", *cp)
		if isSliceHasRace(cp.Nums) {
			slicesWithRaces = append(slicesWithRaces, cp.Nums)
		}
		time.Sleep(time.Nanosecond * 100)
	}

	fmt.Printf("\n\nslices with races: \n")
	for _, sl := range slicesWithRaces {
		fmt.Println(sl)
	}
}

func isSliceHasRace(s []int) bool {
	first := s[0]
	for _, e := range s {
		if e != first {
			return true
		}
	}
	return false
}
