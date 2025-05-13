package main

import (
	"fmt"

	"github.com/zin-min-thu/go-price-calculator/filemanager"
	"github.com/zin-min-thu/go-price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	doneChans := make([]chan bool, len(taxRates))

	// result := make(map[float64][]float64)

	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))

		// cmdm := cmdmanager.New()

		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)

		go priceJob.Process(doneChans[index])

		// if err != nil {
		// 	fmt.Println("could not pricess job")
		// 	fmt.Println(err)
		// }
	}

	for _, doneChan := range doneChans {
		<-doneChan
	}

	// fmt.Println(result)
}
