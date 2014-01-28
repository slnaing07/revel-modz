package controllers

import "github.com/robfig/revel"

func init() {
	revel.OnAppStart(GorpInit)
	revel.InterceptMethod((*GorpController).Begin, revel.BEFORE)
	revel.InterceptMethod((*User).getUserLogin, revel.BEFORE)
	revel.InterceptMethod((*GorpController).Commit, revel.AFTER)
	revel.InterceptMethod((*GorpController).Rollback, revel.FINALLY)

	revel.TemplateFuncs["eq"] = func(a, b interface{}) bool { return a == b }

}
