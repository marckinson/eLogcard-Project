# eLogcard Application 

Link to documentations (pre-requisite and installation): 
- https://github.com/marckinson/eLogcard-Project/blob/master/Docs%20Installation/Pr%C3%A9-requis.md 
- https://github.com/marckinson/eLogcard-Project/blob/master/Docs%20Installation/Deploiement%20de%20l'application.md
 

##Application overview##

This application is designed to demonstrate how assets can be modeled on the Blockchain using part lifecycle scenario. 
In the scenario parts and logs  are modeled using Blockchain technology with the following attributes:


#### Asset 1: Parts 

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

#### Asset 2: Logs 

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
| AH_admin		    | Read/Display (All parts), Transfer Ownership/Responsibility (On their own Parts)   				   				            |
| Supplier   	 	| Create, Read/Display (their Own parts), Transfer Ownership/Responsibility (On their own Parts)					            |
| Manufacturer   	| Create, Read/Display (their Own parts), Transfer Ownership/Responsibility (On their own Parts)        		                |
| Customer		    | Read/Display (their Own parts),  Transfer Ownership/Responsibility (On their own parts),               		                |
| Maintenance_user 	| Read/Display (their Own parts), Perform Activities(on their own parts), Transfer Ownership/Responsibility (on their own Parts)|

#### Stages of the scenario ####


==============================================================================
- Create Users

| Name       | Role                  |
| -----------| ----------------------|
| EASA 		 | Auditor_authority     |
| AH   		 | AH_admin              |
| TurboMeca  | Supplier				 |
| Ben  	     | Manufacturer          |
| Tom	     | Customer              |
| Bob	     | Customer              |
| Harry 	 | Maintenance_user      |
| Aaron 	 | Maintenance_user      |

==============================================================================
- TurboMeca Creates 1 part  
- TurboMeca Transfers responsibility and ownership on this part to AH
- AH displays the part
- AH transfers responsibility and ownership on the  part to Tom the Customer who just bought an Aircraft
- Tom's part needs to be revised. He transfers responsibility and ownership to Harry.
- Harry perform the necessary acts on the part.
- Harry send the part back to Tom.
- Tom Display his  part historic 

==============================================================================
- Ben Creates 1 part
- Ben Transfers responsibility and ownership on this part to AH
- AH displays the part
- AH transfers responsibility and ownership on the  part to Bob the Customer who just bought an Aircraft
- Bob's part needs to be revised. He transfers responsibility and ownership to Aaron.
- Aaron perform the necessary acts on the part.
- Aaron send the part back to Tom.
- Bob Display his  part historic
 
==============================================================================
- EASA Display historic of Bob and Tom Parts 







