module chanteybase-srv

go 1.17

require (
	github.com/jmoiron/sqlx v1.3.4
	github.com/lib/pq v1.10.2
	github.com/mattn/go-sqlite3 v2.0.3+incompatible
	local.dev/actions v0.0.0
	local.dev/models v0.0.0
)

replace local.dev/models => ./models

replace local.dev/actions => ./actions
