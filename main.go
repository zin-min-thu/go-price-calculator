package main

import (
	"fmt"

	"github.com/zin-min-thu/go-price-calculator/cmdmanager"
	"github.com/zin-min-thu/go-price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	// result := make(map[float64][]float64)

	for _, taxRate := range taxRates {
		// fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))

		cmdm := cmdmanager.New()

		priceJob := prices.NewTaxIncludedPriceJob(cmdm, taxRate)
		err := priceJob.Process()

		if err != nil {
			fmt.Println("could not pricess job")
			fmt.Println(err)
		}
	}

	// fmt.Println(result)
}
