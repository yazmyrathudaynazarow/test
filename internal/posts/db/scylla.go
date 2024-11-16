package postsDb

import (
	"fmt"
	"github.com/gocql/gocql"
	"test/project/internal/posts/models"
)

type DbRepo struct {
	db *gocql.Session
}

func NewDbRepo(db *gocql.Session) *DbRepo {
	return &DbRepo{
		db: db,
	}
}

// GetVideos LIST VIDEOS
func (r *DbRepo) GetVideos() ([]models.Video, error) {
	var results []models.Video

	query := r.db.Query("SELECT id, duration, title FROM test.videos").Consistency(gocql.One)
	iter := query.Iter()

	var video models.Video
	for iter.Scan(&video.ID, &video.Duration, &video.Title) {
		results = append(results, video)
	}

	if err := iter.Close(); err != nil {
		fmt.Println("err in getVideos--->", err)
		return nil, err
	}

	return results, nil
}

// CreateVideo CREATE VIDEO
func (r *DbRepo) CreateVideo(video models.Video) error {
	query := r.db.Query("INSERT INTO test.videos (id, title, duration) VALUES (?, ?, ?)", video.ID, video.Title, video.Duration).Consistency(gocql.One)

	err := query.Exec()
	if err != nil {
		fmt.Println("err in createVideo--->", err)
	}

	return err
}

// DeleteVideo Delete VIDEO
func (r *DbRepo) DeleteVideo(id models.ID) error {
	query := r.db.Query("DELETE FROM test.videos WHERE id = ?", id.ID).Consistency(gocql.One)

	err := query.Exec()
	if err != nil {
		fmt.Println("err in deleteVideo--->", err)
		return err
	}

	return nil
}
