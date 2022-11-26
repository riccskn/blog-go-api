package repository

import (
	"blog-go-api/internal/application/dto"
	"blog-go-api/internal/application/model"
	"errors"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func (r *Repository) GetOne(id int64) (*model.Post, error) {

	var post model.Post
	db := r.db.Model(&post)

	query := db.Debug().Select("*").Where("id = ?", id).Find(post)

	if query.RowsAffected == 0 {
		return &post, errors.New("Invalid id")
	}

	return &post, nil
}

func (r *Repository) All() (*[]model.Post, error) {

	var posts []model.Post
	db := r.db.Model(&posts)

	query := db.Debug().Select("*").Find(posts)

	if query.Error != nil {
		return &posts, errors.New("No post found")
	}
	return &posts, nil
}

func (r *Repository) Create(dto *dto.PostDTO) (*model.Post, error) {

	var post model.Post
	db := r.db.Model(&post)

	post.Title = dto.Title
	post.Body = dto.Body

	createTodo := db.Debug().Create(&post)
	db.Commit()

	if createTodo.Error != nil {
		return &post, errors.New("failed to create a new post")
	}

	return &post, nil
}

func (r *Repository) Delete(id int64) (err error) {

	var post model.Post
	db := r.db.Model(&post)

	query := db.Debug().Select("*").Where("id = ?", id).Find(post)

	if query.RowsAffected == 0 {
		return errors.New("Invalid id")
	}

	query.Delete(post)

	return nil
}
