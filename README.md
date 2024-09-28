# gin-authentication
**Standard Go Project Layout:** https://github.com/golang-standards/project-layout/

## About
* Backend: goLang
* Database: PostgreSQL + .env
* Frontend: Gin Template, jQuery, JS, Bootstrap

### Features to Implement:
1. **Login and Registration**:
    - Form validation (both server-side and client-side).
    - Password hashing (using bcrypt).
    - JWT or session-based authentication for secure login.

2. **User Profile Management**:
    - View and update user profiles.
    - Password reset functionality.

3. **Role-Based Access Control**:
    - Differentiate user roles (admin, user, etc.) and restrict access to certain routes.

4. **Secure Routes**:
    - Middleware to protect routes from unauthorized access.

5. **Database Management**:
    - Migrations for user management.
    - Using GORM or another ORM for database interactions.

6. **Frontend Integration**:
    - Responsive UI using Bootstrap.
    - AJAX calls using jQuery for a smoother user experience.

## Downloads
```bash
go get github.com/gin-gonic/gin
go get gorm.io/gorm
go get golang.org/x/crypto/bcrypt
go get github.com/joho/godotenv
go get github.com/jackc/pgx/v4

```
