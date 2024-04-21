## System Requirment

Before you begin working on this project, ensure that your system meets the following requirements:
- [Go](https://go.dev/doc/install) is installed on your machine.
- [PostgreSql](https://www.postgresql.org/download/) is installed on your machine.

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/tafhdytllah/customer-list.git
    ```

2. Navigate to the project directory:

    ```bash
    cd customer-list
    ```
3. Init project
   ```bash
   go mod init github.com/tafhdytllah/customer-list
   ```
4. Install Library :
   - Gorilla mux
     ```bash
     go get -u github.com/gorilla/mux
     ```
   - Gorm
     ```bash
     go get -u gorm.io/gorm
     ```
   - Driver PostgreSQL
     ```bash
     go get -u gorm.io/driver/postgres
     ```
   - Viper
     ```bash
     go get github.com/spf13/viper
     ```
5. Create database schema into your machine
   - [database.sql](/database.sql)

6. Create file .env on root directory

    ```bash
    touch .env
    ```
7. Setup file .env
    ```bash
    PORT=8080
    ADDRESS=localhost
    DB_HOST=localhost
    DB_USER={your_db_user}
    DB_PASSWORD={your_db_password}
    DB_NAME={your_db_name}
    DB_PORT={your_db_port}
    ```
8. Run go project in local machine :

    ```bash
    go run main.go
    ```
9.  Application will be accessible at [http://localhost:8080](http://localhost:8080).
   

### Api Documentation
- [Customer List](/docs/api_spec.md)

### Contact

If you have any questions or comments about this project, please feel free to contact me at [gmail](mailto:taufikhh.97@gmail.com).