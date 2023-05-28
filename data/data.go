package data

import (
	"log"

	"gorm.io/gorm"
)

type StudyBuddy struct {
	gorm.Model
	Word       string `gorm:"size:120"`
	Definition string `gorm:"size:500"`
	Category   string `gorm:"size:60"`
}

func CreateTable() {
	if err := Db.AutoMigrate(&StudyBuddy{}); err != nil {
		log.Fatal(err)
	}
	log.Println("Table created successfully")
}

func CreateNote(note *StudyBuddy) {
	if err := Db.Create(note).Error; err != nil {
		log.Fatalln(err)
	}
	log.Println("Note created successfully ", note.ID)
}

func DisplayAllNotes() {
	var notesList []StudyBuddy
	if err := Db.Find(&notesList).Error; err != nil {
		log.Fatal(err)
	}
	for _, row := range notesList {
		log.Println("[", row.Category, "] ", row.Word, "-", row.Definition)
	}
}
