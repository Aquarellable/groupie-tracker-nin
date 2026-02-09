package main

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/handlers"
	"html/template"
	"net/http"
	"os"
)

// Structures pour l'API
type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Relations    string   `json:"relations"` // URL vers les relations
}

type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

// Structure combinée pour la page de détails
type FullData struct {
	Artist   Artist
	Relation Relation
}

func getArtists(w http.ResponseWriter, r *http.Request) {
	apiURL := "https://groupietrackers.herokuapp.com/api/artists"
	resp, err := http.Get(apiURL)
	if err != nil {
		http.Error(w, "Erreur API", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var artists []Artist
	json.NewDecoder(resp.Body).Decode(&artists)

	tmpl, err := template.ParseFiles("templates/artists.html")
	if err != nil {
		http.Error(w, "Template introuvable : "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, artists)
	if err != nil {
		http.Error(w, "Erreur d'affichage : "+err.Error(), http.StatusInternalServerError)
	}
}

// NOUVELLE FONCTION : Gestion de la page de détails
func getArtistDetails(w http.ResponseWriter, r *http.Request) {
	// 1. Récupérer l'ID de l'URL (ex: /details?id=1)
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID de l'artiste manquant", http.StatusBadRequest)
		return
	}

	// 2. Récupérer les infos de l'artiste
	respArt, _ := http.Get("https://groupietrackers.herokuapp.com/api/artists/" + id)
	var artist Artist
	json.NewDecoder(respArt.Body).Decode(&artist)
	respArt.Body.Close()

	// 3. Récupérer les relations (dates et lieux)
	respRel, _ := http.Get("https://groupietrackers.herokuapp.com/api/relation/" + id)
	var relation Relation
	json.NewDecoder(respRel.Body).Decode(&relation)
	respRel.Body.Close()

	// 4. Fusionner les données
	data := FullData{
		Artist:   artist,
		Relation: relation,
	}

	// 5. Envoyer au template
	tmpl, err := template.ParseFiles("templates/details.html")
	if err != nil {
		http.Error(w, "Créez le fichier templates/details.html !", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)
}

func main() {
	// Débogage dossiers
	dir, _ := os.Getwd()
	fmt.Println("📁 Dossier courant :", dir)

	// Routes
	http.HandleFunc("/", handlers.HandleHome)
	http.HandleFunc("/artists", getArtists)
	http.HandleFunc("/details", getArtistDetails) // La nouvelle route !

	fmt.Println("\n🎸 Serveur démarré !")
	fmt.Println("👉 Accueil : http://localhost:8080")
	fmt.Println("👉 Artistes : http://localhost:8080/artists")

	http.ListenAndServe(":8080", nil)
}