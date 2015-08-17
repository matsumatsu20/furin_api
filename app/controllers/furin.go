package controllers
import (
	"furin_api/app/models"
	"github.com/robfig/revel"
	"fmt"
)


type Furin struct {
	*revel.Controller
}

func (c Furin) Index() revel.Result {
	var furin (*models.Furin)

	rows, _ := DbMap.Select(models.Furin{}, "select * from furin")
	for _, row := range rows {
		furin = row.(*models.Furin)
		fmt.Printf("%d, %s\n", furin.Id, furin.Name, furin.Description)
	}

	return c.RenderJson(rows)
}