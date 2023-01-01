package repository

const (
	// create a new transaction
	createTransaction = `INSERT INTO transactions (amount, datetime) VALUES ($1, $2) RETURNING *`

	// get transactions grouped by hour
	getTransactionsByHour = `
		SELECT DISTINCT ON (datetime) SUM(amount) AS amount, date_trunc('hour', datetime) as datetime
        FROM transactions
        WHERE datetime < $1
        GROUP BY date_trunc('hour', datetime)`

	getBalance = `SELECT SUM(amount) AS total_amount FROM transactions`
)
