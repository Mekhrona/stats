package stats

import (
	"fmt"
	"reflect"

	"github.com/Mekhrona/bank/v2/pkg/types"

	"testing"
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

func TestCategoriesAv(t *testing.T) {
	
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
			Status: "OK",
		},
	}

	expected:= map[types.Category] types.Money {
       
		"food": 72_00,
	   "clothes": 127_00,
	}
	result:=CategoriesAvg(payments)

	if !reflect.DeepEqual(expected, result){
		t.Errorf( "invalid result, expected %v, actual: %v", expected, result)
	}



	
}