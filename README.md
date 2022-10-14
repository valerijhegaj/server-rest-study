# Server golang study
## Tasks

### The main goal
In this project I will try to write a restful server with the authorization database and the queries described below, 
it will be server to store data

### Available http requests
    POST   /api/users                  - create new user
    POST   /api/session                - create new access token for user
    GET    /api/files/{user_id}/{path} - get file from user files
    POST   /api/files/{user_id}/{path} - add new file to cloud
    PUT    /api/files/{user_id}/{path} - update file
    DELETE /api/files/{user_id}/{path} - delete file
### Steps
1. Write minimal golang server with all info in ram, and files in file system (implemented v0.1)
2. Add isolation files from everybody (implemented v0.2)
3. Add rw permissions to files (implemented v0.3)
4. Add database to store info about users
5. Add more http request (a. change password, b. delete user, c. get file structure, d. get user_id)
6. 
