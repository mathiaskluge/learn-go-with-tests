package main

const (
	spanish = "ES"
	french  = "FR"

	enHelloPrefix = "Hello, "
	esHelloPrefix = "Hola, "
	frHelloPrefix = "Bonjour, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = esHelloPrefix
	case french:
		prefix = frHelloPrefix
	default:
		prefix = enHelloPrefix
	}
	return
}

func main() {
	//fmt.Println(Hello("", ""))
}
