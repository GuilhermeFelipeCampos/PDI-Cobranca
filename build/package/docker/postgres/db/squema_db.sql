CREATE SCHEMA pdi_cobranca;

GRANT ALL PRIVILEGES ON DATABASE "pdi_cobranca_db" TO "postgres";

GRANT USAGE ON SCHEMA pdi_cobranca TO "postgres";
ALTER USER "postgres" SET search_path = 'pdi_cobranca';


SET SCHEMA 'pdi_cobranca';
ALTER DEFAULT PRIVILEGES
    IN SCHEMA pdi_cobranca
GRANT SELECT, UPDATE, INSERT, DELETE ON TABLES
    TO "postgres";

ALTER DEFAULT PRIVILEGES
    IN SCHEMA pdi_cobranca
GRANT USAGE ON SEQUENCES
    TO "postgres";

