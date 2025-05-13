package prices

import (
	"fmt"

	"github.com/zin-min-thu/go-price-calculator/conversion"
	"github.com/zin-min-thu/go-price-calculator/iomanager"
)

type TaxIncludedPriceJob struct {
	IOManager         iomanager.IOManager `json:"-"` // not include in json
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
}

func (job *TaxIncludedPriceJob) LoadData() error {

	lines, err := job.IOManager.ReadLines()

	if err != nil {
		return err
	}

	prices, err := conversion.StringToFloat(lines)

	if err != nil {
		return err
	}

	job.InputPrices = prices

	return nil
}

func (job TaxIncludedPriceJob) Process(doneChan chan bool) {

	err := job.LoadData()

	if err != nil {
		// return err
	}

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		tacIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", tacIncludedPrice)
	}

	job.TaxIncludedPrices = result

	// return job.IOManager.WriteResult(job)
	job.IOManager.WriteResult(job)
	doneChan <- true
	// fmt.Println(result)

}

func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:   iom,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
