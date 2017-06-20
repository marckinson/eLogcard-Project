# eLogcard Application 

## Application overview 

This application is designed to demonstrate how assets can be modeled on the Blockchain using a part lifecycle scenario. 
In the scenario parts, aicrafts, assemblies and logs  are modeled using Blockchain technology with the following attributes:

## Assets 

#### Asset 1: Parts 

| Attributes      | Type                   |   
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
| Logs 			  | [] Logs                   |


#### Asset 2: Aircraft 

| Attributes      | Type                   |
| --------------- | ---------------------- |
| AN     		  | String  			   |
| SN      	      | string                 |
| Id_Aircraft     | String  			   |
| Owner           | string                 |
| Parts    	      | []string               |
| Assemblies      | string                 |
| Logs       	  | []Logs                 |

#### Asset 3: Assembly 

| Attributes      | Type                   |
| --------------- | ---------------------- |
| AN     		  | String  			   |
| SN      	      | string                 |
| Id_Assembly     | String  			   |
| Helicopter      | string                 |
| Owner		      | string                 |
| Parts    	      | []string               |
| Logs       	  | []Logs                 |


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
| REMOVE_MANAGER 	| |
| SCRAPPING_MANAGER	| |


## Functionnalities 


![](/desktop/test.png)


| Asset      | Functionnality                   |
| --------------- | ---------------------- |
| User      		  | Create Users  			   |
|       	      | Log in Users
| Part     | CreatePart  			   |
|       | getAllPartsDetails                 |
| 		      | getPartDetails                 |
|     	      | responsibilityTransfer               |
|        	  | ownershipTransfer                 |
|       | performActivities                 |
| 		      | scrappPart                 |
| Aircraft      | createAircraft               |
|        	  | ownershipTransfer                 |
|        	  | getAcDetails
|			  |AcPartsListing
|        	  | addPartToAc                 |
|        	  | RemovePartFromAc                 |
|        	  | ReplacePartOnAircraft                 |
|        	  | AcOwnershipTransfer                 |
|        	  | RemovePartFromAc                 |
|        	  | ReplacePartOnAircraft                 |
|        	  | AddAssemblyToAc                 |
|        	  | AcAssembliesListing                 |
|        	  | RemoveAssemblyFromAc                 |
|        	  | scrappAircraft                 |
|Assembly  	  | createAssembly                 |
|        	  | getAllAssembliesDetails                 |
|        	  | getAssembDetails                 |
|        	  | AssembPartsListing                 |
|        	  | addPartToAssemb                 |
|        	  | RemovePartFromAs                 |
|        	  | AssembOwnershipTransfer                 |
|        	  | ReplacePartOnAssembly                 |
|        	  | scrappAssembly |

Notes:

- (1): According to user's role & rights 
- (2): According to user's role & rights 
- (3): According to user's role & rights 