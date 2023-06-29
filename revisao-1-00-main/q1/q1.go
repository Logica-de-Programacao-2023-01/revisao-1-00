package q1

import "errors"

func CalculateDiscount(currentPurchase float64, purchaseHistory []float64) (float64, error) {
	// Implemente sua solução aqui
	if currentPurchase <= 0 {
		return 0, errors.New("inválido")
	}

	var total float64
	for _, compra := range purchaseHistory {
		total += compra
	}

	mediaCompras := total / float64(len(purchaseHistory))

	var discount float64

	if total > 1000 {
		discount = currentPurchase * 0.10
	} else if total <= 1000 && total > 500 {
		discount = currentPurchase * 0.05
	} else if total <= 500 {
		discount = currentPurchase * 0.02
	}

	if len(purchaseHistory) == 0 {
		discount = currentPurchase * 0.10
	}

	if mediaCompras > 1000 && discount < (currentPurchase*0.20) {
		discount = currentPurchase * 0.20
	}

	return discount, nil
}
