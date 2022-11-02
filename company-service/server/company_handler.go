package server

import (
	"company-service/models"
	"company-service/server/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) HandleCreateCompany() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.Company
		if err := decode(c, &user); err != nil {
			response.JSON(c, "", http.StatusBadRequest, nil, err)
			return
		}
		userResponse, err := s.AuthService.CreateCompany(&user)
		if err != nil {
			err.Respond(c)
			return
		}
		response.JSON(c, "Company Creation successful", http.StatusCreated, userResponse, nil)
	}
}

//
//func (s *Server) HandleUpdateCompany() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		_, user, err := GetValuesFromContext(c)
//		if err != nil {
//			err.Respond(c)
//			return
//		}
//		medicationID, errr := strconv.ParseUint(c.Param("medicationID"), 10, 32)
//		if errr != nil {
//			response.JSON(c, "invalid ID", http.StatusBadRequest, nil, errr)
//			return
//		}
//		var updateMedicationRequest models.UpdateMedicationRequest
//		if err := decode(c, &updateMedicationRequest); err != nil {
//			response.JSON(c, "", http.StatusBadRequest, nil, err)
//			return
//		}
//		err = s.MedicationService.UpdateMedication(&updateMedicationRequest, uint(medicationID), user.ID)
//		if err != nil {
//			err.Respond(c)
//			return
//		}
//		response.JSON(c, "medication updated successfully", http.StatusOK, nil, nil)
//	}
//}
//
//func (s *Server) HandleGetCompanyDetails() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		_, user, err := GetValuesFromContext(c)
//		if err != nil {
//			err.Respond(c)
//			return
//		}
//		id := c.Param("id")
//		userId, errr := strconv.ParseUint(id, 10, 32)
//		if errr != nil {
//			response.JSON(c, "error parsing id", http.StatusBadRequest, nil, errr)
//			return
//		}
//		medication, err := s.MedicationService.GetMedicationDetail(uint(userId), user.ID)
//		if err != nil {
//			response.JSON(c, "", http.StatusInternalServerError, nil, errors.New("internal server error", http.StatusInternalServerError))
//			return
//		}
//		response.JSON(c, "retrieved medications successfully", http.StatusOK, gin.H{"medication": medication}, nil)
//	}
//}
//
//func (s *Server) handleDeleteCompany() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		_, user, err := GetValuesFromContext(c)
//		if err != nil {
//			err.Respond(c)
//			return
//		}
//
//		if err := s.AuthService.DeleteUserByEmail(user.Email); err != nil {
//			err.Respond(c)
//			return
//		}
//
//		response.JSON(c, "user successfully deleted", http.StatusOK, nil, nil)
//	}
//}
