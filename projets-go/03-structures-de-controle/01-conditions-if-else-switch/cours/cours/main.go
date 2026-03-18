package main

import (
	"fmt"
	"strconv"
)

func main() {
	// La variable 'resultat' n'existe que dans le bloc if
	if resultat := 10 * 5; resultat > 40 {
		fmt.Printf("Le résultat %d est supérieur à 40\n", resultat)
	}

	// fmt.Println(resultat) // Erreur ! resultat n'existe plus ici
	test()
	casFallthrough()
	exempleSwitch()
}

// Cas d'usage courant avec les erreurs :
func test() {
	texte := "123"

	// Conversion et vérification d'erreur en une ligne
	if nombre, err := strconv.Atoi(texte); err == nil {
		fmt.Printf("Nombre converti : %d\n", nombre)
	} else {
		fmt.Printf("Erreur de conversion : %v\n", err)
	}
}

func casFallthrough() {
	note := 85

	switch {
	case note >= 80:
		fmt.Println("Très bien")
		fallthrough
	case note >= 60:
		fmt.Println("Vous avez réussi")
	default:
		fmt.Println("Essayez encore")
	}
}

// Utilisation du "early return" pattern
func validateUser(age int, name string) error {
	if age < 0 {
		return fmt.Errorf("âge invalide")
	}

	if name == "" {
		return fmt.Errorf("nom requis")
	}

	// Traitement principal
	return nil
}

func exempleSwitch() {
	age := 16

	if age < 0 {
		fmt.Println("Âge invalide")
	} else if age < 13 {
		fmt.Println("Enfant")
	} else if age < 18 {
		fmt.Println("Adolescent")
	} else if age < 65 {
		fmt.Println("Adulte")
	} else {
		fmt.Println("Senior")
	}
}

// // Difficile à lire
// if condition1 {
//     if condition2 {
//         if condition3 {
//             // code
//         }
//     }
// }

// // Plus lisible
// if !condition1 {
//     return
// }
// if !condition2 {
//     return
// }
// if !condition3 {
//     return
// }
// // code

