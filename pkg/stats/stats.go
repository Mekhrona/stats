package stats

import "github.com/Mekhrona/bank/v2/pkg/types"


//Avg расчитывает среднюю сумму платежа.
func Avg(payments []types.Payment) types.Money{
	numberOfpayments:=0
	sum:=0
	for _, payment := range payments {
		if !(payment.Status=="FAIL"){
		numberOfpayments++
		sum+=int(payment.Amount)
		}	
	}

	avg:=sum/numberOfpayments

	return types.Money(avg)
}

//TotalInCategory находит сумму покупок в определённой категории.
func TotalInCategory (payments []types.Payment, category types.Category) types.Money {
    sum:=0
	for _, payment := range payments {
		if payment.Category==category && !(payment.Status=="FAIL"){
			sum+=int(payment.Amount)
		}
	}
	return types.Money(sum)

}

func FilterByCategory( payments [] types.Payment, category types.Category) []types.Payment{

	var filtered [] types.Payment
	for _, payment := range payments {
		if payment.Category==category {
			filtered = append (filtered,payment)
		}
		
	}

return filtered

}

