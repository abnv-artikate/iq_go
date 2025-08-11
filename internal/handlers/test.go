package handlers

import (
	"iq-go/internal/services"
	"iq-go/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TestHandler struct {
	testService *services.TestService
}

func NewTestHandler(testService *services.TestService) *TestHandler {
	return &TestHandler{
		testService: testService,
	}
}

type SubmitTestRequest struct {
	TestID    uint                  `json:"test_id" binding:"required"`
	Answers   []SubmitAnswerRequest `json:"answers" binding:"required"`
	TimeTaken int                   `json:"time_taken"`
}

type SubmitAnswerRequest struct {
	QuestionID   uint   `json:"question_id" binding:"required"`
	UserAnswer   string `json:"user_answer"`
	ResponseTime int    `json:"response_time"`
}

func (h *TestHandler) GetQuestions(c *gin.Context) {
	testIDStr := c.DefaultQuery("test_id", "1")
	testID, err := strconv.ParseUint(testIDStr, 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid test ID")
		return
	}

	questions, err := h.testService.GetQuestionsByTestID(uint(testID))
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch questions")
		return
	}

	// Remove correct answers from response
	for i := range questions {
		questions[i].CorrectAnswer = ""
	}

	utils.SuccessResponse(c, http.StatusOK, "Questions fetched successfully", questions)
}

func (h *TestHandler) SubmitTest(c *gin.Context) {
	var req SubmitTestRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		utils.ErrorResponse(c, http.StatusUnauthorized, "User not authenticated")
		return
	}

	// Convert handlers types to services types
	serviceAnswers := make([]services.SubmitAnswerRequest, len(req.Answers))
	for i, answer := range req.Answers {
		serviceAnswers[i] = services.SubmitAnswerRequest{
			QuestionID:   answer.QuestionID,
			UserAnswer:   answer.UserAnswer,
			ResponseTime: answer.ResponseTime,
		}
	}

	result, err := h.testService.SubmitTest(userID.(uint), req.TestID, serviceAnswers, req.TimeTaken)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to submit test")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Test submitted successfully", result)
}
