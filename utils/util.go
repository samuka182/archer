package util

import (
	"fmt"
	"strconv"
	"strings"
)

//StringToInt ...
func StringToInt(str string) int {
	nonFractionalPart := strings.Split(str, ".")
	valor, err := strconv.Atoi(nonFractionalPart[0])
	if err != nil {
		panic(err)
	}
	return valor
}

//	if response.PayloadIsArray() {
//		array := response.GetJSONArray()
//
//		printArray(array.Values)
//
//	} else {
//		object := response.GetJSONObject()
//
//		printMap(object.Value)
//	}
//

func printArray(array []map[string]interface{}) {
	for _, obj := range array {
		printMap(obj)
	}
}

func printArrayInterface(array []interface{}) {
	for _, obj := range array {
		printMap(obj.(map[string]interface{}))
	}
}

func printMap(m map[string]interface{}) {
	for k, v := range m {
		_, ok := v.([]interface{})

		if ok {
			printArrayInterface(v.([]interface{}))
		} else {
			_, ok2 := v.(map[string]interface{})
			if ok2 {
				printMap(v.(map[string]interface{}))
			} else {
				fmt.Println(fmt.Sprintf("Field: \t%s Value: \t%s", k, v))
			}
		}
	}
}
