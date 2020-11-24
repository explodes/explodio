Postgres Deployment
===================


### Creating a new database
```
kubectl exec -it postgres-XXXX-XXXX -- bash
root# su postgres
postgres$ psql -h localhost -p 5432 -U postgres-user postgres-db
postgres=# create database "my-service-db";
postgres=# create user "my-service-user" with encrypted password 'password';
postgres=# grant all privileges on database "my-service-db" to "my-service-user";
```