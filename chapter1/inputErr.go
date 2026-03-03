package main
import(
	"fmt"
	"bufio" //for taking input from keyboard
	"os" //for taking input, used with bufio
	"log" //for logging exceptions
	"strings" //for trimming whitespace from input
	"reflect" //for checking type of a variable
	"strconv" //for converting string to float (or other types)
)
func main(){
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin) // reader variable allows easy input multiple times
	input, err := reader.ReadString('\n') // err stores the error that may occur during input
	if err != nil{
		log.Fatal(err) 
	}
	input = strings.TrimSpace(input) // trim the new line \n during inputed
	grade, err := strconv.ParseFloat(input, 64) // 64 indicates conversion from string to float64
	gradeInt, err := strconv.ParseInt(input, 10, 0) // 10 indicates base, 0 indicates int32/64 as resultant type size
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Type of input:", reflect.TypeOf(input)) // input taken directly is always string
	fmt.Println("Type of grade:", reflect.TypeOf(grade)) // string input in converted to float64
	fmt.Println("Type of gradeInt:", reflect.TypeOf(gradeInt)) // string input in converted to float64
	fmt.Println("The grade is",grade)
	var status string // declared here so that it stays in the scope of the entire function
	if grade >= 40 {
		status = "pass"
	} else{
		status = "fail"
	}
	fmt.Println("A grade of", grade, "is", status)
}