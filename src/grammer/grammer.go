package grammer

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

func Slice_use() {
	s0 := make([]string, 3)
	fmt.Println("empty s0: ", s0)

	s0[0] = "string0"
	s0[1] = "string1"
	s0[2] = "string2"
	fmt.Println("init s0", s0)

	s0 = append(s0, "string3")
	s0 = append(s0, "string4")
	fmt.Println("after append", s0)

	s1 := make([]string, len(s0))
	copy(s1, s0)
	fmt.Println("s1: ", s1)

	s2 := s1[2:4]
	fmt.Println("s2: ", s2)

	s3 := make([][]string, 3)
	for i := 0; i < 3; i++ {
		len := i + 1
		s3[i] = make([]string, len)
		for j := 0; j < len; j++ {
			s3[i][j] = strconv.Itoa(i) + strconv.Itoa(j)
		}
	}

	fmt.Println("s3:", s3)
}

func Map_Use() {
	m := make(map[string]int)
	m["v1"] = 1
	m["v2"] = 2
	fmt.Println("map m:", m)

	v1 := m["v1"]
	fmt.Println("map v1: ", v1)

	delete(m, "v1")
	fmt.Println("map m:", m)

	_, prs1 := m["v1"] //第一个值为是否在m里，返回false
	fmt.Println("prs1", prs1)

	v2, prs2 := m["v2"] //v2存在，返回true
	fmt.Println("v2, prs2:", v2, prs2)
}

func VarArgs(str string, nums ...int) {
	fmt.Println("nums:", nums)

	total := 0
	for _, num := range nums {
		total += num
	}

	fmt.Println("total :", total)
}

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.height * r.width
}

func (r rect) perim() float64 {
	return 2 * (r.height + r.width)
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println("g: ", g)
	fmt.Println("g.area():", g.area())
	fmt.Println("g.perim():", g.perim())
}

func UseInterface() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}

//channel
func Worker(done chan bool) {
	fmt.Println("start work")
	time.Sleep(time.Second)
	done <- true
	fmt.Println("end work")
}

func InOut(in chan<- string, out <-chan string) {
	in <- "in1"
	<-out
}

func UseChannel() {
	ch := make(chan string, 2)
	ch <- "chan 1"
	ch <- "chan 2"
	fmt.Println("ch0", <-ch)
	fmt.Println("ch1", <-ch)

	done := make(chan bool, 1)
	Worker(done)
	<-done
	fmt.Println("call Work done")

	ch1 := make(chan string, 1)
	ch2 := make(chan string, 2)
	ch2 <- "out1"
	InOut(ch1, ch2)
	fmt.Println(<-ch1)
	fmt.Println("test inout chan done")

}

func UseChannel_Select() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 5)
		c1 <- "ch1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("case time.after")
	default:
		fmt.Println("no case")
	}
}

func UseChannelRange() {
	list := make(chan string, 3)
	list <- "chan 0"
	list <- "chan 1"
	//close(list)

	fmt.Println("before for")
	for elm := range list {
		fmt.Println("in for")
		fmt.Println(elm)
	}

}
