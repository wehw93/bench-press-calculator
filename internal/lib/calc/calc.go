package calc

import "bench_press_calculator/internal/model"

func CountCalc(weight float32, quantity float32, averageWeight float32) *model.Stat {
	MaxPress1 := (weight*quantity)/30 + weight
	MaxPress2 := weight*(1+0.0333*quantity)
	MaxPress3 := weight/(1.0278 - 0.0278 *quantity)
	MaxPress4 := weight*100/(101.3 - 2.67123*quantity)
	MaxPress5 :=weight*(1+0.025*quantity)
	MaxPress := MaxPress1 +MaxPress2 + MaxPress3 + MaxPress4 + MaxPress5
	MaxPress = MaxPress/5
	PersentBetter := ((MaxPress - averageWeight) / averageWeight) * 100
	return &model.Stat{
		MaxPress: MaxPress,
		PersentBetter: PersentBetter,
	}
}
