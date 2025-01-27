package handlers

import (
	"net/http"
	"fmt"
	"license/models"
	"license/config"
	"github.com/google/uuid"
)

func AddKey(w http.ResponseWriter, req *http.Request) {
	new_uuid := uuid.New().String()
	key := models.LicenseKey{Key: new_uuid}
	config.Db.Create(&key)

	fmt.Fprintf(w, "Key: %s", key.Key)
}

func DeleteKey(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	key_uuid := req.URL.Query().get("uuid")
	var key models.LicenseKey{key: key_uuid}

	if key_uuid == "" {
		http.Error(w, "Missing UUID", http.StatusBadRequest)
	}

	config.Db.First(&product)
	config.Db.Delete(&product)
}

