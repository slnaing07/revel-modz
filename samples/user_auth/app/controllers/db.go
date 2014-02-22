package controllers

import (
	gorm "github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/robfig/revel"

	"github.com/iassic/revel-modz/modules/auth"
	"github.com/iassic/revel-modz/modules/user"
)

var (
	TestDB *gorm.DB

	fill = false
)

type DbController struct {
	*revel.Controller
	Txn *gorm.DB
}

func (c *DbController) Begin() revel.Result {
	txn := TestDB.Begin()
	err := txn.Error
	checkPANIC(err)
	c.Txn = txn
	return nil
}

func (c *DbController) Commit() revel.Result {
	if c.Txn == nil {
		return nil
	}
	err := c.Txn.Commit().Error
	checkPANIC(err)

	c.Txn = nil
	return nil
}

func (c *DbController) Rollback() revel.Result {
	if c.Txn == nil {
		return nil
	}
	err := c.Txn.Rollback().Error
	checkPANIC(err)

	c.Txn = nil
	return nil
}

func InitDB() {

	var driver, spec string
	var found bool
	if driver, found = revel.Config.String("db.driver"); !found {
		revel.ERROR.Fatal("No db.driver found.")
	}
	if spec, found = revel.Config.String("db.spec"); !found {
		revel.ERROR.Fatal("No db.spec found.")
	}

	// Open a connection.
	ndb, err := gorm.Open(driver, spec)
	checkPANIC(err)

	ndb.SetLogger(gorm.Logger{revel.INFO})
	TestDB = &ndb

	revel.INFO.Println("Connection made to DB")
}

func SetupTables() {
	addTables()
}

func SetupDevDB() {

	dropTables()
	addTables()
	fillTables()
	testDevDB()
	return

}

func dropTables() {

	TestDB.DropTable(auth.UserAuth{})
	TestDB.DropTable(user.UserBasic{})

}

func addTables() {

	TestDB.AutoMigrate(auth.UserAuth{})
	TestDB.AutoMigrate(user.UserBasic{})

}

func fillTables() {

	// Start a new transaction
	// trans, err := TestDB.Begin()
	trans := TestDB.Begin()
	err := trans.Error
	checkERROR(err)

	for _, up := range dev_users {

		ub := &user.UserBasic{
			UserId:   up.UserId,
			UserName: up.UserName,
		}
		err = user.AddUserBasic(trans, ub)
		checkERROR(err)

		_, err = auth.AddUserAuth(trans, up)
		checkERROR(err)
	}

	// if the commit is successful, a nil error is returned
	err = trans.Commit().Error
	checkERROR(err)

	if err != nil {
		revel.ERROR.Println("Unable to fill DB")
	} else {
		revel.INFO.Println("Filled DB tables")
	}
}

var dev_users = []*user.UserPass{
	&user.UserPass{UserId: 100001, UserName: "demo1@domain.com", Password: "demopass"},
	&user.UserPass{UserId: 100002, UserName: "demo2@domain.com", Password: "demopass"},
	&user.UserPass{UserId: 100003, UserName: "demo3@domain.com", Password: "demopass"},
	&user.UserPass{UserId: 100004, UserName: "demo4@domain.com", Password: "demopass"},
	&user.UserPass{UserId: 100005, UserName: "demo5@domain.com", Password: "demopass"},
	&user.UserPass{UserId: 100006, UserName: "demo6@domain.com", Password: "demopass"},
	&user.UserPass{UserId: 100007, UserName: "demo7@domain.com", Password: "demopass"},
	&user.UserPass{UserId: 100008, UserName: "demo8@domain.com", Password: "demopass"},
}

func testDevDB() {
	for _, up := range dev_users {
		u := user.GetUserBasicByUserId(TestDB, up.UserId)
		if u == nil {
			revel.ERROR.Println("Failed to look up user by id:", up.UserId)
		}
		u = user.GetUserBasicByName(TestDB, up.UserName)
		if u == nil {
			revel.ERROR.Println("Failed to look up user by name:", up.UserName)
		}

		a, err := auth.Authenticate(TestDB, up)
		checkERROR(err)
		if a == nil {
			revel.ERROR.Printf("Failed to authenticate user: %+v\n", *up)
		}
	}
}
