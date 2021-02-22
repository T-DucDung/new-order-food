package models

import (
	"errors"
	"math"
	"new-order-food/queries"
	"strconv"
	"time"
)

type Rate struct {
	UserId     int   `json:"user_id" xml:"user_id"`
	ProductId  int   `json:"product_id" xml:"product_id"`
	Rate       int   `json:"rate" xml:"rate"`
	LastUpDate int64 `json:"last_up_date" xml:"last_up_date"`
}

func (this *Rate) SetRate(pid, uid int, rate int) error {
	check, err := checkExistRate(pid, uid)
	if err != nil {
		return err
	}
	rate = int(math.Round(float64(rate)))
	r := Rate{
		UserId:    pid,
		ProductId: uid,
		Rate:      rate,
	}

	if check == true {
		curRateUser, err := r.GetRate()
		if err != nil {
			return err
		}

		p := Product{}
		err = p.UpdateRate(pid, strconv.Itoa(curRateUser), strconv.Itoa(rate))
		if err != nil {
			return err
		}
		err = r.UpdateRate()
		if err != nil {
			return err
		}

		return nil
	} else {
		p := Product{}
		err = p.UpdateRate(pid, "0", strconv.Itoa(rate))
		if err != nil {
			return err
		}

		data, err := db.Prepare("insert into Rate (UserId,ProductId,Rate,LastUpDate)  VALUES(?, ?, ?, ?);")
		if err != nil {
			return err
		}
		_, err = data.Exec(r.UserId, r.ProductId, r.Rate, time.Now().Unix())
		if err != nil {
			return err
		}
		return nil

		return nil
	}

	return errors.New("don't know")
}

func (this *Rate) GetRate() (int, error) {
	var rate int
	err := db.QueryRow(queries.GetRate(strconv.Itoa(this.UserId), strconv.Itoa(this.ProductId))).Scan(&rate)
	if err != nil {
		return -1, err
	}
	return rate, nil
}

func (this *Rate) UpdateRate() error {
	data, err := db.Prepare("UPDATE Rate as r SET r.Rate = + ?, r.LastUpDate = ? WHERE r.UserId = ? and r.ProductId = ?;")
	if err != nil {
		return err
	}
	_, err = data.Exec(this.Rate, time.Now().Unix(), this.UserId, this.ProductId)
	if err != nil {
		return err
	}
	return nil
}

func checkExistRate(uid, pid int) (bool, error) {
	var check bool
	err = db.QueryRow(queries.CheckRateExist(strconv.Itoa(uid), strconv.Itoa(pid))).Scan(&check)
	if err != nil {
		return false, err
	}
	return check, nil
}

func countRate(pid int) (int, error) {
	var count int
	err = db.QueryRow(queries.CountRate(strconv.Itoa(pid))).Scan(&count)
	if err != nil {
		return -1, err
	}
	return count, nil
}
