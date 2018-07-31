package main

import "fmt"

type Cmember struct {
	name      string
	age       int
	address   string
	rank      int
	clearance string
}

func main() {
	// types of initialization
	cm := Cmember{name: "test", age: 20}
	fmt.Println("cm=", cm)

	var cm1 Cmember
	cm1.name = "test2"
	cm1.age = 25
	cm1.rank = 1
	cm1.address = "none"

	fmt.Println("cm1=", cm1)

	cm2 := Cmember{
		name: "test3",
		age:  24, //comma is necessary if the closing bracket is in the nextline
	}

	fmt.Println("cm2=", cm2)

	cmp := &cm2
	cmp.age = 20

	fmt.Println("cmp=", *cmp)

	var cmSlice []Cmember
	cmSlice = append(cmSlice, cm, Cmember{name: "test4", age: 30})

	for i, v := range cmSlice {
		fmt.Println(i, v)
	}

	var arr = []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("arr =", arr)
	var brr = arr[1:4]
	fmt.Println("brr =", brr)
	var crr = arr[:3]
	fmt.Println("crr =", crr)
	var drr = arr[3:]
	fmt.Println("drr =", drr)
	drr[0] = 15

	fmt.Println("drr =", drr)
	fmt.Println("arr =", arr) // original will also be modified

	var cMap map[string]Cmember
	cMap = make(map[string]Cmember)
	cMap["test"] = cm

	cMap = map[string]Cmember{
		"test10": Cmember{name: "test10", age: 10},
		"test20": Cmember{name: "test20", age: 20},
		"test40": Cmember{name: "test40", age: 40},
	}

	cMap["test30"] = Cmember{name: "test30", age: 30}

	elem := cMap["test30"]
	fmt.Println(elem)

	v, ok := cMap["test30"]
	fmt.Println(v, ok)

	delete(cMap, "test30")

	for k, v := range cMap {
		fmt.Println(k, v)
		v.isTeenager()
	}
}

func (c Cmember) isTeenager() {
	if c.age <= 20 {
		fmt.Println("member", c.name, " is a teenager")
	} else {
		fmt.Println("member", c.name, " is not a teenager")
	}

}
