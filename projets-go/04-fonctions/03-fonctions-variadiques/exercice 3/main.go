// package main qui créer une fonction variadique qui vérifie si tous les arguments passés respectent une condition.
package main

func main() {
	// Exemple d'utilisation de la fonction verifierCondition
	isAllPositive := func(element interface{}) bool {
		if num, ok := element.(int); ok {
			return num > 0
		}
		return false
	}

	nombres := []interface{}{1, 2, 3, 4, 5}
	if verifierCondition(isAllPositive, nombres...) {
		println("Tous les éléments sont positifs.")
	} else {
		println("Il y a au moins un élément qui n'est pas positif.")
	}
}

// Fonction variadique qui vérifie si tous les arguments passés respectent une condition.
func verifierCondition(condition func(interface{}) bool, elements ...interface{}) bool {
	for _, element := range elements {
		if !condition(element) {
			return false
		}
	}
	return true
}

// Exercice 3 : OK (19.0)
