package grammer

import (
	"fmt"
	"log"
	"math"
	"net/http"
	//"runtime"
	"strconv"
	"sync/atomic"
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
	close(list) //如果不关闭， 后面的range list会一直等待第三个channel, 关闭后，读的时候
	//就不会阻塞，不会等待
	//	fmt.Println("list print")
	//	fmt.Println(<-list)
	//	fmt.Println(<-list)
	//	fmt.Println(<-list)
	return
	fmt.Println("before for")
	for elm := range list {
		fmt.Println("in for")
		fmt.Println(elm)
	}
}

//////////////////////// work thread ///////////////////////
type Job struct {
	name string
	id   int
}

func (j Job) DoJob(id int) {
	fmt.Println("worker id:", id, "job id:", j.id, "is done")
}

func woker(id int, jobs <-chan Job, results chan<- int) {
	for j := range jobs {
		time.Sleep(time.Second) //如果不等待， 迁程切换之前已经把job执行完毕
		j.DoJob(id)
		results <- id * 2
	}
}

func UseWorkThread() {
	jobs := make(chan Job, 100)
	results := make(chan int, 100)
	for w := 1; w <= 3; w++ {
		go woker(w, jobs, results)
	}

	for i := 1; i <= 9; i++ {
		job := Job{name: "job", id: i}
		jobs <- job
	}

	close(jobs)
	for a := 1; a <= 9; a++ {
		<-results
	}
}

////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////
// 原子计数器

func UseAtomic() {
	var ops uint64 = 0
	for i := 0; i < 50; i++ {
		go func() {
			for {
				atomic.AddUint64(&ops, 1)
				// runtime.Gosched() //如果注释掉，会无限循环, 占掉CPU的所有资源
			}
		}()
	}

	time.Sleep(time.Second)
	lastOps := atomic.LoadUint64(&ops)
	fmt.Println("ops: ", lastOps)
	lastOps = atomic.LoadUint64(&ops)
	fmt.Println("ops: ", lastOps)
	//fmt.Println("ops: ", ops)
}

//////////////////////////////////////////////////////////

/////////////////////////////////////////////////////////////
//////////////////web http///////////////////////////////////////////
func SimpleHttpHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", v)
	}

	fmt.Fprintf(w, "hello, it's test web")
}

func UseTttp() {
	http.HandleFunc("/", SimpleHttpHandler)
	fmt.Print("begin web listen")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndserve:", err)
	}
}

/////////////////////////////////////////////////////////////
