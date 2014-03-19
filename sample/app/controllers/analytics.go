package controllers

import (
	"fmt"
	"time"

	"github.com/iassic/revel-modz/modules/analytics"
	"github.com/iassic/revel-modz/modules/user"
	"github.com/revel/revel"
)

func (c App) RecordPageRequest() revel.Result {

	now := time.Now()

	var U *user.UserBasic
	var V *user.Visitor

	u := c.RenderArgs["user_basic"]
	if u != nil {
		U = u.(*user.UserBasic)
		// revel.INFO.Printf("user:\n%+v\n", U)
		err := analytics.SaveUserPageRequest(c.Txn, U.UserId, now, c.Request.Request)
		checkERROR(err)
		return nil
	}
	v := c.RenderArgs["visitor"]
	if v != nil {
		V = v.(*user.Visitor)
		// revel.INFO.Printf("visitor:\n%+v\n", V)
		err := analytics.SaveVisitorPageRequest(c.Txn, V.VisitorId, now, c.Request.Request)
		checkERROR(err)
		return nil
	}

	revel.ERROR.Println("Shouldn't get here, means there was no user or visitor")
	return nil
}

func (c App) RecordPageEvent() revel.Result {
	return nil
}

// helper functions
func (c App) addNewVisitor() (*user.Visitor, error) {

	pr, err := analytics.ParsePageRequest(c.Request.Request)
	checkERROR(err)
	// revel.WARN.Printf("PageReq: \n%+v\n", pr)

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
	// revel.WARN.Printf("PageReq: \n%+v\n", pr)

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

func (c App) updateVisitorWithUserIdPanic() {
	revel.ERROR.Println("PANIC!!! visitorId <-> userId not implemented")
	// panic()
}
