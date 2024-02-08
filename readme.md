# Menu Updatation System

## Introduction

The Menu Updatation System is a project that enables users to create an account in a local PostgreSQL database and add menus that can be viewed after logging in. The project consists of a basic frontend built in Angular and a backend that handles functionalities such as adding new users, user authentication, adding menus, and fetching menus. The backend is implemented in Go, and the database is configured in PostgreSQL. Additionally, the backend is deployed on an EC2 instance, while the frontend is deployed using S3 and CloudFront.

## Future Scope

- Integrate AWS Secret Manager for enhanced security.
- Improve cookie management for better user experience.

### How to run
DB should be configured and using .env file.
```bash
go run main.go
```