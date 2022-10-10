# Server golang study
## Tasks

### The main goal
In this project I will try to write a restful server with the authorization database and the queries described below, 
it will be server to store data

### Available http requests
    GET    /api/users            - get public info about all users
    POST   /api/users            - create new user
    GET    /api/users/{user_id}  - get public info about user_id
    POST   /api/session          - create new access token for user
    GET    /api/{user_id}/{path} - get file from user files
    POST   /api/{user_id}/{path} - add new file to cloud
    PUT    /api/{user_id}/{path} - update file
    DELETE /api/{user_id}/{path} - delete file
### Steps
1. Write golang server with all info in ram, and files in file system
2. Add database to store info about users and define where will be files
