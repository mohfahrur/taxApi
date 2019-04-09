package services

type TaxDetail struct {
	TaxType    string
	Refundable bool
	TaxPay     int
}

// func (TaxDetail) GetTaxDetail(price int, code int) {
// 	var taxDetail TaxDetail
// 	if code == 1 {
// 		tax := price * 10 / 100
// 		taxDetail := TaxDetail{
// 			TaxType:    "food",
// 			Refundable: true,
// 			TaxPay:     tax,
// 		}
// 	}

// 	if code == 2 {
// 		tax := 10 + (price * 2 / 100)
// 		taxDetail := TaxDetail{
// 			TaxType:    "tobacco",
// 			Refundable: false,
// 			TaxPay:     tax,
// 		}
// 	}

// 	if code == 3 {
// 		tax := 0
// 		if price >= 100 {
// 			tax := (price - 100) * 1 / 100
// 		}
// 		taxDetail := TaxDetail{
// 			TaxType:    "entertainment",
// 			Refundable: false,
// 			TaxPay:     tax,
// 		}
// 	}
// 	return taxDetail
// }
