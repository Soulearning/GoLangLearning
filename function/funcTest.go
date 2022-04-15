package main

import (
	"fmt"
)

func main()  {
	fmt.Printf("单返回值：%d\n",a())
	r1,r2 := b()
	fmt.Printf("多返回值：%d,%s\n",r1,r2)
}

func a() int {
	return 10
}

func b() (int,string)  {
	return 100,"TEST"
}