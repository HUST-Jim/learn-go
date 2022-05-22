package main

import (
	"database/sql"
	"fmt"
	"github.com/HUST-Jim/learn-go/week2/service"
	"github.com/pkg/errors"
)

func main() {
	_, err := service.GetUserById(1)
	if err != nil {
		fmt.Printf("%+v", err)
	}
	if errors.Is(err, sql.ErrNoRows) {
		fmt.Println("好果汁")
	}
}
