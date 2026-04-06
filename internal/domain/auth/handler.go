package auth

// "gotron/internal/common"

// "github.com/mfaisal-Ash/gotron/internal/common"

// type Handler struct {
// 	service *Service
// }

// func NewHandler(service *Service) *Handler {
// 	return &Handler{
// 		service: service,
// 	}
// }

// func (h *Handler) Register(c *gin.Context) {
// 	var req RegisterRequest
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		common.JSONErrorResponse(c, http.StatusBadRequest, "Invalid request payload", map[string]interface{}{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	resp, err := h.service.Register(req)
// 	if err != nil {
// 		common.JSONErrorResponse(c, http.StatusBadRequest, err.Error(), nil)
// 		return
// 	}
// 	common.JSONSuccessResponse(c, http.StatusCreated, "User registered successfully", resp, nil)
// }

// func (h *Handler) Login(c *gin.Context) {
// 	var req LoginRequest
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		common.JSONErrorResponse(c, http.StatusBadRequest, "Invalid request payload", map[string]interface{}{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	resp, err := h.service.Login(req)
// 	if err != nil {
// 		common.JSONErrorResponse(c, http.StatusUnauthorized, err.Error(), nil)
// 		return
// 	}
// 	common.JSONSuccessResponse(c, http.StatusOK, "Login successful", resp, nil)
// }

// func (h *Handler) RefreshToken(c *gin.Context) {
// 	var req RefreshTokenRequest
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		common.JSONErrorResponse(c, http.StatusBadRequest, "Invalid request payload", map[string]interface{}{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	resp, err := h.service.RefreshToken(req)
// 	if err != nil {
// 		common.JSONErrorResponse(c, http.StatusUnauthorized, err.Error(), nil)
// 		return
// 	}
// 	common.JSONSuccessResponse(c, http.StatusOK, "Token refreshed successfully", resp, nil)
// }

// func (h *Handler) Logout(c *gin.Context) {
// 	var req LogoutRequest
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		common.JSONErrorResponse(c, http.StatusBadRequest, "Invalid request payload", map[string]interface{}{
// 			"error": err.Error(),
// 		})
// 		return
// 	}
// 	if err := h.service.Logout(req); err != nil {
// 		common.JSONErrorResponse(c, http.StatusInternalServerError, err.Error(), nil)
// 		return
// 	}

// 	common.JSONSuccessResponse(c, http.StatusOK, "Logout successful", nil, nil)
// }

// func (h *Handler) Me(c *gin.Context) {
// 	userID, exists := c.Get("user_id")
// 	if !exists {
// 		common.JSONErrorResponse(c, http.StatusUnauthorized, "User not authenticated", nil)
// 		return
// 	}

// 	resp, err := h.service.GetUserInfo(userID.(string))
// 	if err != nil {
// 		common.JSONErrorResponse(c, http.StatusInternalServerError, err.Error(), nil)
// 		return
// 	}
// 	common.JSONSuccessResponse(c, http.StatusOK, "User info retrieved successfully", resp, nil)
// }
