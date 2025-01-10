# **Task Management REST API with Go**

## **Project Overview**

This project is designed to practice building a **REST API** using **Go** with modern development tools and best practices.

###  **Key Features**
1. **CRUD Operations:** Create, Read, Update, and Delete tasks.  
2. **Database Integration:** Store task data using **PostgreSQL**.
3. **Frontend:** A simple frontend for interacting with the API.
4. **Containerization:** Use **Docker** and **Docker Compose** for simplified deployment.  
5. **Scalability:** Deploy and scale the application with **Kubernetes**.  

---

## **Tech Stack**

- **Programming Language:** Go  
- **Database:** PostgreSQL  
- **Containerization:** Docker, Docker Compose  
- **Orchestration:** Kubernetes  
- **Tools:** Curl  

---


## **Setup Instructions (with docker compose)**

###  1. Clone the Repository


```bash

git clone git@github.com:Mimadfaoussi/GO.git
cd projects/task-manager-api
```

###  2. Build and Start Services with Docker Compose

```bash
docker compose up --build -d
```
<ul>
<li> --build: Ensures images are rebuilt. </li>
<li> -d: Runs containers in detached mode. </li>
</ul>


###  3. Verify Services
List running containers:
```bash
docker ps
```

Verify PostgreSQL connection:
```bash
docker exec -it postgres-taskdb psql -U taskuser -d taskdb
```

To stop the services:

```bash
docker compose stop
```

## **Setup Instructions (To run it manually )**

### 1. Clone the Repository


```bash

git clone git@github.com:Mimadfaoussi/GO.git
cd projects/task-manager-api
```

###  2. Set Up PostgreSQL Using Docker

```bash
sudo docker run --name postgres-taskdb \
  -e POSTGRES_DB=taskdb \
  -e POSTGRES_USER=taskuser \
  -e POSTGRES_PASSWORD=taskpassword \
  -p 5432:5432 -d postgres
```

### 3. Verify Database Connection 

we can verify that by connecting to the container using the following command:

```bash
docker exec -it postgres-taskdb psql -U taskuser -d taskdb
```

### 4. Run Database Schema 

```bash
docker cp /models/schema.sql postgres-taskdb:/schema.sql
docker exec -it postgres-taskdb psql -U taskuser -d taskdb -f /schema.sql
```


##  Run the Application

###  1. Start the API Server
```bash
go run cmd/main.go
```

###  Testing API Endpoints

we can test the endpoints locally using curl


####  Get All Tasks

```bash
curl http://localhost:8080/tasks
```


####  Create a Task

```bash
curl -X POST -H "Content-Type: application/json" \
-d '{"title": "Build API", "description": "Using Go and Docker", "status": "Pending"}' \
http://localhost:8080/tasks/create

```
#### Update a Task

```bash
curl -X PUT -H "Content-Type: application/json" \
-d '{"id": 1, "title": "Learn Go", "description": "Practice APIs", "status": "Done"}' \
http://localhost:8080/tasks/update
```

####  Delete a Task

```bash
curl -X DELETE -H "Content-Type: application/json" \
-d '{"id": 1}' http://localhost:8080/tasks/delete

```

#### Access the frontend

The frontend is served as a separate microservice using Nginx and can be accessed at :


```bash
http://localhost:3000
```

