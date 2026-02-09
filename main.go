package main

import (
    "encoding/json"
    "fmt"
    "groupie-tracker/handlers"
    "net/http"
    "os"
    "time"
)

type Artist struct {
    ID           int      `json:"id"`
    Image        string   `json:"image"`
    Name         string   `json:"name"`
    Members      []string `json:"members"`
    CreationDate int      `json:"creationDate"`
    FirstAlbum   string   `json:"firstAlbum"`
    Locations    string   `json:"locations"`
    ConcertDates string   `json:"concertDates"`
    Relations    string   `json:"relations"`
}

type Location struct {
    ID        int      `json:"id"`
    Locations []string `json:"locations"`
    Dates     string   `json:"dates"`
}

type Date struct {
    ID    int      `json:"id"`
    Dates []string `json:"dates"`
}

type Relation struct {
    ID             int                 `json:"id"`
    DatesLocations map[string][]string `json:"datesLocations"`
}

func getArtists(w http.ResponseWriter, r *http.Request) {
    apiURL := "https://groupietrackers.herokuapp.com/api/artists"
    client := http.Client{
        Timeout: 10 * time.Second,
    }
    
    resp, err := client.Get(apiURL)
    if err != nil {
        http.Error(w, "Erreur lors de la requête externe", http.StatusInternalServerError)
        return
    }
    defer resp.Body.Close()
    
    var artists []Artist
    err = json.NewDecoder(resp.Body).Decode(&artists)
    if err != nil {
        http.Error(w, "Erreur : JSON impossible à lire", http.StatusInternalServerError)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(artists)
}

func main() {
    // DÉBOGAGE : Affiche le dossier courant
    dir, _ := os.Getwd()
    fmt.Println("📁 Dossier courant :", dir)
    
    // Vérifie si les templates existent
    if _, err := os.Stat("templates/home.html"); err == nil {
        fmt.Println("✅ templates/home.html trouvé")
    } else {
        fmt.Println("❌ templates/home.html NOT FOUND")
    }
    
    if _, err := os.Stat("templates/artists.html"); err == nil {
        fmt.Println("✅ templates/artists.html trouvé")
    } else {
        fmt.Println("❌ templates/artists.html NOT FOUND")
    }
    
    // Routes
    http.HandleFunc("/", handlers.HandleHome)
    http.HandleFunc("/artists", getArtists)
    
    // Démarrage du serveur
    fmt.Println("\n🎸 Serveur démarré !")
    fmt.Println("👉 Accueil : http://localhost:8080")
    fmt.Println("👉 Artistes : http://localhost:8080/artists")
    
    http.ListenAndServe(":8080", nil)
}