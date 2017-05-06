# eLogcard Application 
=======

##Application overview##

This application is designed to demonstrate how assets can be modeled on the Blockchain using part lifecycle scenario. 
In the scenario parts and logs  are modeled using Blockchain technology with the following attributes:


Parts 

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

Logs 

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

| Participant       | Permissions                                                                  |
| ------------------| -----------------------------------------------------------------------------|
| Auditor_authority | Read/Display (All parts)                                				       |
| AH_admin		    | Read/Display (All parts)                                           		   |
| Supplier   	 	| Create, Read/Display (Own parts), Transfer (Own Parts)					   |
| Manufacturer   	| Create, Read/Display (Own parts), Transfer (Own Parts)        		       |
| Customer		    | Read/Display (Own parts),  Transfer (own parts),               		       |
| Maintenance_user 	| Read/Display (Own parts), Perform Activities(own parts), Transfer (Own Parts)|


- Only Suppliers and Manufacturers can create Part 
- Suppliers, manufacturers, customers and maintenance users can display details of all the parts they own.
- Auditor_authority and AH_Admin can display details of all the parts ever created.
- Only suppliers, manufacturers, Customers and maintenance_user can Transfer Responsibility on a Part provided that they are currently owner of this part.
- Only suppliers, manufacturers, Customers and maintenance_user can Transfer Ownership on a Part provided that they are currently owner of this part.
- Only maintenance_user can perform acts on a part provided that he/she is the current owner of this part.
- Suppliers, manufacturers, customers and maintenance users can  display details on a specific part only if they own it.
- Auditor_authority and AH_Admin can see details on any specific part they want.


####Stages:####

1) Créer des Parts
2) Récupérer et Afficher les Parts Créées 
3) Effectuer des transferts de responsabilité et de propriété
4) Effectuer des activités de maintenance sur la Part 
5) Visualiser l'Historique d'une Part 








