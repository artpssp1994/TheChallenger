package main

import (
	"fmt"
)

var resultMap = make(map[string] []int)

func main(){
	setSubstringToMap("GAGGCATCATCATCATCATTGCCCATCAT")
	findTandemFromMap()

}

func setSubstringToMap(input string){
	

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
	// fmt.Println(resultMap)
	
}

func setIgnoreValue(key string){
	ignoreKey := key
	for len(ignoreKey) + len(key) <= 10 {
		ignoreKey += key
		resultMap[ignoreKey] = []int{-1}
	}
		
}

func findTandemFromMap(){
	for k, mapVal := range resultMap {
		var startIndex int
		var repeatedTimes int
		if mapVal[0] != -1 && len(mapVal) > 1{
			// fmt.Println("key : ", k, mapVal)
			for i,repeatIndex := range mapVal {
				if startIndex == 0 {
					startIndex = repeatIndex
					repeatedTimes = 1
				}else{
					if (repeatIndex - mapVal[i-1]) == len(k) {
						repeatedTimes += 1
					}else{
						if repeatedTimes > 1{
							fmt.Println(k,startIndex,repeatedTimes)
						}
						startIndex = repeatIndex
						repeatedTimes = 1
					}
				}

			}
		}
		if repeatedTimes > 1{
			fmt.Println(k,startIndex,repeatedTimes)
		}
	}
}