package refresh_repositories

import (
	"context"
	"database/sql"
	"fmt"
	elClient "github.com/olivere/elastic/v7"
)

func TruncateRdbm(db *sql.DB, tables []string) (err error) {
	for _, table := range tables {
		_, err = db.Exec(fmt.Sprintf("TRUNCATE TABLE %s", table))
		if err != nil {
			return
		}
	}
	return
}

func TruncateNoSql(el *elClient.Client, indices []string) (err error) {
	ctx := context.Background()
	for _, index := range indices {
		_, err = el.DeleteIndex(index).Do(ctx)
		if err != nil {
			return
		}
	}
	return
}

func TruncateRepositories(db *sql.DB, el *elClient.Client, tables, indices []string) (err error) {
	if err = TruncateRdbm(db, tables); err != nil {
		return
	}
	if err = TruncateNoSql(el, indices); err != nil {
		return
	}
	return
}
