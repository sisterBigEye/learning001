package main

//func main(){
//	fmt.Println("hello World!")
//}

import "fmt"

const englishHelloPrefix = "Hello,"
const spanishHelloPrefix = "Halo,"
const franchhHelloPrefix = "bonjuer,"

func Hello(name string,language string) string {
	if name == ""{
		name = "World"
	}
	if language == "Spanish"{
		return "Hola, "+name
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("",""))
}