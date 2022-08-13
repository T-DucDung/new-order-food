package models

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"math"
	"sync"
	"time"
)

var (
	onceGetRate sync.Once
	aRate       *mongo.Collection
)

//go:generate easytags $GOFILE json,xml,bson

type Rate struct {
	UserId     int   `json:"user_id" xml:"user_id" bson:"user_id"`
	ProductId  int   `json:"product_id" xml:"product_id" bson:"product_id"`
	Rate       int   `json:"rate" xml:"rate" bson:"rate"`
	LastUpDate int64 `json:"last_up_date" xml:"last_up_date" bson:"last_up_date"`
}

func getRateDocument() *mongo.Collection {
	onceGetRate.Do(func() {
		aRate = db.Database("neworderfood").Collection("Rate")
	})
	return aRate
}

func (this *Rate) SetRate(pid, uid int, rate int) error {
	check,lRate, err := checkExistRate(uid, pid)
	if err != nil {
		return err
	}
	if rate == lRate{
		return nil
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
		err = p.UpdateRate(pid, curRateUser, rate)
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
		err = p.UpdateRate(pid, 0, rate)
		if err != nil {
			return err
		}

		_, err = getRateDocument().InsertOne(context.TODO(), Rate{
			UserId:     r.UserId,
			ProductId:  r.ProductId,
			Rate:       r.Rate,
			LastUpDate: time.Now().Unix(),
		})
		if err != nil {
			log.Println(err.Error(), "err.Error() models/rate.go:77")
			return err
		}
		return nil
	}

	return errors.New("don't know")
}

func (this *Rate) GetRate() (int, error) {
	var rate Rate
	filter := bson.D{{"user_id", this.UserId}, {"product_id", this.ProductId}}
	err = getRateDocument().FindOne(context.TODO(), filter).Decode(&rate)
	if err != nil {
		return 0, err
	}
	return rate.Rate, nil
}

func (this *Rate) UpdateRate() error {
	filter := bson.D{{"user_id", this.UserId}, {"product_id", this.ProductId}}
	_, err := getRateDocument().UpdateOne(context.TODO(), filter, bson.M{"$set": bson.M{"rate": this.Rate}})
	if err != nil {
		log.Println(err.Error(), "err.Error() models/rate.go:107")
		return err
	}

	return nil
}

func checkExistRate(uid, pid int) (bool, int, error) {
	a := Rate{}
	filter := bson.D{{"user_id", uid}, {"product_id", pid}}

	err = getRateDocument().FindOne(context.TODO(), filter).Decode(&a)
	if err != nil {
		return false, 0, nil
	}

	return true, a.Rate, nil
}
