package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	start := time.Now()

	f, err := os.Create("mi_archivo.txt")

	defer f.Close()

	if err != nil {
		panic(err)
	}

	final := 16777215

	for i := 0; i <= final; i++ {
		_, err = f.WriteString(fmt.Sprintf("%06x\n", i))
		if err != nil {
			panic(err)
		}

	}

	elapsed := time.Since(start)
	log.Printf("Sin concurrencia %s", elapsed)

}
