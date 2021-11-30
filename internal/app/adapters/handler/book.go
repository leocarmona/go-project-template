package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/leocarmona/go-project-template/internal/app/domain"
	"github.com/leocarmona/go-project-template/internal/app/transport/inbound"
	"github.com/leocarmona/go-project-template/internal/app/transport/mapper"
	"github.com/leocarmona/go-project-template/internal/app/transport/presenter"
	"net/http"
	"strconv"
)

type BookHandler struct {
	services *domain.Services
}

func NewBookHandler(services *domain.Services) *BookHandler {
	return &BookHandler{
		services: services,
	}
}

func (h *BookHandler) Configure(server *echo.Echo) {
	server.POST("/books", h.CreateBook)
	server.GET("/books/:id", h.ReadBookById)
	server.PUT("/books/:id", h.UpdateBookById)
	server.DELETE("/books/:id", h.DeleteBookById)
	server.GET("/books", h.ListBooks)
}

func (h *BookHandler) CreateBook(c echo.Context) error {
	request := new(inbound.CreateBookRequest)
	if err := c.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	book := mapper.BookFromCreateBookRequest(request)
	err := h.services.Book.Create(c.Request().Context(), book)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, presenter.CreateBook(book))
}

func (h *BookHandler) ReadBookById(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	book, err := h.services.Book.ReadById(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	if book == nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, presenter.ReadBookById(book))
}

func (h *BookHandler) UpdateBookById(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	request := new(inbound.UpdateBookRequest)
	if err := c.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	updated, err := h.services.Book.UpdateById(c.Request().Context(), mapper.BookFromUpdateBookRequest(id, request))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, presenter.UpdateBookById(updated))
}

func (h *BookHandler) DeleteBookById(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	deleted, err := h.services.Book.DeleteById(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, presenter.DeleteBookById(deleted))
}

func (h *BookHandler) ListBooks(c echo.Context) error {
	books, err := h.services.Book.List(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, presenter.ListBooks(books))
}
