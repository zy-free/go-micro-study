package main

import (
	"fmt"
	mProto "github.com/zy-free/micro-study/api/member/proto"
	"time"
)

func main() {
	mClient := mProto.New()
	fmt.Println(time.Now())
	rsp, err := mClient.GetByID(3)
	fmt.Println(time.Now())

	fmt.Println(rsp, err)

}
