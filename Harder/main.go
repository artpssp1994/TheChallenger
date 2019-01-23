package main

import (
	"fmt"
	"strconv"
	"encoding/json"
	"bufio"
	"os"
)

func main(){
	//Receive input from console
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')

	//Call tandem to retrive map that have 
	//tandem word as key and repeated times as value
	ansMap := tandem(text)

	//Set response to json format
	j, _ := json.Marshal(ansMap)
	fmt.Println(string(j))
}

//tandem - Receive text string and return answer map
func tandem(text string) map[string]int {
	mSubStringIndex := createMapSubStringAndIndex(text)
	ansMap := findTandemFromMap(mSubStringIndex)
	return ansMap
}

//createMapSubStringAndIndex - Receive text string and 
//loop in to find Tandem repeat with from size 3 to 10
//return map with tandem word as keys and list of start index of as values
func createMapSubStringAndIndex(input string) map[string] []int{
	mSubStringIndex := make(map[string] []int)
	//loop in to find Tandem repeat from size 3 to 10
	for size := 3; size <= 10; size++ {
		for i := range input {
			if i + size < len(input) + 1 {
				frameValue := input[i:i+size]
				if val, ok := mSubStringIndex[frameValue]; ok {
					if val[0] != -1 {
						mSubStringIndex[frameValue] = append(mSubStringIndex[frameValue],i)
						setIgnoreValue(frameValue,&mSubStringIndex)
					}
				}else{
					mSubStringIndex[frameValue] = []int{i}
					setIgnoreValue(frameValue,&mSubStringIndex)
				}
			}
		}
	}
	return mSubStringIndex
}

//setIgnoreValue - Set tandem word that we dont need to count
//Example case : we count CAT but we dont count CATCAT
func setIgnoreValue(key string, mSubStringIndex *map[string] []int){
	ignoreKey := key
	for len(ignoreKey) + len(key) <= 10 {
		ignoreKey += key
		(*mSubStringIndex)[ignoreKey] = []int{-1}
	}
		
}

//findTandemFromMap - Loop in tadem map and return count of repeat times for each start index
func findTandemFromMap(mSubStringIndex map[string] []int) map[string]int {
	ansMap := make(map[string]int)
	for k, mapVal := range mSubStringIndex {
		var startIndex int
		var repeatedTimes int
		if mapVal[0] != -1 && len(mapVal) > 1{
			for i,repeatIndex := range mapVal {
				if startIndex == 0 {
					startIndex = repeatIndex
					repeatedTimes = 1
				}else{
					if (repeatIndex - mapVal[i-1]) == len(k) {
						repeatedTimes += 1
					}else{
						if repeatedTimes > 1{
							formatKey := setTandemKeyFormat(startIndex,k)
							ansMap[formatKey] = repeatedTimes
						}
						startIndex = repeatIndex
						repeatedTimes = 1
					}
				}

			}
		}
		if repeatedTimes > 1{
			formatKey := setTandemKeyFormat(startIndex,k)
			ansMap[formatKey] = repeatedTimes
		}
	}
	return ansMap
}

//SsetTandemKeyFormat - Set key of json out put format
func setTandemKeyFormat(startIndex int, tandem string) string{
	s := strconv.Itoa(startIndex)
	ans := s +"-"+ tandem
	return ans
	
}