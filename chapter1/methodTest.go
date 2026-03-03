package main
import(
	"fmt"
	"strings"
)

func main() {
	broken := "G# r#cks!"
	fmt.Println("Broken string:", broken)
	replacer := strings.NewReplacer("#","o")
	fmt.Println("'replacer' variable:",replacer)
	fixed := replacer.Replace(broken)
	fmt.Println("Fixed string:",fixed)
	broken = strings.NewReplacer("o","#").Replace(fixed)
	fmt.Println("Broken string:", broken)
}