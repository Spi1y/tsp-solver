package main

import "fmt"

func linterDemo(condition1, condition2 bool) {
	if condition1 {
		fmt.Println("cond1")
	} else if condition1 {
		fmt.Println("cond2")
	}
}
