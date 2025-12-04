package storage

import (
	"errors"
	"log"

	"github.com/ZainaDali/TP1_GO.git/internal/config"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type GORMStore struct {
	DB *gorm.DB
}

func NewGORMStore(db *gorm.DB) *GORMStore {
	return &GORMStore{DB: db}
}

// ConnectDB initialise la connexion à la base de données et exécute les migrations
func ConnectDB() (*gorm.DB, error) {
	var err error
	// Utilise le chemin de la base de données depuis la configuration
	dbName := config.Config.Database.Name
	// log.Printf("Tentative de connexion à la base de données : %s", dbName)

	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		log.Fatalf("Échec de la connexion à la base de données '%s': %v", dbName, err)
	}

	// log.Println("Connexion à la base de données SQLite réussie !")

	// Exécute la migration automatique pour TOUS les modèles
	err = db.AutoMigrate(&Contact{})
	if err != nil {
		log.Fatalf("Échec de la migration de la base de données pour Product: %v", err)
	}
	// log.Println("Migration de la base de données pour Product réussie !")

	return db, nil
}

func (g *GORMStore) GetAll() ([]*Contact, error) {
	var contacts []*Contact
	if err := g.DB.Find(&contacts).Error; err != nil {
		return nil, err
	}
	return contacts, nil
}

func (g *GORMStore) Add(contact *Contact) error {
	return g.DB.Create(contact).Error
}

func (g *GORMStore) Update(id uint, newName, newEmail string) error {
	updates := map[string]interface{}{}

	if newName != "" {
		updates["name"] = newName
	}
	if newEmail != "" {
		updates["email"] = newEmail
	}

	if len(updates) == 0 {
		return ErrContactNotFound
	}

	return g.DB.Model(&Contact{}).Where("id = ?", id).Updates(updates).Error
}

func (g *GORMStore) Delete(id uint) error {
	result := g.DB.Delete(&Contact{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrContactNotFound
	}
	return nil
}

func (g *GORMStore) GetUserById(id uint) (*Contact, error) {
	var contact Contact
	if err := g.DB.First(&contact, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrContactNotFound
		}
		return nil, err
	}
	return &contact, nil
}
