package main

import (
	"encoding/json"
	"fmt"
)

type TestStructTobytes struct {
	UserId string
	Name   string
}

//type SliceMock struct {
//	addr uintptr
//	len  int
//	cap  int
//}

func main() {

	//var testStruct = &TestStructTobytes{"ooo", "ppp"}
	//Len := unsafe.Sizeof(*testStruct)
	//testBytes := &SliceMock{
	//	addr: uintptr(unsafe.Pointer(testStruct)),
	//	cap:  int(Len),
	//	len:  int(Len),
	//}
	//data := *(*[]byte)(unsafe.Pointer(testBytes))
	data := []byte{
		245, 81, 58, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 251, 81, 58, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0,
		0, 0, 0,
	}
	marshal, _ := json.Marshal(data)
	fmt.Println("marshal is : ", marshal)
	var test TestStructTobytes
	json.Unmarshal(marshal, &test)
	fmt.Println("ptestStruct.data is : ", test.UserId, test.Name)
	//fmt.Println("[]byte is : ", data)
	//var ptestStruct *TestStructTobytes = *(**TestStructTobytes)(unsafe.Pointer(&data))
	//fmt.Println("ptestStruct.data is : ", ptestStruct.UserId, ptestStruct.Name)
}
