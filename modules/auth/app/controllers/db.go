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

func (c *AuthController) Begin() r.Result {
	txn, err := AuthDbMap.Begin()
	if err != nil {
		panic(err)
	}
	c.Txn = txn
	return nil
}

func (c *AuthController) Commit() r.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Commit(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}

func (c *AuthController) Rollback() r.Result {
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
	AuthDbMap = &gorp.DbMap{Db: db.Db, Dialect: gorp.PostgresDialect{}}

	// ub := AuthDbMap.AddTable(models.UserBase{}).SetKeys(true, "UserId")
	// ub.ColMap("Email").Unique = true
	// setColumnSizes(ub, map[string]int{
	// 	"UserName": 32,
	// 	"Email":    64,
	// })

	// AuthDbMap.AddTable(models.UserVolatile{}).SetKeys(true, "UserId")
	AuthDbMap.AddTable(models.UserAuth{}).SetKeys(true, "UserId")

	AuthDbMap.TraceOn("[gorp]", log.New(GLogger{glog.Info}, "", 0))
	AuthDbMap.CreateTablesIfNotExists()

	// TODO:  change fill to some check against the db
	// TODO:  add a test fill option for development
	if fill {
		now := time.Now().UnixNano()

		demoUser := &models.UserBase{0, "demo", "demo@domain.com"}
		errU := AuthDbMap.Insert(demoUser)
		checkFail(errU)

		demoVolatile := &models.UserVolatile{demoUser.UserId, now, 0, 0, 0, now}
		errV := AuthDbMap.Insert(demoVolatile)
		checkFail(errV)

		demoPassword, _ := bcrypt.GenerateFromPassword([]byte("demo"), bcrypt.DefaultCost)
		demoAuth := &models.UserAuth{demoUser.UserId, demoPassword, "", 0, 0, 0, 0}
		errA := AuthDbMap.Insert(demoAuth)
		checkFail(errA)
	}

}
