package q4

import "fmt"

func CalculateFinalPrice(basePrice float64, state string, taxCode int) (float64, error) {
	// Implemente sua solução aqui
	if basePrice <= 0 {
		return 0, fmt.Errorf("preço base inválido")
	}

	var imp float64
	switch taxCode {
	case 1:
		imp = 0.05
	case 2:
		imp = 0.1
	case 3:
		imp = 0.15
	default:
		return 0, fmt.Errorf("imposto não encontrado")
	}

	var freightRate float64
	switch state {
	case "SP":
		freightRate = 0.1
	case "RJ":
		freightRate = 0.15
	case "MG":
		freightRate = 0.2
	case "ES":
		freightRate = 0.25
	default:
		freightRate = 0.3
	}

	finalPrice := basePrice + (basePrice * imp) + (basePrice * freightRate)
	return finalPrice, nil
}
