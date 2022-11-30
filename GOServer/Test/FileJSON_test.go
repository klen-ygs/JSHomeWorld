package Test

import (
	"GOServer/DaoModle"
	"encoding/json"
	"os"
	"reflect"
	"testing"
)

func TestFileJSON(t *testing.T) {
	eqArr := []DaoModle.User{
		{
			Phone:    "17777405355",
			Name:     "123456",
			Password: "klen",
		},
		{
			Phone:    "15677644897",
			Name:     "123456",
			Password: "hello",
		},
	}
	file, err := os.ReadFile("user.json")
	if err != nil {
		panic(err)
	}
	var arr []DaoModle.User
	err = json.Unmarshal(file, &arr)
	if err != nil {
		panic(err)
	}
	for i, v := range arr {
		if reflect.DeepEqual(&v, eqArr[i]) {
			t.Error("want:", eqArr[i], "having:", v)
		}
	}
}
