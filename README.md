# 📘 Go Web Projekt – Übersicht
---

## 🔤 Naming Conventions in Go
### PascalCase und camelCase

- **PascalCase** → Public (exported)  
  → Von anderen Packages sichtbar
- **camelCase** → Private  
  → Nur innerhalb des Packages sichtbar

### Beispiel

```go
type WelcomeRepo struct {}     // public
func getMessage() string {}    // private
```
---

# 🧠 Funktionen & Receiver in Go

## Beispiel einer Funktionsdefinition

```go
func (p *WelcomeRepo) GetWelcomeMessage() string {
    return "Hello"
}
```

## Erklärung der Bestandteile

| Bestandteil | Bedeutung |
|------------|------------|
| `(p *WelcomeRepo)` | Receiver |
| `GetWelcomeMessage()` | Methodenname |
| `string` | Rückgabewert |


## Warum brauchen wir einen Receiver?

- Ermöglicht Methoden auf Structs
- Vergleichbar mit `this` in Java
- In Go wird der Receiver **explizit davor geschrieben**


## Warum Pointer-Receiver (`*WelcomeRepo`)?

- Verhindert Kopie des gesamten Structs bei jedem Methodenaufruf
- Effizienter Speicherverbrauch
- Ermöglicht Änderungen am Struct

---

# 🗄️ Database Informationen (PostgreSQL + Go)

## Nützliche Links

- https://go.dev/wiki/SQLInterface  → Dynamische Tables erstellen
- https://www.source-fellows.com/golang-datenbankzugriffe-sql/ → DB komplett aufsetzen
- https://hub.docker.com/_/postgres
- https://www.youtube.com/watch?v=Y7a0sNKdoQk
- https://www.docker.com/products/docker-desktop/
- https://www.youtube.com/watch?v=Hs9Fh1fr5s8   → pgAdmin

---

# 🐳 Docker Commands (PostgreSQL Setup)

### Step 1: Docker starten

Docker Desktop öffnen.


### Step 2: Container erstellen

```bash
docker run --name yourcontainername -e POSTGRES_PASSWORD=mysecretpassword -p 5432:5432 -d postgres
```

`-p 5432:5432` → Container-Port : Dein lokaler Port


### Step 3: Datenbank erstellen

```bash
docker exec -ti yourcontainername createdb -U postgres yourdatabasename
```


### Step 4: pq Package installieren

```bash
go get github.com/lib/pq
```

→ `go.mod` überprüfen (Paket wurde hinzugefügt)


### Step 5: Datenbank in Go implementieren

Connection String Beispiel:

```go
connectionString := "postgres://postgres:mysecretpassword@localhost:5432/yourdatabasename?sslmode=disable"
db, err := sql.Open("postgres", connectionString)
defer db.Close()
```


### Step 6: Mit PostgreSQL verbinden

```bash
docker start yourcontainername
docker exec -ti yourcontainername psql -U postgres
```


### Step 7: Mit Datenbank verbinden

```bash
\c yourdatabasename
```


# 🧰 Befehle, die wir immer brauchen

## Container anzeigen

```bash
docker ps
```


## Container löschen

```bash
docker rm containername
```

Force löschen:

```bash
docker rm -f containername
```

## Weitere Befehle

```bash
docker start containername
docker stop containername
docker restart containername
```


## PostgreSQL Befehle, wenn wir mit der Datenbank verbunden sind

| Befehl                | Bedeutung |
|-----------------------|-----------|
| `\q`                  | Beenden |
| `\dt`                 | Tabellen anzeigen |
| `\c databasename`     | Datenbank wechseln |


## SQL-Abfrage ausführen, wenn wir mit der Datenbank verbunden sind

```sql
SELECT * FROM users;
```
---

# ✅ Typischer Workflow

1. Docker starten
2. Container erstellen
3. Datenbank erstellen
4. pq installieren
5. DB in Go einbinden
6. Tabellen prüfen mit `\dt`
7. Daten prüfen mit `SELECT * FROM users;`
8. In Terminal: go build -o willweb.exe .\cmd\server
9. go run ./cmd/server/