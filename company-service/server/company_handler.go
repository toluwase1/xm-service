package server

import (
	"company-service/models"
	"company-service/server/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) CreateCompany() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.Company
		if err := decode(c, &user); err != nil {
			response.JSON(c, "", http.StatusBadRequest, nil, err)
			return
		}
		userResponse, err := s.AuthService.SignupUser(&user)
		if err != nil {
			err.Respond(c)
			return
		}
		response.JSON(c, "Signup successful, check your email for verification", http.StatusCreated, userResponse, nil)
	}
}

func (s *Server) UpdateCompany() gin.HandlerFunc {

}

func (s *Server) DeleteCompany() gin.HandlerFunc {

}

func (s *Server) GetCompanyByName() gin.HandlerFunc {

}
