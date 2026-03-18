// package main qui calcul l'imc avec les erreurs
package main

import "fmt"

// Définir des erreurs personnalisées pour différentes situations d'erreur
var (
	ErrTailleInvalide = fmt.Errorf("la taille doit être supérieure à zéro")
	ErrPoidsInvalide  = fmt.Errorf("le poids doit être supérieur à zéro")

	ErrPoidsNegatif  = fmt.Errorf("le poids ne peut pas être négatif")
	ErrTailleNegatif = fmt.Errorf("la taille ne peut pas être négative")

	ErrPoidsZero  = fmt.Errorf("le poids ne peut pas être zéro")
	ErrTailleZero = fmt.Errorf("la taille ne peut pas être zéro")

	ErrPoidsExcessif   = fmt.Errorf("le poids est excessif")
	ErrTailleExcessive = fmt.Errorf("la taille est excessive")

	ErrTailleTropPetite = fmt.Errorf("la taille est trop petite")
	ErrTailleTropGrande = fmt.Errorf("la taille est trop grande")
)

func main() {
	// Définir le poids et la taille pour calculer l'IMC
	poids := 70.0  // en kilogrammes
	taille := 1.75 // en mètres
	// Calculer l'IMC et afficher le résultat
	imc, err := calculerIMC(poids, taille)
	if err != nil {
		fmt.Printf("Erreur lors du calcul de l'IMC : %v\n", err)
	} else {
		fmt.Printf("L'IMC pour un poids de %.2f kg et une taille de %.2f m est : %.2f\n", poids, taille, imc)
	}
}

func calculerIMC(poids, taille float64) (float64, error) {
	// Vérifier que la taille est supérieure à zéro pour éviter la division par zéro
	if taille <= 0 {
		return 0, ErrTailleInvalide
	}

	// Vérifier que le poids est supérieur à zéro
	if poids <= 0 {
		return 0, ErrPoidsInvalide
	}

	if poids < 0 {
		return 0, ErrPoidsNegatif
	}

	if taille < 0 {
		return 0, ErrTailleNegatif
	}

	if poids == 0 {
		return 0, ErrPoidsZero
	}

	if taille == 0 {
		return 0, ErrTailleZero
	}

	if poids > 500 {
		return 0, ErrPoidsExcessif
	}

	if taille > 3 {
		return 0, ErrTailleExcessive
	}

	if taille < 0.5 {
		return 0, ErrTailleTropPetite
	}

	if taille > 2.5 {
		return 0, ErrTailleTropGrande
	}

	if poids > 500 {
		return 0, ErrPoidsExcessif
	}

	// Calculer l'IMC en utilisant la formule : IMC = poids / (taille * taille)
	imc := poids / (taille * taille)
	return imc, nil // Retourner l'IMC calculé et nil pour indiquer qu'il n'y a pas d'erreur
}
