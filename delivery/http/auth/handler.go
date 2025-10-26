package auth

import (
	"encoding/json"
	"net/http"
	"time"
	"web-api/delivery/http/response"
	auth_maker "web-api/usecase/auth"
)

type AuthHandler struct {
	tokenMaker auth_maker.TokenMaker
}

func NewAuthHandler(maker auth_maker.TokenMaker) *AuthHandler {
	return &AuthHandler{
		tokenMaker: maker,
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}
	access_token, err := h.tokenMaker.CreateToken(req.Email, "read write", time.Minute, auth_maker.TokenTypeAccessToken)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	refresh_token, err := h.tokenMaker.CreateToken(req.Email, "", time.Hour*24*7, auth_maker.TokenTypeRefreshToken)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	response.Success(w, http.StatusOK, LoginResponse{
		AccessToken:  access_token,
		RefreshToken: refresh_token,
	})

}

func (h *AuthHandler) Register() {

}

func (h *AuthHandler) Logout() {

}

func (h *AuthHandler) RefreshToken() {

}

func (h *AuthHandler) ValidateToken() {

}

func (h *AuthHandler) ForgotPassword() {

}

func (h *AuthHandler) ResetPassword() {

}

func (h *AuthHandler) ChangePassword() {

}

func (h *AuthHandler) VerifyEmail() {

}
