package models

import (
	"github.com/jinzhu/gorm"
	"html"
	"strings"
	"time"
)

type Recipes struct {
	ID             uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Title          string    `gorm:"size:255;not null;unique" json:"title"`
	Description    string    `gorm:"size:255;not null;unique" json:"description"`
	Name           string    `gorm:"size:255;not null;unique" json:"name"`
	Slug           string    `gorm:"size:255;not null;unique" json:"slug"`
	MainImg        string    `gorm:"size:255;not null" json:"main_img"`
	Detailed       string    `gorm:"size:255;not null" json:"detailed"`
	DetailedFull   string    `gorm:"size:255;not null" json:"detailed_full"`
	Ingredients    string    `gorm:"size:255;not null" json:"ingredients"`
	Appointment    string    `gorm:"size:255;not null" json:"appointment"`
	MainIngredient string    `gorm:"size:255;not null" json:"main_ingredient"`
	Dish           string    `gorm:"size:255;not null" json:"dish"`
	Time           string    `gorm:"size:255;not null" json:"time"`
	Diet           string    `gorm:"size:255;not null" json:"diet"`
	CountServing   string    `gorm:"size:255;not null" json:"count_serving"`
	Content        string    `gorm:"size:255;not null;" json:"content"`
	Author         User      `json:"author"`
	AuthorID       uint32    `sql:"type:int REFERENCES users(id)" json:"author_id"`
	CreatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (r *Recipes) Prepare() {
	r.ID = 0
	r.Title = html.EscapeString(strings.TrimSpace(r.Title))
	r.Description = html.EscapeString(strings.TrimSpace(r.Description))
	r.Name = html.EscapeString(strings.TrimSpace(r.Name))
	r.Slug = smartTruncate(r.Name)

	r.Content = html.EscapeString(strings.TrimSpace(r.Content))
	r.Author = User{}
	r.CreatedAt = time.Now()
	r.UpdatedAt = time.Now()
}

func smartTruncate(text string) string {
	var MaxLength = 250
	if len(text) < MaxLength {
		return text
	}

	var truncated string
	words := strings.SplitAfter(text, "-")
	// If MaxLength is smaller than length of the first word return word
	// truncated after MaxLength.
	if len(words[0]) > MaxLength {
		return words[0][:MaxLength]
	}
	for _, word := range words {
		if len(truncated)+len(word)-1 <= MaxLength {
			truncated = truncated + word
		} else {
			break
		}
	}
	return strings.Trim(truncated, "-")
}

func (r *Recipes) SaveReceipt(db *gorm.DB) (*Recipes, error) {
	var err error
	err = db.Debug().Model(&Recipes{}).Create(&r).Error
	if err != nil {
		return &Recipes{}, err
	}
	if r.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", r.AuthorID).Take(&r.Author).Error
		if err != nil {
			return &Recipes{}, err
		}
	}
	return r, nil
}
