Go Web Project - Overview
---



## 🔤 Naming Conventions in Go
### PascalCase und camelCase
-   **PascalCase** → Public (exported)\
    → Visible to other packages

-   **camelCase** → Private\
    → Only visible within the same package

```go
//e.g
type WelcomeRepo struct {}     // public
func getMessage() string {}    // private
```
---

# 🧠 Functions & Receivers in Go

### Example of a Function Definition
```go
func (p *WelcomeRepo) GetWelcomeMessage() string {
    return "Hello"
}
```

### Explanation of the Components
| Component             | Meaning      |
|-----------------------|--------------|
| `(p *WelcomeRepo)`    | Receiver     |
| `GetWelcomeMessage()` | Method name  |
| `string`              | Return value |


### Why Do We Need a Receiver?
-   Comparable to `this` in Java\
-   In Go, the receiver is written **explicitly before the function
    name**


### Why Use a Pointer Receiver (`*WelcomeRepo`)?
-   Prevents copying the entire struct on every method call\
-   More memory efficient\
-   Allows modification of the struct

---

# 🗄️ Database Information (PostgreSQL + Go)
### Useful Links
- https://go.dev/wiki/SQLInterface  → Create dynamic tables 
- https://www.source-fellows.com/golang-datenbankzugriffe-sql/ → Set up the database completely
- https://hub.docker.com/_/postgres
- https://www.youtube.com/watch?v=Y7a0sNKdoQk
- https://www.docker.com/products/docker-desktop/
- https://www.youtube.com/watch?v=Hs9Fh1fr5s8   → pgAdmin

---

# 🐳 Docker Commands (PostgreSQL Setup)
### Step 1: Docker starten
Open Docker Desktop.

### Step 2: Create Container
```bash
docker run --name yourcontainername 
  -e POSTGRES_PASSWORD=mysecretpassword 
  -p 5432:5432 \   # Host-Port : Container-Port
  -d postgres
```


### Step 3: Create Database
```bash
docker exec -ti yourcontainername createdb -U postgres yourdatabasename
```

### Step 4: Install pq Package
```bash
go get github.com/lib/pq
```

→ `go.mod` Check whether the package was added to go.mod


### Step 5: Implement Database in Go
```go
// e.g
connectionString := "postgres://postgres:mysecretpassword@localhost:5432/yourdatabasename?sslmode=disable"
db, err := sql.Open("postgres", connectionString)
defer db.Close()
```


### Step 6: Connect to PostgreSQL 
```bash
docker start yourcontainername
docker exec -ti yourcontainername psql -U postgres
```

### Step 7: Connect to the database 
```bash
\c yourdatabasename
```


# 🧰 Commands 
### Show container
```bash
docker ps
```

### Remove conntainer
```bash
docker rm containername
docker rm -f containername //Force delete
```

### Container management commands
```bash
docker start containername
docker stop containername
docker restart containername
```

### PostgreSQL Commands (When Connected to the Database)
| Command           | Meaning         |
|-------------------|-----------------|
| `\q`              | Quit            |
| `\dt`             | Show tables     |
| `\c databasename` | Switch Database |


### Execute SQL Query (When Connected to the Database)
```sql
SELECT * FROM users;
```
---

## ✅ Typical Workflow 
1.  Start Docker
2.  Create container
3.  Create database
4.  Install pq
5.  Integrate DB in Go
6.  Check tables with `\dt`
7.  Check data with `SELECT * FROM users;`
8.  `go build -o willweb.exe .\cmd\server`
9.  `go run ./cmd/server/`


setup temp server: https://dashboard.ngrok.com/get-started/setup/windows