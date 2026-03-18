package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	chercherElement()

	ages := []int{25, -5, 200, 0}

	// Vérifier les âges et afficher les résultats
	for _, age := range ages {
		err := verifierAge(age)
		if err != nil {
			fmt.Printf("Âge %d : %v\n", age, err)
		} else {
			fmt.Printf("Âge %d : valide\n", age)
		}
	}

	nombres := []float64{16, -4, 0}
	for _, nombre := range nombres {
		resultat, err := calculerRacine(nombre)
		if err != nil {
			fmt.Printf("Erreur pour %f : %v\n", nombre, err)
		} else {
			fmt.Printf("La racine carrée de %f est approximativement %f\n", nombre, resultat)
		}
	}

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
		// Valider l'utilisateur et afficher les résultats
		err := validerUtilisateur(utilisateur.nom, utilisateur.email, utilisateur.age)

		// Afficher les résultats de la validation
		if err != nil {
			fmt.Printf("Erreur pour l'utilisateur %s : %v\n", utilisateur.nom, err)
		} else {
			fmt.Printf("L'utilisateur %s est valide.\n", utilisateur.nom)
		}
	}
}

// PAS DE TRY/CATCH EN GO, ON UTILISE DES VALEURS DE RETOUR POUR GÉRER LES ERREURS

func chercherElement() {
	// Exemple de fonction qui cherche un élément dans une slice et retourne une erreur si l'élément n'est pas trouvé
	fruits := []string{"pomme", "banane", "orange"}
	element := "kiwi"
	trouve, err := trouverElement(fruits, element)
	if err != nil {
		fmt.Printf("Erreur : %s\n", err)
	} else {
		fmt.Printf("L'élément '%s' a été trouvé : %t\n", element, trouve)
	}
}

func trouverElement(slice []string, element string) (bool, error) {
	for _, item := range slice {
		if item == element {
			return true, nil // Élément trouvé, pas d'erreur
		}
	}
	return false, fmt.Errorf("l'élément '%s' n'a pas été trouvé", element) // Élément non trouvé, retourner une erreur
}

type error interface {
	Error() string
}

func verifierAge(age int) error {
	// Exemple de fonction qui vérifie si un âge est valide et retourne une erreur si ce n'est pas le cas
	if age < 0 {
		return fmt.Errorf("âge invalide : %d (doit être positif)", age)
	}
	// Supposons que l'âge maximum raisonnable soit 150 ans
	if age > 150 {
		return fmt.Errorf("âge invalide : %d (trop élevé)", age)
	}
	return nil
}

func calculerRacine(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("impossible de calculer la racine carrée d'un nombre négatif")
	}

	// simulation d'un calcul simple de racine carrée (pour l'exemple)
	return x / 2, nil // Retourne une valeur fictive pour la racine carrée
}

func validerUtilisateur(nom, email string, age int) error {
	// Vérifier le nom
	if nom == "" {
		return errors.New("le nom ne peut pas être vide")
	}

	if len(nom) < 2 {
		return errors.New("le nom doit comporter au moins 2 caractères")
	}

	// Vérifier l'email
	if email == "" {
		return errors.New("l'email ne peut pas être vide")
	}

	if !strings.Contains(email, "@") {
		return errors.New("l'email doit contenir un '@'")
	}

	// Vérifier l'âge
	if age < 0 {
		return errors.New("l'âge ne peut pas être négatif")
	}

	if age > 120 {
		return errors.New("l'âge doit être inférieur ou égal à 120 (tu es trop vieux :) )")
	}

	return nil // Toutes les validations sont passées, pas d'erreur
}
