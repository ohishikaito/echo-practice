# e.g. make m-file name=create_users
gen-migrate:
	migrate create -ext sql -dir ./ddl/db -seq $(name)

# e.g. make migrate
migrate:
	migrate -source file://./ddl/db -database 'mysql://root:password@tcp(localhost:13306)/development' up

# NOTE: n回分migrationが戻るので、実行時は要注意
# e.g. make migrate-drop n=7
migrate-drop:
	migrate -source file://./ddl/db -database 'mysql://root:password@tcp(localhost:13306)/development' down $(n)
