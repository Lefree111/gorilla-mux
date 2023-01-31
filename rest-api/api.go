package restapi

import (
	"encoding/json"
	database "github.com/Lefree111/gorilla-mux/database"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

func Home(w http.ResponseWriter, r *http.Request) {

}
func PostAlbum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var album database.Album
	json.NewDecoder(r.Body).Decode(&album)
	db.Create(&album)
	json.NewEncoder(w).Encode(album)
}

func GetAlbum(w http.ResponseWriter, r *http.Request) {
	var album database.Album
	id := mux.Vars(r)["id"]
	db.First(&album, id)
	if album.ID == 0 {
		json.NewEncoder(w).Encode("album not found!")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(album)
}

func GetAlbums(w http.ResponseWriter, r *http.Request) {
	var albums []*database.Album
	db.Find(&albums)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(albums)
}

func UpdateAlbum(w http.ResponseWriter, r *http.Request) {
	var album database.Album
	id := mux.Vars(r)["id"]
	db.First(&album, id)
	if album.ID == 0 {
		json.NewEncoder(w).Encode("album not found!")
		return
	}
	json.NewDecoder(r.Body).Decode(&album)
	db.Save(&album)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(album)
}

func DeleteAlbum(w http.ResponseWriter, r *http.Request) {
	var album database.Album
	id := mux.Vars(r)["id"]
	db.First(&album, id)
	if album.ID == 0 {
		json.NewEncoder(w).Encode("album not found!")
		return
	}
	db.Delete(&album, id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("album deleted successfully")
}
