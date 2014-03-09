package controllers

import (
	// "github.com/iassic/revel-modz/modules/analytics"
	// "github.com/iassic/revel-modz/modules/user"
	"github.com/revel/revel"
)

func (c App) AnalyticsPost() revel.Result {
	return nil
}

// Admin functions
func (c Admin) AnalyticsView() revel.Result {
	analytic_data := "dummy"

	return c.Render(analytic_data)
}
