package main

import (
	"fmt"
)

type set map[string]struct{}

func main() {
	s := make(set)
	s["item1"] = struct{}{}
	s["item1"] = struct{}{}
	s["item2"] = struct{}{}
	s["item2"] = struct{}{}
	s["item3"] = struct{}{}

	fmt.Println(s.getSetValues())
}

func (s set) getSetValues() []string {
	var setValues []string

	for k := range s {
		setValues = append(setValues, k)
	}

	return setValues

}
