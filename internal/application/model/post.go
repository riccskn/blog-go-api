package model

import "time"

type Post struct {
	ID        int64  `gorm:"primary_key;auto_increment;not_null"`
	Title     string `gorm:"type:varchar(15);unique;not null"`
	Body      string `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time
}

func (p *Post) BeforeCreate() error {
	p.CreatedAt = time.Now()

	return nil
}
