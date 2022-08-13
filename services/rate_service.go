package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"new-order-food/models"
)

func SetRate(uid, pid int, rate int) error {
	r := models.Rate{}
	go func(){
		req_map := map[string]interface{}{
			"user_id": uid,
			"item_id":pid,
			"rating": rate,
		}
		bBody,err := json.Marshal(req_map)
		if err != nil {
			log.Println(err.Error(), "err.Error() services/rate_service.go:22")
			return
		}
		_, err = http.Post(fmt.Sprintf("%s/add-new-rate", N),"application/json", bytes.NewReader(bBody))
		if err != nil {
			log.Println(err.Error(), "err.Error() services/rate_service.go:28")
			return
		}
	}()
	return r.SetRate(pid, uid, rate)
}
