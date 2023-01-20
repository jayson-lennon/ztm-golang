package main

import "fmt"

type Counter struct {
	hits int
}

func increment(counter *Counter) {
	counter.hits++
	fmt.Println("counter is:", counter)
}
func replace(old *string, new string, counter *Counter) {
	*old = new
	increment(counter)
}

func main() {
	counter := Counter{}
	hello := "hello"
	world := "world"
	//hi := "hi"
	fmt.Println(hello, world)
	replace(&hello, "Hi", &counter)
	fmt.Println(hello, world)

	phrase := []string{hello, world}
	fmt.Println(phrase)
	replace(&phrase[1], "go", &counter)
	fmt.Println(phrase)

}
