package main
import(
	"fmt"
	"log"
	"os"
)
func main(){
	fileInfo, err := os.Stat("inputErr.go")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fileInfo.Size())
}