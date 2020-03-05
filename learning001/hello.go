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

	prefix := englishHelloPrefix
	switch language{
	case "french":
		prefix = franchhHelloPrefix
	case "spanish":
		prefix = spanishHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("",""))
}