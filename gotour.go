package main

import (
	"fmt"
	"math"
	"time"

	"rsc.io/quote"
) // use external package

func HelloWorld() {
	fmt.Println("Hello, World")
}

func HelloWorldWithQuote() {
	fmt.Println(quote.Go())
}

// a function
func add(x int, y int) int {
	return x + y
}

// a function that returns two results
func swap(x string, y string) (string, string) {
	return y, x
}

// a function that returns two named return values
func split(input int) (x int, y int) {
	x = input * 2
	y = input / 2
	return // this returns the two named return values (a.k.a naked return)
}

// Fuction values: pass a function just like other values
func double(fn func(int, int) int, x int, y int) int {
	return fn(x, y) * 2
}

// Function closures: a function value that references variables from outside its body
func addOne() func(int) int {
	sum := 1
	return func(x int) int {
		sum += x // the function is "bound" to the variables `sum`
		return sum
	}
}

// declare variables
var i, j int = 1, 2

// declare constants
const Pi = 3.14
const ConstantStr = "HELLO"

// declare variables in a block
var (
	isBool  bool   = false
	isNum   int    = 1
	isStr   string = "abc"
	isBool2 bool   // default value is false
	isNum2  int    // default value is 0
	isStr2  string // default value is ""
)

// declare constants in a block
const (
	Big   = 1 << 10  // shifting a 1 bit left 10 places, i.e. binary 2^10
	Small = Big >> 9 // i.e. 1 << 1 = 2
)

// struct: a collection of fields
type Vertex struct { // a named struct
	x int
	y int
}

// struct method: Go does not have classe, but you can define methods on struct
func (v Vertex) sum() int { // define a method of Vertex v
	// note that the method operates on a copy of the original Vertex value
	// i.e. this will copy the value on each method call (less efficient)
	return v.x + v.y
}

// struct method: Go does not have class, but you can use struct's pointer receiver to modify the struct's fields
func (v *Vertex) double() { // define a method of Vertex v
	// note that if you do not use a pointer receiver (*Vertex), the method will operate on a copy of the original Vertex value
	v.x = v.x * 2
	v.y = v.y * 2
}

// interfaces: an interface type is defined as a set of method signatures
type Animal interface {
	run()
}

type Dog struct {
	name string
}

// type Dog implements the interface Animal, but we don't need to explicitly declare that it does so
//   instead, we do that by declaring a Dog instance as type interface Animal, which will enforce the Dog type to implement the interface method
//   ex. var dog Animal = Dog{"john's dog"}
func (d Dog) run() {
	fmt.Println(d.name + " is running.")
}

type DogPtr struct {
	name string
}

func (d *DogPtr) run() {
	fmt.Println(d.name + " is running.")
}

// built-in interface: fmt.Stringer
//   type Stringer interface {
// 	   String() string
//   }
// implement built-in interface: fmt.Stringer
type IPAddr [4]byte

func (v IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", v[0], v[1], v[2], v[3])
}

// built-in interface: error
//   type error interface {
//   	Error() string
//   }
// implement built-in interface: error
type MyError struct {
	msg string
}

func (err *MyError) Error() string {
	return fmt.Sprintf("message: %s", err.msg)
}

func funcThrowsMyError() error {
	return &MyError{"error inside funcThrowsMyError"}
}

// goroutine: a lightweight thread managed by the Go runtime
func heavyTask(elapsed int) {
	fmt.Printf("started heavy task: %d\n", elapsed)
	time.Sleep(time.Duration(elapsed*1000) * time.Millisecond)
	fmt.Printf("finished heavy task: %d\n", elapsed)
}

// channels: a typed conduit through which you can send and receive values with the channel operator, <-
// ex. ch <- v    // send v to channel ch
//     v := <-ch  // receive from ch and assign the value to v
func send(elapsed int, ch chan int) {
	fmt.Printf("started sending: %d\n", elapsed)
	time.Sleep(time.Duration(elapsed*1000) * time.Millisecond)
	ch <- elapsed // send 100 to ch
	fmt.Printf("finished sending: %d\n", elapsed)
}

func main() {
	HelloWorld()          // Hello, World
	HelloWorldWithQuote() // Don't communicate by sharing memory, share memory by communicating.

	fmt.Println(math.Pi)        // 3.141592653589793
	fmt.Println(add(1, 2))      // 3
	fmt.Println(swap("a", "b")) // b a
	fmt.Println(split(2))       // 4 1

	// type inference: a short hand for function level declaration
	k := 3   // type inferenced
	f := 1.2 // type inferenced
	// note: on package leve, every statement must begin with a keyword, ex. var, func, etc.
	fmt.Println(i, j, k, f, Pi, ConstantStr, Big, Small)       // 1 2 3 1.2 3.14 HELLO 1024 2
	fmt.Println(isBool, isNum, isStr, isBool2, isNum2, isStr2) // false 1 abc false 0

	// for-loop: no () parentheses in Go
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum) // 45

	// while-loop: use for in Go
	for sum < 1000 {
		sum *= 2
	}
	fmt.Println(sum) // 1440

	// if-condition: no () parentheses in Go
	if sum < 1000 {
		fmt.Println(sum) // not printed
	}

	// switch: unlike C++/Java, it will break automatically at the end of each case
	x := 100
	switch x % 2 {
	case 0:
		fmt.Println("is even") // is even
	default:
		fmt.Println("is odd")
	}

	// pointer
	var p *int // default: nil pointer
	//fmt.Println(*p) // panic: runtime error: invalid memory address or nil pointer dereference
	i := 42
	p = &i
	fmt.Println(*p) // 42, dereference, read i through the pointer p

	var v = Vertex{1, 2}
	fmt.Println(v)        // {1 2}
	fmt.Println(v.x, v.y) // 1 2
	v.x = 10
	fmt.Println(v) // {10 2}

	// struct method: Go does not have class, but you can define methods on struct
	fmt.Println(v.sum()) // 12
	v.double()
	fmt.Println(v) // {20 4}

	// struct fields can be accessed through a struct pointer
	var u *Vertex = &v
	fmt.Println(*u)     // {20 4}
	fmt.Println((*u).x) // 20
	// note that Go provides syntax suger for dereferencing a struct pointer, i.e. * can be omitted
	fmt.Println(u)   // {20 4}
	fmt.Println(u.x) // 20

	// struct fields can be partially constructed, with default values assigned to fields
	var v2 = Vertex{}
	fmt.Println(v2) // {0 0}
	var v3 = Vertex{x: 2}
	fmt.Println(v3) // {2 0}

	// Arrays
	var arr [10]int  // note: an array's length is part of its type, so arrays cannot be resized
	fmt.Println(arr) // [0 0 0 0 0 0 0 0 0 0]

	// 2-dimensional array
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	fmt.Println(board) // [[_ _ _] [_ _ _] [_ _ _]]

	// iterate an array
	nums := [6]int{1, 2, 3, 4, 5, 6} // an array of integers
	for i, v := range nums {
		fmt.Println(i, v)
	}
	// 0 1
	// 1 2
	// 2 3
	// 3 4
	// 4 5
	// 5 6

	// Slices: it is like a reference to the underlying array, i.e. a slice does not store any data
	//   an array has a fixed size; a slice is a dynamically-sized view into the elements of an array
	fmt.Println(nums)                         // [1 2 3 4 5 6]
	fmt.Println(nums[1:])                     // [2 3 4 5 6]
	fmt.Println(nums[:1])                     // [1]
	fmt.Println(cap(nums[1:]), len(nums[1:])) // capacity=5, length=5

	var subarr []int = nums[1:4] // a slice of the array
	fmt.Println(subarr)          // [2 3 4]
	subarr[0] = -1
	fmt.Println(subarr) // [-1 3 4]
	fmt.Println(nums)   // [1 -1 3 4 5 6]: note the original array is changed when the slice's element is changed

	// Appending to a slice
	subarr2 := append(subarr, 5)
	// note: if subarr is too small to fit the appended value, a bigger array will be allocated & returned
	fmt.Println(subarr)  // [-1 3 4]
	fmt.Println(subarr2) // [-1 3 4 5]

	arrStruct := []struct { // an array of struct: an anonymous struct
		i int
		b bool
	}{
		{1, true},
		{2, false},
		{3, true},
	}
	fmt.Println(arrStruct) // [{1 true} {2 false} {3 true}]

	// Map
	var mp = map[int]string{
		1: "one",
		2: "two",
	}
	fmt.Println(mp) // map[1:one 2:two]
	mp[1] = "uno"   // mutating a Map
	fmt.Println(mp) // map[1:uno 2:two]

	// make vs. new: memory allocation and value initialization
	var newNum []int = make([]int, 5, 10) // newNum has type: []int, allocate an array of integers of length=5 & capacity=10
	// make() takes a type T, which must be a slice, map or channel type, ex. int is not allowed
	fmt.Println(newNum, len(newNum), cap(newNum)) // [0 0 0 0 0] 5 10
	// the above make is equivalent to:
	var newNum2 []int = new([10]int)[0:5]
	fmt.Println(newNum2, len(newNum2), cap(newNum2)) // [0 0 0 0 0] 5 10

	var newNumPtr *int = new(int) // newNumPtr has type: *int
	fmt.Println(*newNumPtr)       // 0

	// iterate a Map
	mp2 := map[string]int{
		"key1": 1,
		"key2": 2,
	}
	for k, v := range mp2 {
		fmt.Printf("%v: %v\n", k, v) // key1: 1 & key2: 2
	}

	// Function values: pass a function just like other values
	var result = double(add, 1, 2)
	fmt.Println(result) // add(1, 2) * 2 = 6

	// Function closures: a function value that references variables from outside its body
	fn := addOne()
	fmt.Println(fn(10)) // 11

	// interfaces
	var dog Animal = Dog{"john's dog"}
	fmt.Printf("(%v, %T)\n", dog, dog) // ({john's dog}, main.Dog)
	// note: interface values can be thought of as a tuple of a value and a concrete type: (value, type)
	dog.run() // john's dog is running

	var dogPtr Animal = &DogPtr{"john's dogPtr"}
	// note: if you do not wirte &DogPtr here, you will receive the following compile error:
	//   DogPtr does not implement Animal (run method has pointer receiver)
	fmt.Printf("(%v, %T)\n", dogPtr, dogPtr) // (&{john's dogPtr}, *main.DogPtr)
	dogPtr.run()                             // john's dog is running

	var nullPtr Animal
	// nullPtr.run()
	// note: if call a method on a nil interface, you will receive the following run-time error:
	//   panic: runtime error: invalid memory address or nil pointer dereference
	fmt.Printf("(%v, %T)\n", nullPtr, nullPtr) // (<nil>, <nil>)
	// note: the nil interface has no concrete type inside the interface tuple to indicate which concrete method to call

	// empty interface: it specifies zero methods
	var emptyItf interface{}
	fmt.Printf("(%v, %T)\n", emptyItf, emptyItf) // (<nil>, <nil>)

	// type assertion: it provides access to an interface value's underlying concrete value
	var itf interface{} = "abc"
	value := itf.(string) // access the underlying concrete value as type string
	fmt.Println(value)    // abc

	value, ok := itf.(string)
	fmt.Println(value, ok) // abc true

	fail, ok := itf.(float64) // fails to access the underlying concrete value as type float64
	fmt.Println(fail, ok)     // 0 false

	// type switch: it is a construct that permits several type assertions in series.
	switch v := itf.(type) {
	case string:
		fmt.Printf("string type matched: value=%v, type=%T\n", v, v) // string type matched: value=abc, type=string
	case float64:
		fmt.Printf("float64 type matched: value=%v, type=%T\n", v, v)
	default:
		fmt.Printf("no type matched: value=%v, type=%T\n", v, v)
	}

	ip := IPAddr{127, 0, 0, 1}
	fmt.Printf("%v\n", ip) // 127.0.0.1

	err := funcThrowsMyError()
	fmt.Println(err) // message: error inside funcThrowsMyError

	// goroutine: a lightweight thread managed by the Go runtime
	go heavyTask(3) // this starts a new goroutine running
	time.Sleep(100 * time.Millisecond)
	heavyTask(1)
	// started heavy task: 3
	// started heavy task: 1
	// finished heavy task: 1
	// finished heavy task: 3
	time.Sleep(3000 * time.Millisecond)

	// channels: a typed conduit through which you can send and receive values with the channel operator, <-
	ch := make(chan int)
	go send(2, ch)
	go send(3, ch)
	go send(1, ch)
	x, y, z := <-ch, <-ch, <-ch // receive from ch
	// note: sends and receives block until the other side is ready
	//       this allows goroutines to synchronize without explicit locks or condition variables

	fmt.Println(x, y, z)
	// started sending: 1
	// started sending: 2
	// started sending: 3
	// finished sending: 1
	// finished sending: 2
	// finished sending: 3
	// 1 2 3

	// defer: deferred function calls are pushed onto a stack
	//        when a function returns, its deferred calls are executed in last-in-first-out order
	fmt.Println("begin")
	for i := 0; i < 3; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("end")
	// begin, end, 2, 1, 0
}

// go mod init gotour // enable dependency tracking for your code by creating a go.mod file
// go run gotour.go
// go build gotour.go & ./gotour
