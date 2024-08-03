package functions

func CalculateFinalBill(costPerMessage float64, numMessages int) float64 {
	costForMessages := calculateBaseBill(costPerMessage, numMessages)
	discount := calculateDiscount(numMessages)
	return float64(costForMessages) - discount*float64(costForMessages)
}

func calculateDiscount(messagesSent int) float64 {
	if messagesSent > 5000 {
		return 0.2
	}
	if messagesSent > 1000 {
		return 0.1
	}
	return 0.0
}

func calculateBaseBill(costPerMessage float64, messagesSent int) float64 {
	return costPerMessage * float64(messagesSent)
}
