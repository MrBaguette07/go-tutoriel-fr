// Package todo fournit les fonctionnalités de gestion de tâches.
//
// Cette bibliothèque permet de créer, modifier et organiser des tâches
// avec support des priorités et des échéances.
//
// Exemple d'utilisation :
//
//	gestionnaire := todo.NouveauGestionnaire()
//	tache, err := gestionnaire.CreerTache("Acheter du lait", todo.PrioriteHaute)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	gestionnaire.MarquerComplete(tache.ID)
package todo

import (
	"errors"
	"time"
)

// Constantes pour les priorités
const (
	PrioriteBasse   = 1
	PrioriteMoyenne = 2
	PrioriteHaute   = 3
)

// Erreurs prédéfinies
var (
	ErrTacheInexistante = errors.New("tâche inexistante")
	ErrTitreVide        = errors.New("le titre ne peut pas être vide")
	ErrPrioriteInvalide = errors.New("priorité invalide")
)

// Tache représente une tâche à accomplir
type Tache struct {
	ID           int        `json:"id"`
	Titre        string     `json:"titre"`
	Description  string     `json:"description,omitempty"`
	Priorite     int        `json:"priorite"`
	EstComplete  bool       `json:"est_complete"`
	DateCreation time.Time  `json:"date_creation"`
	DateEcheance *time.Time `json:"date_echeance,omitempty"`
}

// GestionnaireTaches gère une collection de tâches
type GestionnaireTaches struct {
	taches     map[int]*Tache
	prochainID int
}

// NouveauGestionnaire crée une nouvelle instance de gestionnaire de tâches
func NouveauGestionnaire() *GestionnaireTaches {
	return &GestionnaireTaches{
		taches:     make(map[int]*Tache),
		prochainID: 1,
	}
}

// CreerTache ajoute une nouvelle tâche avec le titre et la priorité spécifiés.
// Retourne la tâche créée ou une erreur si les paramètres sont invalides.
func (g *GestionnaireTaches) CreerTache(titre string, priorite int) (*Tache, error) {
	// Validation des paramètres
	if titre == "" {
		return nil, ErrTitreVide
	}

	if priorite < PrioriteBasse || priorite > PrioriteHaute {
		return nil, ErrPrioriteInvalide
	}

	// Création de la tâche
	tache := &Tache{
		ID:           g.prochainID,
		Titre:        titre,
		Priorite:     priorite,
		EstComplete:  false,
		DateCreation: time.Now(),
	}

	// Ajout au gestionnaire
	g.taches[tache.ID] = tache
	g.prochainID++

	return tache, nil
}

// ObtenirTache retourne la tâche avec l'ID spécifié.
// Retourne nil et ErrTacheInexistante si la tâche n'existe pas.
func (g *GestionnaireTaches) ObtenirTache(id int) (*Tache, error) {
	tache, existe := g.taches[id]
	if !existe {
		return nil, ErrTacheInexistante
	}
	return tache, nil
}

// ListerTaches retourne toutes les tâches, optionnellement filtrées par statut.
// Si seulementEnCours est true, ne retourne que les tâches non complétées.
func (g *GestionnaireTaches) ListerTaches(seulementEnCours bool) []*Tache {
	var resultat []*Tache

	for _, tache := range g.taches {
		if seulementEnCours && tache.EstComplete {
			continue // Ignorer les tâches complétées
		}
		resultat = append(resultat, tache)
	}

	return resultat
}

// MarquerComplete marque une tâche comme terminée.
// Retourne une erreur si la tâche n'existe pas.
func (g *GestionnaireTaches) MarquerComplete(id int) error {
	tache, existe := g.taches[id]
	if !existe {
		return ErrTacheInexistante
	}

	tache.EstComplete = true
	return nil
}

// DefinirEcheance définit une date d'échéance pour une tâche.
// Retourne une erreur si la tâche n'existe pas.
func (g *GestionnaireTaches) DefinirEcheance(id int, echeance time.Time) error {
	tache, existe := g.taches[id]
	if !existe {
		return ErrTacheInexistante
	}

	tache.DateEcheance = &echeance
	return nil
}

// EstEnRetard vérifie si la tâche a dépassé sa date d'échéance.
// Retourne false si aucune échéance n'est définie.
func (t *Tache) EstEnRetard() bool {
	if t.DateEcheance == nil || t.EstComplete {
		return false
	}
	return time.Now().After(*t.DateEcheance)
}

// ObtenirPrioriteTexte retourne la priorité sous forme de texte lisible.
func (t *Tache) ObtenirPrioriteTexte() string {
	switch t.Priorite {
	case PrioriteBasse:
		return "Basse"
	case PrioriteMoyenne:
		return "Moyenne"
	case PrioriteHaute:
		return "Haute"
	default:
		return "Inconnue"
	}
}

// CompterTaches retourne le nombre total de tâches et le nombre de tâches complétées.
func (g *GestionnaireTaches) CompterTaches() (total int, completees int) {
	total = len(g.taches)

	for _, tache := range g.taches {
		if tache.EstComplete {
			completees++
		}
	}

	return total, completees
}
