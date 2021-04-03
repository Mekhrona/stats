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

//FilterByCategory 
func FilterByCategory( payments [] types.Payment, category types.Category) []types.Payment{

	var filtered [] types.Payment
	for _, payment := range payments {
		if payment.Category==category {
			filtered = append (filtered,payment)
		}
		
	}

return filtered

}

func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
    var i int
	m := make(map[types.Category]types.Money)

	for i=0; i<len(payments); i++ {
		category:=payments[i].Category
		paymentsInCat:=FilterByCategory(payments, category)
		averageIncat:=Avg(paymentsInCat)
        m[category]=averageIncat
		
	}
	return m
}

func PeriodsDynamic(first map[types.Category]types.Money, second map[types.Category]types.Money,) map[types.Category]types.Money {
  
  difference := make(map[types.Category]types.Money)	
  length:=first
  if len(second)>len(first){
	length=second
   }
   for key :=range length  {
		
		periodDiff:=second[key]-first[key]
		difference[key]=periodDiff
  }
   
  return difference

   }


