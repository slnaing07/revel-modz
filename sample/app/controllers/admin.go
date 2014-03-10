package controllers

// TEMPLATE FILE

import (
	"github.com/iassic/revel-modz/modules/analytics"
	"github.com/revel/revel"

	"github.com/iassic/revel-modz/sample/app/routes"
)

type Admin struct {
	User
}

// moving towards RBAC here...
func (c Admin) CheckLoggedIn() revel.Result {
	u := c.userConnected()
	if u == nil {
		c.Flash.Error("Please log in first")
		return c.Redirect(routes.App.Login())
	}

	// look up role in RBAC module
	isAdmin := u.UserName == "admin@domain.com"

	if !isAdmin {
		return c.Redirect(routes.App.Index())
	}

	// set up things for an admin role
	c.Session["admin"] = "true"

	return nil
}

func (c Admin) Index() revel.Result {
	return c.Render()
}

// defined in maillist.go
// func (c Admin) MaillistView() revel.Result
// func (c Admin) MaillistCompose() revel.Result {

// Admin functions
func (c Admin) AnalyticsView() revel.Result {
	return c.Render()
}

func (c Admin) AnalyticsFilter(group, id string) revel.Result {
	results, err := analytics.GetAllUserPageRequests(c.Txn)
	checkERROR(err)

	revel.WARN.Println("len(page_requests) =", len(results))

	return c.RenderJson(results)
}
