package point_learning

import(
	"fmt"
)

func init()  {
	fmt.Println("Excuted pointTest Init!")
}

/**
	a *int :表示a是一个指针类型的变量
	&a :表示a的地址
	*a :表示访问指针指向的变量的值
*/

func Out(){
	fmt.Println(">>>>>>>>>>>>>>>>>>>>")
	var a int=1
	fmt.Printf("初始值：%d\n",a)
	pA(&a)
	fmt.Println(&a)
	fmt.Printf("调用pA后：%d\n",a)
	fmt.Println(&a)
	nA(a)
	fmt.Printf("调用nA后：%d\n",a)

	var f int = 100
	var l int = 200

	fmt.Printf("调用swap前：f = %d,l = %d\n",f,l)
	swap(&f,&l)
	fmt.Printf("调用swap后：f = %d,l = %d\n",f,l)

}

func pA(a *int) {
	*a=10
}

func nA(a int) {
	a = 20
}

func swap(a,b *int){
	c := *a
	*a = *b
	*b = c
}
