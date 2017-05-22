# eLogcard Application 

## Scenario overview 

#### Create users  

| Name       | Role                  |
| -----------| ----------------------|
| TurboMeca  | Supplier				 |
| Ben  	     | Manufacturer          |
| Tom	     | Customer              |
| Harry 	 | Maintenance_user      |
| AH   		 | AH_admin              |
| EASA 		 | Auditor_authority     |



==============================================================================

- Supplier Creates a part.
- Supplier Transfers responsibility and ownership on this part to AH.
- AH_Admin displays historic of the part.
- AH transfers responsibility and ownership on the part the Customer who bought the Aircraft.
- Customer needs his part to be revised. He transfers responsibility to Maintenance_user and sends the part to Maintenance_user.
- Maintenance_user performs the necessary acts on the part.
- Maintenance_user transfers responsibility on the part to Customer and send the part back to Customer.
- Customer Displays his part historic to see performed acts on the part.

==============================================================================

- Manufacturer Creates a part
- Manufacturer Transfers responsibility and ownership on this part to AH
- AH_Admin displays historic of the part.
- AH transfers responsibility and ownership on the  part to the Customer who bought the Aircraft
- Customer needs his part to be revised. He transfers responsibility to Maintenance_user and sends the part to Maintenance_user.
- Maintenance_user performs the necessary acts on the part.
- Maintenance_user transfers responsibility on the part to Customer and send the part back to Customer.
- Customer Displays his part historic to see performed acts on the part.
 
==============================================================================

 - EASA Displays historic of Customers Parts 


==============================================================================

#### Link to documentations (pre-requisite and installation): 

- https://github.com/marckinson/eLogcard-Project/blob/master/Installation/Pre%20requisite.md
- https://github.com/marckinson/eLogcard-Project/blob/master/Installation/Deployment%20of%20the%20Application.md


