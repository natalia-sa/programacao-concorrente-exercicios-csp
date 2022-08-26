// em nenhum momenyo era dito que tipo de int eec retornava
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
	requests_ch := make(chan request, len(requests)) // verificar se no pseudocodigo esse canal foi criado com esse len msm
	join_ch := make(chan int)

	for i := 0; i < len(requests); i ++ {
		requests_ch <-requests[i]
	}

	for i := 0; i < max_workers; i++ {
		go exec(requests, requests_ch, join_ch)
	}

	for i := 0; i < len(requests); i++ {
		<-join_ch
	}
	// originalmente o close tava aq
	close(requests_ch)
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
		fmt.Printf("request %d: Processing request\n", request.id)

		for {
			fmt.Printf("request %d: checking if pred is free\n", request.id)
			time.Sleep(time.Duration(2) * time.Second)
			
			pred := request_func(request.pred_id, requests)
			fmt.Printf("request %d: %d is executed? %d\n", request.id, pred.id, pred.executed)
			if pred.id == -1 || pred.executed {
				fmt.Printf("request %d: %d its free, ready to exec\n", request.id, pred.id)
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
				// nao ta atualizando o statys das requests
				// talvez na prova essa linha estivesse depois do join_ch
				request.executed = true
				join_ch <- result
				break // foi o que faltou na prova
			}
		}
	}
}


func do_exec(result_ch chan int, req request) {
	v := 2
	randomNumber := rand.Intn(6)
	time.Sleep(time.Duration(randomNumber) * time.Second)
	fmt.Printf("%d executed\n", req.id)
	result_ch <- v
}

// talvez ele quisesse que toda routine ficasse rodando verificando se havia uma request livre para ser executada
// esse codigo so n cumpre com um requisito que é as threads nao ficarem ociosas