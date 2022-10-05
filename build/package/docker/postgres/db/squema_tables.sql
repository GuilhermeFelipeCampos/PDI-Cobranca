CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
DROP TABLE IF EXISTS users CASCADE;
create table users (
  id uuid DEFAULT(public.uuid_generate_v4()),
  name text,
  email text unique not null,
  keyword text not null
);

DROP TABLE IF EXISTS clients CASCADE;

create table clients (
  id uuid DEFAULT(public.uuid_generate_v4()),
  name text not null,
  email text unique not null,
  cpf varchar(15) unique not null
);

DROP TABLE IF EXISTS charges CASCADE;

CREATE TYPE status_charge AS ENUM ('EXPIRED', 'PENDDING', 'PAY');

CREATE TABLE charges (
  id uuid DEFAULT(public.uuid_generate_v4()),
  client_id uuid not null, 
  description text not null,
  due_date date not null,
  ammount varchar(15) not null,
  status status_charge not null
);