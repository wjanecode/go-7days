package geeorm

import (
	"database/sql"
	"geeorm/log"
	"geeorm/session"
)

// Engine main struct ,manage all db session
type Engine struct {
	db *sql.DB
}

func NewEngine(driver, source string) (e *Engine, err error) {

	db, err := sql.Open(driver, source)
	if err != nil {
		geeLog.Error(err)
		return
	}

	if err = db.Ping(); err != nil {
		geeLog.Error(err)
		return
	}
	e = &Engine{db: db}
	geeLog.Info("connect database success")
	return
}

func (engine *Engine) close() {
	if err := engine.db.Close(); err != nil {
		geeLog.Error("Failed to close database")
	}
	geeLog.Info("close database success")
}

func (engine *Engine) NewSession() *session.Session {
	return session.New(engine.db)
}
