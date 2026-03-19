// package main qui créer une fonction qui retourne une closure permettant d'accumuler des valeurs avec une opération personnalisée (addition, multiplication, etc.).
package main

import "fmt"

func main() {
	// Créer un operateur d'addition
	additionAccumulateur := creerOperateur(0, addition)
	// Créer un operateur de multiplication
	multiplicationAccumulateur := creerOperateur(1, multiplication)
	// Créer un operateur de division
	divisionAccumulateur := creerOperateur(100, division)
	// Utiliser les accumulateurs
	fmt.Printf("Addition: %d\n", additionAccumulateur(5))             // 5
	fmt.Printf("Addition: %d\n", additionAccumulateur(10))            // 15
	fmt.Printf("Multiplication: %d\n", multiplicationAccumulateur(2)) // 2
	fmt.Printf("Multiplication: %d\n", multiplicationAccumulateur(3)) // 6
	fmt.Printf("Division: %d\n", divisionAccumulateur(2))             // 50
	fmt.Printf("Division: %d\n", divisionAccumulateur(5))             // 10

}

func creerOperateur(initial int, operation func(int, int) int) func(int) int {
	accumulateur := initial

	return func(val int) int {
		accumulateur = operation(accumulateur, val)
		return accumulateur
	}
}

func addition(a, b int) int {
	return a + b
}

func multiplication(a, b int) int {
	return a * b
}

func division(a, b int) int {
	if b == 0 {
		return 0 // Éviter la division par zéro
	}

	return a / b
}

// Exercice 1 : OK (19.5)
