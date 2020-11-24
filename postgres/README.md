Postgres Deployment
===================


### Creating a new database
```
kubectl exec -it postgres-XXXX-XXXX -- bash
postgres$ psql -h localhost -p 5432 -U postgres-user postgres-db
postgres=# create database "hello-db"; 
create user "hello-user" with encrypted password 'password'; 
grant all privileges on database "hello-db" to "hello-user";
```