package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	start := time.Now()

	f, err := os.Create("mi_archivo_concurrente.txt")

	defer f.Close()

	if err != nil {
		panic(err)
	}

	numGoRutines := 100
	doneCh := make(chan struct{})

	final := 16777215

	for i := 0; i <= final; i = i + (final / numGoRutines) + 1 {

		paso := i + (final / numGoRutines)
		if paso > final {
			paso = final
		}

		fmt.Printf("ejecutando %d %d\n", i, paso)
		go calcNum(i, paso, f, doneCh)
	}

	doneNum := 0
	for doneNum < numGoRutines {
		<-doneCh
		doneNum++
		fmt.Println("termino ", doneNum)
	}

	fmt.Println("listo!!")

	elapsed := time.Since(start)
	log.Printf("Concurrencia  V1 %s", elapsed)
}

func calcNum(start, end int, f *os.File, doneCh chan struct{}) {
	for i := start; i <= end; i++ {
		_, err := f.WriteString(fmt.Sprintf("%06x\n", i))
		if err != nil {
			panic(err)
		}
	}
	doneCh <- struct{}{}
}
