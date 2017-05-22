# eLogcard Application 

## Application overview 

This application is designed to demonstrate how assets can be modeled on the Blockchain using a part lifecycle scenario. 
In the scenario parts, aicrafts, assemblies and logs  are modeled using Blockchain technology with the following attributes:

## Assets 

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

#### Asset 3: Aircraft 

| Attribute       | Type                   |
| --------------- | ---------------------- |
| Id_Aircraft     | String  			   |
| Owner           | string                 |
| Id_Parts        | []string               |

#### Asset 4: Assembly 

| Attribute       | Type                   |
| --------------- | ---------------------- |
| Id_Assembly     | String  			   |
| Owner           | string                 |
| Id_Parts        | []string               |


## Participants 
 
The application is designed to allow participants to interact with the part assets creating, 
updating and transferring them as their permissions allow. The participants included in the application are as follows:

| Participant       | Rights                                                                                                                   |
| ------------------| ------------------------------------------------------------------------------------------------------------------------------|
| Auditor_authority | Read/Display (All parts)                                				      						  				            |
| AH_admin		    | Read/Display (All parts), Transfer Ownership/Responsibility (On their own Parts)   				   				            |
| Supplier   	 	| Create, Read/Display (their Own parts), Transfer Ownership/Responsibility (On their own Parts)					            |
| Manufacturer   	| Create, Read/Display (their Own parts), Transfer Ownership/Responsibility (On their own Parts)        		                |
| Customer		    | Read/Display (their Own parts),  Transfer Ownership/Responsibility (On their own parts),               		                |
| Maintenance_user 	| Read/Display (their Own parts), Perform Activities(on their own parts), Transfer Ownership/Responsibility (on their own Parts)|


## Functionnalities 

| Functionnalities        |
| ------------------|
| Create Users | 
| Log in Users		    | 
| Create Parts   	 	| 
| Create Aircrafts   	| 
| Create Assemblies		    | 
| Display Parts (1)	|
| Display a Part History | 
| Display Assemblies (2)	| 
| Display an assembly history|
| Display Aircrafts(3)	| 
| Display an aircraft history| 
| Transfer Responsibility	| 
| Transfer Ownership	| 
| Perform Activities 	| 


Notes:

- (1): According to user's role & rights 
- (2): According to user's role & rights 
- (3): According to user's role & rights 








