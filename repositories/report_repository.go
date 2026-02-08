package repositories

import (
	"database/sql"
	"crudapi/models"
)

type ReportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

func (r *ReportRepository) GetTodayReport() (*models.DailyReport, error) {
	report := &models.DailyReport{}

	// total revenue
	err := r.db.QueryRow(`
		SELECT COALESCE(SUM(total_amount), 0)
		FROM transactions
		WHERE created_at::date = CURRENT_DATE
	`).Scan(&report.TotalRevenue)
	if err != nil {
		return nil, err
	}

	// total transaksi
	err = r.db.QueryRow(`
		SELECT COUNT(*)
		FROM transactions
		WHERE created_at::date = CURRENT_DATE
	`).Scan(&report.TotalTransaksi)
	if err != nil {
		return nil, err
	}

	// produk terlaris
	err = r.db.QueryRow(`
		SELECT product_name, SUM(quantity) AS total_qty
		FROM transaction_details td
		JOIN transactions t ON t.id = td.transaction_id
		WHERE t.created_at::date = CURRENT_DATE
		GROUP BY product_name
		ORDER BY total_qty DESC
		LIMIT 1
	`).Scan(
		&report.ProdukTerlaris.Nama,
		&report.ProdukTerlaris.QtyTerjual,
	)

	// kalau belum ada transaksi hari ini
	if err == sql.ErrNoRows {
		report.ProdukTerlaris = models.BestSeller{}
		return report, nil
	}

	if err != nil {
		return nil, err
	}

	return report, nil
}

