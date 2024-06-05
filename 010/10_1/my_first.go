package first

import "fmt"

const name = "Hello, Dolly!"

func Hello() string {
	return name
}

func main() {
	fmt.Println(Hello())
}
