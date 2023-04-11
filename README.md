## Setting Up Projects
1. Clone this repository
2. Create and set up your own environment variables by using the **.env.example file**. Input your postgres database configuration and your jwt secret key
3. And now you're ready to go
4. Run this project `go run main.go`

## Package or Library Used
- Go Framework: **Gin**
- ORM: **GORM**
- Database: **PostgreSQL**
- Viper

## API Documentation
https://documenter.getpostman.com/view/18726863/2s93Xu363y

## Authentication and Authorization
1. To create, read, update and delete product, all users have to provide token
2. Admin can do CRUD on all product data (product authorization all)
3. User can only create, read and update their own product data (product authorization by id)