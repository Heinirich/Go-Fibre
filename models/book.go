package models

import (
	"github.com/Heinirich/gorm-mysql/protocol"
	"gorm.io/gorm"
)

// GORM model
type BookModel struct {
	gorm.Model
	Author    string `gorm:"size:255;not null" json:"author"`
	Title     string `gorm:"size:255;not null" json:"title"`
	Publisher string `gorm:"size:255" json:"publisher"`
}

// ToProto Convert BookModel (DB) -> Book (Proto)
func (b *BookModel) ToProto() *protocol.Book {
	return &protocol.Book{
		Id:        uint32(b.ID),
		Author:    b.Author,
		Title:     b.Title,
		Publisher: b.Publisher,
	}
}

// FromProto Convert Book (Proto) -> BookModel (DB)
func FromProto(pb *protocol.Book) *BookModel {
	return &BookModel{
		Model:     gorm.Model{ID: uint(pb.Id)},
		Author:    pb.Author,
		Title:     pb.Title,
		Publisher: pb.Publisher,
	}
}
