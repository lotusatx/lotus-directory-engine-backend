# Lotus Directory Engine - Backend API# Lotus Directory Engine Backend API



RESTful API for managing users, groups, and roles in the Lotus Directory Engine.A RESTful API for managing users, groups, and roles in a directory system built with Go, GORM, and PostgreSQL.



## Features## Getting Started



- **User Management**: CRUD operations for users with role assignment### Prerequisites

- **Group Management**: Create and manage groups with member operations- Go 1.19+

- **Role Management**: Define roles and assign them to users or groups- PostgreSQL database

- **GORM Integration**: PostgreSQL database with auto-migration- Required Go modules (see `go.mod`)

- **SSL/TLS Support**: Production-ready with certificate configuration

- **CORS Configuration**: Flexible cross-origin resource sharing### Environment Variables

- **Health Check**: Built-in health monitoring endpointCopy `.env.template` to `.env` and fill in your values:

```bash

## Quick Startcp .env.template .env

# Edit .env with your configuration

### Prerequisites```



- Go 1.19 or laterEssential variables:

- PostgreSQL database```

- (Optional) SSL/TLS certificates for HTTPSCONNECTION_STRING=postgresql://username:password@localhost:5432/database_name

PORT=8080

### Environment Configuration# Optional SSL/TLS configuration for API server

TLS_CERT_FILE=path/to/cert.pem

Only **one required setting** - the database connection string. Everything else is optional with sensible defaults.TLS_KEY_FILE=path/to/key.pem

# OR for PFX/P12 certificates

Copy `.env.template` to `.env` and configure:TLS_PFX_FILE=path/to/certificate.pfx

TLS_PFX_PASSWORD=pfx_password

```bash```

# REQUIRED

CONNECTION_STRING=postgres://user:password@host:port/dbname### Database SSL Configuration

The application handles PostgreSQL SSL in three modes:

# OPTIONAL (defaults shown)

SERVER_PORT=80801. **Development (Default)**: SSL disabled (`sslmode=disable`)

CORS_ORIGINS=http://localhost:3000,http://localhost:80802. **Production with SSL**: Use connection string with SSL parameters

3. **Custom SSL**: Specify SSL mode explicitly

# OPTIONAL - Enable HTTPS

TLS_CERT_FILE=/path/to/cert.pem#### Database SSL Examples:

TLS_KEY_FILE=/path/to/key.pem```bash

# OR# Development - SSL disabled (default behavior)

PFX_FILE=/path/to/cert.pfxCONNECTION_STRING=postgresql://user:pass@localhost:5432/db

PFX_PASSWORD=your-pfx-password

```# Production - SSL required

CONNECTION_STRING=postgresql://user:pass@prod-server:5432/db?sslmode=require

### Build and Run

# SSL with custom certificate

```bashCONNECTION_STRING=postgresql://user:pass@server:5432/db?sslmode=require&sslcert=client-cert.pem&sslkey=client-key.pem&sslrootcert=ca-cert.pem

# Install dependencies```

go mod download

### Running the Server

# Build```bash

go build -o lde# HTTP (Development)

go run main.go config.go

# Run

./lde# HTTPS (Production with certificates)

```TLS_CERT_FILE=cert.pem TLS_KEY_FILE=key.pem go run main.go config.go

```

The server will start on the configured port (default: 8080).

### API Server SSL/TLS

## API DocumentationThe server supports both HTTP and HTTPS:

- **HTTP**: Default for development

**📖 Complete API reference is available in [API.md](./API.md)**- **HTTPS**: Enabled when TLS_CERT_FILE and TLS_KEY_FILE are provided



The API documentation includes:The API will be available at:

- All endpoint definitions with request/response examples- HTTP: `http://localhost:8080/api/v1/`

- cURL examples for testing- HTTPS: `https://localhost:8080/api/v1/`

- Complete workflow examples

- Error handling and status codesHealth check: 

- HTTP: `http://localhost:8080/health`

### Quick Reference- HTTPS: `https://localhost:8080/health`



Base URL: `http://localhost:8080/api/v1`---



| Resource | Endpoint | Description |## API Documentation

|----------|----------|-------------|

| Users | `/users` | Create, read, update, delete users |### Base URL

| Groups | `/groups` | Manage groups and members |```

| Roles | `/roles` | Define and assign roles |http://localhost:8080/api/v1

| Health | `/health` | Service health check |```



## SSL/TLS Configuration### Response Format

All responses return JSON. Error responses include an error message in the response body.

### Using PEM Certificates

---

Set both environment variables:

```bash## User Endpoints

TLS_CERT_FILE=/path/to/certificate.pem

TLS_KEY_FILE=/path/to/privatekey.pem### Create User

``````http

POST /users

### Using PFX CertificatesContent-Type: application/json



Convert PFX to PEM format first:{

```bash  "id": "UI000001",

# Extract certificate  "email": "john.doe@company.com",

openssl pkcs12 -in certificate.pfx -clcerts -nokeys -out cert.pem  "name": "John Doe",

  "roles": [],

# Extract private key  "group_ids": []

openssl pkcs12 -in certificate.pfx -nocerts -out key-encrypted.pem}

```

# Decrypt private key (optional, for automation)

openssl rsa -in key-encrypted.pem -out key.pem**Response:** `201 Created`

``````json

{

Then use the PEM files as shown above.  "id": "UI000001",

  "email": "john.doe@company.com",

### PostgreSQL SSL Configuration  "name": "John Doe",

  "roles": [],

The database handler automatically configures SSL based on your environment. For local development without SSL:  "group_ids": []

- The connection will use `sslmode=disable` automatically}

- For production, use a properly configured PostgreSQL connection string with SSL enabled```



## Database Schema### Get All Users

```http

The application automatically creates/migrates these tables:GET /users

- `users` - User accounts with email and name```

- `groups` - Groups with descriptions and member lists

- `roles` - Roles with descriptions and group associations**Response:** `200 OK`

- `user_roles` - Many-to-many relationship between users and roles```json

[

## Kubernetes Deployment  {

    "id": "UI000001",

### Using Environment Variables (Simplest)    "email": "john.doe@company.com",

    "name": "John Doe",

Create a Kubernetes secret:    "roles": [],

```yaml    "group_ids": []

apiVersion: v1  }

kind: Secret]

metadata:```

  name: lde-secrets

type: Opaque### Get User by ID

stringData:```http

  CONNECTION_STRING: "postgres://user:password@postgres-service:5432/lde"GET /users/{id}

  TLS_CERT_FILE: "/certs/tls.crt"```

  TLS_KEY_FILE: "/certs/tls.key"

```**Example:**

```http

Reference in your deployment:GET /users/UI000001

```yaml```

apiVersion: apps/v1

kind: Deployment**Response:** `200 OK`

metadata:```json

  name: lde-backend{

spec:  "id": "UI000001",

  template:  "email": "john.doe@company.com",

    spec:  "name": "John Doe",

      containers:  "roles": [],

      - name: lde  "group_ids": []

        image: your-registry/lde-backend:latest}

        envFrom:```

        - secretRef:

            name: lde-secrets### Update User

        volumeMounts:```http

        - name: tls-certsPUT /users/{id}

          mountPath: /certsContent-Type: application/json

          readOnly: true

      volumes:{

      - name: tls-certs  "email": "john.updated@company.com",

        secret:  "name": "John Updated",

          secretName: tls-certificate  "roles": [],

```  "group_ids": []

}

### Using Azure Key Vault (Advanced)```



If you need Key Vault integration, use the Secrets Store CSI Driver:**Response:** `200 OK`

```json

```yaml{

apiVersion: secrets-store.csi.x-k8s.io/v1  "id": "UI000001",

kind: SecretProviderClass  "email": "john.updated@company.com",

metadata:  "name": "John Updated",

  name: azure-lde-secrets  "roles": [],

spec:  "group_ids": []

  provider: azure}

  parameters:```

    keyvaultName: "your-keyvault"

    objects: |### Delete User

      array:```http

        - objectName: "db-connection-string"DELETE /users/{id}

          objectType: secret```

        - objectName: "tls-certificate"

          objectType: secret**Response:** `204 No Content`

```

---

Or use external-secrets operator:

```yaml## Group Endpoints

apiVersion: external-secrets.io/v1beta1

kind: ExternalSecret### Create Group

metadata:```http

  name: lde-secretsPOST /groups

spec:Content-Type: application/json

  secretStoreRef:

    name: azure-keyvault-store{

    kind: SecretStore  "id": "GRP001",

  target:  "name": "Engineering Team",

    name: lde-secrets  "description": "Software engineering team",

  data:  "members": []

  - secretKey: CONNECTION_STRING}

    remoteRef:```

      key: db-connection-string

```**Response:** `201 Created`



## Development### Get All Groups

```http

### Project StructureGET /groups

```

```

lotus_directory_engine_backend/**Response:** `200 OK`

├── main.go                 # Application entry point

├── config.go              # Configuration loading### Get Group by ID

├── handlers/              # Database operations```http

│   ├── db_handler.go     # DB connection & migrationGET /groups/{id}

│   ├── user_handler.go   # User CRUD```

│   ├── group_handler.go  # Group management

│   └── role_handler.go   # Role management**Example:**

├── api/                   # REST endpoints```http

│   ├── server.go         # HTTP/HTTPS server setupGET /groups/GRP001

│   ├── user_endpoints.go```

│   ├── group_endpoints.go

│   └── role_endpoints.go### Update Group

└── models/               # Data models```http

    └── types.goPUT /groups/{id}

```Content-Type: application/json



### Running Tests{

  "name": "Updated Engineering Team",

```bash  "description": "Updated description",

go test ./...  "members": ["UI000001", "UI000002"]

```}

```

## Security Considerations

### Delete Group

- **Never commit `.env`** - It contains sensitive credentials```http

- **Use Kubernetes secrets** for production deploymentsDELETE /groups/{id}

- **Enable SSL/TLS** for production PostgreSQL connections```

- **Configure CORS** appropriately for your frontend domains

- **Rotate credentials** regularly via your secret management solution**Response:** `204 No Content`



## Troubleshooting### Add User to Group

```http

### Database Connection IssuesPOST /groups/{id}/users

- Verify `CONNECTION_STRING` format: `postgres://user:password@host:port/dbname`Content-Type: application/json

- Check PostgreSQL is running and accessible

- For local development, ensure `sslmode=disable` is appropriate{

  "user_id": "UI000001"

### SSL/TLS Issues}

- Verify certificate files exist at the specified paths```

- Ensure certificate and key match

- Check file permissions (readable by the application)**Response:** `204 No Content`



### CORS Issues### Add Multiple Users to Group

- Verify `CORS_ORIGINS` includes your frontend URL```http

- Check browser console for specific CORS errorsPOST /groups/{id}/users/bulk

Content-Type: application/json

## License

{

MIT  "user_ids": ["UI000001", "UI000002", "UI000003"]

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

## Role Endpoints

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

### Common Error Status Codes
- `400 Bad Request` - Invalid JSON format or missing required fields
- `404 Not Found` - Resource not found
- `500 Internal Server Error` - Server error

### Error Response Format
```json
HTTP/1.1 404 Not Found
Content-Type: text/plain

user not found: UI000999
```

---

## Examples

### Complete User Management Flow

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

---

## Secret Management

### Development vs Production

**Development:**
- Use `.env` files with environment variables
- Acceptable for local development and testing

**Production (Kubernetes):**
- Use Kubernetes secrets for sensitive data
- Integrate with external secret management via operators
- Never store secrets in container images or config files

### Kubernetes Secret Management

The application uses environment variables for configuration, making it perfect for Kubernetes deployments.

#### Basic Kubernetes Secrets

```yaml
# Create secret
apiVersion: v1
kind: Secret
metadata:
  name: app-secrets
type: Opaque
data:
  DB_PASSWORD: <base64-encoded-password>
  JWT_SECRET: <base64-encoded-jwt-secret>
  TLS_PASSWORD: <base64-encoded-tls-password>
```

```yaml
# Use in deployment
spec:
  template:
    spec:
      containers:
      - name: lotus-directory-engine
        env:
        - name: CONNECTION_STRING
          value: "postgresql://user:@db:5432/lotus_directory?sslmode=require"
        - name: DB_PASSWORD
          valueFrom:
            secretKeyRef:
              name: app-secrets
              key: DB_PASSWORD
        - name: JWT_SECRET
          valueFrom:
            secretKeyRef:
              name: app-secrets
              key: JWT_SECRET
```

#### Azure Key Vault Integration

Use the Azure Key Vault CSI driver to sync secrets from Azure Key Vault to Kubernetes secrets:

```yaml
# Install Azure Key Vault CSI driver
helm repo add csi-secrets-store-provider-azure https://azure.github.io/secrets-store-csi-driver-provider-azure/charts
helm install csi csi-secrets-store-provider-azure/csi-secrets-store-provider-azure

# SecretProviderClass
apiVersion: secrets-store.csi.x-k8s.io/v1
kind: SecretProviderClass
metadata:
  name: app-secrets
spec:
  provider: azure
  secretObjects:
  - data:
    - key: DB_PASSWORD
      objectName: database-password
    - key: JWT_SECRET
      objectName: jwt-signing-key
    secretName: app-secrets
    type: Opaque
  parameters:
    keyvaultName: "your-keyvault"
    tenantId: "your-tenant-id"
```

#### External Secrets Operator

Alternative approach using external-secrets operator:

```yaml
# Install external-secrets operator
helm repo add external-secrets https://charts.external-secrets.io
helm install external-secrets external-secrets/external-secrets

# SecretStore for Azure Key Vault
apiVersion: external-secrets.io/v1beta1
kind: SecretStore
metadata:
  name: azure-keyvault
spec:
  provider:
    azurekv:
      vaultUrl: "https://your-vault.vault.azure.net/"
      authType: ManagedIdentity

# ExternalSecret
apiVersion: external-secrets.io/v1beta1
kind: ExternalSecret
metadata:
  name: app-secrets
spec:
  refreshInterval: 1h
  secretStoreRef:
    kind: SecretStore
    name: azure-keyvault
  target:
    name: app-secrets
  data:
  - secretKey: DB_PASSWORD
    remoteRef:
      key: database-password
  - secretKey: JWT_SECRET
    remoteRef:
      key: jwt-signing-key
```

### Other Secret Management Solutions

#### HashiCorp Vault
Use the Vault CSI provider or Vault Agent injector:
```bash
helm repo add hashicorp https://helm.releases.hashicorp.com
helm install vault hashicorp/vault
```

#### AWS Secrets Manager
Use the AWS Load Balancer Controller with External Secrets:
```yaml
provider:
  aws:
    service: SecretsManager
    region: us-east-1
```

### Security Best Practices

1. **Never commit secrets** to version control
2. **Use least privilege** access for secret retrieval
3. **Rotate secrets regularly** (30-90 day cycle)
4. **Monitor secret access** and set up alerts
5. **Use separate secrets** for different environments
6. **Encrypt secrets at rest** in your secret management system

### Development Setup

```bash
# Copy template and configure for local development
cp .env.template .env

# Edit .env with your development values
CONNECTION_STRING=postgresql://postgres:password@localhost:5432/lotus_directory
JWT_SECRET=your-development-jwt-secret-min-256-bits
```

### Production Deployment

```bash
# Create Kubernetes secret
kubectl create secret generic app-secrets \
  --from-literal=DB_PASSWORD='production_password' \
  --from-literal=JWT_SECRET='production_jwt_secret'

# Deploy with secret references in your Helm chart or deployment YAML
```

---

## SSL/TLS Configuration

### Database SSL Modes
PostgreSQL supports several SSL modes. The application automatically handles SSL configuration:

| Mode | Description | Security Level |
|------|-------------|----------------|
| `disable` | No SSL (default for development) | None |
| `require` | SSL required, no certificate verification | Basic |
| `verify-ca` | SSL required, verify CA certificate | Medium |
| `verify-full` | SSL required, verify CA and hostname | High |

### Database SSL Configuration Examples

#### Development (Local PostgreSQL)
```bash
# SSL disabled - default behavior
CONNECTION_STRING=postgresql://user:password@localhost:5432/database

# The application automatically adds: ?sslmode=disable
```

#### Production (Cloud/Remote PostgreSQL)
```bash
# Basic SSL - just encryption
CONNECTION_STRING=postgresql://user:password@prod-server:5432/database?sslmode=require

# Full SSL verification with certificates
CONNECTION_STRING=postgresql://user:password@prod-server:5432/database?sslmode=verify-full&sslcert=/path/to/client.crt&sslkey=/path/to/client.key&sslrootcert=/path/to/ca.crt
```

#### Azure PostgreSQL Example
```bash
CONNECTION_STRING=postgresql://username%40servername:password@servername.postgres.database.azure.com:5432/database?sslmode=require
```

#### AWS RDS PostgreSQL Example
```bash
CONNECTION_STRING=postgresql://username:password@mydb.123456789.us-east-1.rds.amazonaws.com:5432/database?sslmode=require
```

### API Server HTTPS Configuration

#### Generate Self-Signed Certificate (Development)
```bash
# Generate private key
openssl genrsa -out server.key 2048

# Generate certificate signing request
openssl req -new -key server.key -out server.csr

# Generate self-signed certificate
openssl x509 -req -days 365 -in server.csr -signkey server.key -out server.crt

# Convert to PEM format (if needed)
cp server.crt cert.pem
cp server.key key.pem
```

#### Production Certificate Formats

**PEM Format (Linux/Apache)**
```bash
# Generate or obtain PEM certificates
TLS_CERT_FILE=/path/to/certificate.pem
TLS_KEY_FILE=/path/to/private-key.pem
```

**PFX/P12 Format (Windows/IIS)**
```bash
# Use PFX certificate with password
TLS_PFX_FILE=/path/to/certificate.pfx
TLS_PFX_PASSWORD=your_pfx_password

# Note: PFX support requires converting to PEM format:
# openssl pkcs12 -in certificate.pfx -out cert.pem -clcerts -nokeys
# openssl pkcs12 -in certificate.pfx -out key.pem -nocerts -nodes
```
**Let's Encrypt (Free SSL)**
```bash
# Install certbot
sudo apt-get install certbot

# Generate certificate
sudo certbot certonly --standalone -d yourdomain.com

# Certificates will be in:
# /etc/letsencrypt/live/yourdomain.com/fullchain.pem
# /etc/letsencrypt/live/yourdomain.com/privkey.pem
```

#### Environment Variables for HTTPS
```bash
# Set certificate paths
export TLS_CERT_FILE="/path/to/cert.pem"
export TLS_KEY_FILE="/path/to/key.pem"

# Or in .env file
TLS_CERT_FILE=/etc/ssl/certs/server.crt
TLS_KEY_FILE=/etc/ssl/private/server.key
```

### SSL Best Practices

#### Development
- Use `sslmode=disable` for local PostgreSQL
- Use self-signed certificates for API server testing
- Test SSL configuration before deploying to production

#### Production
- **Always use SSL** for database connections (`sslmode=require` minimum)
- Use valid certificates from trusted CA for API server
- Consider certificate auto-renewal (Let's Encrypt)
- Monitor certificate expiration dates

#### Security Considerations
1. **Database**: Never send credentials over unencrypted connections in production
2. **API**: Always use HTTPS in production to protect API keys and user data
3. **Certificates**: Store private keys securely with proper file permissions (600)
4. **Headers**: Consider adding security headers (HSTS, CSP, etc.)

### Troubleshooting SSL Issues

#### Common Database SSL Errors
```bash
# Error: server does not support SSL
# Solution: Check if PostgreSQL has SSL enabled
```

```bash
# Error: certificate verify failed
# Solution: Use sslmode=require instead of verify-full for basic SSL
```

```bash
# Error: connection refused
# Solution: Check if SSL port (usually 5432) is open and SSL is enabled
```

#### Common API Server SSL Errors
```bash
# Error: certificate signed by unknown authority
# Solution: Use proper CA-signed certificate or add to trust store
```

```bash
# Error: cannot load certificate/key files
# Solution: Check file paths and permissions
```

### Certificate Management

#### File Permissions
```bash
# Set correct permissions for certificate files
chmod 600 /path/to/private.key
chmod 644 /path/to/certificate.crt
chown app:app /path/to/certificate.*
```

#### Certificate Renewal Script
```bash
#!/bin/bash
# renew-cert.sh
certbot renew --quiet
systemctl restart lotus-directory-engine
```

#### Monitoring Certificate Expiry
```bash
# Check certificate expiration
openssl x509 -in cert.pem -text -noout | grep "Not After"

# Set up monitoring alert 30 days before expiry
```

---

## Development

### Database Models
The API uses three main models:
- **User**: ID, Email, Name, Roles, GroupIDs
- **Group**: ID, Name, Description, Members
- **Role**: ID, Name, Description, Groups

### Database Migration
Database tables are automatically created/updated on server startup using GORM's AutoMigrate feature.

### CORS
CORS is enabled for all origins in development. Configure appropriately for production use.