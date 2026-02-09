package handlers

import (
    "html/template"
    "net/http"
)

func HandleHome(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("templates/home.html")
    if err != nil {
        // Affiche l'erreur DÉTAILLÉE
        http.Error(w, "Erreur: " + err.Error(), http.StatusInternalServerError)
        return
    }
    
    tmpl.Execute(w, nil)
}