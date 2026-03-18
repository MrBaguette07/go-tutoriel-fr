// package main qui calcul un système de note
package main

import "fmt"

func main() {
	// Exemple de note à convertir
	note := 85

	// Conversion de la note numérique en lettre
	lettre := ConversionNoteEnLettre(note)
	println("La note", note, "correspond à la lettre:", lettre)

	correctionConversionNoteEnLettre()

}

// ConversionNoteEnLettre d'une note numérique en une lettre correspondante.
func ConversionNoteEnLettre(note int) string {
	switch {
	case note >= 90:
		return "A"
	case note >= 80:
		return "B"
	case note >= 70:
		return "C"
	case note >= 60:
		return "D"
	default:
		return "F"
	}
}

// Correction :
func obtenirLettre(note int) (string, string) {
	// Validation de la note
	if note < 0 || note > 100 {
		return "INVALIDE", "La note doit être entre 0 et 100"
	}

	// Conversion avec commentaires
	switch {
	case note >= 90:
		return "A", "Excellent"
	case note >= 80:
		return "B", "Très bien"
	case note >= 70:
		return "C", "Bien"
	case note >= 60:
		return "D", "Passable"
	default:
		return "F", "Insuffisant"
	}
}

func correctionConversionNoteEnLettre() {
	// Demander une note à l'utilisateur
	var note int
	fmt.Print("Entrez une note (0-100) : ")
	fmt.Scanln(&note)

	lettre, commentaire := obtenirLettre(note)

	if lettre == "INVALIDE" {
		fmt.Printf("Erreur : %s\n", commentaire)
	} else {
		fmt.Printf("Note : %d\n", note)
		fmt.Printf("Lettre : %s\n", lettre)
		fmt.Printf("Commentaire : %s\n", commentaire)
	}
}

// Exercice 2 : Semi-ok (15)
