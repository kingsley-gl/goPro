package main

import (
	"net/http"
	"encoding/json"
	"fmt"
	"strconv"
)

type DataBean struct {
	Data1 string `json:"data_1"`
}

type BaseJsonBean struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func NewBaseJsonBean() *BaseJsonBean {
	return &BaseJsonBean{}
}

func Test(w http.ResponseWriter, req *http.Request) {
	result := &BaseJsonBean{}
	result.Code = 100
	result.Message = "成功"
	data := [] DataBean{}
	for a := 0; a < 5; a++ {
		data = append(data[:a], DataBean{"test " + strconv.Itoa(a)})
	}
	result.Data = data
	w.Header().Set("Content-Type","application/json;charset=UTF-8")
	if bytes, err := json.Marshal(result); err == nil {
		fmt.Fprint(w, string(bytes))
	} else {
		fmt.Fprint(w, string("Error"))
	}

}

func main() {
	http.HandleFunc("/", Test)
	http.ListenAndServe(":8001", nil)

}
