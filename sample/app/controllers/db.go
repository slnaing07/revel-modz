package controllers

import (
	gorm "github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/revel/revel"

	"github.com/iassic/revel-modz/modules/auth"
	"github.com/iassic/revel-modz/modules/maillist"
	"github.com/iassic/revel-modz/modules/user"
	"github.com/iassic/revel-modz/modules/user-files"
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
	revel.INFO.Println("Setting up Prod DB")
	addTables()
}

func SetupDevDB() {
	revel.INFO.Println("Setting up Dev DB")
	dropTables()
	addTables()
	fillTables()
	testDevDB()
}

func dropTables() {
	revel.INFO.Println("Dropping tables")
	auth.DropTables(TestDB)
	user.DropTables(TestDB)
	maillist.DropTables(TestDB)
	userfiles.DropTables(TestDB)
}

func addTables() {
	revel.INFO.Println("AutoMigrate tables")
	auth.AddTables(TestDB)
	user.AddTables(TestDB)
	maillist.AddTables(TestDB)
	userfiles.AddTables(TestDB)
}

func fillTables() {

	var err error

	for _, up := range dev_users {

		ub := &user.UserBasic{
			UserId:   up.UserId,
			UserName: up.UserName,
		}
		err = user.AddUserBasic(TestDB, ub)
		checkERROR(err)

		created_at := ub.CreatedAt
		updated_at := ub.UpdatedAt

		if created_at.IsZero() {
			revel.ERROR.Println("Should have created_at after create")
		}
		if updated_at.IsZero() {
			revel.ERROR.Println("Should have updated_at after create")
		}

		_, err = auth.AddUserAuth(TestDB, up)
		checkERROR(err)
	}

	revel.INFO.Println("Filled DBs")
}

var dev_users = []*user.UserPass{
	&user.UserPass{UserId: 100001, UserName: "demo1@domain.com", Password: "demopass"},
	&user.UserPass{UserId: 100002, UserName: "demo2@domain.com", Password: "demopass"},
	&user.UserPass{UserId: 100003, UserName: "demo3@domain.com", Password: "demopass"},
	&user.UserPass{UserId: 100004, UserName: "demo4@domain.com", Password: "demopass"},
	&user.UserPass{UserId: 200001, UserName: "admin@domain.com", Password: "adminpass"},
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
