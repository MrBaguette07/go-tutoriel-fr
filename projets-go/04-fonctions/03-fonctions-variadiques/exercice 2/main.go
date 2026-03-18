// package main qui créer une fonction variadique pour créer des utilisateurs avec des informations optionnelles.
package main

import "fmt"

type Utilisateur struct {
	Nom    string
	Prenom string
	Age    int
	Email  string
	Taille float64
}

func main() {
	utilisateur1 := creationUtilisateur("Gauthier", "Corion", 23, "gauthier.corion@seapalm.fr", 1.78)

	fmt.Printf("Le nom de l'utilisateur est %s %s, âge %d, email %s, taille %.2f\n", utilisateur1.Nom, utilisateur1.Prenom, utilisateur1.Age, utilisateur1.Email, utilisateur1.Taille)
}

func creationUtilisateur(nom string, informations ...interface{}) Utilisateur {
	utilisateur := Utilisateur{
		Nom:    nom,
		Prenom: "",
		Age:    0,
		Email:  "",
		Taille: 0.0,
	}

	for _, info := range informations {
		switch v := info.(type) {
		case string:
			if utilisateur.Prenom == "" {
				utilisateur.Prenom = v
			} else if utilisateur.Email == "" {
				utilisateur.Email = v
			}
		case int:
			utilisateur.Age = v
		case float64:
			utilisateur.Taille = v
		}
	}

	return utilisateur
}

// Exercice 2 : OK (19.0)
