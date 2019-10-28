package mysql

import (
	"database/sql"
	"errors"

	models "github.com/tonyguesswho/upld/pkg"
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

// Get a specific upload based on its id.
func (m *UploadModel) Get(id int) (*models.Upload, error) {
	stmt := `SELECT id, filesize, filename, filetype, initialfilename, created FROM uploads    
	WHERE id = ?`

	row := m.DB.QueryRow(stmt, id)

	u := &models.Upload{}
	err := row.Scan(&u.ID, &u.Filesize, &u.Filename, &u.Filetype, &u.Initialfilename, &u.Created)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		}
		return nil, err
	}
	return u, nil
}
