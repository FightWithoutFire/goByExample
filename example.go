package main

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"embed"
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"errors"
	"flag"
	"fmt"
	"io"
	"math"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	s "strings"
	"sync"
	"sync/atomic"
	"syscall"
	"text/template"
	"time"
	"unicode/utf8"
)

func PrintHelloWorld() {
	fmt.Printf("hello world")
}

func Values() {
	fmt.Println("hello " + "world")
	fmt.Println("1+1= ", 1+2)
	fmt.Println("7.0 / 3.0 =", 7.0/3.0)
	//fmt.Println("true && false = ", true && true)
	//fmt.Println("false || true = ", false || true)
	//fmt.Println("!ture = ", !true)
}

func Variables() {

	var a = "initial"
	fmt.Println(a)

	var b, c = 2, 3
	fmt.Println(b, c)

	var d int
	fmt.Println(d)

	e := "apple"
	fmt.Println(e)

}

const a = "constants"

func Constants() {
	fmt.Println(a)

	const n = 500000000
	fmt.Println(n)

	const c = 3e20 / n
	fmt.Println(c)

	fmt.Println(int64(c))

	fmt.Println(math.Sin(c))

}

func For() {
	i := 1
	for i < 3 {
		fmt.Println(i)
		i++
	}

	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	for {
		fmt.Println(2)
		break
	}

	for n := 0; n < 10; n++ {
		if n%3 == 1 {
			continue
		}
		fmt.Println(n)
	}
}

func IfElse() {

	n := 10
	if n%2 == 0 {
		fmt.Println("n is even")
	}

	if x := 11; x < 0 {
		fmt.Println(" x < 0 ")
	} else if x < 10 {
		fmt.Println(" x < 10 ")
	} else {
		fmt.Println(" x > 10 ")
	}

}

func Switch() {
	i := 2
	switch i {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	}

	d := time.Now().Weekday()
	switch d {
	case time.Saturday, time.Sunday:
		fmt.Println("weekend")
	default:
		fmt.Println(" not weekend ")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("before Noon")
	default:
		fmt.Println("after Noon")

	}

	whatAmI := func(any interface{}) {
		switch any.(type) {
		case string:
			fmt.Println("string")
		case int:
			fmt.Println("int")
		case bool:
			fmt.Println("boolean")
		}
	}
	whatAmI(1)
	whatAmI(true)
	whatAmI("abcd")

}

func Array() {
	var a [4]int
	fmt.Println(a)

	a[3] = 100
	fmt.Println(a)

	m := [3]int{1, 2, 3}
	fmt.Println(m)

	var t [2][3]int
	for x := 0; x < 2; x++ {
		for y := 0; y < 3; y++ {
			t[x][y] = x + y
		}
	}
	fmt.Println(t)
}

func Slice() {
	i := make([]int, 3)
	fmt.Println("init", i)

	i[1] = 2
	i[2] = 3
	fmt.Println("set ", i)
	fmt.Println("get ", i[2])

	fmt.Println("len ", len(i))

	i = append(i, 3)
	i = append(i, 4, 5)

	fmt.Println(i)

	c := make([]int, len(i))
	copy(c, i)
	fmt.Println("c", c)
	fmt.Println("i", i)

	l := i[:4]
	fmt.Println("l1", l)

	l = i[1:]
	fmt.Println("l2", l)

	l = i[2:3]
	fmt.Println("l3", l)

	t := []string{"g", "h", "j"}
	fmt.Println("t", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		twoD[i] = make([]int, i+1)
		for j := 0; j < i+1; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("twoD ", twoD)
}

func Exit() {
	defer fmt.Println("!")
	os.Exit(3)

}

func Map() {
	m := make(map[string]int)
	m["TES"] = 150
	m["EDG"] = 100
	m["RNG"] = 130
	m["JDG"] = 80
	fmt.Println("map ", m)

	v := m["TES"]

	fmt.Println(v)

	fmt.Println(len(m))

	delete(m, "TES")

	_, prs := m["TES"]
	fmt.Println(prs)
	fmt.Println(m)

	fmt.Println(m["TES"])

	fmt.Println(v)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map", n)
}

func Range() {
	v := []string{"1", "2", "3"}

	for _, num := range v {
		fmt.Println(num)
	}

	for i, num := range v {
		if num == "3" {
			fmt.Println(i)
		}
	}

	m := map[string]string{"1": "a", "2": "b"}
	for key, val := range m {
		fmt.Printf("%s -> %s\n", key, val)
	}

	for key := range m {
		fmt.Println("key : ", key)
	}

	for i, char := range "good" {
		fmt.Printf("%d -> %d\n", i, char)
	}
}

func StringAndRune() {

	const s = "สวัสดี"

	fmt.Println("Len: ", len(s))

	for i := 0; i < len(s); i++ {

		fmt.Printf("%x ", s[i])

	}

	fmt.Println("")

	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	for index, val := range s {
		fmt.Printf("%#U starts at %d \n", val, index)
	}

	fmt.Println("USING DecodeRuneInString")

	for i, w := 0, 0; i < len(s); i += w {
		inString, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d \n", inString, i)
		w = size
		examineRune(inString)
	}

}

func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}

func takeTwo1(a int, b int) (int, string) {
	fmt.Println(a + b)
	return a + b, "ok"
}
func takeTwo2(a, b int) (int, string) {
	fmt.Println(a + b)
	return a + b, "ok"
}

func takeMany(a ...int) int {
	fmt.Println("len:", len(a))
	sum := 0
	for _, n := range a {
		sum += n
	}
	fmt.Println("sum", sum)
	return sum
}

func Closure() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func fact01(n int) int {
	if n == 1 {
		return n
	}
	return n * fact01(n-1)

}

func fact22() {
	var fib func(int) int

	fib = func(i int) int {
		if i < 2 {
			return i
		}
		return fib(i-1) + fib(i-2)
	}

	fmt.Println(fib(10))

}

func zeroval(i int) {
	i = 3
}

func zeroprt(i *int) {
	*i = 4
}

func Point() {
	a, b := 1, 2
	zeroval(a)
	zeroprt(&b)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(&b)

}

type Person struct {
	name string
	age  int
}

func GetPerson() *Person {
	p := Person{"JackyLove", 22}
	fmt.Println("new Person: ", p)
	p2 := Person{name: "UZI"}
	p2.age = 25
	fmt.Println("new Person2", p2)
	return &p
}

func CopyPerson(p *Person) {
	p.name = "P1"
	fmt.Println("in 1 ", p)
}

func CopyPerson2(p Person) {
	p.name = "P2"
	fmt.Println("in 2 ", p)
}

func TPerson() {
	p := Person{name: "p", age: 2}
	CopyPerson(&p)
	fmt.Println(p)

	CopyPerson2(p)
	fmt.Println(p)
}

func (receiver *Person) BMI1() {
	receiver.name = "abc"

}

func (receiver Person) BMI2() {
	receiver.name = "abd"

}

func Meth() {
	p := Person{
		name: "e",
		age:  21,
	}

	p.BMI1()
	fmt.Println(p)
	p.BMI2()
	fmt.Println(p)
}

type geometry interface {
	area() float64
	prim() float64
}

type rect struct {
	with   float64
	length float64
}

type circle struct {
	radius float64
}

func (receiver rect) area() float64 {
	fmt.Println(receiver.with * receiver.length)
	return receiver.with * receiver.length
}

func (receiver rect) prim() float64 {
	fmt.Println(2*receiver.with + 2*receiver.length)
	return 2*receiver.with + 2*receiver.length
}

func Interface() {
	var graph geometry
	graph = rect{with: 5.0, length: 3.0}
	fmt.Println(graph.area())
	fmt.Println(graph.prim())
}

type base struct {
	num int
}

type container struct {
	base
	name string
}

func (c container) describe() string {
	fmt.Println(c.name)
	return c.name
}

func EmbStruct() {
	co := container{
		base{
			num: 10,
		},
		"name",
	}
	fmt.Println(co)
	fmt.Println(co.base.num)

	type describer interface {
		describe() string
	}
	var d describer
	d = co
	fmt.Println(d.describe())
}

func MapKeys[K comparable, V any](m map[K]V) []K {

	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

type List[T any] struct {
	head, tail *element[T]
}
type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) GetAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}
func Generic() {
	var m = map[int]string{1: "2", 2: "4", 4: "8"}

	fmt.Println("keys:", MapKeys(m))

	_ = MapKeys[int, string](m)
	var lst List[int]
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.GetAll())
}

func f1(args int) (int, error) {
	if args == 43 {
		return -1, errors.New("error no 43")
	}

	return args, nil
}

func Errors() {

	arr := []int{7, 43}
	for _, val := range arr {
		if i, err := f1(val); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(i)
		}
	}

	for _, val := range arr {
		if i, err := f2(val); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(i)
		}
	}

	_, e := f2(43)
	if ae, ok := e.(*SelfMakeError); ok {
		fmt.Println(ae.msg)
		fmt.Println(ae.code)
	}
}

type SelfMakeError struct {
	msg  string
	code int
}

func (s *SelfMakeError) Error() string {
	return fmt.Sprintf("%s -> %d", s.msg, s.code)
}

func f2(args int) (int, error) {
	if args == 43 {
		return -1, &SelfMakeError{"error no 43", 43}
	}

	return args, nil
}

func f(msg string, group *sync.WaitGroup) {
	for i := 0; i < 3; i++ {
		fmt.Println(msg, i)
	}
	if group != nil {
		group.Done()
	}
}

func Goroutines() {
	group := sync.WaitGroup{}
	group.Add(2)
	f("f1", nil)

	go f("f2", &group)

	go func(msg string) {
		fmt.Println(msg)
		group.Done()
	}("f3")

	group.Wait()
	fmt.Println("done")
}

func Channel() {

	ch := make(chan string)

	go func() {
		ch <- "abcd"
		ch <- "efg"
	}()

	c := <-ch
	fmt.Println(c)

	c = <-ch
	fmt.Println(c)

}

func ChannelBuffering() {

	ch := make(chan string, 2)

	ch <- "aa"
	ch <- "bb"

	c := <-ch
	fmt.Println(c)

	c = <-ch
	fmt.Println(c)

}

func c1(done chan bool) {

	fmt.Println("do 1")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true

}
func ChannelSync() {

	ch := make(chan bool, 1)

	fmt.Println("start")
	go c1(ch)

	<-ch
}

func ping(pi chan<- string) {
	pi <- "good"
}

func pong(po <-chan string, pops chan<- string) {
	msg := <-po
	pops <- msg
}

func ChannelDirection() {
	p1 := make(chan string, 1)
	p2 := make(chan string, 1)

	go ping(p1)

	go pong(p1, p2)
	fmt.Println(<-p2)
}

func Select() {

	c1 := make(chan string, 1)
	c2 := make(chan string, 1)

	go func() {
		time.Sleep(time.Second)
		c1 <- "hello c1"

	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "hello c2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-c1:
			fmt.Println("c1 finish", m1)
		case m2 := <-c2:
			fmt.Println("c2 finish", m2)

		}

	}

}

func Timeout() {

	ch := make(chan string, 1)

	go func() {
		fmt.Println("start")
		time.Sleep(3 * time.Second)
		ch <- "check"
	}()
	select {
	case s := <-ch:
		fmt.Println("not timeout", s)
	case <-time.After(2 * time.Second):
		fmt.Println("time out 2s")
	}

	ch = make(chan string, 1)

	go func() {
		fmt.Println("start")
		time.Sleep(3 * time.Second)
		ch <- "check"
	}()
	select {
	case s := <-ch:
		fmt.Println("not timeout", s)
	case <-time.After(4 * time.Second):
		fmt.Println("time out 2s")
	}
}

func MultiSelect() {

	ch1 := make(chan string)
	ch2 := make(chan string)

	select {
	case msg := <-ch1:
		fmt.Println("select 1", msg)
	default:
		fmt.Println("select default")
	}

	msg := "msg"
	select {
	case ch2 <- msg:
		fmt.Println("into select2 ")
	default:
		fmt.Println("into select2 default")
	}

	select {
	case m1 := <-ch1:
		fmt.Println("select3 m1", m1)
	case m2 := <-ch2:
		fmt.Println("select3 m2", m2)
	default:
		fmt.Println("no select 3")

	}

}

func ClosingChannel() {

	ch1 := make(chan string, 2)
	ch2 := make(chan string)

	go func() {
		for {
			s1, more := <-ch1
			if more {
				fmt.Println("get from ch1", s1)
			} else {
				fmt.Println("no more")
				ch2 <- "done"
				return
			}
		}
	}()

	for i := 0; i < 4; i++ {
		ch1 <- fmt.Sprintf("send job %d", i)
		fmt.Println("send job ", i)
	}

	close(ch1)
	fmt.Println("done send ")
	<-ch2

}

func RangeOverChannel() {

	ch1 := make(chan string, 5)
	for i := 0; i < 5; i++ {
		ch1 <- fmt.Sprintf("hello%d", i)
	}
	close(ch1)

	for s := range ch1 {
		fmt.Println(s)
	}

}

func Timer() {

	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C
	fmt.Println("time up")

	timer2 := time.NewTimer(2 * time.Second)
	go func() {
		<-timer2.C
		fmt.Println("finish in goroutine")
	}()

	stop := timer2.Stop()
	if stop {
		fmt.Println("stop")
	}
	time.Sleep(5 * time.Second)
}

func Ticker() {

	ticker := time.NewTicker(time.Second)

	done := make(chan bool)

	go func() {
		for {
			select {
			case c := <-ticker.C:
				fmt.Println("get ticker ", c)
			case <-done:
				fmt.Println("goroutines out")
				return
			}
		}
	}()

	time.Sleep(16 * time.Second)
	ticker.Stop()
	done <- true
	fmt.Println("main finish")

}

func worker(id int, jobs <-chan int, result chan<- int) {
	for job := range jobs {
		fmt.Printf("%d start %d\n", id, job)
		time.Sleep(time.Second)
		fmt.Printf("%d finish %d\n", id, job)
		result <- job * 2
	}
}

func WorkerPool() {

	jobs := make(chan int, 5)
	results := make(chan int, 5)
	jobNumber := 25
	workerNumber := 3

	for i := 0; i < workerNumber; i++ {
		go worker(i, jobs, results)
	}

	go func(jobsIn chan int) {
		for i := 0; i < jobNumber; i++ {
			jobsIn <- i
		}
		close(jobs)
	}(jobs)

	for i := 0; i < jobNumber; i++ {
		fmt.Println(<-results)

	}
}

func w1(str int) {
	fmt.Println("work group start ", str)
	time.Sleep(time.Second)
	fmt.Println("work group finish ", str)
}

func WaitGroup() {

	group := sync.WaitGroup{}

	for i := 0; i < 3; i++ {
		group.Add(1)
		go func(i int) {
			defer group.Done()
			w1(i)
		}(i)
	}

	group.Wait()
}

func RateLimiter() {

	requests := make(chan int, 3)
	for i := 0; i < 3; i++ {
		requests <- i
	}
	close(requests)
	limiter := time.Tick(200 * time.Millisecond)
	for req := range requests {
		<-limiter
		fmt.Println("limit ", req, time.Now())
	}

	timeRateLimit := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		timeRateLimit <- time.Now()
	}
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			timeRateLimit <- t
		}
	}()

	burstRequest := make(chan int, 3)
	for i := 0; i < 3; i++ {
		burstRequest <- i
	}
	close(burstRequest)

	for req := range burstRequest {
		<-timeRateLimit
		fmt.Println("request: ", req, time.Now())
	}
}

func AtomicCount() {

	var ops1 uint32
	var ops2 uint32

	group := sync.WaitGroup{}
	for i := 0; i < 40; i++ {
		group.Add(1)
		go func(ops1, ops2 *uint32) {
			defer group.Done()
			for i := 0; i < 10000; i++ {
				atomic.AddUint32(ops1, 1)
				*ops2++
			}
		}(&ops1, &ops2)
	}
	group.Wait()
	fmt.Println("result 1:", ops1)
	fmt.Println("result 2:", ops2)

}

type Container struct {
	mux sync.Mutex

	counters map[string]int
}

func (c *Container) inc(name string) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.counters[name]++
}

func Mutexes() {

	c := Container{counters: map[string]int{"a": 0, "b": 0}}

	group := sync.WaitGroup{}
	doInc := func(c *Container, name string, count int) {
		for i := 0; i < count; i++ {
			c.inc(name)
		}
		group.Done()
	}
	for i := 0; i < 10; i++ {
		group.Add(1)
		go doInc(&c, "a", 20000)

	}

	group.Wait()
	fmt.Println(c.counters["a"])
}

type readOp struct {
	key  int
	reap chan int
}

type writeOp struct {
	key  int
	val  int
	reap chan bool
}

func StatefulGoroutines() {
	var readOps uint64
	var writeOps uint64

	readCh := make(chan readOp)
	writeCh := make(chan writeOp)

	go func() {
		var state = make(map[int]int)
		for true {
			select {
			case read := <-readCh:
				read.reap <- state[read.key]
			case write := <-writeCh:
				state[write.key] = write.val
				write.reap <- true
			}
		}

	}()

	for r := 0; r < 100; r++ {
		go func() {
			for true {
				read := readOp{
					key:  rand.Intn(64),
					reap: make(chan int),
				}
				readCh <- read
				<-read.reap
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for true {
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					reap: make(chan bool),
				}
				writeCh <- write
				<-write.reap
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()

	}
	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOpsFinal:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOpsFinal:", writeOpsFinal)

}

func sorted() {

	strLs := []string{"3", "a", "b", "c", "d"}
	sort.Strings(strLs)
	fmt.Println(strLs)

	intLs := []int{2, 3, 5, 7, 3, 2}
	sort.Ints(intLs)
	fmt.Println(intLs)

	areSorted := sort.IntsAreSorted(intLs)
	fmt.Println(areSorted)
}

type byLength []string

func (b byLength) Len() int {
	return len(b)
}

func (by byLength) Swap(a, b int) {
	by[a], by[b] = by[b], by[a]
}

func (by byLength) Less(a, b int) bool {
	return len(by[a]) < len(by[b])
}
func SortByFunction() {

	length := []string{"aaa", "be", "1123", "3", "5"}
	sort.Sort(byLength(length))
	fmt.Println(length)

}

func Panic() {

	//panic("a problem")

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic("/temp/file create fail")
	}
}

func CreateFile(path string) *os.File {
	create, err := os.Create(path)
	if err != nil {
		panic("create file fail")
	}

	return create

}

func CloseFile(file *os.File) {
	err := file.Close()
	if err != nil {
		panic("close file error")
	}
}

func WriteFile(file *os.File) {
	_, err := file.WriteString("good morning\n")
	if err != nil {
		panic("file write error")
	}
}
func Defer() {
	file := CreateFile("D:\\codeProject\\go\\goByExample\\hello.txt")
	defer CloseFile(file)
	WriteFile(file)

}

func myPanic() {
	panic("panic")
}
func Recover() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover")

		}
	}()

	myPanic()
	fmt.Println("end")
}

func SFunction() {
	p := fmt.Println
	p("Contains:", s.Contains("helo", "a"))
	p("ToLower:", s.ToLower("hello"))
	p("ToUpper:", s.ToUpper("hello"))
	p("Split:", s.Split("hello", "e"))
	p("Count:", s.Count("hello", "l"))
	p("HasPrefix:", s.HasPrefix("hello", "he"))
	p("HasSuffix:", s.HasSuffix("hello", "ll"))
	p("Index:     ", s.Index("test", "e"))
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", s.Repeat("a", 5))
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	p("Replace:   ", s.Replace("foooo", "o", "0", 3))
	p("Split:     ", s.Split("a-b-c-d-e", "-"))

}

type point struct {
	x int
	y int
}

func FormatPrint() {

	p := point{3, 5}

	fmt.Printf("only value v 1: %v\n", p)

	fmt.Printf("with struct v 2: %+v\n", p)

	fmt.Printf("full name v 3 : %#v\n", p)

	fmt.Printf("type 4 : %T\n", p)

	fmt.Printf("boolean 5 : %t\n", false)
	fmt.Printf("decimal 6 : %d\n", 432)
	fmt.Printf("bin 7 : %b\n", 43)
	fmt.Printf("char 8 : %c\n", 'a')
	fmt.Printf("hex x : %x\n", 32.5)

	fmt.Printf("float 1 : %f\n", 32.5)
	fmt.Printf("float 2 : %e\n", 154444000000.5)
	fmt.Printf("float 3 : %E\n", 154444000000.5)
	fmt.Printf("string 1 : %s\n", "\"strings\"")
	fmt.Printf("string 2 : %q\n", "\"strings\"")
	fmt.Printf("string 3 : %x\n", "\"strings\"")
	fmt.Printf("point 1 : %p\n", &p)
	fmt.Printf("fomat 1 : |%6d|%6d|\n", 32, 25)
	fmt.Printf("fomat 2 : |%6.2f|%6.2f|\n", 32.2, 25.4)
	fmt.Printf("fomat 3 : |%-6.2f|%-6.2f|\n", 32.2, 25.4)
	fmt.Printf("fomat 4 : |%-6s|%-6s|\n", "abc", "def")
	fmt.Printf("fomat 5 : |%6s|%6s|\n", "abc", "def")
	s := fmt.Sprintf("fomat 6 : |%-6s|%-6s|\n", "abc", "def")
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "helo god %s\n", "hehehe")

}

func TextTemplate() {
	t1 := template.New("t1")
	t1, err := t1.Parse("Value IS {{ . }} \n")
	if err != nil {
		panic("error")
	}

	t1 = template.Must(t1.Parse("Values : {{ . }}\n"))

	t1.Execute(os.Stdout, "helo")
	t1.Execute(os.Stdout, 6)
	t1.Execute(os.Stdout, []string{
		"good",
		"noon",
		"let",
		"us",
		"go",
	})

	Create := func(name, temp string) *template.Template {
		return template.Must(template.New(name).Parse(temp))
	}

	create := Create("t2", "value, {{.Name}} {{.}}\n")
	create.Execute(os.Stdout, struct {
		Name string
	}{"john"})

	create.Execute(os.Stdout, map[string]string{
		"Name": "Ben",
	})

	t3 := Create("t3",
		"{{if . -}} yes {{else -}} no {{end}}\n")
	t3.Execute(os.Stdout, "not empty")
	t3.Execute(os.Stdout, "")

	t4 := Create("t4",
		"Range: {{range .}}{{.}} {{end}}\n")
	t4.Execute(os.Stdout,
		[]string{
			"Go",
			"Rust",
			"C++",
			"C#",
		})
}

func Reg() {
	matchString, err := regexp.MatchString("p([a-z]+)ch", "pe1ach")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(matchString)

	compile, err := regexp.Compile("p([a-z]+)ch")
	fmt.Println(compile.MatchString("peach"))
	fmt.Println(compile.FindString("punch peach punk"))
	fmt.Println(compile.FindAllString("punch peach punk", -1))
	fmt.Println(compile.FindAllStringIndex("punch peach punk", -1))
	fmt.Println(compile.FindAllStringSubmatchIndex("punch peach punk", -1))

	fmt.Println(compile.Match([]byte("punch peach punk")))

	mustCompile := regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(mustCompile)

	fmt.Println(mustCompile.ReplaceAllString("punch peach punk", "bacd"))

	temp := []byte("punch peach punk")
	fmt.Println(string(mustCompile.ReplaceAllFunc(temp, bytes.ToUpper)))
}

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"date"`
}

func Json() {
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	floB, _ := json.Marshal(3.5)
	fmt.Println(string(floB))

	intB, _ := json.Marshal(43)
	fmt.Println(string(intB))

	strB, _ := json.Marshal("what why")
	fmt.Println(string(strB))

	sliB, _ := json.Marshal([]string{"TES", "LNG", "RNG", "V5", "EDG"})
	fmt.Println(string(sliB))

	mapB, _ := json.Marshal(map[string]string{"name": "TES", "game": "lol", "goal": "S12"})
	fmt.Println(string(mapB))

	res1 := response1{
		Page:   4,
		Fruits: []string{"A", "B", "C", "D"},
	}
	res1B, _ := json.Marshal(&res1)
	fmt.Println(string(res1B))

	res2 := response2{
		Page:   4,
		Fruits: []string{"E", "F", "G", "H"},
	}
	res2B, _ := json.Marshal(&res2)
	fmt.Println(string(res2B))

	var data map[string]interface{}

	byt := []byte(`{"num": 4.3, "str": ["1", "2", "3"]}`)
	json.Unmarshal(byt, &data)
	fmt.Println(data)

	f3 := data["num"].(float64)
	fmt.Println(f3)

	str := data["str"].([]interface{})
	s2 := str[0].(string)
	fmt.Println(s2)

	resB := `{"page":25,"date":["I","J","K","L"]}`
	res := response2{}
	json.Unmarshal([]byte(resB), &res)
	fmt.Printf("%#v\n", res)

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
	enc.Encode([]string{"abcd", "efgh", "ijkl"})

}

type Plant struct {
	XMLName xml.Name `xml:"plant" json:"xml_name"`
	Name    string   `xml:"name" json:"name,omitempty"`
	Id      int      `xml:"id,attr" json:"id,omitempty"`
	Origin  []string `xml:"origin" json:"origin,omitempty"`
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v", p.Id, p.Name, p.Origin)
}

func XmlRead() {

	coffee := &Plant{Name: "coffee", Id: 32}
	coffee.Origin = []string{"Ethiopia", "Brazil"}
	fmt.Println(coffee)

	out, _ := xml.MarshalIndent(coffee, " ", " ")

	fmt.Println(xml.Header + string(out))

	var anther Plant

	if err := xml.Unmarshal(out, &anther); err != nil {
		panic(err)
	}

	fmt.Println(anther)

	tomato := &Plant{Name: "tomato", Id: 32, Origin: []string{"a", "b"}}
	marshal, _ := xml.Marshal(tomato)
	fmt.Println(string(marshal))

	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants  []*Plant `xml:"parent>child>plant"`
	}

	nesting := &Nesting{}
	nesting.Plants = []*Plant{tomato, coffee}

	out2, _ := xml.MarshalIndent(nesting, " ", " ")

	fmt.Println(string(out2))
}

func Time() {
	p := fmt.Println
	t1 := time.Now()
	fmt.Println(t1)

	then := time.Date(
		2022, 5, 3, 18, 11, 23, 33, time.UTC)
	fmt.Println(then)

	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	p(then.Weekday())
	p(then.Before(t1))
	p(then.After(t1))
	p(then.Equal(t1))

	diff := t1.Sub(then)
	p(diff)

	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	p(then.Add(diff))
	p(then.Add(-diff))
}

func UnixTime() {

	now := time.Now()
	fmt.Println(now.Unix())
	fmt.Println(now.UnixMicro())
	fmt.Println(now.UnixMilli())
	fmt.Println(now.UnixNano())

	fmt.Println(time.Unix(1661937327, 323))
	fmt.Println(time.UnixMilli(1661937372173))
	fmt.Println(time.UnixMicro(1661937372173245))

}

func TimeFormat() {

	t := time.Now()
	p := fmt.Println
	p(t)
	p(t.Format(time.RFC822))

	t1, _ := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")
	p(t1)

	p(t.Format("3:04PM"))
	p(t.Format("Mon Jan _2 15:04:05 2006"))
	p(t.Format("2006-01-02 15:04:05"))
	p(t.Format("2006-01-02T15:04:05.999999-07:00"))
	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	p(t2)

	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM")
	p(e)
}

func RandomNumber() {
	fmt.Print(rand.Intn(50), ",")
	fmt.Print(rand.Intn(50), ",")
	fmt.Print(rand.Intn(50), ",")
	fmt.Println()
	fmt.Print(rand.Float64()*5.0+5.0, ",")
	fmt.Print(rand.Float64()*5.0+5.0, ",")
	fmt.Print(rand.Float64()*5.0+5.0, ",")
	fmt.Print(rand.Float64()*5.0+5.0, ",")
	fmt.Println()

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	fmt.Println(r.Intn(3))

	s1 := rand.NewSource(32)
	r1 := rand.New(s1)
	fmt.Print(r1.Intn(100), ",")
	fmt.Print(r1.Intn(100), ",")
	fmt.Print(r1.Intn(100), ",")
	fmt.Println()

	s2 := rand.NewSource(32)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ",")
	fmt.Print(r2.Intn(100), ",")
	fmt.Print(r2.Intn(100), ",")
	fmt.Println()

	s3 := rand.NewSource(1200)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100), ",")
	fmt.Print(r3.Intn(100), ",")
	fmt.Print(r3.Intn(100), ",")

}

func NumberParse() {
	float, _ := strconv.ParseFloat("321.3", 64)
	fmt.Println(float)

	int, _ := strconv.ParseInt("321", 0, 64)
	fmt.Println(int)

	int3, _ := strconv.ParseInt("0x325", 0, 64)
	fmt.Println(int3)

	bool, _ := strconv.ParseBool("true")
	fmt.Println(bool)

	uint, _ := strconv.ParseUint("532", 0, 64)
	fmt.Println(uint)

	_, err := strconv.Atoi("true")
	fmt.Println(err)

	int2, _ := strconv.Atoi("32")
	fmt.Println(int2)
}

func URLParse() {

	s := "postgres://user:pass@host.com:5432/path?k=v#f"
	parse, _ := url.Parse(s)

	fmt.Println(parse.Host)
	fmt.Println(parse.Scheme)
	fmt.Println(parse.User.Username())
	p, _ := parse.User.Password()
	fmt.Println(p)

	host, port, _ := net.SplitHostPort(parse.Host)
	fmt.Println(host, port)

	fmt.Println(parse.Path)
	fmt.Println(parse.Fragment)

	fmt.Println(parse.RawQuery)
	query, _ := url.ParseQuery(parse.RawQuery)
	fmt.Println(query)
	fmt.Println(query["k"][0])
}

func Sha256() {

	s := "sha256 is string"

	hash := sha256.New()
	hash.Write([]byte(s))

	bs := hash.Sum(nil)
	fmt.Printf("%x\n", bs)
}

func Base64() {

	data := "abc123!?$*&()'-=@~"

	enc := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Printf("%v \n", enc)

	decodeString, _ := base64.StdEncoding.DecodeString(enc)
	fmt.Printf("%v \n", string(decodeString))

	uEnc := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := base64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}

func Read() {
	file, err := os.ReadFile("./hello.txt")
	check(err)
	fmt.Println(string(file))

	open, err := os.Open("./hello.txt")
	defer open.Close()
	check(err)
	b1 := make([]byte, 6)
	r1, err := open.Read(b1)
	check(err)
	fmt.Println(string(b1[:r1]))

	open.Seek(2, 0)
	check(err)
	b2 := make([]byte, 6)
	r2, err := open.Read(b2)
	check(err)
	fmt.Println(string(b2[:r2]))

	b3 := make([]byte, 4)

	least, err := io.ReadAtLeast(open, b3, 4)
	check(err)
	fmt.Println(string(b3[:least]))

	open.Seek(0, 0)
	reader := bufio.NewReader(open)
	peek, err := reader.Peek(5)
	check(err)
	fmt.Println(string(peek))

}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}

}

func Write() {
	i := []byte("hello\n go\n")

	os.WriteFile("./helo.txt", i, 0644)

	file, err := os.Create("./hello1.txt")
	defer file.Close()
	check(err)
	file.Write(i)

	file.WriteString("write another book")
	file.Sync()

	writer := bufio.NewWriter(file)
	writer.WriteString("write book")
	writer.Flush()

}

func Scanner() {

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		text := scanner.Text()
		if text == "exit" {
			break
		} else {
			fmt.Println(text)
		}
	}

}

func FilePath() {

	p := filepath.Join("heo", "good", "noway")
	fmt.Println(p)

	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	fmt.Println("Dir(p)", filepath.Dir(p))
	fmt.Println("Base(p)", filepath.Base(p))

	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("C:/dir/file"))

	filename := "config.json"
	ext := filepath.Ext(filename)
	fmt.Println(ext)

	fmt.Println(s.TrimSuffix(filename, ext))

	rel, err := filepath.Rel("a/b/c", "a/b/c/d")
	check(err)
	fmt.Println(rel)

	rel, err = filepath.Rel("a/b/c", "a/c/d")
	check(err)
	fmt.Println(rel)

}

func Directory() {
	err := os.Mkdir("subdir", 0755)
	check(err)

	defer os.RemoveAll("subdir")

	createEmptyFile := func(name string) {
		d := []byte("   ")
		check(os.WriteFile(name, d, 0644))

	}

	createEmptyFile("subdir/file1")

	err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)

	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")

	entries, err := os.ReadDir("subdir/parent")
	check(err)
	for _, dir := range entries {
		fmt.Println(dir.Name(), dir.IsDir())
	}

	err = os.Chdir("subdir/parent/child")
	check(err)

	c, err := os.ReadDir(".")
	check(err)
	for _, dir := range c {
		fmt.Println(dir.Name(), dir.IsDir())
	}

	err = os.Chdir("../../..")
	check(err)

	_, err = os.ReadDir(".")
	check(err)

	err = filepath.Walk("subdir", visit)
	check(err)
	time.Sleep(5 * time.Second)

}

func visit(p string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", p, info.IsDir())
	return nil
}

func Temp() {

	temp, err := os.CreateTemp("", "temp")
	defer os.Remove(temp.Name())
	check(err)
	fmt.Println(temp.Name())

	temp.Write([]byte{3, 1, 4, 5})
	temp.Sync()

	mkdirTemp, err := os.MkdirTemp("", "temp")
	defer os.RemoveAll(mkdirTemp)
	fmt.Println(mkdirTemp)

	fname := filepath.Join(mkdirTemp, "file1")
	err = os.WriteFile(fname, []byte{1, 2}, 0666)
	check(err)
}

//go:embed hello.txt
var fileString string

//go:embed hello.txt
var fileByte []byte

//go:embed hello1.txt
var folder embed.FS

func Embed() {

	fmt.Println(fileString)
	fmt.Println(string(fileByte))

	content1, _ := folder.ReadFile("folder/file1.hash")
	fmt.Println(string(content1))
	content2, _ := folder.ReadFile("folder/file2.hash")
	fmt.Println(string(content2))

}

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Args() {
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	arg := os.Args[3]
	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)

}

func Flags() {
	wordPtr := flag.String("word", "foo", "a string")

	numbPtr := flag.Int("numb", 42, "an int")
	forkPtr := flag.Bool("fork", false, "a bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *forkPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}

func Subcommand() {

	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnable := fooCmd.Bool("enable", false, "fooEnable")
	fooName := fooCmd.String("name", "", "fooName")

	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 5, "barLevel")

	args := os.Args
	if 2 > len(args) {
		fmt.Println("expect foo or bar subcommand")
		os.Exit(1)
	}

	switch args[1] {
	case "foo":
		fooCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'foo'")
		fmt.Println("  enable:", *fooEnable)
		fmt.Println("  name:", *fooName)
		fmt.Println("  tail:", fooCmd.Args())
	case "bar":
		barCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'bar'")
		fmt.Println("  level:", *barLevel)
		fmt.Println("  tail:", barCmd.Args())
	default:
		fmt.Println("subcommand not support")

	}

}

func SetEnv() {
	err := os.Setenv("gogo", "good")
	check(err)

	gogo := os.Getenv("gogo")
	fmt.Println(gogo)
	notExist := os.Getenv("notExist")
	fmt.Println(notExist)

	fmt.Println()
	for _, val := range os.Environ() {
		fmt.Println(val)
		keyVal := s.SplitN(val, "=", 2)
		fmt.Println(keyVal[0])

	}

}

func HttpClient() {

	response, err := http.Get("http://www.baidu.com")
	check(err)
	defer response.Body.Close()

	//all, err := io.ReadAll(response.Body)
	check(err)

	scanner := bufio.NewScanner(response.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(i, scanner.Text())
	}
	//fmt.Println(string(all))
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")

}

func get(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "get all")

}
func HttpService() {

	http.HandleFunc("/hello", handleHello)
	http.HandleFunc("/get", get)

	http.ListenAndServe(":9099", nil)
}
func ContextHandle(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	fmt.Println("start context")

	select {
	case <-time.After(3 * time.Second):
		fmt.Println("timeout done")
	case s1 := <-ctx.Done():
		err := ctx.Err()
		check(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println(s1)
	}

}
func Context() {

	http.HandleFunc("/hello", ContextHandle)
	http.ListenAndServe(":9099", nil)

}

func CmdProcess() {
	command := exec.Command("bash", "-c", "date")

	output, err := command.Output()
	fmt.Println(err)
	fmt.Println(string(output))

	gitVersion := exec.Command("bash", "-c", "date")
	gitVersionOutput, err := gitVersion.Output()
	fmt.Println(string(gitVersionOutput))

	grepCmd := exec.Command("findstr", "hello")

	inPipe, err := grepCmd.StdinPipe()
	outPipe, err := grepCmd.StdoutPipe()
	grepCmd.Start()

	inPipe.Write([]byte("hello grep\ngoodbye grep"))
	inPipe.Close()
	grepOutput, err := io.ReadAll(outPipe)
	grepCmd.Wait()
	fmt.Println("> grep hello")
	fmt.Println(string(grepOutput))

	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	check(err)
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}

func ExecProcess() {

	path, err := exec.LookPath("ls")
	check(err)
	//fmt.Println(path)

	//lsCommand := exec.Command("bash", "-c", "ls -a -l -h")
	//lsOutput, err := lsCommand.Output()
	//check(err)
	//fmt.Println(string(lsOutput))

	lsA := []string{"ls", "-a", "-l", "-h"}

	env := os.Environ()
	syscall.Exec(path, lsA, env)

}

func Signal() {

	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	done := make(chan bool, 1)
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}
