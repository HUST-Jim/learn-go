package dao

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

type User struct {
	Id   int64
	Name string
}

// GetUserById
// dao层遇到sql.ErrNoRows需要将其wrap并抛给上层，因为
// 1. 这一层没法处理，而上层需要知道是哪一条sql造成的错误才能处理，也需要知道调用栈
// 2. 上层有可能会做" == sql.ErrNoRows"
func GetUserById(Id int64) (*User, error) {
	user, err := mockDataBase(Id)
	if err != nil {
		return nil, errors.Wrapf(err, "GetUserById err, Id=%d", Id)
	}
	return user, nil
}

func mockDataBase(Id int64) (*User, error) {
	if Id == 1 {
		return nil, sql.ErrNoRows
	} else {
		return &User{
			Id:   Id,
			Name: fmt.Sprintf("name_%d", Id),
		}, nil
	}
}
