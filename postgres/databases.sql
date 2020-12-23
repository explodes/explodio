CREATE DATABASE "hello-db";
CREATE USER "hello-user" WITH ENCRYPTED PASSWORD 'password';
GRANT ALL PRIVILEGES ON DATABASE "hello-db" TO "hello-user";

CREATE DATABASE "tokyo-db";
CREATE USER "tokyo-user" WITH ENCRYPTED PASSWORD 'password';
GRANT ALL PRIVILEGES ON DATABASE "tokyo-db" TO "tokyo-user";