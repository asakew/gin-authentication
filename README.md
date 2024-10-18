# gin-authentication
Standard Go Project Layout: https://github.com/golang-standards/project-layout/

## Google OAuth2 Authentication in Golang v1:
- Backend: goLang + gin + gorilla/mux + JWT token
- Database: PostgreSQL + GORM + .env file


________________________________________________________________
## Downloading
```bash
go get github.com/gin-gonic/gin
go get github.com/gorilla/mux
go get github.com/dgrijalva/jwt-go
go get gorm.io/gorm
go get gorm.io/driver/postgres
go get github.com/joho/godotenv
```

### Update all modules:
```bash
go get -u
go mod tidy
```

## Configuration: .env File
Create a .env file:
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=yourusername
DB_PASSWORD=yourpassword
DB_NAME=yourdbname

JWT_SECRET=yourjwtsecretkey

GOOGLE_CLIENT_ID=yourgoogleclientid
GOOGLE_CLIENT_SECRET=yourgoogleclientsecret
```
