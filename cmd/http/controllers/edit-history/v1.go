package edit_history

import (
	"github.com/gin-gonic/gin"
	edit_history "mono-base/internal/usecases/edit-history"
	"net/http"
	"strconv"
)

var _ Controller = (*ControllerV1)(nil)

type ControllerV1 struct {
	getInfoUseCase edit_history.GetInfoInputUseCase
}

func NewEditHistoryControllerV1(getInfoUseCase edit_history.GetInfoInputUseCase) *ControllerV1 {
	return &ControllerV1{getInfoUseCase: getInfoUseCase}
}

func (u *ControllerV1) GetInfo(ctx *gin.Context) {
	modelIdStr := ctx.Query("modelId")
	if modelIdStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "modelId is required"})
		return
	}

	modelId, err := strconv.ParseUint(modelIdStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid modelId"})
		return
	}

	input := edit_history.GetInfoInput{
		ModelId: uint(modelId),
	}

	results, err := u.getInfoUseCase.Execute(ctx.Request.Context(), input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, results)
}
