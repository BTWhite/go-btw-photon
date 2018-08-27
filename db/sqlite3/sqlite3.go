// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3

package sqlite3

import (
	"database/sql"

	"github.com/mattn/go-sqlite3"
)

type DB struct {
	*sql.DB
}

func Init(filename string) *DB {

	sql.Register("sqlite3driver",
		&sqlite3.SQLiteDriver{
			ConnectHook: func(conn *sqlite3.SQLiteConn) error {
				return nil
			},
		},
	)

	db, err := sql.Open("sqlite3driver", filename)

	if err != nil {
		panic(err.Error())
	}

	return &DB{db}
}
