package bookmarks

import (
	"context"

	"github.com/mrusme/xbsapi/ent/bookmark"

	"github.com/gofiber/fiber/v2"
)

type BookmarkShowResponse struct {
	Success  bool               `json:"success"`
	Bookmark *BookmarkShowModel `json:"bookmark"`
	Message  string             `json:"message"`
}

// Show godoc
// @Summary      Show a bookmark
// @Description  Get bookmark by ID
// @Tags         bookmarks
// @Accept       json
// @Produce      json
// @Param        id   path      string true "Bookmark ID"
// @Success      200  {object}  BookmarkShowResponse
// @Failure      400  {object}  BookmarkShowResponse
// @Failure      404  {object}  BookmarkShowResponse
// @Failure      500  {object}  BookmarkShowResponse
// @Router       /bookmarks/{id} [get]
// @security     BasicAuth
func (h *handler) Show(ctx *fiber.Ctx) error {
	var err error

	id, err := h.getUUIDFromID(ctx.Params("id"))
	if err != nil {
		return ctx.
			Status(fiber.StatusBadRequest).
			JSON(fiber.Map{
				"success": false,
				"message": err.Error(),
			})
	}

	dbBookmark, err := h.entClient.Bookmark.
		Query().
		Where(
			bookmark.ID(id),
		).
		Only(context.Background())
	if err != nil {
		return ctx.
			Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{
				"success": false,
				"message": err.Error(),
			})
	}

	showBookmark := BookmarkShowModel{
		Bookmarks:   dbBookmark.Bookmarks,
		LastUpdated: dbBookmark.LastUpdated.Format(LAST_UPDATED_FORMAT),
		Version:     dbBookmark.Version,
	}

	return ctx.
		Status(fiber.StatusOK).
		JSON(showBookmark)
}
