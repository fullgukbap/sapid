package main

import (
	"context"
	"fmt"

	"github.com/carlmjohnson/requests"
)

const targetURL = "http://49.50.175.242:8080/v1/api/school?year=2024&month=04&day=11"

type Sapid struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Data    []string `json:"data"`
}

func main() {
	var res Sapid
	err := requests.
		URL(targetURL).
		ToJSON(&res).
		Fetch(context.Background())

	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
