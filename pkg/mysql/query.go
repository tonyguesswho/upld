package mysql

import (
	"database/sql"
)

// UploadModel : type which wraps a sql.DB connection pool.
type UploadModel struct {
	DB *sql.DB
}

// Insert a new upload into the database.
func (m *UploadModel) Insert(filesize int64, filename, filetype, initialfilename string) (int, error) {

	stmt := `INSERT INTO uploads (filesize, filename, filetype, initialfilename, created)
	VALUES(?, ?, ?, ?, UTC_TIMESTAMP())`
	result, err := m.DB.Exec(stmt, filesize, filename, filetype, initialfilename)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
