package main

import (
	"fmt"
	"runtime"
	"strconv"
)

const (
	A = 5
	B = A << 1
	C = A >> 1
)

func main() {
	sayHello("test", 1)
	fmt.Println(add(5, 3))
	fmt.Println(addSubtactMultiply(5, 3))
	fmt.Println(addSubtactMultiplyV2(5, 3))

	addV2 := func(a, b int) int {
		return a + b
	}
	fmt.Println(addV2(5, 3))
	fmt.Println(multiplyWithAdd(3, addV2))

	inc := incrementer()
	fmt.Println(inc(), inc(), inc(), inc())
	//arrays
	var arr [3]int
	arr[0], arr[1] = 3, 4
	fmt.Println("arr = ", arr, len(arr))

	brr := [3]int{3, 4}
	fmt.Println("brr = ", brr, len(brr))

	//slice
	srr := []int{}
	fmt.Println("srr length =  ", len(srr))
	srr = append(srr, 4)
	srr = append(srr, 5)

	fmt.Println("srr =", srr, "len =", len(srr))

	//pointers
	var p *int
	i := 4
	p = &i
	// change(&p) both are same
	change(p)

	fmt.Println(*p)

	//type conversions
	n := 32
	f := float64(n)
	fmt.Println(f)

	//constants
	const pi = 3.14
	fmt.Println(pi * f)

	fmt.Println("A binary = ", strconv.FormatInt(A, 2))
	fmt.Println("B binary = ", strconv.FormatInt(B, 2))
	fmt.Println("C binary = ", strconv.FormatInt(C, 2))
	fmt.Println("A =", A)
	fmt.Println("B =", B)
	fmt.Println("C =", C)

	if i = inc(); i < 0 {
		fmt.Println("i is negative")
	} else if i == 0 {
		fmt.Println("i is zero")
	} else {
		fmt.Println("i is positive")
	}

	switch i = inc(); {
	case i < 0:
		fmt.Println("i is negative")
	case i == 0:
		fmt.Println("i is zero")
	default:
		fmt.Println("i is positive")
	}
	// defer
	defer fmt.Println("Defer last execution statement")

	switch os := runtime.GOOS; os {
	//experimental
	case "osx":
	case "darwin":
		fmt.Println("Mac OSX")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Println("Unknown")
	}

	//loops

	for k := 10; k > 0; k-- {
		fmt.Print(k)
	}
	fmt.Println()

	n = 1
	for n < 10 {
		n++
	}

	fmt.Println(n)

	for i = 1; i <= 4; i++ {
		defer fmt.Println("Defer call =", i)
	}

	fmt.Println("End of Function")
}

func change(i *int) {
	*i++
}

func multiplyWithAdd(a int, fn func(a, b int) int) int {
	return a * fn(a, a)
}
func sayHello(s string, i int) {
	fmt.Println(s, i)
}

func add(a, b int) int {
	return a + b
}

func addSubtactMultiply(a, b int) (addition, subtraction, multiply int) {
	addition = add(a, b)
	subtraction = a - b
	multiply = a * b
	return
}

func addSubtactMultiplyV2(a, b int) (int, int, int) {
	return add(a, b), a - b, a * b
}

func incrementer() func() int {
	i := 0

	return func() int {
		i++
		return i
	}
}
