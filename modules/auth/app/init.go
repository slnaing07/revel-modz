package controllers

import (
	"github.com/robfig/revel"

	// models "github.com/iassic/revel-modz/modules/auth/app/models"
	ctrls "github.com/iassic/revel-modz/modules/auth/app/controllers"
)

func init() {
	// add interceptors
	revel.InterceptMethod((*ctrls.AuthController).GetUserLogin, revel.BEFORE)
	// revel.InterceptMethod(Application.AddUser, revel.BEFORE)
	// revel.InterceptMethod(Hotels.checkUser, revel.BEFORE)

	revel.InterceptMethod((*ctrls.AuthController).Begin, revel.BEFORE)
	revel.InterceptMethod((*ctrls.AuthController).Commit, revel.AFTER)
	revel.InterceptMethod((*ctrls.AuthController).Rollback, revel.FINALLY)

	// add module Init funcs to startup phase
	revel.OnAppStart(InitAuthDb)

}
