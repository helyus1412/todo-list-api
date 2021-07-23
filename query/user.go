package query

const (
	QueryAddUser = `
		INSERT INTO users (name, email, password) values (?,?,?)
	`

	QueryGetUserDetail = `
		SELECT
			u.name,
			u.email,
			u.password,
		FROM users u
		WHERE u.id = ?
	`
)
