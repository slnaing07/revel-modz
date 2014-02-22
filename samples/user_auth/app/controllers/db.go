package controllers

import (
	// "database/sql"

	"code.google.com/p/go.crypto/bcrypt"
	"github.com/robfig/revel"

	gorm "github.com/jinzhu/gorm"
	// "github.com/coopernurse/gorp"
	_ "github.com/lib/pq"
	// _ "github.com/mattn/go-sqlite3"

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
	// txn, err := TestDB.Begin()
	if err != nil {
		panic(err)
	}
	c.Txn = txn
	return nil
}

func (c *DbController) Commit() revel.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Commit().Error; err != nil {
		panic(err)
	}
	c.Txn = nil
	return nil
}

func (c *DbController) Rollback() revel.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Rollback().Error; err != nil {
		panic(err)
	}
	c.Txn = nil
	return nil
}

func InitTestDB() {

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
	if err != nil {
		revel.ERROR.Fatal(err)
	}

	ndb.SetLogger(gorm.Logger{revel.INFO})
	TestDB = &ndb

	revel.INFO.Println("Connection made to DB")

	initTestDB_Tables()
	fillTestDB_AuthTable()
	testTestDB_AuthTable()
	return

}

func initTestDB_Tables() {

	TestDB.AutoMigrate(auth.UserAuth{})
	TestDB.AutoMigrate(UserBasic{})

}

func fillTestDB_AuthTable() {
	bp, _ := bcrypt.GenerateFromPassword([]byte("demo"), bcrypt.DefaultCost)
	ua := &auth.UserAuth{UserId: 100001, HashedPassword: bp}
	ub := &UserBasic{UserId: 100001, Email: "demo@domain.com"}

	// Start a new transaction
	// trans, err := TestDB.Begin()
	trans := TestDB.Begin()
	err := trans.Error
	if err != nil {
		revel.ERROR.Println(err)
	}

	// err = trans.Insert(ua)
	err = trans.Save(ua).Error
	if err != nil {
		revel.ERROR.Println(err)
	}

	err = trans.Save(ub).Error
	if err != nil {
		revel.ERROR.Println(err)
	}

	// if the commit is successful, a nil error is returned
	err = trans.Commit().Error
	if err != nil {
		revel.ERROR.Println(err)
	}

	if err != nil {
		revel.ERROR.Println("Unable to add demo to AuthTable")
	} else {
		revel.INFO.Println("ADDED demo to AuthTable")
	}
}

type UserTest struct {
	UserId   int64
	Password string
}

func (u UserTest) AuthId() int64 {
	return u.UserId
}

func (u UserTest) AuthSecret() string {
	return u.Password
	// bp, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	// return bp
}

func testTestDB_AuthTable() {
	user := UserTest{100001, "demo"}
	revel.WARN.Println(user)

	// Start a new transaction
	// trans, err := TestDB.Begin()
	trans := TestDB.Begin()
	err := trans.Error
	if err != nil {
		revel.ERROR.Println(err)
	}

	user_auth := auth.CheckUserAuth(trans, user)
	revel.WARN.Println(user_auth)

	// if the commit is successful, a nil error is returned
	err = trans.Commit().Error
	if err != nil {
		revel.ERROR.Println(err)
	}

}
