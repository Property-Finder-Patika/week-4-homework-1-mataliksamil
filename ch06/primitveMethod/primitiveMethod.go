package main

import (
	"errors"
	"fmt"
)

type PositiveInt int32
type String string

func main() {
	var i PositiveInt = 5
	fmt.Println(i.valid())
	i = -7
	fmt.Println(i.valid())

	var s String = "heyy, naber?"
	fmt.Println(s.whoAreYou())
}

//func (i int) whoAreYou() string {
//	return "I am a number nd my value is " + fmt.Sprintf(i)
//}
//
//func (s string) whoAreYou() string {
//	return "I am a number nd my value is " + fmt.Sprintf(i)
//}

func NewI(arg int32) (PositiveInt, error) {
	if arg > 0 {
		var i PositiveInt = PositiveInt(arg)
		return i, nil
	}
	return PositiveInt(-1), errors.New("oh you pass an invalid value")
}

func (i PositiveInt) valid() bool {
	if i > 0 {
		return true
	}
	return false
}

func (s String) whoAreYou() string {
	return "I am a string and my value is " + string(s) + ".\n"
}
