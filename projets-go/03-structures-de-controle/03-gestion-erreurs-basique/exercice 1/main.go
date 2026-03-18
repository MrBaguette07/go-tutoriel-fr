// package main qui valide un email
package main

import (
	"fmt"
	"strings"
)

func main() {
	// Exemple d'utilisateurs à valider
	utilisateurs := []struct {
		nom   string
		email string
		age   int
	}{
		{"Alice", "alice.example.com", 30},
		{"Bob", "bob.marlé@example.com", -5},
		{"Charlie", "chilie.cahpelineuh", 25},
		{"Diana", "diana.dore", 150},
	}

	// Valider les utilisateurs et afficher les résultats
	for _, utilisateur := range utilisateurs {
		err := validerUtilisateur(utilisateur.nom, utilisateur.email, utilisateur.age)
		if err != nil {
			fmt.Printf("Erreur pour l'utilisateur %s : %v\n", utilisateur.nom, err)
		} else {
			fmt.Printf("L'utilisateur %s est valide.\n", utilisateur.nom)
		}
	}

}

func validerUtilisateur(nom, email string, age int) error {
	// Vérifier le nom
	if nom == "" {
		return fmt.Errorf("le nom ne peut pas être vide")
	}
	// Vérifier l'email
	if email == "" {
		return fmt.Errorf("l'email ne peut pas être vide")
	}

	// Vérifier que l'email contient un '@'
	if !strings.Contains(email, "@") {
		return fmt.Errorf("l'email doit contenir un '@'")
	}

	// Verifier si l'email est sur le bon format prenom.nom@domain.fr
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return fmt.Errorf("l'email doit contenir exactement un '@'")
	}

	// Vérifier que la partie locale et la partie domaine contiennent un point
	localPart := parts[0]
	domainPart := parts[1]
	if !strings.Contains(localPart, ".") {
		return fmt.Errorf("la partie locale de l'email doit contenir un '.'")
	}

	// Vérifier que la partie domaine contient un point
	if !strings.Contains(domainPart, ".") {
		return fmt.Errorf("la partie domaine de l'email doit contenir un '.'")
	}

	// Vérifier l'âge
	if age < 0 {
		return fmt.Errorf("l'âge ne peut pas être négatif")
	}

	if age > 120 {
		return fmt.Errorf("l'âge doit être inférieur ou égal à 120 (tu es trop vieux :) )")
	}

	return nil // Toutes les validations sont passées, pas d'erreur
}
