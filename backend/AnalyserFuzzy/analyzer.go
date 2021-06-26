package AnalyserFuzzy

type analyserFuzzy struct{}

func New() *analyserFuzzy {
	return new(analyserFuzzy)
}

// Analyse 处理被测者的答案，返回分析结果
func (a *analyserFuzzy) Analyse(ans []int) (attachmentStyle, description string, rateOfIntimacy, rateOfAnxiety float64, err error) {
	return "", "", 0, 0, nil
}
