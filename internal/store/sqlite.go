package store

import (
	"database/sql"

	_ "modernc.org/sqlite" // CGO-free driver
)

type DataProvider struct {
	db *sql.DB
}

func NewDataProvider(dbPath string) (*DataProvider, error) {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, err
	}
	return &DataProvider{db: db}, nil
}

func (p *DataProvider) FetchData() ([]map[string]interface{}, error) {
	rows, err := p.db.Query("SELECT title, project FROM TMTask WHERE trashed = 0 AND status = 0 AND start = 1 AND type = 0 AND title IS NOT NULL AND title<>'' ORDER BY todayIndex, userModificationDate;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns, _ := rows.Columns()
	var results []map[string]interface{}

	for rows.Next() {
		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))
		for i := range values {
			valuePtrs[i] = &values[i]
		}

		if err := rows.Scan(valuePtrs...); err != nil {
			return nil, err
		}

		entry := make(map[string]interface{})
		for i, col := range columns {
			val := values[i]
			entry[col] = val
		}
		results = append(results, entry)
	}

	return results, nil
}

func (p *DataProvider) Close() {
	p.db.Close()
}
