package controllers

import (
	models "github.com/iassic/revel-modz/moduls/auth/app/models"
)

func init() {
	// add interceptors
	revel.InterceptMethod((*User).getUserLogin, revel.BEFORE)
	revel.InterceptMethod(Application.AddUser, revel.BEFORE)
	revel.InterceptMethod(Hotels.checkUser, revel.BEFORE)

	revel.InterceptMethod((*GorpController).Begin, revel.BEFORE)
	revel.InterceptMethod((*GorpController).Commit, revel.AFTER)
	revel.InterceptMethod((*GorpController).Rollback, revel.FINALLY)

	// add module Init funcs to startup phase
	revel.OnAppStart(InitAuthDb)

}
