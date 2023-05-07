# goapicrud

This is a golang API based on Gin Framework and GORM ORM. 
It uses Postgres backend based off of free ElephantSQL managed Postgres.
The API supports create/Update/delete/GetByID and GetAll features for a post object (located in Models folder postModel.go) which contains a title and body.
I used GORM to to migrate from Model to table creation in postgres using migrate.go
A controller folder has postController.go which helps with implementing the functions for each of controller action of CRUD operation along with Gin framework's context.
The Main.go uses gin for route mapping of end points to function handlers in controller using gin framework,
the Initializers 2 files
  loadEnvVariables.go has code to read environment variables like port no and db connection string from .env file using godotenv package
  database.go has code to connect to postgres sql using gorm's postgres driver
