package controllers

import (
	"fmt"

	"github.com/iassic/revel-modz/modules/analytics"
	"github.com/iassic/revel-modz/modules/user"
	"github.com/revel/revel"
)

func (c App) RecordPageRequest() revel.Result {

	pr, err := analytics.ParsePageRequest(c.Request.Request)
	checkERROR(err)

	var U *user.UserBasic
	var V *user.Visitor

	u := c.RenderArgs["user_basic"]
	if u != nil {
		U = u.(*user.UserBasic)
	}
	v := c.RenderArgs["visitor"]
	if v != nil {
		V = v.(*user.Visitor)
	}

	revel.INFO.Printf("PageReq:\n%+v\n", pr)
	revel.INFO.Printf("user:\n%+v\n", U)
	revel.INFO.Printf("visitor:\n%+v\n", V)

	return nil
}

func (c App) AnalyticsPost() revel.Result {
	return nil
}

// Admin functions
func (c Admin) AnalyticsView() revel.Result {
	analytic_data := "dummy"

	return c.Render(analytic_data)
}

// helper functions
func (c App) addNewVisitor() (*user.Visitor, error) {

	pr, err := analytics.ParsePageRequest(c.Request.Request)
	checkERROR(err)
	revel.WARN.Printf("PageReq: \n%+v\n", pr)

	vid, err := user.GenerateNewVisitorId(c.Txn)
	checkERROR(err)

	ip := "missing"
	if pr.XRealIp != "" {
		ip = pr.XRealIp
	}

	err = user.AddVisitor(c.Txn, vid, ip)
	checkERROR(err)

	v := &user.Visitor{
		VisitorId: vid,
		VisitorIp: ip,
	}

	c.Session["v"] = fmt.Sprint(v.VisitorId)
	c.RenderArgs["visitor"] = v

	return v, nil
}

func (c App) updateVisitor(v *user.Visitor) error {

	pr, err := analytics.ParsePageRequest(c.Request.Request)
	checkERROR(err)
	revel.WARN.Printf("PageReq: \n%+v\n", pr)

	// check ip addresses and do something
	ip := "missing"
	if pr.XRealIp != "" {
		ip = pr.XRealIp
	}

	if ip != v.VisitorIp {
		revel.INFO.Println("visitor connecting from new ip")

		v.VisitorIp = ip

		err = user.UpdateVisitor(c.Txn, v)
		checkERROR(err)
	}

	c.RenderArgs["visitor"] = v

	return nil
}
