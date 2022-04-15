package vars_learning
/*
	四种变量声明方式
*/


import(
	"fmt"	
)

func Out(){
	// 1、无赋值声明
	var a int
	fmt.Printf("a = %d, a's type is %T\n",a,a)

	//2、正常声明
	var b,c int = 3,4
	fmt.Printf("b = %d, b's type is %T\n",b,b)
	fmt.Printf("c = %d, c's type is %T\n",c,c)

	//3、自动推断
	var d = "zifuchuang"
	fmt.Printf("d = %s, d's type is %T\n",d,d)

	//4、最简赋值
	//注意，此种方式无法对全局变量赋值
	f := 1.52
	fmt.Printf("f = %f, f's type is %T\n",f,f)

}
