# Recruitment Management System

This is a backend system to manage recruitment using golang.

To create a db in postgres
```
createdb -U <user_name> <db_name>
```

To run the app:

create the .env file
``` 
PORT=<portnumber> //eg: 4000
DB_URL=<db_url> //eg: postgres://<username>:<password>@<host>:<port>/<db_name>
```

run the server

```
make run
```
