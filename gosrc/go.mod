module chanteybase-srv

go 1.17

require (
	github.com/go-delve/delve v1.4.0 // indirect
	github.com/jmoiron/sqlx v1.2.0
	github.com/lib/pq v1.3.0
	github.com/mattn/go-sqlite3 v2.0.3+incompatible
	local.dev/actions v0.0.0
	local.dev/models v0.0.0
)

replace local.dev/models => ./models

replace local.dev/actions => ./actions
