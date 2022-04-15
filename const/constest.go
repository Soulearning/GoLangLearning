package main

import(
	"fmt"
)

const(
	//可定义特殊变量iota,默认首行为0，后续每行+1.iota只会在枚举中递增
	aa,b = iota * 10, iota+5
	c,d
	e,f
	g,h = iota - 10,iota -10
	j,k
)

func main()  {
	const a int = 10
	fmt.Println("a = ",a)
	fmt.Printf("aa = %d, b=%d\n",aa,b)
	fmt.Printf("c = %d, d=%d\n",c,d)
	fmt.Printf("e = %d, f=%d\n",e,f)
	fmt.Printf("g = %d, h=%d\n",g,h)
	fmt.Printf("j = %d, k=%d\n",j,k)
}