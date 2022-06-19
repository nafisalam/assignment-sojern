# Project

This application provides solution of the below problems

  1. Util - 'compare versions' 
  2. Application - 'tracking web server'.
  
Note:  The assignment code of /img and /ping APIs are in the api.go file of the server folder. api_test.go has the test cases for the api.go. The version check util code is in the version.go file of the util folder. version_test.go file has the test cases for version check.

This application has below folders structure.

1. .conf : It is a configuration related folder and it has all the environment and application related config parameters.

2. config : It is a store  for the config data

3. image : It has all the images related to this application

4. logs : Application logs are dumped in this folder.

5. server : It has code of the apis. There are two apis.(GET:/img and GET:/ping)

6. util : It has the version check code file and other files related to the application.

7. tmp : It has the ok file.

Prerequisite:

1. Install Go library in you system.
2. Setup the Go path.

Instructions to run the application locally:

 1. Check out the code from github.

 2. Go to the root folder of the application

 3. Run the command go mod vendor

 4. Run the command go run main.go

 5. Call the APIs. The web server run on port 8080. There are two apis. You can use postman or browser to call the APIs

    Open Postman and copy and paste the urls as below. You can also call the Apis from the web broowser.

    curl --location --request GET 'http://localhost:8080/img'
    curl --location --request GET 'http://localhost:8080/ping'

 6. To run the version check go to the util folder , Please run below commands
    1. go test -run TestCompareVersions 
    2. go test ./... -coverprofile cover.out
    3. go tool cover -func cover.out

 
