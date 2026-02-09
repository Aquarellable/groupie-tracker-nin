package handlers

import (
    "groupie-tracker/api"
    "html/template"
    "net/http"
	"path/filepath"
)

func HandleArtists(w http.ResponseWriter, r *http.Request) {
    // Récupère les artistes
    artists, err := api.FetchArtists()
    if err != nil {
        http.Error(w, "Erreur lors de la récupération des artistes", http.StatusInternalServerError)
        return
    }

    // Charge le template avec chemin absolu
    tmplPath := filepath.Join("templates", "artists.html")
    tmpl, err := template.ParseFiles(tmplPath)
    if err != nil {
        http.Error(w, "Erreur de chargement du template: "+err.Error(), http.StatusInternalServerError)
        return
    }
    
    // Affiche le template
    err = tmpl.Execute(w, artists)
    if err != nil {
        http.Error(w, "Erreur d'exécution du template: "+err.Error(), http.StatusInternalServerError)
    }
}