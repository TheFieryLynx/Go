package main
import "fmt"
 
func main() {
	fmt.Println("Hello Go!")
	var hello string
	hello = "test"
	var Hello string = "Test"
	fmt.Println(hello, Hello)

	//определение нескольких переменных
	var (
        name string = "Andrew"
		age int = 19
	)
	fmt.Println(name, age)
	//можно как в python (почти)
	name1 := "Lena"
	fmt.Println(name1)
}