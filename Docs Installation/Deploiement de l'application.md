# Installation de l'application

https://github.com/marckinson/eLogcard-Project


I- Clone the repository :

command line: git clone https://github.com/marckinson/eLogcard-Project.git

II- Init Environments

The cloned directory must contains 3 main folders:
-	folder 		Docker 
-	folder	 	e-logcard-application
-	folder		Docs Installation


1) Launch your Docker environment

In the Folder named Docker:
a) # ./start.sh -t
Tag the latest version of fabric-baseimage
....
b) # ./start.sh
Rebuild composed containers
....

Help with docker commands
- docker-compose up				: launch the containers with the logs 
- docker-compose up -d			: launch the containers without the logs 
- docker-compose down  		    : removed launched containers 
- docker rm [container name]    : delete a peer container 

2) Launch e-logcard application 

In the folder named e-logcard-application 

a) npm install 
Once you did this, you should see a new folder named node_modules

b) node app.js 
This command start the application

III - Interact with the Application 

http://localhost:3000/


Link to the scenario : https://github.com/marckinson/eLogcard-Project/blob/master/README.md


