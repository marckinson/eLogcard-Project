# eLogcard Application 


##Application overview##

This application is designed to demonstrate how assets can be modeled on the Blockchain using part lifecycle scenario. 
In the scenario parts and logs  are modeled using Blockchain technology with the following attributes:


Asset 1: Parts 

| Attribute       | Type                   |   
| --------------- | ---------------------- |   
| PN           	  | String  			   |    
| SN              | String                 |
| Id              | UUID                   |
| PartName        | String                 |
| Type            | String                 |
| Owner           | String                 |
| Responsible     | String                 |
| Helicopter      | String                 |
| Assembly        | String                 |
| Logs 			  | Logs                   |

Asset 2: Logs 

| Attribute       | Type                   |
| --------------- | ---------------------- |
| Ltype           | String  			   |
| Vdate           | TimeStamps             |
| Owner           | String                 |
| Owner        	  | String                 |
| Responsible     | String                 |
| ModType         | String                 |
| Description     | String                 |


The application is designed to allow participants to interact with the part assets creating, 
updating and transferring them as their permissions allow. The participants included in the application are as follows:

| Participant       | Permissions                                                                                                                   |
| ------------------| ------------------------------------------------------------------------------------------------------------------------------|
| Auditor_authority | Read/Display (All parts)                                				      						  				            |
| AH_admin		    | Read/Display (All parts)                                           		   						   				            |
| Supplier   	 	| Create, Read/Display (their Own parts), Transfer Ownership/Responsibility (On their own Parts)					            |
| Manufacturer   	| Create, Read/Display (their Own parts), Transfer Ownership/Responsibility (On their own Parts)        		                |
| Customer		    | Read/Display (their Own parts),  Transfer Ownership/Responsibility (On their own parts),               		                |
| Maintenance_user 	| Read/Display (their Own parts), Perform Activities(on their own parts), Transfer Ownership/Responsibility (on their own Parts)|

####Stages:####

1) Créer des Parts
2) Récupérer et Afficher les Parts Créées 
3) Effectuer des transferts de responsabilité et de propriété
4) Effectuer des activités de maintenance sur la Part 
5) Visualiser l'Historique d'une Part 








