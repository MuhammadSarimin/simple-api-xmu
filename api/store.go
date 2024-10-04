package api

import (
	"github.com/muhammadsarimin/simple-api-xmu/types"
	"gorm.io/gorm"
)

type store struct {
	db *gorm.DB
}

type Store interface {
	FindAll() ([]types.Movie, error)
	Create(m *types.Movie) error
	FindByID(id int) (*types.Movie, error)
	Update(m *types.Movie) error
	Delete(id int) error
}

func NewStore(db *gorm.DB) Store {
	return &store{db}
}

func (s *store) FindAll() ([]types.Movie, error) {

	var movies []types.Movie
	if err := s.db.Find(&movies).Error; err != nil {
		return nil, err
	}

	return movies, nil
}

func (s *store) Create(m *types.Movie) error {
	return s.db.Create(m).Error
}

func (s *store) FindByID(id int) (*types.Movie, error) {
	var movie types.Movie
	if err := s.db.First(&movie, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &movie, nil
}

func (s *store) Update(m *types.Movie) error {

	if err := s.db.First(&types.Movie{}, "id = ?", m.ID).Error; err != nil {
		return err
	}

	return s.db.Updates(m).Error
}

func (s *store) Delete(id int) error {

	if err := s.db.First(&types.Movie{}, "id = ?", id).Error; err != nil {
		return err
	}

	return s.db.Delete(&types.Movie{}, "id = ?", id).Error
}
