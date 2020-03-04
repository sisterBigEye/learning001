package main

//func main(){
//	fmt.Println("hello World!")
//}

import "fmt"

func Hello(name string) string {
	return "Hello," + name
}

func main() {
	fmt.Println(Hello("hello"))
}