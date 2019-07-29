package main

import (
	"fmt"
)

func Names() (string, string){
	return "Foo","Bar"
}
func  main()  {
	n1,n2:=Names()
	fmt.Println(n1,n2)

	n3,n5:=Names()
	fmt.Println(n3)
    fmt.Println(n5)
}

