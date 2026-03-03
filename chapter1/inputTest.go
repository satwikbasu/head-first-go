package main
import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Enter 1st input: ")
	reader := bufio.NewReader(os.Stdin)
	input1,_ := reader.ReadString('\n')
	fmt.Println("1st input entered:", input1)
	fmt.Print("Enter 2nd input: ")
	input2,_ := bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Println("Input entered:", input2)
}