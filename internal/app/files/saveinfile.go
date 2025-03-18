package files

import (
	"encoding/json"
	"go/rest/internal/models"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func SaveInFile(ctx *gin.Context) {

	data, err := json.MarshalIndent(models.Tasklist, "", "\t")
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
	}

	dataYaml := OpenYaml()

	err = os.WriteFile(dataYaml["filename"], data, 0644)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
	}
}
