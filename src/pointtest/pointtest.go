package pointtest

import "fmt"

func  Ptest()  {
	var p int=5
	var p1 *int=&p
	var p2 **int=&p1

	fmt.Println(p,p1,p2)
	fmt.Println(p,*p1,*p2,**p2)
}
