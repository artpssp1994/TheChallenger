package main

import (
	"fmt"
)

var resultMap = make(map[string] []int)

func main(){
    Tandem("GAGGCATCATCATCATCATTGCC")
}

func Tandem(input string){
	

	for i := range input {
		if i + 3 < len(input) + 1 {
			frameValue := input[i:i+3]
			if val, ok := resultMap[frameValue]; ok {
				if val[0] != -1 {
					resultMap[frameValue] = append(resultMap[frameValue],i)
					setIgnoreValue(frameValue)
				}
			}else{
				resultMap[frameValue] = []int{i}
				setIgnoreValue(frameValue)
			}
			// fmt.Println(frameValue)
			
		}
	}
	fmt.Println(resultMap)
	
}

func setIgnoreValue(key string){
	ignoreKey := key
	for len(ignoreKey) + len(key) <= 10 {
		ignoreKey += key
		resultMap[ignoreKey] = []int{-1}
	}
		
}