package stats

import (
	"fmt"

	"github.com/Mekhrona/bank/pkg/types")


func ExampleAvg() {

	payments := [] types.Payment {

		{
			ID: 56,
			Amount: 55_00,
			Category: "food",
		},
		{
			ID: 79,
			Amount: 165_00,
			Category: "clothes",
		},
	}
	
	fmt.Println(Avg(payments))
	//Output:11000

}

func ExampleTotalInCategory() {


	payments := [] types.Payment {

		{
			ID: 56,
			Amount: 55_00,
			Category: "food",
		},
		{
			ID: 79,
			Amount: 165_00,
			Category: "clothes",
		},
		{
			ID: 102,
			Amount: 89_00,
			Category: "clothes",
		},
		{
			ID: 65,
			Amount: 89_00,
			Category: "food",
		},
	}
	 totalSpending:=TotalInCategory(payments,"food")
	 fmt.Println(totalSpending)

	//Output:14400
	
}