# eLogcard Application 
=======

##Application overview##
This application is designed to demonstrate how assets can be modeled on the Blockchain using a car leasing scenario. 
In the scenario vehicles are modeled using Blockchain technology with the following attributes:

| Attribute       | Type                                                                                                  |
| --------------- | ----------------------------------------------------------------------------------------------------- |
| V5cID           | Unique string formed of two chars followed by a 7 digit int, used as the key to identify the vehicle  |
| VIN             | 15 digit int                                                                                          |
| Make            | String                                                                                                |
| Model           | String                                                                                                |
| Colour          | String                                                                                                |
| Reg             | String                                                                                                |
| Owner           | Identity of participant                                                                               |
| Scrapped        | Boolean                                                                                               |
| Status          | Int between 0 and 4                                                                                   |
| LeaseContractID | ChaincodeID, currently unused but will store the ID of the lease contract for the vehicle             |


The application is designed to allow participants to interact with the part assets creating, 
updating and transferring them as their permissions allow. The participants included in the application are as follows:

| Participant       | Permissions                                                          |
| ------------------| ---------------------------------------------------------------------|
| Auditor_authority | Read (All parts)                                				       |
| AH_admin		    | Read (All parts)                                           		   |
| Supplier   	 	| Create, Read (Own parts), Transfer (Own Parts)					   |
| Manufacturer   	| Create, Read (Own parts), Transfer (Own Parts)        		       |
| Customer		    | Read (Own parts),  Transfer (own parts),               		       |
| Maintenance_user 	| Read (Own parts), Perform Activities(own parts), Transfer (Own Parts)|




####Stages:####

1) Créer des Parts

- Only Suppliers and Manufacturers can create Part 

2) Récupérer et Afficher les Parts Créées 

- Suppliers, manufacturers, customers and maintenance users can display details of all the parts they own.
- Auditor_authority and AH_Admin can display details of all the parts ever created.

3) Effectuer des transferts de responsabilité et de propriété

- Only suppliers, manufacturers, Customers and maintenance_user can Transfer Responsibility on a Part provided that they are currently owner of this part.
- Only suppliers, manufacturers, Customers and maintenance_user can Transfer Ownership on a Part provided that they are currently owner of this part.

4) Effectuer des activités de maintenance sur la Part 

- Only maintenance_user can perform acts on a part provided that he/she is the current owner of this part.

5) Visualiser l'Historique d'une Part 

- Suppliers, manufacturers, customers and maintenance users can  display details on a specific part only if they own it.
- Auditor_authority and AH_Admin can see details on any specific part they want.






