// package main qui créer une fonction calculer qui prend deux nombres et une opération (string) et retourne le résultat et une erreur éventuelle.
package main

import "fmt"

func main() {
	// Définir une série d'opérations à tester
	opperations := []struct {
		a, b      float64
		operation string
	}{
		{a: 10, b: 5, operation: "addition"},
		{a: 10, b: 5, operation: "soustraction"},
		{a: 10, b: 5, operation: "multiplication"},
		{a: 10, b: 5, operation: "division"},
		{a: 10, b: 0, operation: "division"},
		{a: 10, b: 5, operation: "inconnue"},
	}

	// Tester les différentes opérations
	for _, op := range opperations {
		resultat, err := calculer(op.a, op.b, op.operation)
		if err != nil {
			fmt.Printf("Erreur pour l'opération %s: %v\n", op.operation, err)
		} else {
			fmt.Printf("Résultat de %s entre %.2f et %.2f: %.2f\n", op.operation, op.a, op.b, resultat)
		}
	}
}

// Fonction qui prend deux nombres et une opération (string) et retourne le résultat et une erreur éventuelle.
func calculer(a, b float64, operation string) (float64, error) {
	var resultat float64
	switch operation {
	case "addition":
		resultat = a + b
	case "soustraction":
		resultat = a - b
	case "multiplication":
		resultat = a * b
	case "division":
		if b == 0 {
			return 0, fmt.Errorf("division par zéro")
		}
		resultat = a / b
	default:
		return 0, fmt.Errorf("opération inconnue")
	}
	return resultat, nil
}

// Exercice 1 : OK (19.5)
