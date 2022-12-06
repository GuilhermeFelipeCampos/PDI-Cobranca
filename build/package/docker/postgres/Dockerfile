FROM postgres:14

ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=admin
ENV POSTGRES_DB=pdi_cobranca_db

ARG $DB_SQL=db/squema_db.sql
ARG $TABLE_SQL=db/squema_tables.sql

EXPOSE 5432:5432
COPY $DB_SQL /docker-entrypoint-initdb.d/
COPY $TABLE_SQL /docker-entrypoint-initdb.d/
