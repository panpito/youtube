package cake

import "log"

func Cakes(recipe, available map[string]int) int {
	var cakes int
	var enoughIngredient = true

	for enoughIngredient {
		for ingredient, requiredAmount := range recipe {
			log.Printf("checking for %v (%v)", ingredient, requiredAmount)
			amount, ok := available[ingredient]
			if !ok {
				enoughIngredient = false
				break
			}

			newAmount := amount - requiredAmount
			if newAmount < 0 {
				enoughIngredient = false
				break
			}

			available[ingredient] = newAmount
		}

		if enoughIngredient {
			cakes++
		}
	}

	return cakes
}

//func Cakes(recipe, available map[string]int) (c int) {
//	for {
//		for ingredient, amount := range recipe {
//			if available[ingredient] >= amount {
//				available[ingredient] -= amount
//			} else {
//				return  // we are out of x
//			}
//		}
//		c++
//	}
//}
