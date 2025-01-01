The goal of this project is to practice using Go to build rest api's

Project Overview:

The task management rest api should allow users to :

1 - Create, Read, Update, and Delete (CRUD) tasks.
2 - Store task data in a PostgreSQL database.
3 - Containerize the application using Docker and Docker Compose.
4 - Deploi and Scale the application using kubernetes.

in order to test the Created EndPoints locally : 

GetAll : curl http://localhost:8080/tasks

Create : curl -X POST -H "Content-Type: application/json" -d '{"title": "Build API", "description": "Using Go and Docker", "status": "Pending"}' http://localhost:8080/tasks/create

Update : curl -X PUT -H "Content-Type: application/json" -d '{"id": 1, "title": "Learn Go", "description": "Practice APIs", "status": "Done"}' http://localhost:8080/tasks/update

Delete : curl -X DELETE -H "Content-Type: application/json" -d '{"id": 1}' http://localhost:8080/tasks/delete

