# e-Logcard Application 

## Description of the APIs

#### Non-secured methods

- Registration 

| Criteria      | Value                                                                                                                   |
| --------------| ------------------------------------------------------------------------------------------------------------------------------|
| Host 			| http://localhost:3000                               				      						  				            |
| Method		| POST   			   				            |
| URI		    | /registration			   				            |
| Header   	 	| 					            |
| Parameters   	|         		                |
| Body		    | { "username":"", "password":"", "role":"" }              		                |

- Login 

| Criteria      | Value                                                                                                                   |
| --------------| ------------------------------------------------------------------------------------------------------------------------------|
| Host 			| http://localhost:3000                               				      						  				            |
| Method		| POST  			   				            |
| URI		    | /login				   				            |
| Header   	 	| 					            |
| Parameters   	|         		                |
| Body		    | {	"username":"","password":"" }              		                |


#### Secured methods

- CreatePart

| Criteria      | Value                                                                                                                   |
| --------------| ------------------------------------------------------------------------------------------------------------------------------|
| Host 			| http://localhost:3000                               				      						  				            |
| Method		| POST   			   				            |
| URI		    | /logcard/parts			   				            |
| Header   	 	| 					            |
| Parameters   	|         		                |
| Body		    | { "pn":"", "sn":"", "partName":"", "type":"", "responsible":"", "helicopter":"", "assembly":""}              		                |

- Transfer Ownership 

| Criteria      | Value                                                                                                                   |
| --------------| ------------------------------------------------------------------------------------------------------------------------------|
| Host 			| http://localhost:3000                               				      						  				            |
| Method		| PUT   			   				            |
| URI		    | /parts/:id/:action				   				            |
| Header   	 	| 					            |
| Parameters   	|         		                |
| Body		    | { "owner":"" }             		                |

- Transfer Responsibility 

| Criteria      | Value                                                                                                                   |
| --------------| ------------------------------------------------------------------------------------------------------------------------------|
| Host 			| http://localhost:3000                               				      						  				            |
| Method		| PUT   			   				            |
| URI		    | /parts/:id/:action				   				            |
| Header   	 	| 					            |
| Parameters   	|         		                |
| Body		    | {	"responsible":"" }              		                |

- Perform Activities 

| Criteria      | Value                                                                                                                   |
| --------------| ------------------------------------------------------------------------------------------------------------------------------|
| Host 			| http://localhost:3000                               				      						  				            |
| Method		| PUT   			   				            |
| URI		    | /parts/:id/:action				   				            |
| Header   	 	| 					            |
| Parameters   	|         		                |
| Body		    | {	"modType":"", "description": "" }           		                |

- Display part Historic 

| Criteria      | Value                                                                                                                   |
| --------------| ------------------------------------------------------------------------------------------------------------------------------|
| Host 			| http://localhost:3000                               				      						  				            || Method		    			   				            |
| Method		| GET  			   				            |
| URI		    | /parts/:id/historic				   				            |
| Header   	 	| 					            |
| Parameters   	|         		                |
| Body		    |               		                |


- Display all parts 

| Criteria      | Value                                                                                                                   |
| --------------| ------------------------------------------------------------------------------------------------------------------------------|
| Host 			| http://localhost:3000                               				      						  				            |
| Method		| GET    			   				            |
| URI		    | /logcard/parts				   				            |
| Header   	 	| 					            |
| Parameters   	|         		                |
| Body		    |               		                |