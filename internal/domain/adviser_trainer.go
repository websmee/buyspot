package domain

import "sort"

type AdviserTrainer struct {
	incIndex       int
	initParameters []any
	mapParameters  func(parameters []any) *Adviser
	incParameter   func(parameters []any, index int) bool
	checkResult    func(result *AdviserTrainerResult) bool
}

type AdviserTrainerResult struct {
	AdviceFrequencyPerDay float64
	SuccessRatePercent    float64
}

func NewAdviserTrainer(
	initParameters []any,
	mapParameters func(parameters []any) *Adviser,
	incParameter func(parameters []any, index int) bool,
	checkResult func(result *AdviserTrainerResult) bool,
) *AdviserTrainer {
	return &AdviserTrainer{
		0,
		initParameters,
		mapParameters,
		incParameter,
		checkResult,
	}
}

func (r *AdviserTrainer) Train(historyBySymbols map[string][]Kline) []Adviser {
	var effectiveAdvisers []Adviser
	advisers := r.getAllPossibleAdvisers()
	for _, adviser := range advisers {
		result := r.runAdviser(&adviser, historyBySymbols)
		if r.checkResult(result) {
			effectiveAdvisers = append(effectiveAdvisers, adviser)
		}
	}

	return r.filterByFrequency(r.filterBySuccessRate(effectiveAdvisers))
}

func (r *AdviserTrainer) filterBySuccessRate(advisers []Adviser) []Adviser {
	filteredMap := make(map[int][]Adviser)
	for _, adviser := range advisers {
		filteredMap[adviser.SuccessRatePercent] = append(filteredMap[adviser.SuccessRatePercent], adviser)
	}

	var filtered []Adviser
	for _, advisers := range filteredMap {
		sort.Slice(advisers, func(i, j int) bool {
			return advisers[i].AdviceFrequencyPerDay > advisers[j].AdviceFrequencyPerDay
		})
		filtered = append(filtered, advisers[0])
	}

	return filtered
}

func (r *AdviserTrainer) filterByFrequency(advisers []Adviser) []Adviser {
	sort.Slice(advisers, func(i, j int) bool {
		return advisers[i].AdviceFrequencyPerDay > advisers[j].AdviceFrequencyPerDay
	})

	var filtered []Adviser
	sr := 0
	for _, adviser := range advisers {
		if adviser.SuccessRatePercent > sr {
			filtered = append(filtered, adviser)
			sr = adviser.SuccessRatePercent
		}
	}

	return filtered
}

func (r *AdviserTrainer) getAllPossibleAdvisers() []Adviser {
	var advisers []Adviser
	var res [][]any
	cp := make([]any, len(r.initParameters))
	copy(cp, r.initParameters)
	r.incrementRecursive(cp, &res, 0)
	advisers = append(advisers, *r.mapParameters(r.initParameters))
	for _, parameters := range res {
		advisers = append(advisers, *r.mapParameters(parameters))
	}

	return advisers
}

func (r *AdviserTrainer) incrementRecursive(parameters []any, res *[][]any, index int) {
	for r.incParameter(parameters, index) {
		cp := make([]any, len(parameters))
		copy(cp, parameters)
		*res = append(*res, cp)
		if index < len(parameters)-1 {
			r.incrementRecursive(parameters, res, index+1)
		}
	}

	parameters[index] = r.initParameters[index]
}

func (r *AdviserTrainer) runAdviser(adviser *Adviser, historyBySymbols map[string][]Kline) *AdviserTrainerResult {
	var checksCount, adviceCount, successCount int
	for _, marketData := range historyBySymbols {
		for i := 0; i < len(marketData); i++ {
			if i < 25 || i > len(marketData)-9 {
				continue
			}

			checksCount++
			if advice := adviser.GetAdvice(marketData[i-25 : i]); advice != nil {
				adviceCount++
				c := r.checkAdvice(advice, marketData[i:i+8])
				if c == 1 {
					successCount++
				}
				if c == 0 { // no result
					adviceCount--
				}
			}
		}
	}

	res := &AdviserTrainerResult{
		AdviceFrequencyPerDay: float64(adviceCount) / (float64(checksCount) / float64(len(historyBySymbols))) * 24,
		SuccessRatePercent:    float64(successCount) / float64(adviceCount) * 100,
	}

	adviser.AdviceFrequencyPerDay = res.AdviceFrequencyPerDay
	adviser.SuccessRatePercent = int(res.SuccessRatePercent)

	return res
}

func (r *AdviserTrainer) checkAdvice(advice *Advice, followingMarketData []Kline) int {
	startPrice := followingMarketData[0].Open
	for i := range followingMarketData {
		highDiff := followingMarketData[i].High - startPrice
		lowDiff := followingMarketData[i].Low - startPrice
		highPercent := highDiff / startPrice * 100
		lowPercent := lowDiff / startPrice * 100
		isTakeProfit := highPercent >= advice.BuyOrderSettings.TakeProfit
		isStopLoss := lowPercent <= advice.BuyOrderSettings.StopLoss
		if isTakeProfit && !isStopLoss {
			return 1
		}
		if !isTakeProfit && isStopLoss {
			return -1
		}
		if isTakeProfit && isStopLoss {
			return 0
		}
	}

	return 0
}
