// package main qui permet de faire un gestionnaire de mots de passe avec des erreurs
package main

import "fmt"

// Définir des erreurs personnalisées pour différentes situations d'erreur liées au mot de passe
var (
	ErrMotDePasseTropCourt     = fmt.Errorf("le mot de passe doit contenir au moins 8 caractères")
	ErrMotDePasseTropLong      = fmt.Errorf("le mot de passe doit contenir au maximum 64 caractères")
	ErrMotDePasseSansMajuscule = fmt.Errorf("le mot de passe doit contenir au moins une majuscule")
	ErrMotDePasseSansMinuscule = fmt.Errorf("le mot de passe doit contenir au moins une minuscule")
	ErrMotDePasseSansChiffre   = fmt.Errorf("le mot de passe doit contenir au moins un chiffre")
	ErrMotDePasseSansSpecial   = fmt.Errorf("le mot de passe doit contenir au moins un caractère spécial")
	ErrRepetitionCaractere     = fmt.Errorf("le mot de passe ne doit pas contenir de caractères répétés plus de 2 fois consécutivement")
	ErrMotDePasseTropCommun    = fmt.Errorf("le mot de passe est trop commun et facilement devinable")
	ErrMotDePasseContientNom   = fmt.Errorf("le mot de passe ne doit pas contenir le nom de l'utilisateur")
)

// NiveauSecurite représente les différents niveaux de sécurité pour les mots de passe
type NiveauSecurite int

// Définir des niveaux de sécurité pour les mots de passe
const (
	NiveauFaible NiveauSecurite = iota
	NiveauMoyen
	NiveauFort
)

// CriteresValidation représente les critères de validation pour les mots de passe
type CriteresValidation struct {
	LongueurMinimale      int
	LongueurMaximale      int
	ExigerMajuscule       bool
	ExigerMinuscule       bool
	ExigerChiffre         bool
	ExigerSpecial         bool
	RepetitionMaximale    int
	NiveauSecuriteMinimum NiveauSecurite
	VerifierMotsCourants  bool
}

// ResultatValidation représente le résultat de la validation d'un mot de passe
type ResultatValidation struct {
	Valide      bool
	Erreurs     []error
	Score       int
	Niveau      NiveauSecurite
	Suggestions []string
}

func main() {
	// Exemple d'utilisation du gestionnaire de mots de passe
	motDePasse := "P@sw0rd123gpoijsdojgi"
	niveau := NiveauFort
	resultat := validerMotDePasse(motDePasse, niveau)
	if resultat.Valide {
		fmt.Printf("Le mot de passe '%s' est valide pour le niveau de sécurité %d.\n", motDePasse, niveau)
	} else {
		fmt.Printf("Le mot de passe '%s' est invalide pour le niveau de sécurité %d. Erreurs :\n", motDePasse, niveau)
		for _, err := range resultat.Erreurs {
			fmt.Printf("- %s\n", err)
		}
	}
}

// Fonction qui simule la validation d'un mot de passe en fonction de différents critères et retourne un résultat de validation
func obtenirCriteresValidation(niveau NiveauSecurite) CriteresValidation {
	switch niveau {
	case NiveauFaible:
		return CriteresValidation{
			LongueurMinimale:      8,
			LongueurMaximale:      64,
			ExigerMajuscule:       true,
			ExigerMinuscule:       true,
			ExigerChiffre:         true,
			ExigerSpecial:         false,
			RepetitionMaximale:    2,
			NiveauSecuriteMinimum: NiveauFaible,
			VerifierMotsCourants:  false,
		}
	case NiveauMoyen:
		return CriteresValidation{
			LongueurMinimale:      12,
			LongueurMaximale:      64,
			ExigerMajuscule:       true,
			ExigerMinuscule:       true,
			ExigerChiffre:         true,
			ExigerSpecial:         true,
			RepetitionMaximale:    2,
			NiveauSecuriteMinimum: NiveauMoyen,
			VerifierMotsCourants:  true,
		}
	case NiveauFort:
		return CriteresValidation{
			LongueurMinimale:      16,
			LongueurMaximale:      64,
			ExigerMajuscule:       true,
			ExigerMinuscule:       true,
			ExigerChiffre:         true,
			ExigerSpecial:         true,
			RepetitionMaximale:    1,
			NiveauSecuriteMinimum: NiveauFort,
			VerifierMotsCourants:  true,
		}
	default:
		return obtenirCriteresValidation(NiveauMoyen) // Par défaut, utiliser les critères du niveau moyen
	}
}

var motsDePasseCourants = []string{
	"123456",
	"password",
	"123456789",
	"12345",
	"12345678",
	"azerty",
	"qwerty",
	"abc123",
	"football",
	"monkey",
	"letmein",
}

func estMotDePasseCommun(motDePasse string) bool {
	for _, mot := range motsDePasseCourants {
		if motDePasse == mot {
			return true
		}
	}
	return false
}

// Fonction qui simule la validation d'un mot de passe en fonction de différents critères et retourne un résultat de validation
func analyserComposition(motDePasse string) (bool, bool, bool, bool) {
	hasMajuscule := false
	hasMinuscule := false
	hasChiffre := false
	hasSpecial := false

	for _, char := range motDePasse {
		switch {
		case char >= 'A' && char <= 'Z':
			hasMajuscule = true
		case char >= 'a' && char <= 'z':
			hasMinuscule = true
		case char >= '0' && char <= '9':
			hasChiffre = true
		case (char >= 32 && char <= 47) || (char >= 58 && char <= 64) || (char >= 91 && char <= 96) || (char >= 123 && char <= 126):
			hasSpecial = true
		}
	}
	return hasMajuscule, hasMinuscule, hasChiffre, hasSpecial
}

// Fonction qui vérifie si un mot de passe contient des caractères répétés plus de maxRepetition fois consécutivement
func verifierRepetition(motDePasse string, maxRepetition int) bool {
	if maxRepetition <= 0 {
		return true // Pas de limite de répétition
	}

	count := 1
	for i := 1; i < len(motDePasse); i++ {
		if motDePasse[i] == motDePasse[i-1] {
			count++
			if count > maxRepetition {
				return false
			}
		} else {
			count = 1
		}
	}
	return true
}

func calculerScoreSecurite(motDePasse string) int {
	score := 0

	longueur := len(motDePasse)
	switch {
	case longueur >= 16:
		score += 3
	case longueur >= 12:
		score += 2
	case longueur >= 8:
		score++
	}

	hasMajuscule, hasMinuscule, hasChiffre, hasSpecial := analyserComposition(motDePasse)
	if hasMajuscule {
		score++
	}
	if hasMinuscule {
		score++
	}
	if hasChiffre {
		score++
	}
	if hasSpecial {
		score++
	}

	// Bonus pour les mots de passe qui ne sont pas communs
	if !estMotDePasseCommun(motDePasse) {
		score++
	}

	// Malus pour les mots de passe avec des caractères répétés
	if !verifierRepetition(motDePasse, 2) {
		score--
	}

	// Points bonus pour la complexité du mot de passe
	if hasChiffre && hasSpecial && hasMajuscule && hasMinuscule {
		score++
	}

	if score < 0 {
		score = 0
	}
	if score > 10 {
		score = 10
	}
	return score
}

func determinerNiveauSecurite(score int) NiveauSecurite {
	switch {
	case score >= 8:
		return NiveauFort
	case score >= 5:
		return NiveauMoyen
	default:
		return NiveauFaible
	}
}

func genererSuggestions(motDePasse string, criteres CriteresValidation) []string {
	suggestions := []string{}

	hasMajuscule, hasMinuscule, hasChiffre, hasSpecial := analyserComposition(motDePasse)
	if len(motDePasse) < criteres.LongueurMinimale {
		suggestions = append(suggestions, fmt.Sprintf("Ajouter au moins %d caractères", criteres.LongueurMinimale-len(motDePasse)))
	}
	if len(motDePasse) > criteres.LongueurMaximale {
		suggestions = append(suggestions, fmt.Sprintf("Réduire à au maximum %d caractères", criteres.LongueurMaximale))
	}
	if criteres.ExigerMajuscule && !hasMajuscule {
		suggestions = append(suggestions, "Ajouter au moins une majuscule")
	}
	if criteres.ExigerMinuscule && !hasMinuscule {
		suggestions = append(suggestions, "Ajouter au moins une minuscule")
	}
	if criteres.ExigerChiffre && !hasChiffre {
		suggestions = append(suggestions, "Ajouter au moins un chiffre")
	}
	if criteres.ExigerSpecial && !hasSpecial {
		suggestions = append(suggestions, "Ajouter au moins un caractère spécial")
	}
	if !verifierRepetition(motDePasse, criteres.RepetitionMaximale) {
		suggestions = append(suggestions, fmt.Sprintf("Éviter les caractères répétés plus de %d fois consécutivement", criteres.RepetitionMaximale))
	}
	if criteres.VerifierMotsCourants && estMotDePasseCommun(motDePasse) {
		suggestions = append(suggestions, "Éviter les mots de passe courants")
	}
	return suggestions
}

func validerMotDePasse(motDePasse string, niveau NiveauSecurite) ResultatValidation {
	resultat := ResultatValidation{
		Valide:  true,
		Erreurs: []error{},
	}

	criteres := obtenirCriteresValidation(niveau)

	if motDePasse == "" {
		resultat.Valide = false
		resultat.Erreurs = append(resultat.Erreurs, fmt.Errorf("le mot de passe ne peut pas être vide"))
		return resultat
	}

	longueur := len(motDePasse)
	if longueur < criteres.LongueurMinimale {
		resultat.Valide = false
		resultat.Erreurs = append(resultat.Erreurs, ErrMotDePasseTropCourt)
	}

	if longueur > criteres.LongueurMaximale {
		resultat.Valide = false
		resultat.Erreurs = append(resultat.Erreurs, ErrMotDePasseTropLong)
	}

	hasMajuscule, hasMinuscule, hasChiffre, hasSpecial := analyserComposition(motDePasse)
	if criteres.ExigerMajuscule && !hasMajuscule {
		resultat.Valide = false
		resultat.Erreurs = append(resultat.Erreurs, ErrMotDePasseSansMajuscule)
	}

	if criteres.ExigerMinuscule && !hasMinuscule {
		resultat.Valide = false
		resultat.Erreurs = append(resultat.Erreurs, ErrMotDePasseSansMinuscule)
	}

	if criteres.ExigerChiffre && !hasChiffre {
		resultat.Valide = false
		resultat.Erreurs = append(resultat.Erreurs, ErrMotDePasseSansChiffre)
	}

	if criteres.ExigerSpecial && !hasSpecial {
		resultat.Valide = false
		resultat.Erreurs = append(resultat.Erreurs, ErrMotDePasseSansSpecial)
	}

	if !verifierRepetition(motDePasse, criteres.RepetitionMaximale) {
		resultat.Valide = false
		resultat.Erreurs = append(resultat.Erreurs, ErrRepetitionCaractere)
	}

	if criteres.VerifierMotsCourants && estMotDePasseCommun(motDePasse) {
		resultat.Valide = false
		resultat.Erreurs = append(resultat.Erreurs, ErrMotDePasseTropCommun)
	}

	resultat.Score = calculerScoreSecurite(motDePasse)
	resultat.Niveau = determinerNiveauSecurite(resultat.Score)

	resultat.Suggestions = genererSuggestions(motDePasse, criteres)
	return resultat
}
