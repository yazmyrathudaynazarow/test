package posts

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
	"test/project/internal/handlers"
	"test/project/internal/posts/models"
	"test/project/pkg/config"
)

const (
	getVideosURL   = "/get-videos"
	createVideoURL = "/create-video"
	deleteVideoURL = "/delete-video/:id"
)

type handler struct {
	repository *Repository
	cfg        *config.Configs
}

func NewHandler(repository *Repository, cfg *config.Configs) handlers.Handler {
	return &handler{
		repository: repository,
		cfg:        cfg,
	}
}

func (h *handler) Register(router fiber.Router) {
	router.Get(getVideosURL, h.GetVideos)
	router.Post(createVideoURL, h.CreateVideo)
	router.Post(deleteVideoURL, h.DeleteVideo)
}

// GetVideos GetAllVideos GetVideos
// @Summary Get all videos
// @Description Retrieve a list of videos from the database
// @Tags Videos
// @Produce json
// @Success	200 {object} string "Success"
// @Failure 500 {object} string "Success"
// @Router /videos/get-videos [get]
func (h *handler) GetVideos(c *fiber.Ctx) error {
	videos, err := h.repository.GetVideos()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": videos,
	})
}

// CreateVideo AddVideo CreateVideo
// @Summary Create Video
// @Description Add New Video
// @Tags Videos
// @Produce json
// @Param video body models.Video true "Video Data"
// @Success 200 {object} string "Success"
// @Failure 500 {object} string "Success"
// @Router /videos/create-video [post]
func (h *handler) CreateVideo(c *fiber.Ctx) error {
	var video models.Video

	if err := c.BodyParser(&video); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}
	err := h.repository.CreateVideo(video)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "SUCCESS",
	})
}

// DeleteVideo Delete DeleteVideo
// @Summary Delete Video
// @Description Delete Video
// @Tags Videos
// @Produce json
// @Param id path int true "Video ID"
// @Success 200 {object} string "Success"
// @Failure 500 {object} string "Internal Server Error"
// @Router /videos/delete-video/{id} [post]
func (h *handler) DeleteVideo(c *fiber.Ctx) error {
	var (
		id  models.ID
		err error
	)

	id.ID, err = strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid id",
		})
	}

	err = h.repository.DeleteVideo(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "DELETE SUCCESS",
	})
}
