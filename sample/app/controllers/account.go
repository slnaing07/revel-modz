package controllers

import (
	// "github.com/iassic/revel-modz/modules/user"
	"github.com/revel/revel"
)

func (c User) Account() revel.Result {
	// u := c.userConnected()

	// // get stuff from DB
	// userbasic := getU

	// // create & file in UserRegister struct
	// userregister := &UserRegister {
	// 	...
	// }

	// c.RenderArgs["userregister"] = userregister

	return c.Render()
}
