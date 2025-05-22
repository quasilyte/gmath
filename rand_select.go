package gmath

type RandSelectConfig[T any] struct {
	Rand *Rand

	// Whether higher weights should end up in
	// the beginning of the slice.
	HigherFirst bool

	Options []RandSelectOption[T]
}

type RandSelectOption[T any] struct {
	Weight float64
	Data   T
}

func ToWeightThresholds[T any](options []RandSelectOption[T]) float64 {
	totalWeight := 0.0
	for i := range options {
		o := &options[i]
		totalWeight += o.Weight
		o.Weight = totalWeight
	}
	return totalWeight
}

func RandSelect[T any](config RandSelectConfig[T]) {
	options := config.Options

	totalWeight := 0.0
	for _, o := range options {
		totalWeight += o.Weight
	}

	if config.HigherFirst {
		randSelectHigherFirst(config.Rand, totalWeight, options)
	} else {
		randSelectLowerFirst(config.Rand, totalWeight, options)
	}
}

func randSelectHigherFirst[T any](rand *Rand, totalWeight float64, options []RandSelectOption[T]) {
	for i := 0; i < len(options)-1; i++ {
		r := rand.Float() * totalWeight
		j := 0

		sum := 0.0
		for j = i; j < len(options); j++ {
			sum += options[j].Weight
			if sum >= r {
				break
			}
		}

		w := options[j].Weight
		options[i], options[j] = options[j], options[i]
		totalWeight -= w
	}
}

func randSelectLowerFirst[T any](rand *Rand, totalWeight float64, options []RandSelectOption[T]) {
	for i := len(options) - 1; i > 0; i-- {
		r := rand.Float() * totalWeight
		j := 0

		sum := options[0].Weight
		for sum < r {
			j++
			sum += options[j].Weight
		}

		w := options[j].Weight
		options[i], options[j] = options[j], options[i]
		totalWeight -= w
	}
}
