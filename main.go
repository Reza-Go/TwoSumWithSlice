package main 

import "fmt"

    



func main (){

	nums := []int { 3 , 5 ,5 , 0}
	target := 10 
    result := twoSum(nums , target)
    fmt.Println(result)

}

func twoSum(numbers []int , target int) []int {

	for i , item :=range numbers{

		complement := target - item 
		for j := i+1 ; j <len(numbers) ; j ++ {
			if i != j && complement == numbers[j] {
				return []int{i , j }
			}
		}

	
	}

	return []int{0,0}
}