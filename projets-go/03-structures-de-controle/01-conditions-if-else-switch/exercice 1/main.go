// package main qui calcule le prix TTC en fonction du prix HT et du taux de TVA selon le pays.
package main

import "fmt"

func main() {
	prixTTC := CalculPrixTTC(100.0, tauxTVAParPays("France"))
	println("Prix TTC du vélo en France:", prixTTC)

	prixTTC = CalculPrixTTC(100.0, tauxTVAParPays("Belgique"))
	println("Prix TTC du vélo en Belgique:", prixTTC)

	correctionConversionNoteEnLettre()
}

// CalculPrixTTC calcule le prix TTC en fonction du prix HT et du taux de TVA.
func CalculPrixTTC(prixHT float64, tauxTVA float64) float64 {
	// Calcul du prix TTC en fonction du prix HT et du taux de TVA
	return prixHT * (1 + tauxTVA/100)
}

func tauxTVAParPays(pays string) float64 {
	// Retourne le taux de TVA en fonction du pays
	switch pays {
	case "France":
		return 20.0
	case "Belgique":
		return 21.0
	case "Suisse":
		return 7.7
	default:
		return 0.0 // Taux de TVA inconnu pour ce pays
	}
}

// Vrai correction :

func correctionConversionNoteEnLettre() {
	// Tests avec différents pays
	tests := []struct {
		prix float64
		pays string
	}{
		{100.0, "France"},
		{50.0, "Belgique"},
		{200.0, "Suisse"},
		{75.0, "Canada"},
	}

	for _, test := range tests {
		prixTTC, tauxTVA := calculerPrixTTC(test.prix, test.pays)
		fmt.Printf("--- %s ---\n", test.pays)
		fmt.Printf("Prix HT : %.2f€\n", test.prix)
		fmt.Printf("TVA : %.1f%%\n", tauxTVA*100)
		fmt.Printf("Prix TTC : %.2f€\n\n", prixTTC)
	}
}

// Fonction pour calculer le prix TTC
func calculerPrixTTC(prixHT float64, pays string) (float64, float64) {
	var tauxTVA float64

	switch pays {
	case "France":
		tauxTVA = 0.20
	case "Belgique":
		tauxTVA = 0.21
	case "Suisse":
		tauxTVA = 0.077
	default:
		tauxTVA = 0.0
	}

	prixTTC := prixHT * (1 + tauxTVA)
	return prixTTC, tauxTVA
}

// Exercice 1 : Semi-ok (15)
