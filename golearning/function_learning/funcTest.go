package function_learning

import (
	"fmt"
)
/**
	执行流程：import --> const --> const --> var --> init() --> main() --> exit
**/

func init()  {
	
}

func Out()  {
	fmt.Printf("单返回值：%d\n",a())
	r1,r2 := b()
	fmt.Printf("多返回值：%d,%s\n",r1,r2)
	r3,r4 := c()
	fmt.Printf("多返回值：%d,%d\n",r3,r4)
}

func a() int {
	return 10
}

func b() (int,string)  {
	return 100,"TEST"
}

func c() (r3 int,r4 int) {
	r3=100
	r4=500
	return
}