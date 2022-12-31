package repository

const (
	createTransaction = `INSERT INTO transaction (amount, datetime) VALUES ($1, $2) RETURNING *`
)
