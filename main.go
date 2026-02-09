package main

import (
	"fmt"

	"dkds.com/tax-calculator/filemanager"
	"dkds.com/tax-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChannels := make([]chan bool, len(taxRates))

	for i, taxRate := range taxRates {
		doneChannels[i] = make(chan bool)

		// cmd := cmdmanager.New()
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go priceJob.Process(doneChannels[i])
		// if err != nil {
		// 	fmt.Println("Could not process job")
		// 	fmt.Println(err)
		// }
	}

	for _, doneChannel := range doneChannels {
		<-doneChannel
	}
}
