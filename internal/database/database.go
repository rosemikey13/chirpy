package database

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"
)


type DB struct {
	path string
	mux  *sync.RWMutex
}

type DBStructure struct {
	Chirps map[int]Chirp `json:"chirps"`
}

type Chirp struct {
	Id int `json:"id"`
	Body string `json:"body"`
}

var dbId = 1

// NewDB creates a new database connection
// and creates the database file if it doesn't exist
func NewDB(path string) (*DB, error){
	

	mux := &sync.RWMutex{}

	return &DB{path, mux}, nil
}

// CreateChirp creates a new chirp and saves it to disk
func (db *DB) CreateChirp(body string) (Chirp, error) {

	length := len(body)
		
	
	if length > 140 {
		return Chirp{}, fmt.Errorf("too long")
	}


		resStrings := strings.Split(body, " ")

		cleanedMsg := []string{}

		for _, bodyStr := range resStrings{

			loweredStr := strings.ToLower(bodyStr)

			if loweredStr == "kerfuffle" || loweredStr == "sharbert" || loweredStr == "fornax"{
				cleanedMsg = append(cleanedMsg, "****")
				continue
			}
			cleanedMsg = append(cleanedMsg, bodyStr)
		}

		body = strings.Join(cleanedMsg, " ")
		

	chirp := Chirp{dbId, body}

	dbId++

	data, err := json.Marshal(chirp)
	if err != nil {
		return Chirp{}, err
	}

	err = os.WriteFile(db.path,data,0600)
	if err != nil {
		return Chirp{}, fmt.Errorf("Unable to save chirp to DB: %v", err)  
	}

	return chirp, nil

}

// GetChirps returns all chirps in the database
func (db *DB) GetChirps() ([]Chirp, error)

// ensureDB creates a new database file if it doesn't exist
func (db *DB) ensureDB() error{
	if _, exists := os.ReadFile(db.path); os.IsNotExist(exists) {
		err := os.WriteFile(db.path, []byte{}, 0600)
		if err != nil {
			return fmt.Errorf("Unable to create DB file: %v", err)
		}
	}
	return nil
}

// loadDB reads the database file into memory
func (db *DB) loadDB() (DBStructure, error)

// writeDB writes the database file to disk
func (db *DB) writeDB(dbStructure DBStructure) error 