package lib

import (
	"fmt"
	"log"
	"time"
)

func Foo() string {
	return "hi"
}

func ShowTime(f func() error) error {
	t := time.Now()
	defer func() {
		t2 := time.Now()
		fmt.Printf("Ran in %v\n", t2.Sub(t))
	}()
	return f()
}

func TimeRepeatedly(n int, f func()) {
	trials := make([]time.Duration, n)
	for i, _ := range trials {
		t := time.Now()
		f()
		t2 := time.Now()
		fmt.Printf(".")
		trials[i] = t2.Sub(t)
	}

	fmt.Printf("\nmean = %s\n", computeStats(trials).Mean)
}

type stats struct {
	Mean time.Duration
}

func computeStats(l []time.Duration) stats {
	total := 0.0
	for _, e := range l {
		total += float64(e)
	}

	ret := stats{
		Mean: time.Duration(int64(total / float64(len(l)))),
	}
	return ret
}

func MainWrapper(f func() error) {
	err := f()
	if err != nil {
		log.Fatal(err)
	}
}

func MainWrapperWithTime(f func() error) {
	err := ShowTime(f)
	if err != nil {
		log.Fatal(err)
	}
}
