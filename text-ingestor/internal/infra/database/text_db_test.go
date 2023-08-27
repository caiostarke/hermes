package database

import (
	"fmt"
	"testing"

	"github.com/caiostarke/hermes/text-ingestor/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupInMemoryDB(autoMigrateValue any) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(autoMigrateValue)

	return db, nil
}

func TestCreateTex(t *testing.T) {
	db, err := setupInMemoryDB(&entity.Text{})
	if err != nil {
		t.Error(err)
	}

	txt, _ := entity.NewText("I got it!")
	textDB := NewText(db)

	err = textDB.Create(txt)
	assert.Nil(t, err)

	var textFound entity.Text
	err = db.Find(&textFound, "id = ?", txt.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, txt.ID, textFound.ID)
	assert.Equal(t, txt.Text, textFound.Text)
	assert.NotNil(t, textFound.Text)
}

func TestFindAllText(t *testing.T) {
	db, err := setupInMemoryDB(&entity.Text{})
	if err != nil {
		t.Error(err)
	}

	for i := 0; i < 10; i++ {
		text, err := entity.NewText(fmt.Sprintf("text %d", i))
		assert.NoError(t, err)
		db.Create(text)
	}

	textDB := NewText(db)
	texts, err := textDB.FindAll()
	assert.NoError(t, err)
	assert.Len(t, texts, 10)
	assert.Equal(t, "text 0", texts[0].Text)
	assert.Equal(t, "text 9", texts[9].Text)
}
