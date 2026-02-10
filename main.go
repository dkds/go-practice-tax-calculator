package main

import (
	"fmt"

	"dkds.com/tax-calculator/filemanager"
	"dkds.com/tax-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChannels := make([]chan bool, len(taxRates))
	errorChannels := make([]chan error, len(taxRates))

	for i, taxRate := range taxRates {
		doneChannels[i] = make(chan bool)
		errorChannels[i] = make(chan error)

		// cmd := cmdmanager.New()
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go priceJob.Process(doneChannels[i], errorChannels[i])
		// if err != nil {
		// 	fmt.Println("Could not process job")
		// 	fmt.Println(err)
		// }
	}

	for i := range taxRates {
		select {
		case err := <-errorChannels[i]:
			if err != nil {
				fmt.Println("Could not process job")
				fmt.Println(err)
			}
		case <-doneChannels[i]:
			fmt.Println("Done!")
		}
	}
}
