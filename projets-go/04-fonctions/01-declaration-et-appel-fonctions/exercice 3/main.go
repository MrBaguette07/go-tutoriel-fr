// package main qui créer une fonction estMajeur qui prend un âge et retourne true si la personne est majeure (18 ans ou plus), false sinon.
package main

func main() {
	age := []int{17, 18, 20}
	for _, a := range age {
		if estMajeur(a) {
			println("La personne est majeure.")
		} else {
			println("La personne n'est pas majeure.")
		}
	}
}

// estMajeur prend un âge en entrée et retourne true si la personne est majeure (18 ans ou plus), false sinon.
func estMajeur(age int) bool {
	if age >= 18 {
		return true
	}
	return false
}
