// Package repository
package repository

import (
	"database/sql"

	"github.com/Viggie1/gourl/internal/models"
)

type URLRepository struct {
	db *sql.DB
}

func NewURLRepository(db *sql.DB) *URLRepository {
	return &URLRepository{db: db}
}

func (r *URLRepository) Create(url *models.ShortenedURL) error {
	query := `INSERT INTO shortened_urls (url, shortcode, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id`

	return r.db.QueryRow(query, url.URL, url.ShortCode, url.CreatedAt, url.UpdatedAt).Scan(&url.ID)
}

func (r *URLRepository) CheckURL(searchURL string) (*models.ShortenedURL, error) {
	query := `SELECT * FROM shortened_urls WHERE url=$1`

	var shortURL models.ShortenedURL

	err := r.db.QueryRow(query, searchURL).Scan(&shortURL.ID, &shortURL.URL, &shortURL.ShortCode, &shortURL.CreatedAt, &shortURL.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &shortURL, nil
}

func (r *URLRepository) Retrieve(shortenedURL string) (*models.ShortenedURL, error) {
	query := `SELECT * FROM shortened_urls WHERE shortcode = $1`

	var returnURL models.ShortenedURL

	err := r.db.QueryRow(query, shortenedURL).Scan(&returnURL.ID, &returnURL.URL, &returnURL.ShortCode, &returnURL.CreatedAt, &returnURL.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &returnURL, nil
}

func (r *URLRepository) Delete(shortcode string) (bool, error) {
	query := `DELETE FROM shortened_urls WHERE shortcode = $1`

	rowRes, err := r.db.Exec(query, shortcode)
	if err != nil {
		return false, err
	}

	rowsAffected, _ := rowRes.RowsAffected()

	if rowsAffected == 0 {
		return false, sql.ErrNoRows
	}

	return true, nil
}
