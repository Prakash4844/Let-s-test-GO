package main

import (
	"fmt"
)

func main() {
    fmt.Println(Hello("", ""))
}

const (
	englishHelloPrefix = "Hello, "
	japaneseHelloPrefix = "Mosi mosi, "
	hindiHelloPrefix = "Namaste, "
	cavemanHelloPrefix = "Ouga, "

	altLanguageJP = "Japanese"
	altLanguageHI = "Hindi"
	altLanguageCaveman = "Ouga"
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string){
	switch language {
		case altLanguageJP:
			prefix = japaneseHelloPrefix
		case altLanguageHI:
			prefix = hindiHelloPrefix
		case altLanguageCaveman:
			prefix = cavemanHelloPrefix
		default:
			prefix = englishHelloPrefix
	}
	return prefix
}

// func greetingPrefix(language string) (prefix string){
// 	if language == altLanguageJP{
// 		prefix = japaneseHelloPrefix
// 	} else if language == altLanguageHI{
// 		prefix = hindiHelloPrefix
// 	} else if language == altLanguageCaveman{
// 		prefix = cavemanHelloPrefix
// 	}
// 	return prefix
// }