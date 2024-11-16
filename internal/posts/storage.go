package posts

import (
	"github.com/gocql/gocql"
	postsDb "test/project/internal/posts/db"
	"test/project/internal/posts/models"
)

type Db interface {
	GetVideos() ([]models.Video, error)
	CreateVideo(video models.Video) error
	DeleteVideo(id models.ID) error
}

type Repository struct {
	Db
}

func NewRepository(db *gocql.Session) *Repository {
	return &Repository{
		Db: postsDb.NewDbRepo(db),
	}
}
