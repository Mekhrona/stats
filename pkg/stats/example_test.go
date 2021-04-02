package stats

import (
	"fmt"

	"github.com/Mekhrona/bank/v2/pkg/types"

)


func ExampleAvg() {

	payments := [] types.Payment {

		{
			ID: 56,
			Amount: 55_00,
			Category: "food",
			Status: "FAIL",
		},
		{
			ID: 79,
			Amount: 165_00,
			Category: "clothes",
			Status: "OK",
		},
	}
	
	fmt.Println(Avg(payments))
	//Output:16500

}

func ExampleTotalInCategory() {


	payments := [] types.Payment {

		{
			ID: 56,
			Amount: 55_00,
			Category: "food",
			Status: "OK",
		},
		{
			ID: 79,
			Amount: 165_00,
			Category: "clothes",
			Status: "OK",
		},
		{
			ID: 102,
			Amount: 89_00,
			Category: "clothes",
			Status: "OK",
		},
		{
			ID: 65,
			Amount: 89_00,
			Category: "food",
			Status: "FAIL",
		},
	}
	 totalSpending:=TotalInCategory(payments,"food")
	 fmt.Println(totalSpending)

	//Output:5500
	
}