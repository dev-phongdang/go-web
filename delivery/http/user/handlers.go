package user

import (
	"errors"
	"net/http"
	"web-api/delivery/http/response"
	"web-api/usecase/user"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type UserHandler struct {
	uc *user.UserUsecase
}

// NewUserHandler creates a new user handler.
func NewUserHandler(uc *user.UserUsecase) *UserHandler {
	return &UserHandler{
		uc: uc,
	}
}

// GetUserByID retrieves a user by their ID.
func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}
	u, err := h.uc.GetUserByID(id)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	if u == nil {
		response.Error(w, http.StatusNotFound, errors.New("user not found"))
		return
	}
	userDTO := UserResponse{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}

	response.Success(w, http.StatusOK, userDTO)
}
