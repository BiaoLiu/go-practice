package main

import "fmt"

func panic_test2(){

}

func recover_test(){
	defer func() {
		if r:=recover();r!=nil{

		}
	}()
}

func defer_test2(){
	fmt.Printf("defer_test2...")
}