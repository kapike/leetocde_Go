package main

import "fmt"

func _findNumberIn2DArray(matrix [][]int, target int, width int) bool{
	if width == 0 || len(matrix)==0{
		return false
	}
	temp := matrix[0][width-1]
	if target == temp{
		return true
	}else if temp > target{
		return _findNumberIn2DArray(matrix,target,width-1)
	}else{
		return _findNumberIn2DArray(matrix[1:],target,width)
	}
}
func findNumberIn2DArray(matrix [][]int, target int) bool {

	heigh := len(matrix)
	if heigh == 0{
		return  false
	}
	width := len(matrix[0])

	return _findNumberIn2DArray(matrix,target,width)
}

func main()  {
	fmt.Println(findNumberIn2DArray())
}
