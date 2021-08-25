package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	f, err := os.Create("mi_archivo_concurrenteV2.txt")

	if err != nil {
		panic(err)
	}
	outCh := make(chan string)
	doneWrite := make(chan struct{})

	// write
	go func() {
		for s := range outCh {
			_, err := f.WriteString(s)
			if err != nil {
				panic(err)
			}
		}
		doneWrite <- struct{}{}
	}()

	numGoRutines := 10
	doneCh := make(chan string)
	final := 16777215

	for i := 0; i <= final; i = i + (final / numGoRutines) + 1 {

		paso := i + (final / numGoRutines)
		if paso > final {
			paso = final
		}

		fmt.Printf("ejecutando %d %d\n", i, paso)
		go calcNum(i, paso, outCh, doneCh)
	}

	doneNum := 0
	for doneNum < numGoRutines {
		rangoTerminado := <-doneCh
		fmt.Println("termino rango ", rangoTerminado)
		doneNum++
	}

	close(outCh)
	<-doneWrite
	fmt.Println("listo!!")

	elapsed := time.Since(start)
	log.Printf("Concurrencia V2 %s", elapsed)
}

func calcNum(start, end int, resultCh chan string, doneCh chan string) {
	var sBuilder strings.Builder
	for i := start; i <= end; i++ {
		fmt.Fprint(&sBuilder, fmt.Sprintf("%06x\n", i))
	}
	resultCh <- sBuilder.String()
	doneCh <- fmt.Sprintf("%v al %v", start, end)
}
