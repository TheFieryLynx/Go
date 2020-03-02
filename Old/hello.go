package main
import "fmt"
 

func helloprint() {
	fmt.Println("Hello World")
}

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

	//массивы
	var numbers [5]int = [5]int{1,2,3,4,5}
	fmt.Println(numbers)
	helloprint()
	helloprint()
}