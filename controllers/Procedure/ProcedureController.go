package procedure_controllers

import (
	"his_apis_go/initializers"
	procedure_models "his_apis_go/models/Procedure"

	"github.com/gin-gonic/gin"
)

func ProcedureInfoIndex(c *gin.Context) {
	var procedures []procedure_models.ProcedureInfo

	result := initializers.DB.Find(&procedures)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data":  procedures,
		"count": len(procedures),
	})
}
