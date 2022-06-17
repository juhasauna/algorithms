package concurrent

import (
	"algorithms/utils"
	"fmt"
	"time"
)

func main() {

	data := []int{}
	for i := 0; i < 10000; i++ {
		data = append(data, initData...)
	}
	// for i := 0; i < 5; i++ {
	inParal(data)
	// consequtively(data)
	// }
}
func consequtively(data []int) {

	defer utils.TimeTrack(time.Now(), "consequtively")
	sum := 0
	for v := range data {
		sum += v
	}
	fmt.Println(sum)
}

func inParal(data []int) {
	defer utils.TimeTrack(time.Now(), "in paral")
	l := len(data)

	ch := make(chan int, 3)
	go func() {
		tempSum := 0
		for _, v := range data[:l/3] {
			tempSum += v
		}
		ch <- tempSum
	}()
	go func() {
		tempSum := 0
		for _, v := range data[l/3 : 2*l/3] {
			tempSum += v
		}
		ch <- tempSum
	}()
	go func() {
		tempSum := 0
		for _, v := range data[2*l/3:] {
			tempSum += v
		}
		ch <- tempSum
	}()
	sum := 0
	for i := 0; i < 3; i++ {
		v := <-ch
		sum += v
	}
	fmt.Println(sum)
}
