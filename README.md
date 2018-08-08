# go-gin-orm

CRUD API with Golang, Gin Framework and PostgreSQL ORM

### **Usage**

URL : localhost:3000

Get the User in JSON format.

```http
GET /readUsers
```

Create User.

```http
GET /createUser
```
```json
{
    "user_id": [auto],
    "user_name": "[your username]",
    "email": "[your email]",
    "password": "[your password]"
}
```

Update User

```http
GET /updateUser
```
```json
{
    "user_name": "[update your username]",
    "email": "[parameter email]",
    "password": "[update your password]"
}
```

Delete User.

```http
GET /deleteUser
```
```json
{
    "email": "[parameter email]",
}
```

Response format.

```json
{
    "users": [
        {
            "user_id": 1,
            "user_name": "yudhiz99",
            "email": "yudhis_tiro@yahoo.com",
            "password": "6763389141h4lh3418y48194h34y189"
        },
        {
            "user_id": 2,
            "user_name": "bangke",
            "email": "mahasiswateknologi@gmail.com",
            "password": "ncisqnckslqncskqi7896767969341nkln"
        }
    ]
}
```

### **Run**

```bash
go run mainorm.go