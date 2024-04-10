package internal

import (
	"context"
	"time"

	"github.com/carlmjohnson/requests"
)

// const targetURL = "http://49.50.175.242:8080/v1/api/school?year=2024&month=04d&day=11"

type Sapid struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Data    []string `json:"data"`
}

func GetMeals(t time.Time) ([]string, error) {
	// fmt.Println(fmt.Sprintf(targetURL, t.Year(), int(t.Month()), t.Day()))
	// url := fmt.Sprintf(targetURL, t.Year(), int(t.Month()), t.Day())

	var res Sapid
	err := requests.
		URL("http://49.50.175.242:8080/v1/api/school?year=2024&month=04&day=11").
		ToJSON(&res).
		Fetch(context.Background())

	if err != nil {
		return nil, err
	}

	return res.Data, nil
}
