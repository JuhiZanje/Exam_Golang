package main

import "fmt"

func main(){
	var disk int
	fmt.Println("Enetr the number of disk:")
	fmt.Scan(&disk)
	towerOfHanoi(disk,"A","C","B")
}

func towerOfHanoi(disk int,rF,rD,rM string){
	if (disk == 1){
		fmt.Printf("Move disk:1 from rod:%v to rod:%v\n",rF,rD)
        return
	}
	towerOfHanoi(disk-1,rF,rM,rD)
	fmt.Printf("Move disk:%d from rod:%v to rod:%v\n",disk,rF,rD)
	towerOfHanoi(disk-1,rM,rD,rF)
}