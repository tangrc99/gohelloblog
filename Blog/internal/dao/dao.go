package dao

import (
	"github.com/tangrc99/gohelloblog/global"
)

// StartTransaction Package database operations and other operations into a single transaction.
func StartTransaction(f func() error) error {
	tx := global.MySQL.Begin()
	err := f()
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
