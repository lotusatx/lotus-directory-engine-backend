# API Endpoints Documentation

Complete reference for all API endpoints in the Lotus Directory Engine.

## Base URL
```
http://localhost:8080/api/v1
```

## Table of Contents
- [Users](#users)
- [Groups](#groups)
- [Roles](#roles)
- [Health Check](#health-check)

---

## Users

### Create User
```http
POST /users
Content-Type: application/json

{
  "id": "UI000001",
  "email": "john.doe@company.com",
  "name": "John Doe",
  "roles": [],
  "group_ids": []
}
```

**Response:** `201 Created`
```json
{
  "id": "UI000001",
  "email": "john.doe@company.com",
  "name": "John Doe",
  "roles": [],
  "group_ids": []
}
```

### Get All Users
```http
GET /users
```

**Response:** `200 OK`
```json
[
  {
    "id": "UI000001",
    "email": "john.doe@company.com",
    "name": "John Doe",
    "roles": [],
    "group_ids": []
  }
]
```

### Get User by ID
```http
GET /users/{id}
```

**Example:**
```http
GET /users/UI000001
```

**Response:** `200 OK`
```json
{
  "id": "UI000001",
  "email": "john.doe@company.com",
  "name": "John Doe",
  "roles": [],
  "group_ids": []
}
```

### Update User
```http
PUT /users/{id}
Content-Type: application/json

{
  "email": "john.updated@company.com",
  "name": "John Updated",
  "roles": [],
  "group_ids": []
}
```

**Response:** `200 OK`

### Delete User
```http
DELETE /users/{id}
```

**Response:** `204 No Content`

---

## Groups

### Create Group
```http
POST /groups
Content-Type: application/json

{
  "id": "GRP001",
  "name": "Engineering Team",
  "description": "Software engineering team",
  "members": []
}
```

**Response:** `201 Created`

### Get All Groups
```http
GET /groups
```

**Response:** `200 OK`

### Get Group by ID
```http
GET /groups/{id}
```

**Example:**
```http
GET /groups/GRP001
```

### Update Group
```http
PUT /groups/{id}
Content-Type: application/json

{
  "name": "Updated Engineering Team",
  "description": "Updated description",
  "members": ["UI000001", "UI000002"]
}
```

### Delete Group
```http
DELETE /groups/{id}
```

**Response:** `204 No Content`

### Add User to Group
```http
POST /groups/{id}/users
Content-Type: application/json

{
  "user_id": "UI000001"
}
```

**Response:** `204 No Content`

### Add Multiple Users to Group
```http
POST /groups/{id}/users/bulk
Content-Type: application/json

{
  "user_ids": ["UI000001", "UI000002", "UI000003"]
}
```

**Response:** `204 No Content`

### Remove User from Group
```http
DELETE /groups/{id}/users/{userId}
```

**Example:**
```http
DELETE /groups/GRP001/users/UI000001
```

**Response:** `204 No Content`

### Remove Multiple Users from Group
```http
DELETE /groups/{id}/users/bulk
Content-Type: application/json

{
  "user_ids": ["UI000001", "UI000002"]
}
```

**Response:** `204 No Content`

### Get Group Members
```http
GET /groups/{id}/members
```

**Response:** `200 OK`
```json
{
  "members": ["UI000001", "UI000002", "UI000003"]
}
```

### Get User's Groups
```http
GET /users/{userId}/groups
```

**Example:**
```http
GET /users/UI000001/groups
```

**Response:** `200 OK`
```json
[
  {
    "id": "GRP001",
    "name": "Engineering Team",
    "description": "Software engineering team",
    "members": ["UI000001", "UI000002"]
  }
]
```

---

## Roles

### Create Role
```http
POST /roles
Content-Type: application/json

{
  "id": "ROLE001",
  "name": "Admin",
  "description": "Administrator role with full permissions",
  "groups": []
}
```

**Response:** `201 Created`

### Get All Roles
```http
GET /roles
```

**Response:** `200 OK`

### Get Role by ID
```http
GET /roles/{id}
```

**Example:**
```http
GET /roles/ROLE001
```

### Update Role
```http
PUT /roles/{id}
Content-Type: application/json

{
  "name": "Super Admin",
  "description": "Updated admin role",
  "groups": ["GRP001"]
}
```

### Delete Role
```http
DELETE /roles/{id}
```

**Response:** `204 No Content`

### Add Group to Role
```http
POST /roles/{id}/groups
Content-Type: application/json

{
  "group_id": "GRP001"
}
```

**Response:** `204 No Content`

### Add Multiple Groups to Role
```http
POST /roles/{id}/groups/bulk
Content-Type: application/json

{
  "group_ids": ["GRP001", "GRP002", "GRP003"]
}
```

**Response:** `204 No Content`

### Remove Group from Role
```http
DELETE /roles/{id}/groups/{groupId}
```

**Example:**
```http
DELETE /roles/ROLE001/groups/GRP001
```

**Response:** `204 No Content`

### Remove Multiple Groups from Role
```http
DELETE /roles/{id}/groups/bulk
Content-Type: application/json

{
  "group_ids": ["GRP001", "GRP002"]
}
```

**Response:** `204 No Content`

### Get Role's Groups
```http
GET /roles/{id}/groups
```

**Response:** `200 OK`
```json
{
  "groups": ["GRP001", "GRP002", "GRP003"]
}
```

---

## User-Role Relationships

### Assign Role to User
```http
POST /users/{userId}/roles
Content-Type: application/json

{
  "role_id": "ROLE001"
}
```

**Response:** `204 No Content`

### Assign Multiple Roles to User
```http
POST /users/{userId}/roles/bulk
Content-Type: application/json

{
  "role_ids": ["ROLE001", "ROLE002", "ROLE003"]
}
```

**Response:** `204 No Content`

### Remove Role from User
```http
DELETE /users/{userId}/roles/{roleId}
```

**Example:**
```http
DELETE /users/UI000001/roles/ROLE001
```

**Response:** `204 No Content`

### Remove Multiple Roles from User
```http
DELETE /users/{userId}/roles/bulk
Content-Type: application/json

{
  "role_ids": ["ROLE001", "ROLE002"]
}
```

**Response:** `204 No Content`

### Get User's Roles
```http
GET /users/{userId}/roles
```

**Response:** `200 OK`
```json
[
  {
    "id": "ROLE001",
    "name": "Admin",
    "description": "Administrator role",
    "groups": ["GRP001"]
  }
]
```

---

## Bulk Operations

### Bulk Assign Role to Users
```http
POST /roles/{id}/users/bulk
Content-Type: application/json

{
  "user_ids": ["UI000001", "UI000002", "UI000003"]
}
```

**Response:** `204 No Content`

### Bulk Remove Role from Users
```http
DELETE /roles/{id}/users/bulk
Content-Type: application/json

{
  "user_ids": ["UI000001", "UI000002", "UI000003"]
}
```

**Response:** `204 No Content`

---

## Health Check

### Health Check
```http
GET /health
```

**Response:** `200 OK`
```json
{
  "status": "ok",
  "service": "lotus-directory-engine"
}
```

---

## Error Responses

### Common Status Codes
- `200 OK` - Request successful
- `201 Created` - Resource created successfully
- `204 No Content` - Request successful, no content to return
- `400 Bad Request` - Invalid JSON format or missing required fields
- `404 Not Found` - Resource not found
- `500 Internal Server Error` - Server error

### Error Response Format
```
HTTP/1.1 404 Not Found
Content-Type: text/plain

user not found: UI000999
```

---

## cURL Examples

### Complete Workflow Example

1. **Create a user:**
```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "id": "UI000001",
    "email": "john@company.com",
    "name": "John Doe"
  }'
```

2. **Create a group:**
```bash
curl -X POST http://localhost:8080/api/v1/groups \
  -H "Content-Type: application/json" \
  -d '{
    "id": "GRP001",
    "name": "Engineering",
    "description": "Engineering team"
  }'
```

3. **Add user to group:**
```bash
curl -X POST http://localhost:8080/api/v1/groups/GRP001/users \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "UI000001"
  }'
```

4. **Create a role:**
```bash
curl -X POST http://localhost:8080/api/v1/roles \
  -H "Content-Type: application/json" \
  -d '{
    "id": "ROLE001",
    "name": "Developer",
    "description": "Software developer role"
  }'
```

5. **Assign role to user:**
```bash
curl -X POST http://localhost:8080/api/v1/users/UI000001/roles \
  -H "Content-Type: application/json" \
  -d '{
    "role_id": "ROLE001"
  }'
```

6. **Get user with roles:**
```bash
curl http://localhost:8080/api/v1/users/UI000001
```