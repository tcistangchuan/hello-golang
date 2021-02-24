package main

import (
	"encoding/json"
	"log"
)

type S struct {
	Name string `json:"name2"`
	// omitempty 如果 struct old 字段为空 生成的json没有这个字段
	// string 是指生成的json字段为string
	Old int `json:"old,string,omitempty"`
}

func main() {
	s := S{
		Name: "tc",
		Old:  2,
	}
	// struct转化为字符串
	data, err := json.Marshal(s)
	log.Println(string(data))
	log.Println(err)
	// 字符串转化为struct
	//s2:=new(S)
	var s2 S
	err = json.Unmarshal(data, &s2)
	log.Println(err)
	log.Println(s2)
}
