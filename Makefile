#!make
ifneq (,$(wildcard ./.env))
    include ./.env
    export $(shell sed 's/=.*//' ./.env)
endif

GO111MODULE=on

docker-up:
	docker-compose -f docker-compose.yml up -d

schema-up:
	psql "host=localhost port=${POSTGRES_PORT_EXPOSE} dbname=${POSTGRES_DB} user=${POSTGRES_USER} password=${POSTGRES_PASSWORD}" -a -f schema.sql