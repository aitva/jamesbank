package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("data/ville.csv")
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(f)
	for i := 0; i < 10; i++ {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(record[2], record[3], record[8])
	}
}
