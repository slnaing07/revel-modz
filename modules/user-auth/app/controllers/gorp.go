package controllers

import (
	"database/sql"
	"log"
	"time"

	"code.google.com/p/go.crypto/bcrypt"
	"github.com/coopernurse/gorp"
	"github.com/golang/glog"
	_ "github.com/lib/pq"
	r "github.com/robfig/revel"
	db "github.com/robfig/revel/modules/db/app"

	"github.com/iassic/revel-modz/modules/user-auth/app/models"
)

var (
	Dbm *gorp.DbMap

	fill = false
)

func setColumnSizes(t *gorp.TableMap, colSizes map[string]int) {
	for col, size := range colSizes {
		t.ColMap(col).MaxSize = size
	}
}

func GorpInit() {
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

// GLogger is an adapter from log.Logger to glog
type GLogger struct {
	out func(args ...interface{})
}

func (g GLogger) Write(p []byte) (n int, err error) {
	g.out(string(p))
	glog.Flush()
	return len(p), nil
}

type GorpController struct {
	*r.Controller
	Txn *gorp.Transaction
}

func (c *GorpController) Begin() r.Result {
	txn, err := Dbm.Begin()
	if err != nil {
		panic(err)
	}
	c.Txn = txn
	return nil
}

func (c *GorpController) Commit() r.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Commit(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}

func (c *GorpController) Rollback() r.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Rollback(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}
