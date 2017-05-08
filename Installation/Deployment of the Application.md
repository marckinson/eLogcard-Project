# Installation de l'application




## Clone the repository 

- $ git clone https://github.com/marckinson/eLogcard-Project.git

## Init Environments

The cloned directory must contains 3 main folders:
-	folder 		Docker 
-	folder	 	e-logcard-application
-	folder		Docs Installation


#### 1) Launch your Docker environment

In the Folder named Docker
- $ start.sh -t
- $ start.sh

Help with docker commands
- $ docker-compose up				: launch the containers with the logs 
- $ docker-compose up -d			: launch the containers without the logs 
- $ docker-compose down  		    : removed launched containers 
- $ docker rm [container name]    : delete a peer container 

#### 2) Launch e-logcard application 

In the folder named e-logcard-application 

- $ npm install (You should now see a new folder named node_modules)
- $ node app.js (Starts the application)

## Interact with the Application 

- http://localhost:3000/




