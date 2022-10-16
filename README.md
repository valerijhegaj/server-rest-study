# Server golang study
## Tasks

### The main goal
In this project I will try to write a restful server with the authorization database and the queries described below, 
it will be server to store data

### Available http requests
    POST   /api/user                    - create new user
    POST   /api/session                 - create new access token for user
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
5. Add more http request (a. change password, b. delete user, c. get file structure, d. get user_id)
6. Add database to store info about users
7. Add support big files

## Details of http requests
### /api/user
#### POST
    request:
      Cookie: not required
      body: {
               "username":your_username, 
               "password":your_password
            }
    response:
      201 - success creted user
      400 - bad request, offen when you sent bad body or username is empty
      403 - permission denied, user already exist, even password is right
### /api/session
#### POST
    request:
      Cookie: not required
      body: {
               "username":your_username, 
               "password":your_password,
               "max_age":cookies_lifetime
            }
    response:
      201 - success log in
        Set-Cookies: token=your_access_token
      400 - bad request, offen when you sent bad body or username is empty
      403 - permission denied, username or password incorrect
### /api/files/{username}/{path}
#### GET
    request:
      Cookie: token=your_access_token
    response:
      200 - success get file
      400 - bad request
      403 - permission denied, you have no rights to read this file
      404 - file don't exist
      500 - something went wrong inside server
#### POST
    request:
      Cookie: token=your_access_token
      body: {
               "file_data":data
            }
    response:
      201 - success created file
      400 - bad request
      403 - permission denied, you have no rights to create this file
      500 - something went wrong inside server
#### PUT
    request:
      Cookie: token=your_access_token
      body: {
               "file_data":data
            }
    response:
      201 - success created file
      400 - bad request
      403 - permission denied, you have no rights to create this file
      500 - something went wrong inside server
#### DELETE
    request:
      Cookie: token=your_access_token
    response:
      200 - success deleted file
      400 - bad request
      403 - permission denied, you have no rights to create this file
      500 - something went wrong inside server
### /api/give_access 
#### POST
    request:
      Cookie: token=your_access_token
      body: {
               "username":friend,
               "rights":"rw" (or "r" or ""),
               "path":path_to_file
            }
    response:
      201 - success created right
      400 - bad request
      401 - unauthorized
      500 - something went wrong inside server

