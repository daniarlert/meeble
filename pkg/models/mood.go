package models

import (
	"gorm.io/gorm"
)

// A Mood represents a new record of the user mood with an optional
// reason/explanation. It also keeps track of when the record was created.
type Mood struct {
	gorm.Model
	Mood   string `json:"mood"`
	Reason string `json:"reason,omitempty"`
}

func CreateMood(db *gorm.DB, mood *Mood) error {
	if result := db.Create(mood); result.Error != nil {
		return result.Error
	}

	return nil
}

func GetFirstMood(db *gorm.DB) (*Mood, error) {
	m := &Mood{}
	if result := db.First(m); result.Error != nil {
		return nil, result.Error
	}

	return m, nil
}

func GetLastMood(db *gorm.DB) (*Mood, error) {
	m := &Mood{}
	if result := db.Last(m); result.Error != nil {
		return nil, result.Error
	}

	return m, nil
}

func GetAllMoods(db *gorm.DB) ([]*Mood, error) {
	var moods []*Mood
	if result := db.Find(moods); result.Error != nil {
		return nil, result.Error
	}

	return moods, nil
}
