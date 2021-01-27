package helps

import (
	"encoding/json"
	"fmt"
)

func ByteToString(data []byte) string {
	return string(data[:])
}

func ByteToInt(data []byte) (int, error) {
	var num int
	err := json.Unmarshal(data, &num)
	return num, err
}

func ByteToInt64(data []byte) (int64, error) {
	var num int64
	err := json.Unmarshal(data, &num)
	return num, err
}

func ByteToFloat32(data []byte) (float32, error) {
	var num float32
	err := json.Unmarshal(data, &num)
	return num, err
}

func ByteToFloat64(data []byte) (float64, error) {
	var num float64
	err := json.Unmarshal(data, &num)
	return num, err
}

func ByteToPrint(data map[string][]byte) map[string]string {
	mapString := map[string]string{}
	for k, v := range data {
		mapString[k] = string(v[:])
	}
	return mapString
}

func ByteToPrints(data []map[string][]byte) []map[string]string {
	mapString := []map[string]string{}
	for _, v := range data {
		mapString = append(mapString, ByteToPrint(v))
	}
	return mapString
}

func InterToString(data map[string]interface{}) map[string]string {
	mapString := map[string]string{}
	for k, v := range data {
		mapString[k] = fmt.Sprintf("%v", v)
	}
	return mapString
}
