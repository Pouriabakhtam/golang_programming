package prices

import (
	"fmt"

	"tax_calc.ca/conversion"
	"tax_calc.ca/iomanager"
)

type Taxincludedpricejob struct {
	IOmanager         iomanager.IOManager `json:"-"`
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]float64  `json:"tax_include_price"`
}

func New(iom iomanager.IOManager, taxr float64) *Taxincludedpricejob {
	return &Taxincludedpricejob{
		TaxRate:     taxr,
		InputPrices: []float64{10.0, 20.0},
		IOmanager:   iom,
	}
}

func (job *Taxincludedpricejob) Process(donechannel chan bool) {
	job.LoadData()
	result := make(map[string]float64)
	for _, PriceValues := range job.InputPrices {
		result[fmt.Sprintf("%.2f==>", PriceValues)] = PriceValues * (1 + job.TaxRate)
	}
	fmt.Println(result)
	job.TaxIncludedPrices = result
	job.IOmanager.WriteJSONFile(job)
	donechannel <- true
}

func (job *Taxincludedpricejob) LoadData() {
	sprices, err := job.IOmanager.ReadFile()
	if err != nil {
		panic(err)
	}
	prices, err := conversion.Conversion(sprices)
	if err != nil {
		panic(err)
	}
	job.InputPrices = prices
	// fmt.Println("for", job.TaxRate, "and prices", job.InputPrices, "result is", job.TaxIncludedPrices)
}
