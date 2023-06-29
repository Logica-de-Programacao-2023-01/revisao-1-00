package q3

import "errors"

func FindMinMaxAverage(numbers []int) (int, int, float64, error) {
	// Implemente sua solução aqui
	if len(numbers) == 0 {
		return 0, 0, 0, errors.New("lista vazia")
	}

	max := numbers[0]
	min := numbers[0]
	soma := 0

	for _, q := range numbers {
		if q > max {
			max = q
		}

		if q < min {
			min = q
		}

		soma += q
	}

	average := float64(soma-max-min) / float64(len(numbers)-2)
	return max, min, average, nil
}
