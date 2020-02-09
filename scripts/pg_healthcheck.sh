#!/bin/sh

if select=$(echo 'SELECT 1' | psql --username ${POSTGRES_USER} --dbname ${POSTGRES_DB} --quiet --no-align --tuples-only ) && [ ${select} = '1' ];
then
  exit 0;
fi;

exit 1