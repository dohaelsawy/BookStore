package api

import (
	"net/http"
	"regexp"
	"strconv"
	db "github.com/dohaelsawy/bookStore/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)


type BookRequest struct {
	Name        string `json:"name" validate:"required"`
	PublishDate string `json:"publish_date" validate:"required"`
	Price       int32  `json:"price" validate:"required"`
	SKU         string `json:"sku" validate:"required,sku"`
	Description string `json:"description" validate:"required"`
	Author      string `json:"author" validate:"required"`
}

func validateSKU(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := re.FindAllString(fl.Field().String(), -1)
	if len(matches) != 1 {
		return false
	}
	return true
}

func (b *BookRequest) Validate() error {
validate := validator.New(validator.WithRequiredStructEnabled())
	validate.RegisterValidation("sku", validateSKU)
	if err := validate.Struct(b) ; err != nil {
		return err	
	}
	return nil 
}

func (server *Server) addBook(ctx *gin.Context) {
	server.l.Println("handle adding new book ... you're good")
	var req BookRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponce(err))
		return
	}
	if err := req.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponce(err))
		return
	}

	arg := db.CreateBookParams{
		Name:        req.Name,
		PublishDate: req.PublishDate,
		Price:       req.Price,
		Sku:         req.SKU,
		Description: req.Description,
		Author:      req.Author,
	}
	

	book, err := server.store.CreateBook(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponce(err))
		return
	}

	ctx.JSON(http.StatusOK, book)
}

func (server *Server) updateBook(ctx *gin.Context) {
	server.l.Println("handle updating book ... you're good")
	bookIDReq := ctx.Param("id")
	bookID, err := strconv.ParseInt(bookIDReq, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponce(err))
		return
	}

	var req db.UpdateBookParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponce(err))
		return
	}
	val := BookRequest{
		Name:        req.Name,
		PublishDate: req.PublishDate,
		Price:       req.Price,
		SKU:         req.Sku,
		Description: req.Description,
		Author:      req.Author,
	}
	if err := val.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponce(err))
		return
	}
	arg := db.UpdateBookParams{
		BookID:      int32(bookID),
		Name:        req.Name,
		PublishDate: req.PublishDate,
		Price:       req.Price,
		Sku:         req.Sku,
		Description: req.Description,
		Author:      req.Author,
	}

	book, err := server.store.UpdateBook(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponce(err))
		return
	}

	ctx.JSON(http.StatusOK, book)
}

func (server *Server) deleteBook(ctx *gin.Context) {
	server.l.Println("handle deleting book ... you're good")
	bookIDReq := ctx.Param("id")
	bookID, err := strconv.ParseInt(bookIDReq, 10, 32)
	if err != nil {
		// Handle the error (e.g., invalid integer format)
		ctx.JSON(http.StatusBadRequest, errorResponce(err))
		return
	}

	err = server.store.DeleteBook(ctx, int32(bookID))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponce(err))
		return
	}

	ctx.JSON(http.StatusOK, "DONE DELETED ... you're good")
}

func (server *Server) getBook(ctx *gin.Context) {
	server.l.Println("handle getting book ... you're good")
	bookIDReq := ctx.Param("id")
	bookID, err := strconv.ParseInt(bookIDReq, 10, 32)
	if err != nil {
		// Handle the error (e.g., invalid integer format)
		ctx.JSON(http.StatusBadRequest, errorResponce(err))
		return
	}

	book, err := server.store.GetBook(ctx, int32(bookID))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponce(err))
		return
	}

	ctx.JSON(http.StatusOK, book)
}

func (server *Server) listBooks(ctx *gin.Context) {
	server.l.Println("handle getting booksss ... you're good")
	books, err := server.store.ListBooks(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponce(err))
		return
	}

	ctx.JSON(http.StatusOK, books)
}
