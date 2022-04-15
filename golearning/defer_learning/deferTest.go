package defer_learning


import(
	"fmt"	
)

func init()  {
	fmt.Println("Excuted defertTest Init!")
}

/**
	defer在函数中，在函数生命周期结束之前执行，多个 defer 进行压栈操作，先声明，后执行
	再有返回值的函数时，先执行返回操作，在执行 defer
*/
func Out(){
	fmt.Println(">>>>>>>>>>>>>>>>>>>>")
	defer fmt.Println("第一次定义defer")
	defer fmt.Println("第二次定义defer")
	deferT()
}

func deferT() int {
	defer fmt.Println("执行 defer 操作！")
	return number()
}

func number() int {
	fmt.Println("执行返回值操作！")
	return 0
}