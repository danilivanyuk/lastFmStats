package main
import "fmt"

func main(){
	var valueToPrint string = "Say hi"
	printGreetings(valueToPrint)
}
func printGreetings(value string) {
	fmt.Println(value)
}
