package calculatrice

import (
	"errors"
	"fmt"
)

// Erreurs personnalisées pour la calculatrice
var (
	ErrDivisionParZero  = errors.New("division par zéro impossible")
	ErrDenominateurZero = errors.New("le dénominateur ne peut pas être zéro")
)

// Fraction représente une fraction mathématique avec un numérateur et un dénominateur.
type Fraction struct {
	Numerateur   int `json:"numerateur"`
	Denominateur int `json:"denominateur"`
}

func main() {
	addition := Addition(5, 3)
	soustraction := Soustraction(5, 3)
	multiplication := Multiplication(5, 3)
	division, err := Division(5, 3)
	if err != nil {
		fmt.Println("Erreur lors de la division:", err)
		return
	}

	fmt.Println("Addition:", addition)
	fmt.Println("Soustraction:", soustraction)
	fmt.Println("Multiplication:", multiplication)
	fmt.Println("Division:", division)
}

// NouvelleFraction crée une nouvelle instance de Fraction avec le numérateur et le dénominateur spécifiés.
// Retourne une erreur si le dénominateur est zéro.
func NouvelleFraction(numerateur, denominateur int) (*Fraction, error) {
	if denominateur == 0 {
		return nil, ErrDenominateurZero
	}
	return &Fraction{Numerateur: numerateur, Denominateur: denominateur}, nil
}

// Addition additionne deux nombres entiers et retourne le résultat.
func Addition(a, b int) int {
	return a + b
}

// Soustraction soustrait deux nombres entiers et retourne le résultat.
func Soustraction(a, b int) int {
	return a - b
}

// Multiplication multiplie deux nombres entiers et retourne le résultat.
func Multiplication(a, b int) int {
	return a * b
}

// Division divise deux nombres entiers et retourne le résultat.
// Retourne une erreur si le diviseur est zéro.
func Division(dividende, diviseur float64) (float64, error) {
	if diviseur == 0 {
		return 0, ErrDivisionParZero
	}
	return dividende / diviseur, nil
}

// Addition additionne deux fractions et retourne le résultat sous forme de nouvelle fraction.
// Le résultat n'est pas automatiquement simplifié.
func (f *Fraction) Addition(autre *Fraction) *Fraction {
	// Pour additionner a/b + c/d = (a*d + c*b) / (b*d)
	nouveauNumerateur := f.Numerateur*autre.Denominateur + autre.Numerateur*f.Denominateur
	nouveauDenominateur := f.Denominateur * autre.Denominateur

	// Pas de vérification d'erreur car on sait que les dénominateurs sont non-zéro
	resultat, _ := NouvelleFraction(nouveauNumerateur, nouveauDenominateur)
	return resultat
}

// Multiplication multiplie deux fractions et retourne le résultat sous forme de nouvelle fraction.
// Le résultat n'est pas automatiquement simplifié.
func (f *Fraction) Multiplication(autre *Fraction) *Fraction {
	// Pour multiplier a/b * c/d = (a*c) / (b*d)
	nouveauNumerateur := f.Numerateur * autre.Numerateur
	nouveauDenominateur := f.Denominateur * autre.Denominateur

	resultat, _ := NouvelleFraction(nouveauNumerateur, nouveauDenominateur)
	return resultat
}

// Exercice 2 : Semi-ok
