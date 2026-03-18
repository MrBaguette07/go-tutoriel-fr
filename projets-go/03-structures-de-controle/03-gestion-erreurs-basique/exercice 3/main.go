// package main qui lis un CSV et gère les erreurs
package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
)

// Définir des erreurs personnalisées pour différentes situations d'erreur
var (
	ErrFichierInexistant  = errors.New("fichier inexistant")
	ErrFichierVide        = errors.New("fichier vide")
	ErrFormatCSV          = errors.New("format CSV invalide")
	ErrColonnesManquantes = errors.New("colonnes manquantes")
	ErrDonneesInvalides   = errors.New("données invalides")
)

// DonneesCSV représente les données extraites d'un fichier CSV
type DonneesCSV struct {
	Entetes    []string
	Lignes     [][]string
	NbLignes   int
	NbColonnes int
}

func main() {
	fmt.Println("Création du fichier test")
	if err := creerFichierTest(); err != nil {
		fmt.Printf("Erreur lors de la création du fichier de test : %v\n", err)
		return
	}

	fmt.Println("Lecture du fichier CSV")
	donnees, err := lireFichierCSV("utilisateurs.csv")
	if err != nil {
		fmt.Printf("Erreur lors de la lecture du fichier CSV : %v\n", err)
		return
	}
	fmt.Printf("Données lues depuis le fichier CSV : %v\n", donnees)
	fmt.Println("Extraction de la colonne 'age'")
	ages, err := extraireColonneNumerique(donnees, "age")
	if err != nil {
		fmt.Printf("Erreur lors de l'extraction de la colonne 'age' : %v\n", err)
		return
	}
	fmt.Printf("Ages extraits : %v\n", ages)
}

// Fonction qui simule la lecture d'un fichier CSV et retourne une liste d'utilisateurs ou une erreur
func lireFichierCSV(nomFichier string) (*DonneesCSV, error) {
	// lire le fichier CSV et gérer les erreurs
	if nomFichier == "" {
		return nil, fmt.Errorf("le nom du fichier ne peut pas être vide")
	}

	os.Stat(nomFichier) // Vérifier si le fichier existe
	if _, err := os.Stat(nomFichier); os.IsNotExist(err) {
		return nil, fmt.Errorf("le fichier '%s' n'existe pas", nomFichier)
	}

	fichier, error := os.Open(nomFichier)
	if error != nil {
		return nil, fmt.Errorf("erreur lors de l'ouverture du fichier : %v", error)
	}
	defer fichier.Close()

	lecture := csv.NewReader(fichier)

	// Lire toutes les lignes
	lignes, err := lecture.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("%w dans %s : %v", ErrFormatCSV, nomFichier, err)
	}

	// Vérifier que le fichier n'est pas vide
	if len(lignes) == 0 {
		return nil, fmt.Errorf("%w : %s", ErrFichierVide, nomFichier)
	}

	// Extraire les en-têtes
	entetes := lignes[0]
	donneesLignes := lignes[1:]

	// Vérifier la cohérence des colonnes
	nbColonnesAttendu := len(entetes)
	for i, ligne := range donneesLignes {
		if len(ligne) != nbColonnesAttendu {
			return nil, fmt.Errorf("%w : ligne %d a %d colonnes, %d attendues",
				ErrColonnesManquantes, i+2, len(ligne), nbColonnesAttendu)
		}
	}

	return &DonneesCSV{
		Entetes:    entetes,
		Lignes:     donneesLignes,
		NbLignes:   len(donneesLignes),
		NbColonnes: nbColonnesAttendu,
	}, nil
}

// Fonction pour créer un fichier CSV de test (à utiliser une seule fois pour générer le fichier)
func creerFichierTest() error {
	contenu := `id,nom,email,age
1,Jean Dupont,jean.dupont@example.com,30
2,Marie Curie,marie.curie@example.com,35
3,Albert Einstein,albert.einstein@example.com,40
`
	return os.WriteFile("utilisateurs.csv", []byte(contenu), 0644)
}

func extraireColonneNumerique(donnees *DonneesCSV, nomColonne string) ([]float64, error) {
	// Trouver l'index de la colonne
	indexColonne := -1
	for i, entete := range donnees.Entetes {
		if entete == nomColonne {
			indexColonne = i
			break
		}
	}
	if indexColonne == -1 {
		return nil, fmt.Errorf("%w : colonne '%s' non trouvée", ErrColonnesManquantes, nomColonne)
	}

	// Extraire les données de la colonne et les convertir en float64
	var valeurs []float64
	for i, ligne := range donnees.Lignes {
		valeurStr := ligne[indexColonne]
		var valeur float64
		_, err := fmt.Sscanf(valeurStr, "%f", &valeur)
		if err != nil {
			return nil, fmt.Errorf("%w : ligne %d, valeur '%s' n'est pas un nombre valide", ErrDonneesInvalides, i+2, valeurStr)
		}
		valeurs = append(valeurs, valeur)
	}
	return valeurs, nil
}
