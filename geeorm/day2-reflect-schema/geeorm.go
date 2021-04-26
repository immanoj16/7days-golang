package geeorm

import (
	"database/sql"
	"geeorm/dialect"
	"geeorm/log"
	"geeorm/session"
)

// Engine referes to geeorm engine
type Engine struct {
	db      *sql.DB
	dialect dialect.Dialect
}

// NewEngine creates Engine instance
func NewEngine(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
		return
	}

	if err = db.Ping(); err != nil {
		log.Error(err)
		return
	}

	// make sure the specific dialect exist
	dial, ok := dialect.GetDialect(driver)
	if !ok {
		log.Errorf("dialect %s Not found", driver)
	}

	e = &Engine{db: db, dialect: dial}
	log.Info("Successfully connected to database")
	return
}

// Close closes the db connection
func (e *Engine) Close() {
	if err := e.db.Close(); err != nil {
		log.Error("Failed to close the database")
	}

	log.Info("Successfully closed database")
}

// NewSession create new session instance
func (e *Engine) NewSession() *session.Session {
	return session.New(e.db, e.dialect)
}
