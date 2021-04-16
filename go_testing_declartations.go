package main
import(
	"fmt"
)

func main() {
	fmt.Println("test 1")
	print("Testing if func main need to have brakets ()\n")
	print("Testing if func main need to have brakets () needs brackets and package has to be used otherwise we will get these errors: \n")
	fmt.Println("ERROR: 1 -- per missing brackets in func main () --- ")
	fmt.Println("# command-line-arguments \n./go_variables.go:6:11: syntax error: unexpected {, expecting (\n./go_variables.go:7:9: syntax error: unexpected literal \"Testing if func main need to have brakets ()\\n\", expecting type\n./go_variables.go:8:1: syntax error: unexpected } after top level declaration")	
	print("\n")
	fmt.Println("fmt.Println(\"ERROR: 2 -- per missing not used package fmt")
	fmt.Println("package main: \ngo_variables.go:2:7: expected 'STRING', found '{'\ngo_variables.go:3:2: expected ';', found \"fmt\"")
}