package gateway

import (
	"database/sql"
	"errors"
	"log/slog"

	_ "github.com/mattn/go-sqlite3"

	"github.com/aetherlink/aetherlink/contract"
)

type EventListener struct {
	db             *sql.DB
	aetherContract *contract.AETHER
}

func NewEventListener(db *sql.DB, aetherContract *contract.AETHER) *EventListener {
	return &EventListener{
		db:             db,
		aetherContract: aetherContract,
	}
}

func (e *EventListener) Run() {
	sink := make(chan *contract.AETHERFileCIDRegistered)
	sub, err := e.aetherContract.WatchFileCIDRegistered(nil, sink, nil)
	if err != nil {
		slog.Error("Error watching events", "error", err)
	}

	slog.Info("Subscribed to events")

	for {
		select {
		case event := <-sink:
			if err := e.handleEvent(event); err != nil {
				slog.Error("Error handling event", "error", err)
			}

		case err := <-sub.Err():
			slog.Error("Error watching events", "error", err)
		}
	}
}

func (e *EventListener) handleEvent(event *contract.AETHERFileCIDRegistered) error {
	slog.Info("File CID registered", "name", event.FileName, "cid", event.Cid, "userAddress", event.User)

	username, err := e.aetherContract.GetUserName(nil, event.User)
	if err != nil {
		return err
	}

	res, err := e.db.Exec("INSERT INTO users (useraddress, username, cid, filename) VALUES (?, ?, ?, ?)",
		event.User,
		username,
		event.Cid,
		event.FileName,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		return errors.New("expected 1 row to be affected")
	}

	slog.Info("Inserted data", "rowsAffected", rowsAffected)

	return nil
}

func (e *EventListener) Close() {
	e.db.Close()
}
