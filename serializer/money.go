package serializer

import (
	"test_mysql/model"
	"test_mysql/pkg/util"
)

type Money struct {
	UserID    uint   `json:"user_id"`
	UserName  string `json:"user_name"`
	UserMoney string `json:"user_money"`
}

func BuildMoney(item *model.User, key string) Money {
	util.Encrypt.SetKey(key)
	return Money{
		UserID:    item.ID,
		UserName:  item.Username,
		UserMoney: util.Encrypt.AesDecoding(item.Money),
	}
}
