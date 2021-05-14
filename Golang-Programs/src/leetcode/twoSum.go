// index of two numbers in array who sum up to target
package main

import "fmt"

func main() {
	arr := []int {1,2,3,4,5}
	var sum int = 7

	fmt.Println("Given array:",arr)
	fmt.Printf("Give a (valid) target: ")
	fmt.Scanf("%d", &sum)
	
	fmt.Println("\nTarget:",sum)

	result := twoSum(arr,sum)
	fmt.Println("TwoSum solution array:",result)
}

func twoSum(nums []int, target int) []int {
    var pos = []int {-1,-1}

	for i, val := range nums {
		if val <= target {
            
            pos[0] = i
            for j := i+1; j < len(nums); j++ {
                
                if nums[j] + nums[pos[0]] == target {
                    pos[1] = j
                    goto END
                }
                
            }
            
        }
	}
    
    END:
    return pos
}