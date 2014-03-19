package controllers

import (
	// "github.com/iassic/revel-modz/modules/user"
	"github.com/revel/revel"
	"github.com/iassic/revel-modz/modules/user"
	"github.com/iassic/revel-modz/sample/app/models"
	"github.com/iassic/revel-modz/modules/auth"
	"github.com/iassic/revel-modz/modules/maillist"

	"github.com/iassic/revel-modz/sample/app/routes"


)

func (c User) Account() revel.Result {

	u:= c.userConnected()
	UA, err := user.GetUserAddressesById(c.Txn, u.UserId)
	checkERROR(err)
	UD, err := user.GetUserDetailById(c.Txn, u.UserId)
	checkERROR(err)
	UP, err := user.GetUserPhonesById(c.Txn, u.UserId)
	checkERROR(err)

	ur := &models.UserRegister{
		Fname:UD.FirstName,
		Minit:UD.Middle,
		Lname:UD.LastName,
		Dob:UD.Dob,
		Sex:UD.Sex,
		Address1:UA.AddressLine1,
		Address2:UA.AddressLine2,       
		City:UA.City,            
		State:UA.State,           
		Zipcode:UA.Zip,        
		Country:UA.Country,        
		PhoneNumber:UP.PhoneNumber,               
	}

	c.RenderArgs["ur"] = ur


	return c.Render()
}
