package mocks

import (
	"time"

	"github.com/Fairuzzzzz/clipvault/internal/models"
)

var mockClip = &models.Clip{
	ID:      1,
	Title:   "Test",
	Content: "Test",
	Created: time.Now(),
	Expires: time.Now(),
}

type ClipModel struct{}

func (m *ClipModel) Insert(title, content string, expires int) (int, error) {
	return 2, nil
}

func (m *ClipModel) Get(id int) (*models.Clip, error) {
	switch id {
	case 1:
		return mockClip, nil
	default:
		return nil, models.ErrNoRecord
	}
}

func (m *ClipModel) Latest() ([]*models.Clip, error) {
	return []*models.Clip{mockClip}, nil
}
