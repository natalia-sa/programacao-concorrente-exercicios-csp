package main
 
import ( 
	"fmt" 
	"os" 
)
 
func main() {
    fmt.Println("Hello, world!") 
    fmt.Println(os.Getenv("USER"), ", Let's be friends!") 
}