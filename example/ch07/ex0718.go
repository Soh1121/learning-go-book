package main

import "io"

func doTypeSwitch(i any) {
	switch j := i.(type) {
	case nil:
		// iはnil
	case int:
		// jの型はint
	case MyInt:
		// jの形はMyInt
	case io.Reader:
		// jの型はio.Reader
	case string:
		// jの型はstring
	case bool, rune:
		// iはブール値かrune
	default:
		// その他
	}
}
