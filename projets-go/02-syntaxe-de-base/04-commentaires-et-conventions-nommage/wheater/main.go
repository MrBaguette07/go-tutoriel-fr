// Package weather fournit des fonctionnalités de récupération et d'analyse
// de données météorologiques.
//
// Ce package permet d'obtenir des informations météo depuis différentes sources
// et de les analyser pour fournir des prédictions et des alertes.
//
// Utilisation basique :
//
//	client := weather.NewClient("votre-api-key")
//	meteo, err := client.ObtenirMeteoActuelle("Paris")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Température: %.1f°C\n", meteo.Temperature)
//
// Utilisation avancée avec prédictions :
//
//	predictions, err := client.ObtenirPredictions("Lyon", 7) // 7 jours
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	for _, pred := range predictions {
//	    fmt.Printf("%s: %s, %.1f°C\n",
//	               pred.Date.Format("2006-01-02"),
//	               pred.Description,
//	               pred.TemperatureMoyenne)
//	}
//
// Configuration :
//
// Le package peut être configuré via des variables d'environnement :
//   - WEATHER_API_KEY : Clé API pour le service météo
//   - WEATHER_TIMEOUT : Timeout des requêtes en secondes (défaut: 30)
//   - WEATHER_CACHE_TTL : Durée de vie du cache en minutes (défaut: 10)
package weather

import (
	"context"
	"errors"
	"time"
)

// Constantes de configuration par défaut
const (
	// TimeoutParDefaut est le timeout par défaut pour les requêtes HTTP
	TimeoutParDefaut = 30 * time.Second

	// CacheTTLParDefaut est la durée de vie par défaut du cache
	CacheTTLParDefaut = 10 * time.Minute

	// URLAPIParDefaut est l'URL de base de l'API météo
	URLAPIParDefaut = "https://api.meteo.fr/v1"
)

// Erreurs prédéfinies du package
var (
	// ErrCleAPIManquante indique qu'aucune clé API n'a été fournie
	ErrCleAPIManquante = errors.New("clé API manquante")

	// ErrVilleInexistante indique que la ville demandée n'a pas été trouvée
	ErrVilleInexistante = errors.New("ville inexistante")

	// ErrServiceIndisponible indique que le service météo est temporairement indisponible
	ErrServiceIndisponible = errors.New("service météo indisponible")

	// ErrDonneesInvalides indique que les données reçues de l'API sont invalides
	ErrDonneesInvalides = errors.New("données météo invalides")
)

// ConditionMeteo représente les conditions météorologiques possibles
type ConditionMeteo int

const (
	// ConditionInconnue indique que la condition météo n'est pas déterminée
	ConditionInconnue ConditionMeteo = iota

	// Ensoleille indique un temps ensoleillé
	Ensoleille

	// NuageuxPartiel indique un temps partiellement nuageux
	NuageuxPartiel

	// Nuageux indique un temps nuageux
	Nuageux

	// Pluvieux indique de la pluie
	Pluvieux

	// Orageux indique des orages
	Orageux

	// Neigeux indique de la neige
	Neigeux
)

// String retourne la représentation textuelle de la condition météo
func (c ConditionMeteo) String() string {
	switch c {
	case Ensoleille:
		return "Ensoleillé"
	case NuageuxPartiel:
		return "Partiellement nuageux"
	case Nuageux:
		return "Nuageux"
	case Pluvieux:
		return "Pluvieux"
	case Orageux:
		return "Orageux"
	case Neigeux:
		return "Neigeux"
	default:
		return "Inconnu"
	}
}

// MeteoActuelle représente les conditions météorologiques actuelles pour une ville
type MeteoActuelle struct {
	// Ville est le nom de la ville
	Ville string `json:"ville"`

	// Pays est le code pays (ISO 3166-1 alpha-2)
	Pays string `json:"pays"`

	// Temperature est la température actuelle en degrés Celsius
	Temperature float64 `json:"temperature"`

	// TemperatureRessentie est la température ressentie en degrés Celsius
	TemperatureRessentie float64 `json:"temperature_ressentie"`

	// Humidite est le pourcentage d'humidité (0-100)
	Humidite int `json:"humidite"`

	// Pression est la pression atmosphérique en hPa
	Pression float64 `json:"pression"`

	// VitesseVent est la vitesse du vent en km/h
	VitesseVent float64 `json:"vitesse_vent"`

	// DirectionVent est la direction du vent en degrés (0-360)
	DirectionVent int `json:"direction_vent"`

	// Condition décrit les conditions météorologiques actuelles
	Condition ConditionMeteo `json:"condition"`

	// Description est une description textuelle des conditions
	Description string `json:"description"`

	// DateMiseAJour est la date de dernière mise à jour des données
	DateMiseAJour time.Time `json:"date_mise_a_jour"`
}

// EstChaud retourne true si la température est considérée comme chaude (>= 25°C)
func (m *MeteoActuelle) EstChaud() bool {
	return m.Temperature >= 25.0
}

// EstFroid retourne true si la température est considérée comme froide (<= 5°C)
func (m *MeteoActuelle) EstFroid() bool {
	return m.Temperature <= 5.0
}

// EstVenteux retourne true si le vent est fort (>= 50 km/h)
func (m *MeteoActuelle) EstVenteux() bool {
	return m.VitesseVent >= 50.0
}

// Configuration contient les paramètres de configuration du client météo
type Configuration struct {
	// CleAPI est la clé d'authentification pour l'API météo
	CleAPI string

	// URLAPI est l'URL de base de l'API météo
	URLAPI string

	// Timeout est le timeout pour les requêtes HTTP
	Timeout time.Duration

	// CacheTTL est la durée de vie des données en cache
	CacheTTL time.Duration

	// ActiverCache indique si le cache doit être utilisé
	ActiverCache bool
}

// ConfigurationParDefaut retourne une configuration avec les valeurs par défaut
func ConfigurationParDefaut() *Configuration {
	return &Configuration{
		URLAPI:       URLAPIParDefaut,
		Timeout:      TimeoutParDefaut,
		CacheTTL:     CacheTTLParDefaut,
		ActiverCache: true,
	}
}

// Client représente un client pour l'API météo
type Client struct {
	config *Configuration
	cache  map[string]*entreeCache // Cache simple en mémoire
}

// entreeCache représente une entrée du cache avec sa date d'expiration
type entreeCache struct {
	donnees    *MeteoActuelle
	expiration time.Time
}

// NouveauClient crée un nouveau client météo avec la clé API fournie.
// Utilise la configuration par défaut.
func NouveauClient(cleAPI string) (*Client, error) {
	if cleAPI == "" {
		return nil, ErrCleAPIManquante
	}

	config := ConfigurationParDefaut()
	config.CleAPI = cleAPI

	return NouveauClientAvecConfig(config)
}

// NouveauClientAvecConfig crée un nouveau client météo avec une configuration personnalisée.
func NouveauClientAvecConfig(config *Configuration) (*Client, error) {
	if config.CleAPI == "" {
		return nil, ErrCleAPIManquante
	}

	return &Client{
		config: config,
		cache:  make(map[string]*entreeCache),
	}, nil
}

// ObtenirMeteoActuelle récupère les conditions météorologiques actuelles pour la ville spécifiée.
// Utilise le cache si disponible et non expiré.
//
// Paramètres :
//   ville : nom de la ville (ex: "Paris", "Lyon", "Marseille")
//
// Retourne :
//   *MeteoActuelle : les conditions météorologiques actuelles
//   error : erreur en cas de problème de récupération
//
// Erreurs possibles :
//   - ErrVilleInexistante si la ville n'est pas trouvée
//   - ErrServiceIndisponible si l'API est indisponible
//   - ErrDonneesInvalides si les données reçues sont corrompues
func (c *Client) ObtenirMeteoActuelle(ville string) (*MeteoActuelle, error) {
	return c.ObtenirMeteoActuelleAvecContexte(context.Background(), ville)
}

// ObtenirMeteoActuelleAvecContexte récupère les conditions météorologiques actuelles
// avec support d'annulation via le contexte.
func (c *Client) ObtenirMeteoActuelleAvecContexte(ctx context.Context, ville string) (*MeteoActuelle, error) {
	// Vérifier le cache en premier
	if c.config.ActiverCache {
		if meteo := c.obtenirDuCache(ville); meteo != nil {
			return meteo, nil
		}
	}

	// Récupérer depuis l'API
	meteo, err := c.recupererDepuisAPI(ctx, ville)
	if err != nil {
		return nil, err
	}

	// Mettre en cache
	if c.config.ActiverCache {
		c.mettreEnCache(ville, meteo)
	}

	return meteo, nil
}

// ViderCache supprime toutes les entrées du cache
func (c *Client) ViderCache() {
	c.cache = make(map[string]*entreeCache)
}

// StatistiquesCache retourne des statistiques sur l'utilisation du cache
func (c *Client) StatistiquesCache() (entrees int, expirees int) {
	maintenant := time.Now()

	for _, entree := range c.cache {
		entrees++
		if maintenant.After(entree.expiration) {
			expirees++
		}
	}

	return entrees, expirees
}

// Fonctions privées pour l'implémentation

func (c *Client) obtenirDuCache(ville string) *MeteoActuelle {
	entree, existe := c.cache[ville]
	if !existe {
		return nil
	}

	// Vérifier l'expiration
	if time.Now().After(entree.expiration) {
		delete(c.cache, ville)
		return nil
	}

	return entree.donnees
}

func (c *Client) mettreEnCache(ville string, meteo *MeteoActuelle) {
	c.cache[ville] = &entreeCache{
		donnees:    meteo,
		expiration: time.Now().Add(c.config.CacheTTL),
	}
}

func (c *Client) recupererDepuisAPI(ctx context.Context, ville string) (*MeteoActuelle, error) {
	// Implémentation de la récupération depuis l'API
	// Cette fonction ferait un appel HTTP réel en production

	// Simulation pour l'exemple
	return &MeteoActuelle{
		Ville:                ville,
		Pays:                 "FR",
		Temperature:          22.5,
		TemperatureRessentie: 24.0,
		Humidite:             65,
		Pression:             1013.25,
		VitesseVent:          15.0,
		DirectionVent:        230,
		Condition:            Ensoleille,
		Description:          "Ensoleillé avec quelques nuages",
		DateMiseAJour:        time.Now(),
	}, nil
}
