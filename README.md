# eLogcard Application 


Scenario: 

Parties prenantes:
Users: 
- Supplier 
- Manufacturer
- Customer
- Maintenance_user

Authorities: 
- AH_admin
- Auditor_authority


Desciption 
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






