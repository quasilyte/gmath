package gmath

type RandSelectConfig[T any] struct {
	Rand *Rand

	Options []RandSelectOption[T]
}

type RandSelectOption[T any] struct {
	Weight float64
	Data   T
}

func RandSelect[T any](config RandSelectConfig[T]) {
	options := config.Options

	totalWeight := 0.0
	for _, o := range config.Options {
		totalWeight += o.Weight
	}

	for i := len(options) - 1; i > 0; i-- {
		r := config.Rand.Float() * totalWeight
		j := 0

		// Options are not sorted. Have to use a linear search.
		sum := options[0].Weight
		for sum < r {
			j++
			sum += options[j].Weight
		}

		options[i], options[j] = options[j], options[i]
		totalWeight -= options[i].Weight
	}
}
