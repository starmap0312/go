package main

import (
	"fmt"
	"math"

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
	return v.x + v.y
}

// struct method: Go does not have classe, but you can define methods on struct
func (v *Vertex) double() { // define a method of Vertex v
	// note that use a pointer receiver (*Vertex) so that the method can change the Vertex value
	v.x = v.x * 2
	v.y = v.y * 2
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

	// struct method: Go does not have classe, but you can define methods on struct
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

	// Slices: it is like a reference to the underlying array, i.e. a slice does not store any data
	//   an array has a fixed size; a slice is a dynamically-sized view into the elements of an array
	nums := [6]int{1, 2, 3, 4, 5, 6}          // an array of integers
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

	// Function values: pass a function just like other values
	var result = double(add, 1, 2)
	fmt.Println(result) // add(1, 2) * 2 = 6

	// Function closures: a function value that references variables from outside its body
	fn := addOne()
	fmt.Println(fn(10)) // 11

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
