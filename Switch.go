package main

import "fmt"

func main(){

	num:=11111100
	
	zeros:=findMaxConsecutiveZero(num)
	ones:=findMaxConsecutiveOne(num)
	
	switch{
		case (zeros>=3 && isOdd(num)):
			fmt.Println("000-ODD")
			
		case (zeros<3 && isOdd(num)):
			fmt.Println(">3 0s-ODD")
			
		case (ones>=5 && !isOdd(num)):
			fmt.Println(">5 1s-EVEN")
	}
}

func findMaxConsecutiveZero(num int) int{
	max,count:=0,0
	for(num>0){
		bit:=num%10
		num=int(num/10)
		if(bit==0){
			count = count + 1
		}else{
			count=0
		}			
		if(max<count){
			max=count
		}		
	}
	return max
}

func findMaxConsecutiveOne(num int) int{
	max,count:=0,0
	for(num>0){
		bit:=num%10
		num=int(num/10)
		if(bit==1){
			count = count + 1
		}else{
			count=0
		}			
		if(max<count){
			max=count
		}		
	}
	return max
}

func isOdd(num int) bool{
	if(num%10==1){
		return true
	}else{
		return false
	}	
}

