package http

import (
	"SuperPanel/domain"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

type ResponseError struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

type UserController struct {
	UUseCase domain.UserUsecase
}

func NewUserController(e *fiber.App, us domain.UserUsecase) {
	handler := &UserController{
		UUseCase: us,
	}

	e.Get("/users/:id", handler.GetByID)
	//e.POST("/articles", handler.Store)
	//e.GET("/articles/:id", handler.GetByID)
	//e.DELETE("/articles/:id", handler.Delete)
}
func (this *UserController) GetByID(c *fiber.Ctx) error {
	i, err := strconv.ParseInt(c.Params("id"), 10, 32)
	if err != nil {
		panic(err)
	}
	Id := int32(i)
	ctx := c.Context()
	user, err := this.UUseCase.GetByID(ctx, Id)
	if err != nil {
		return c.JSON(ResponseError{Message: err.Error()})
	}
	b, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}
	return c.Status(fiber.StatusOK).SendString(string(b))

}
func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
