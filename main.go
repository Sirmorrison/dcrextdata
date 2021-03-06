// Copyright (c) 2018-2019 The Decred developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/raedahgroup/dcrextdata/exchanges"
	"github.com/raedahgroup/dcrextdata/postgres"
	"github.com/raedahgroup/dcrextdata/pow"
	"github.com/raedahgroup/dcrextdata/version"
	"github.com/raedahgroup/dcrextdata/vsp"
	"github.com/raedahgroup/dcrextdata/web"
)

// const dcrlaunchtime int64 = 1454889600
// var opError error
// var beginShutdown = make(chan bool)

func main() {
	// Create a context that is cancelled when a shutdown request is received
	// via requestShutdown.
	ctx := withShutdownCancel(context.Background())
	// Listen for both interrupt signals and shutdown requests.
	go shutdownListener()

	if err := _main(ctx); err != nil {
		if logRotator != nil {
			log.Error(err)
		} else {
			fmt.Println(err)
		}
		os.Exit(1)
	}
	os.Exit(0)
}

func _main(ctx context.Context) error {
	cfg, _, err := loadConfig()
	if err != nil {
		return err
	}

	defer func() {
		if logRotator != nil {
			logRotator.Close()
		}
	}()

	// Display app version.
	log.Infof("%s version %v (Go version %s)", version.AppName,
		version.Version(), runtime.Version())

	db, err := postgres.NewPgDb(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPass, cfg.DBName)
	defer func(db *postgres.PgDb) {
		err := db.Close()
		if err != nil {
			log.Error("Could not close database connection: %v", err)
		}
	}(db)

	if err != nil {
		return err
	}

	if cfg.Reset {
		log.Info("Dropping tables")
		err = db.DropAllTables()
		if err != nil {
			db.Close()
			log.Error("Could not drop tables: ", err)
			return err
		}
		log.Info("Tables dropped")
	}

	if cfg.HttpMode {
		go web.StartHttpServer(cfg.HTTPHost, cfg.HTTPPort, db)
	}

	wg := new(sync.WaitGroup)

	collectData := func() error {
		if !cfg.DisableVSP {
			if exists := db.VSPInfoTableExits(); !exists {
				if err := db.CreateVSPInfoTables(); err != nil {
					log.Error("Error creating vsp info table: ", err)
					return err
				}
			}

			if exists := db.VSPTickTableExits(); !exists {
				if err := db.CreateVSPTickTables(); err != nil {
					log.Error("Error creating vsp data table: ", err)
					return err
				}

				if err := db.CreateVSPTickIndex(); err != nil {
					log.Error("Error creating vsp data index: ", err)
					return err
				}
			}

			vspCollector, err := vsp.NewVspCollector(cfg.VSPInterval, db)
			if err == nil {
				wg.Add(1)
				go vspCollector.Run(ctx, wg)
			} else {
				log.Error(err)
			}
		}
		if !cfg.DisableExchangeTicks {
			if exists := db.ExchangeTableExits(); !exists {
				if err := db.CreateExchangeTable(); err != nil {
					log.Error("Error creating exchange table: ", err)
					return err
				}
			}

			if exists := db.ExchangeTickTableExits(); !exists {
				if err := db.CreateExchangeTickTable(); err != nil {
					log.Error("Error creating exchange tick table: ", err)
					return err
				}

				if err := db.CreateExchangeTickIndex(); err != nil {
					log.Error("Error creating exchange tick index: ", err)
					return err
				}
			}

			ticksHub, err := exchanges.NewTickHub(ctx, cfg.DisabledExchanges, db)
			if err == nil {
				wg.Add(1)
				ticksHub.Run(ctx, wg)
			} else {
				log.Error(err)
			}
		}

		if !cfg.DisablePow {
			if exists := db.PowDataTableExits(); !exists {
				if err := db.CreatePowDataTable(); err != nil {
					log.Error("Error creating PoW data table: ", err)
					return err
				}
			}

			powCollector, err := pow.NewCollector(cfg.DisabledPows, cfg.PowInterval, db)
			if err == nil {
				wg.Add(1)
				powCollector.Collect(ctx, wg)
			} else {
				log.Error(err)
			}

		}

		return nil
	}

	if err = collectData(); err != nil {
		return err
	}

	ticker := time.NewTicker(300 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if err := collectData(); err != nil {
				log.Error(err)
				log.Info("Goodbye")
				os.Exit(1)
			}
		case <-ctx.Done():
			log.Infof("Shutting down collector")
			log.Info("Goodbye")
			return nil
		}
	}
	
	return nil
}
