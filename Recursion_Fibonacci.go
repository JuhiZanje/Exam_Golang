package main

import "fmt"

func fibonacci(num int) int{
	
	defer fmt.Printf("iteration for=%d\n",num)
	if (num<=1){		
		return num		
	}		
	return fibonacci(num-2)+fibonacci(num-1)
}


func main(){	
	n:=5	
	for i:=0;i<n;i++{
		fmt.Println(fibonacci(i)," ")
	}
}