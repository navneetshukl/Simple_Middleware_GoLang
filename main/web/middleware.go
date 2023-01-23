package main

import (
	"fmt"
	"net/http"
)

func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Enter the sum of 5 and 7")
		var c int
		fmt.Scanln(&c)
		if c==12{
			fmt.Println("Hit The Page")
		next.ServeHTTP(w,r)
		}else{
			fmt.Println("Please Enter the correct Answer")
		}
		
	})
}