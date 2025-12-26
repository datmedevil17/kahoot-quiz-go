package user

import (
	"net/http"
	"strconv"

	"github.com/datmedevil17/kahoot-quiz-go/internal/services/user"

	"github.com/datmedevil17/kahoot-quiz-go/internal/utils"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service   *user.Service
	jwtSecret string
}

func NewHandler(service *user.Service, jwtSecret string) *Handler {
	return &Handler{
		service:   service,
		jwtSecret: jwtSecret,
	}
}

func (h *Handler) SignUp(c *gin.Context) {
	var req SignupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Enter valid email")
	}
	user, err := h.service.CreateUser(req.Email, req.Name, req.Password, req.Role)
	if err != nil {
		if err.Error() == "user already exists" {
			utils.ErrorResponse(c, http.StatusConflict, err.Error())
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	token, err := utils.GenerateToken(user.Email, user.ID, h.jwtSecret)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Error generating token")
		return
	}
	c.String(http.StatusOK, token)
}

func (h *Handler) SignIn(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Enter valid email")
	}
	user, err := h.service.AuthenticateUser(req.Email, req.Password)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	token, err := utils.GenerateToken(user.Email, user.ID, h.jwtSecret)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Error generating token")
		return
	}
	c.String(http.StatusOK, token)
}

func (h *Handler) GetUserById(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid userId")
		return
	}

	user, err := h.service.GetUserById(uint(userID))
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "User not found")
		return
	}

	c.JSON(http.StatusOK, user)

}

func (h *Handler) GetCurrentUserId(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid userId")
		return
	}
	user, err := h.service.GetUserById(uint(userID))
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "User not found")
		return
	}
	c.JSON(http.StatusOK, user)

}

func (h *Handler) ViewProfile(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid userId")
		return
	}
	user, err := h.service.GetUserProfile(uint(userID))
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "User not found")
		return
	}
	c.JSON(http.StatusOK, user)

}

func (h *Handler) GetProfile(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Error occured in profile.")
		return
	}
	user, err := h.service.GetUserProfile(uint(userID))
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "User not found")
		return
	}
	c.JSON(http.StatusOK, user)

}
