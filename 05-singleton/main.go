package main

import (
	"fmt"
	"time"

	"github.com/design-pattern/05-singleton/singleton"
)

func main() {
	go func() {
		time.Sleep(time.Millisecond * 300)
		fmt.Println(*singleton.Instance())
	}()

	for i := 0; i < 50; i++ {
		go func(ix int) {
			time.Sleep(time.Millisecond * 60)
			singleton.Instance().Count = singleton.Instance().Count + 1
			fmt.Println(ix, singleton.Instance().Count)
		}(i)
	}

	fmt.Scanln()

	// 14 2
	// 23 4
	// 21 5
	// 22 6
	// 19 7
	// 18 8
	// 24 3
	// 20 9
	// 1 10
	// 32 12
	// 16 11
	// 31 13
	// 15 14
	// 30 15
	// 17 16
	// 29 17
	// 13 18
	// 27 19
	// 25 20
	// 26 21
	// 28 22
	// 46 23
	// 34 24
	// 4 25
	// 48 26
	// 47 27
	// 45 28
	// 44 29
	// 43 30
	// 42 31
	// 41 32
	// 40 33
	// 39 34
	// 35 35
	// 36 36
	// 37 37
	// 33 38
	// 38 39
	// 6 40
	// 9 41
	// 0 43
	// 12 44
	// 11 45
	// 10 46
	// 3 47
	// 7 48
	// 2 49
	// 8 42
	// 5 50
	// 49 51
	// {51}
}
