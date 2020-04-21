package util

import "math/rand"

func RandomString(n int) string{
var strList=[]byte("qwewqrqjfdssifsuiDISGJFDSGFJGFIJSOIFGKFLCVKB")
    result:=make([]byte,n)
    for i:=range  result{
    	result[i]=strList[rand.Intn(len(strList))]
    }
    return  string(result)
}
