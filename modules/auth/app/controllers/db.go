package controllers

import (
	"database/sql"
	"log"
	"time"

	"code.google.com/p/go.crypto/bcrypt"
	"github.com/coopernurse/gorp"
	"github.com/lib/pq"
	"github.com/robfig/revel"
	db "github.com/robfig/revel/modules/db/app"

	auth "github.com/iassic/revel-modz/modules/auth/app/models"
	user "github.com/iassic/revel-modz/modules/user/app/models"
)

var (
	AuthDbMap *gorp.DbMap

	fill = false
)

type AuthDbController struct {
	*revel.Controller
	Txn *gorp.Transaction
}

func (c *AuthDbController) Begin() r.Result {
	txn, err := Dbm.Begin()
	if err != nil {
		panic(err)
	}
	c.Txn = txn
	return nil
}

func (c *AuthDbController) Commit() r.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Commit(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}

func (c *AuthDbController) Rollback() r.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Rollback(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}

func setColumnSizes(t *gorp.TableMap, colSizes map[string]int) {
	for col, size := range colSizes {
		t.ColMap(col).MaxSize = size
	}
}

func InitAuthDb() {
	db.Init()
	Dbm = &gorp.DbMap{Db: db.Db, Dialect: gorp.PostgresDialect{}}

	ub := Dbm.AddTable(models.UserBase{}).SetKeys(true, "UserId")
	ub.ColMap("Email").Unique = true
	setColumnSizes(ub, map[string]int{
		"UserName": 32,
		"Email":    64,
	})

	Dbm.AddTable(models.UserVolatile{}).SetKeys(true, "UserId")
	Dbm.AddTable(models.UserAuth{}).SetKeys(true, "UserId")

	Dbm.TraceOn("[gorp]", log.New(GLogger{glog.Info}, "", 0))
	Dbm.CreateTablesIfNotExists()

	// TODO:  change fill to some check against the db
	// TODO:  add a test fill option for development
	if fill {
		now := time.Now().UnixNano()

		demoUser := &models.UserBase{0, "demo", "demo@domain.com"}
		errU := Dbm.Insert(demoUser)
		checkFail(errU)

		demoVolatile := &models.UserVolatile{demoUser.UserId, now, 0, 0, 0, now}
		errV := Dbm.Insert(demoVolatile)
		checkFail(errV)

		demoPassword, _ := bcrypt.GenerateFromPassword([]byte("demo"), bcrypt.DefaultCost)
		demoAuth := &models.UserAuth{demoUser.UserId, demoPassword, "", 0, 0, 0, 0}
		errA := Dbm.Insert(demoAuth)
		checkFail(errA)
	}

}
