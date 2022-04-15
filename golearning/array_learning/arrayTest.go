package array_learning

import(
	"fmt"
)

func init()  {
	fmt.Println("Excuted arrayTest Init!")
}

func Out() {
	fmt.Println(">>>>>>>>>>>>>>>>>>>>")
	/**
		静态数组，传入函数，相当于值拷贝，不对入参发生修改
	*/
	var a [3]int = [3]int{}
	fmt.Println("printStaticArray调用之前：")
	printStaticArray(a)
	fmt.Println("printStaticArray调用之后：")
	printStaticArray(a)

	/**
		动态数组，相当于传入引用
	*/
	b := []int{1,2,3}
	fmt.Println("printArray调用之前：")
	printArray(b)
	fmt.Println("printArray调用之后：")
	printArray(b)
	
}

func printStaticArray(array [3]int){
	for index,value := range array{
		fmt.Printf("index = %d,values = %d\n",index,value)
	}
	array[0] = 100
}

func printArray(array []int){
	for index,value := range array{
		fmt.Printf("index = %d,values = %d\n",index,value)
	}
	array[0] = 100
}


