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
const nums = "1234567890"

func randomString(n int , Word string) (string){
	var builder strings.Builder
	k := len(Word)
	for i := 0 ; i < n ;i++ {
		c := Word[rand.Intn(k)]
		builder.WriteByte(c)
	}
	return builder.String()
}


func RandomString(length int) string {
	return randomString(int(length) , alphapet)
}

func RandomNumber(min int64 , max int64) int64 {
	return randomNumber(min , max)
}

func RandomEmail() string {
	var builder strings.Builder
	email:= "@gmail.com"
	s := randomString(6 , alphapet)
	builder.WriteString(s)
	builder.WriteString(email)
	return builder.String()
}


func RandomPssword() string {
	return randomString(8 , alphapet)
}

func RandomPhoneNumber() string {
	return randomString(11 , nums)
}

func RandomShippingMethod() string {
	methods := []string{"paypal" , "visa" , "criditcard"}
	k := len(methods)
	return methods[rand.Intn(k)]
}



