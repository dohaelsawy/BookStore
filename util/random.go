package util

import (
	"math/rand"
	"strings"
	"time"
)

func init (){
	rand.Seed(time.Now().UnixNano())
}
func randomNumber(min int64 , max int64) int64{
	return int64(rand.Intn(int(max))) + min
}

const alphapet = "abcdefghijklmnopqrstuvwxyz"


func randomString(n int) (string){
	var builder strings.Builder
	k := len(alphapet)
	for i := 0 ; i < n ;i++ {
		c := alphapet[rand.Intn(k)]
		builder.WriteByte(c)
	}
	return builder.String()
}


func RandomString(length int) string {
	return randomString(int(length))
}

func RandomNumber(min int64 , max int64) int64 {
	return randomNumber(min , max)
}
