### Setup

- Create .env file and provide the database url(DBURL) and token(TOKEN) following the .env.example pattern

- run app using 
    ```
    go run ./web
    ```
- The default port is 4000 
- use the addr flag to set a different port
example
```
go run ./cmd/web -addr=":9000" 

for port 9000
```