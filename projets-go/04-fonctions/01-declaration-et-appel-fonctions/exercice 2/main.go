// package main qui créer une fonction diviser qui prend deux nombres et retourne le résultat de la division et le reste.
package main

import "fmt"

func main() {
	// Définir deux nombres pour la division
	dividende := 10
	diviseur := 3
	// Appeler la fonction diviser et afficher le résultat
	quotient, reste := diviser(dividende, diviseur)
	fmt.Printf("Le résultat de la division de %d par %d est : %d avec un reste de %d.\n", dividende, diviseur, quotient, reste)
}

// diviser prend deux entiers en entrée (dividende et diviseur) et retourne le quotient et le reste de la division. Si le diviseur est zéro, elle affiche un message d'erreur et retourne zéro pour les deux valeurs.
func diviser(dividende, diviseur int) (int, int) {
	if diviseur == 0 {
		fmt.Println("Erreur : division par zéro")
		return 0, 0
	}
	quotient := dividende / diviseur
	reste := dividende % diviseur
	return quotient, reste
}
