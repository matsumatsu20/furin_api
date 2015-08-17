package controllers

import (
	"fmt"
	"github.com/robfig/revel"
	"furin_api/app/models"
)

type User struct {
	*revel.Controller
}

func (c User) Index() revel.Result {
	for i := 0; i < 10; i++ {
		DbMap.Insert(&models.User{0, fmt.Sprintf("user%d", i)})
	}

	rows, _ := DbMap.Select(models.User{}, "select * from user")
	for _, row := range rows {
		user := row.(*models.User)
		fmt.Printf("%d, %s\n", user.Id, user.Name)
	}

	return c.Render()
}