package api

import (
    "encoding/json"
    "groupie-tracker/models"
    "net/http"
    "time"
)

const BaseURL = "https://groupietrackers.herokuapp.com/api"

// FetchArtists récupère tous les artistes
func FetchArtists() ([]models.Artist, error) {
    client := http.Client{Timeout: 10 * time.Second}
    
    resp, err := client.Get(BaseURL + "/artists")
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    
    var artists []models.Artist
    err = json.NewDecoder(resp.Body).Decode(&artists)
    return artists, err
}

// FetchLocations récupère toutes les locations
func FetchLocations() ([]models.Location, error) {
    client := http.Client{Timeout: 10 * time.Second}
    
    resp, err := client.Get(BaseURL + "/locations")
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    
    var result struct {
        Index []models.Location `json:"index"`
    }
    err = json.NewDecoder(resp.Body).Decode(&result)
    return result.Index, err
}