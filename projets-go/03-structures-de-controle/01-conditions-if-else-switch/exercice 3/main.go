// Package qui permet de valider basiquement un email
package main

import (
	"fmt"
)

// Calcul de l'aire d'un cercle avec un rayon donné
func main() {
	// Test avec un email valide
	email := "gauthier.corion@seapalm.fr"
	if CheckerEmail(email) {
		fmt.Println("L'email est valide.")
	} else {
		fmt.Println("L'email n'est pas valide.")
	}

	// Test avec un email invalide
	emailNul := "gauthier.corionseapalm.fr"
	if CheckerEmail(emailNul) {
		fmt.Println("L'email est valide.")
	} else {
		fmt.Println("L'email n'est pas valide.")
	}
}

// CheckerEmail vérifie si l'email est valide en vérifiant la présence d'un '@' et d'un '.' et une longueur minimale.
func CheckerEmail(email string) bool {
	// Vérifie si l'email contient un '@' et un '.'
	if len(email) < 5 || !containsAtSign(email) || !containsDot(email) {
		return false
	}
	return true
}

// containsAtSign vérifie si une chaîne contient le caractère '@'.
func containsAtSign(s string) bool {
	for _, char := range s {
		if char == '@' {
			return true
		}
	}
	return false
}

// containsDot vérifie si une chaîne contient le caractère '.'.
func containsDot(s string) bool {
	for _, char := range s {
		if char == '.' {
			return true
		}
	}
	return false
}

// Exercice 3 : OK (20)
