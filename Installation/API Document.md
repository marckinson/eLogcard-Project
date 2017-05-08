# e-Logcard Application 

## Description of the APIs

#### Non-secured methods

- Registration 

| Criteria      | Value                                                                                                                   |
| --------------| ------------------------------------------------------------------------------------------------------------------------------|
| Host 			| http://localhost:3000                               				      						  				            |
| Method		|    			   				            |
| URI		    | 				   				            |
| Header   	 	| 					            |
| Parameters   	|         		                |
| Body		    | { "username":"", "password":"", "role":"" }              		                |

- Login 

| Criteria      | Value                                                                                                                   |
| --------------| ------------------------------------------------------------------------------------------------------------------------------|
| Host 			| http://localhost:3000                               				      						  				            |
| Method		|   			   				            |
| URI		    | 				   				            |
| Header   	 	| 					            |
| Parameters   	|         		                |
| Body		    | {	"username":"","password":"" }              		                |


#### Secured methods

- CreatePart

| Criteria      | Value                                                                                                                   |
| --------------| ------------------------------------------------------------------------------------------------------------------------------|
| Host 			| http://localhost:3000                               				      						  				            |
| Method		| POST   			   				            |
| URI		    | 				   				            |
| Header   	 	| 					            |
| Parameters   	|         		                |
| Body		    | { "pn":"", "sn":"", "partName":"", "type":"", "responsible":"", "helicopter":"", "assembly":""}              		                |

- Transfer Ownership 

| Criteria      | Value                                                                                                                   |
| --------------| ------------------------------------------------------------------------------------------------------------------------------|
| Host 			| http://localhost:3000                               				      						  				            |
| Method		| PUT   			   				            |
| URI		    | 				   				            |
| Header   	 	| 					            |
| Parameters   	|         		                |
| Body		    | { "owner":"" }             		                |

- Transfer Responsibility 

| Criteria      | Value                                                                                                                   |
| --------------| ------------------------------------------------------------------------------------------------------------------------------|
| Host 			| http://localhost:3000                               				      						  				            |
| Method		| PUT   			   				            |
| URI		    | 				   				            |
| Header   	 	| 					            |
| Parameters   	|         		                |
| Body		    | {	"responsible":"" }              		                |

- Perform Activities 

| Criteria      | Value                                                                                                                   |
| --------------| ------------------------------------------------------------------------------------------------------------------------------|
| Host 			| http://localhost:3000                               				      						  				            |
| Method		| PUT   			   				            |
| URI		    | 				   				            |
| Header   	 	| 					            |
| Parameters   	|         		                |
| Body		    | {	"modType":"", "description": "" }           		                |

- Display part Historic 

| Criteria      | Value                                                                                                                   |
| --------------| ------------------------------------------------------------------------------------------------------------------------------|
| Host 			| http://localhost:3000                               				      						  				            || Method		    			   				            |
| Method		| GET  			   				            |
| URI		    | 				   				            |
| Header   	 	| 					            |
| Parameters   	|         		                |
| Body		    |               		                |


- Display all parts 

| Criteria      | Value                                                                                                                   |
| --------------| ------------------------------------------------------------------------------------------------------------------------------|
| Host 			| http://localhost:3000                               				      						  				            |
| Method		| GET    			   				            |
| URI		    | 				   				            |
| Header   	 	| 					            |
| Parameters   	|         		                |
| Body		    |               		                |