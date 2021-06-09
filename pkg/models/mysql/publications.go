package mysql

import (
	"database/sql"
	"errors"

	"github.com/Linus4/menoci_go/pkg/models"
)

type PublicationModel struct {
	DB *sql.DB
}

func (m *PublicationModel) Insert(title, content, expires string) (int, error) {
	stmt := `INSERT INTO publications (title, content, created, expires)
	VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *PublicationModel) Get(id int) (*models.Publication, error) {
	stmt := `SELECT id, title, content, created, expires FROM publications
	WHERE expires > UTC_TIMESTAMP() AND id = ?`

	p := &models.Publication{}
	err := m.DB.QueryRow(stmt, id).Scan(&p.ID, &p.Title, &p.Content, &p.Created, &p.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return p, nil
}

func (m *PublicationModel) Latest() ([]*models.Publication, error) {
	return nil, nil
}
