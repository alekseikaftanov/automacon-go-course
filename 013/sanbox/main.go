package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	str := `{"number":1, "sign_date":"2023-09-11", "landlord":"БендерОстап", "tenat":"БалагановШура"}`

	var c any

	err := json.Unmarshal([]byte(str), &c)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v", c)
}
