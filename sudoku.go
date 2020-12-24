package main

import "fmt"

func main(){
	sudo:=[9][9]int{
		{7, 9, 2, 1, 5, 4, 3, 8, 6}, 
             {6, 4, 3, 8, 2, 7, 1, 5, 9},
             {8, 5, 1, 3, 9, 6, 7, 2, 4},
             {2, 6, 5, 9, 7, 3, 8, 4, 1},
             {4, 8, 9, 5, 6, 1, 2, 7, 3},
             {3, 1, 7, 4, 8, 2, 9, 6, 5},
             {1, 3, 6, 7, 4, 8, 5, 9, 2},
             {9, 7, 4, 2, 1, 5, 6, 3, 8},
             {5, 2, 8, 6, 3, 9, 4, 1, 7},
	}
	if check_sudoku(sudo){
		fmt.Println("The sudoku is correct")
	}else{
		fmt.Println("oops! check the sudoku again!!")
	}
}

func check_sudoku(sudo [9][9]int) bool{
		for i:=0;i<9;i++{
			for j:=0;j<9;j++{
				num:=sudo[i][j]								
				
				if isInRow(num,i,j,sudo){					
					return false
				}
				
				if isInCol(num,i,j,sudo){					
					return false
				}
				
				if isInBox(num,i,j,sudo){					
					return false
				}
								
			}
		}
		return true
}

func isInRow(num int,row int,col int,sudo [][]int) bool{
	for i:=col+1;i<9;i++{
		if sudo[row][i]==num{
			return true
		}		
	}
	return false
}

func isInCol(num int,row int,col int,sudo [][]int) bool{
	for i:=row+1;i<9;i++{
		if sudo[i][col]==num{
			return true
		}		
	}
	return false
}

func isInBox(num int,row int,col int,sudo [][]int) bool{
	r,c:=(row/3)*3,(col/3)*3	
	
	for i:=r;i<r+3;i++{
		for j:=c;j<c+3;j++{
			if (i!=row || j!=col) && sudo[i][j]==num{
				return true
			}
		}
	}
	return false
}