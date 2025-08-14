include .env

MIGRATION_FOLDER=db/migration

migrate-create:
	goose -dir ${MIGRATION_FOLDER} create ${name} sql

migrate-up: 
	goose -dir ${MIGRATION_FOLDER} mysql ${DB_URL} up

status:
	goose -dir ${MIGRATION_FOLDER} mysql ${DB_URL} status

migrate-down:
	goose -dir ${MIGRATION_FOLDER} mysql ${DB_URL} down