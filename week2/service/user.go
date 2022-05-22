package service

import "github.com/HUST-Jim/learn-go/week2/dao"

func GetUserById(Id int64) (*dao.User, error) {
	return dao.GetUserById(Id)
}
