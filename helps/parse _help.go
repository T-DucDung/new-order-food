package helps

import "encoding/json"

func ToString(data []byte) string {
	return string(data[:])
}

func ToInt(data []byte) (int, error) {
	var num int
	err := json.Unmarshal(data, &num)
	return num, err
}

func ToInt64(data []byte) (int64, error) {
	var num int64
	err := json.Unmarshal(data, &num)
	return num, err
}

func ToFloat32(data []byte) (float32, error) {
	var num float32
	err := json.Unmarshal(data, &num)
	return num, err
}

func ToFloat64(data []byte) (float64, error) {
	var num float64
	err := json.Unmarshal(data, &num)
	return num, err
}

func ToPrint(data map[string][]byte) (map[string]string, error) {
	mapString := map[string]string{}
	for k, v := range data {
		mapString[k] = string(v[:])

	}
	return mapString, nil
}

func ToPrints(data []map[string][]byte) ([]map[string]string, error) {
	mapString := []map[string]string{}
	for _, v := range data {
		mapStr, err := ToPrint(v)
		if err != nil {
			return nil, err
		}
		mapString = append(mapString, mapStr)
	}
	return mapString, nil
}
