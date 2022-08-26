// construa um pipeline em que uma goroutine gere strings aleatorias, enquanto uma segunda filtra as strings
// e uma terceira escreve os valores filtrados na saida padrao
package main 

// 2 canais, um pra receber todas as strings, um pra guardar so as filtradas
// é necessario usar canal bufferizado para n perder string? n, elas n serao perdidas, as routines travam na hora certa
// o for deve estar no main ou em cada uma das routines? ou tanto faz?
import (
	"fmt"
	"strings"
)

func main() {
	general := make(chan string)
	filtered := make(chan string)
	for {
		go generateRandomString(general)
		go filterString(general, filtered)
		go printString(filtered)
	}
}

func generateRandomString(general chan string) {
	x := "blab"
	fmt.Println("The random string is", x) 
	general <- x
}

func filterString(general chan string, filtered chan string) {
	x := <- general
	if strings.Contains(x, "bl") {
		fmt.Println("has string") 
		filtered <- x
	}
}

func printString(filtered chan string) {
	x := <- filtered
	fmt.Println("Filtered: ", x) 
}