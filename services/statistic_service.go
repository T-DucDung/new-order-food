package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"new-order-food/models"
	"new-order-food/responses"
)

var N = "http://192.168.0.104:5000"

func TotalRevenu() (float32, error) {
	s := models.Statistics{}
	return s.TotalRevenu()
}

func TotalOrder() (int, error) {
	s := models.Statistics{}
	return s.TotalOrder()
}

func TopSale(num int) ([]responses.ProductRes, error) {
	s := models.Statistics{}
	return s.TopSale(num)
}

func TotalNewAcc() (int, error) {
	s := models.Statistics{}
	return s.TotalNewAcc()
}

func ExOrder(startTime, endTime int64) (string, error) {
	s := models.Statistics{}
	return s.ExOrder(startTime, endTime)
}

func ExImport(startTime, endTime int64) (string, error) {
	s := models.Statistics{}
	return s.ExImport(startTime, endTime)
}

func GetRecommend(uid int) ([]responses.ProductRes, error) {
	resp, err := http.Get(fmt.Sprintf("%s/get-recommend/%d", N, uid))
	if err != nil {
		log.Println(err.Error(), "err.Error() services/statistic_service.go:46")
		return make([]responses.ProductRes, 0), responses.UnSuccess
	}
	bBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error(), "err.Error() services/statistic_service.go:53")
		return make([]responses.ProductRes, 0), responses.UnSuccess
	}

	body := make(map[string]interface{})
	err = json.Unmarshal(bBody, &body)
	if err != nil {
		log.Println(err.Error(), "err.Error() services/statistic_service.go:58")
		return make([]responses.ProductRes, 0), responses.UnSuccess
	}
	data, exists := body["data"]
	if !exists {
		log.Println("response error, data not exsited services/statistic_service.go:63")
		return make([]responses.ProductRes, 0), responses.UnSuccess
	}

	itemArr, ok := data.([]interface{})
	if !ok {
		log.Println("response error, data not ok services/statistic_service.go:71")
		return make([]responses.ProductRes, 0), responses.UnSuccess
	}
	responses := make([]responses.ProductRes, 0)
	for index, item := range itemArr {
		product, err := (&models.Product{}).GetProduct(int(item.(float64)))
		if err != nil {
			log.Println(err.Error(), "err.Error() services/statistic_service.go:76")
			continue
		}
		responses = append(responses, product)
		if index == 3 {
			break
		}
	}

	if len(responses) < 4 {
		topsale, _ := TopSale(4 - len(responses))
		responses = append(responses, topsale...)
	}

	return responses, nil
}
