package main

//func main(){
//	fmt.Println("hello World!")
//}

import "fmt"

const englishHelloPrefix = "Hello,"
const spanishHelloPrefix = "Hola,"
const franchhHelloPrefix = "bonjuer,"

func Hello(name string,language string) string {
	if name == ""{
		name = "World"
	}

	//prefix := englishHelloPrefix
	//switch language{
	//case "french":
	//	prefix = franchhHelloPrefix
	//case "spanish":
	//	prefix = spanishHelloPrefix
	//}
	//
	//return prefix + name
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string){  // 命名返回值（prefix string）
	switch language {
	case "french":
		prefix = franchhHelloPrefix
	case "Spanish":
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return  //你只需调用 return 而不是 return prefix 即可返回所设置的值
}

//函数名称以小写字母开头。在 Go 中，公共函数以大写字母开始，私有函数以小写字母开头。我们不希望我们算法的内部结构暴露给外部，所以我们将这个功能私有化。
func main() {
	fmt.Println(Hello("",""))
}