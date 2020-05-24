#!/bin/sh -e

psql --variable=ON_ERROR_STOP=1 --username "postgres" <<-EOSQL
    CREATE ROLE events WITH LOGIN PASSWORD 'events';
    CREATE DATABASE "coke_development" OWNER = events;
    GRANT ALL PRIVILEGES ON DATABASE "events-api" TO events;
EOSQL