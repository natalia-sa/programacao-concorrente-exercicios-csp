// prova periodo 2021.2: implementar as funções handle e exec de forma que handle recebe uma lista de requests
// e um numero maximo de workers e deve criar routines para processar essas requisições.
// cada requisição, exceto a primeira, tem um pai e so pode executar quando o pai tiver executado.
// routines nao devem ficar ociosas e o maximo de routines estabelecido deve ser seguido
// caso o exec demore mais de 5 segundos para executar a requisição deve ser descartada e o resultado -1 deve ser retornado
// caso a requisição seja processada com sucesso o reusltado 1 deve ser retornado.
package main

import (
	"fmt"
	"time"
	"math/rand"
)

type request struct {
	id int
	pred_id int
	executed bool
}

// originalmente se chaamria handle
func main() {
	max_workers := 2

	requests := []request { 
		request {
			id: 0, 
			pred_id: -1, 
			executed: false, 
		},
		request {
			id: 1, 
			pred_id: 0, 
			executed: false, 
		},
		request {
			id: 2, 
			pred_id: 1, 
			executed: false, 
		},
		request {
			id: 3, 
			pred_id: 2, 
			executed: false, 
		},
	}

	// começa o trabalho
	requests_ch := make(chan request, max_workers)
	join_ch := make(chan int)

	for i := 0; i < len(requests); i ++ {
		requests_ch <-requests[i]
	}

	close(requests_ch)

	for i := 0; i < max_workers; i++ {
		go exec(requests, requests_ch, join_ch)
	}

	for i := 0; i < len(requests); i++ {
		<-join_ch
	}
	// originalmente o close tava aq
	close(join_ch)
	fmt.Printf("done")
}

// essa função ja podiamos considerar que existia e, portanto, nao precisavamos implementar
func request_func(request_id int, requests []request) request {
	for i := 0; i < len(requests); i++ {
		request := requests[i]
		if request.id == request_id {
			return request
		}
	}
	return request{-1, -1, false}
}

func exec(requests []request, requests_ch chan request, join_ch chan int) {
	result_ch := make(chan int)

	for request := range requests_ch {
		fmt.Printf("Processing request: : %d", request.id)
		for {
			fmt.Printf("checking if pred is free")
			time.Sleep(time.Duration(2) * time.Second)
			
			pred := request_func(request.pred_id, requests)
			// o que acontece se n tiver pred?
			if (pred.id != -1 && pred.executed) {

				fmt.Printf("%d its free, ready to exec", pred.id)
				time.Sleep(time.Duration(2) * time.Second)

				go do_exec(result_ch, request)

				timer := time.Tick(time.Second * time.Duration(5))
				var result int

				select {
				case <-timer:
					result = -1
				case <-result_ch:
					result = 0
				}

				join_ch <- result
				request.executed = true
				break // foi o que faltou colocar na prova
			}

			
		}
	}
}

func do_exec(result_ch chan int, req request) {
	v := 2
	randomNumber := rand.Intn(6)
	time.Sleep(time.Duration(randomNumber) * time.Second)
	fmt.Printf("%d executed", req.id)
	result_ch <- v
}