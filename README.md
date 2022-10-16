# Server golang study
## Tasks

### The main goal
In this project I will try to write a restful server with the authorization database and the queries described below, 
it will be server to store data

### Available http requests
    POST   /api/users                   - create new user (body {"username":your_username, "password":your_password})
    POST   /api/session                 - create new access token for user (body {"username":your_username, "password":your_password}) in cookies returns access token
    GET    /api/files/{username}/{path} - get file from user files
    POST   /api/files/{username}/{path} - add new file to cloud
    PUT    /api/files/{username}/{path} - update file
    DELETE /api/files/{username}/{path} - delete file
    POST   /api/give_access             - give access to other user (body {"username":your_friend, "path":path, "rights": "rw"/"r"/""}), "rw" - your friend can write delete update and read this file, "r" - only read, "" - only for you, if you write "" in username it will be access rights for everybody
### Steps
1. Write minimal golang server with all info in ram, and files in file system (implemented v0.1)
2. Add isolation files from everybody (implemented v0.2)
3. Add rw permissions to files (implemented v0.3)
4. Add cookies support and api tests (implemented v0,4)
5. Add database to store info about users
6. Add more http request (a. change password, b. delete user, c. get file structure, d. get user_id)
