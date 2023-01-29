package handlers

import (
	"context"
	"hamkorbank/api/http"
	"hamkorbank/models"
	"math/rand"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

//@Router /v1/phone [GET]
//@Summary This api for getting all phones
//@Description This api for getting all phones from db
//@Accept json
//@Produce json
//@Param limit query int false "limit"
//@Param page query int false "page"
//Success 200 {object} http.Response{data=string} "Response data"
//Failure 400 {object} http.Response
func (h *Handler) GetAllPhones(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", ""))
	page, _ := strconv.Atoi(c.DefaultQuery("page", ""))

	res, err := h.storage.Phone().GetAllPhones(context.Background(), &models.GetAllPhonesRequest{
		Limit:  limit,
		Offset: page,
	})

	if err != nil {
		h.handleResponse(c, http.InternalServerError, err)
		return
	}

	h.handleResponse(c, http.OK, res)
}

//@Router /v1/phone/{id} [GET]
//@Summary Get Phone number by id
//@Description Getting phone number by their id
//@Accept json
//@Produce json
//@Param id path string true "id"
//@Success 200 {object} http.Response{data=string} "Response data"
//@Success 500 {object} http.Response
func (h *Handler) GetPhoneById(c *gin.Context) {
	id := c.Param("id")

	res, status, err := h.storage.Phone().GetPhoneById(context.Background(), &models.GetByIdPhoneRequest{
		ID: id,
	})
	if status {
		h.handleResponse(c, http.NotFound, err)
	}

	if err != nil {
		h.handleResponse(c, http.InternalServerError, err)
	}

	rand.Seed(time.Now().UnixNano())
	random := rand.Int()

	if random%5 == 0 {
		h.handleResponse(c, http.InternalServerError, err)
		return
	}

	h.handleResponse(c, http.OK, res)
}
