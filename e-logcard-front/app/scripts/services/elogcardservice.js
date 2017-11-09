'use strict';

/**
 * @ngdoc service
 * @name eLogcardFrontApp.eLogcardService
 * @description
 * # eLogcardService
 * Service in the eLogcardFrontApp.
 */
app.service('eLogcardService', function ($http, $q, userService,$rootScope) {

        var self = this;
        this.debug = false;
        this.baseproxyUri = "/blockchain/";

        this.addAuthorizationHttp = function () {
            $http.defaults.headers.common.Authorization = 'Bearer ' + userService.getToken();

        };
        /*
                this.checkConnected = function () {
                    var defered = $q.defer();

                    var request = {
                        'answer': false,
                        'status': "403",
                        'message': "your are not connected"

                    };

                    if (userService.getState()) {
                        request.answer: true;
                        request.status: "200";
                        request.message: "your are connected ";

                        defered.reject(request);
                    } else {
                        defered.reject(request);
                    }

                    return defered.promise;
                };*/



        var factory = {
            //TEST
            gettest: function () {
                console.log("call service eLogcardService verifiez si bien execute ");
            },

            // USER
            subscribe: function (userName, password, role) {
                var defered = $q.defer();

                var request = {
                    'answer': false,
                    'status': ""
                };

                let registrationUri = self.baseproxyUri + "registration";

                if (self.debug) {
                    console.log(registrationUri);
                }

                var Userdata = {
                    "username": userName,
                    "password": password,
                    "role": role
                };

                $http.post(registrationUri, Userdata)
                    .then(
                        function (response) {
                            request.answer = response.data;
                            request.status = response.status;
                            // ajout authentification 
                            self.addAuthorizationHttp();

                            userService.setState(true);
                            userService.setToken(response.data);
                            userService.setUser(userName);
                            userService.setRole(role);


                            defered.resolve(request);
                        },
                        function (error) {
                            self.answer = error.data || 'Request failed';
                            self.status = error.status;
                            userService.clearValues;
                            defered.reject(request);
                        }
                    );

                return defered.promise;
            },
			findAllUsers:function(){
				 var defered = $q.defer();
				 $http.get(self.baseproxyUri +"users").then(function(response){
					defered.resolve(response.data);
				 
				 });
				 return defered.promise;
			},
            login: function (userName, password) {
                var defered = $q.defer();

                var request = {
                    'answer': false,
                    'status': ""
                };

                var header = {
                    "username": userName,
                    "password": password
                };
                if (self.debug) {
                    console.log("hearder");
                    console.log(header);
                }

                // construction requete 
                let loginUri = self.baseproxyUri + "login";

                if (self.debug) {
                    console.log(loginUri);
                    console.log(header);
                }

                $http.post(loginUri, header)
                    .then(
                        function (response) {
                            request.answer = response.data;
                            request.status = response.status;

                            userService.setState(true);
                            userService.setToken(response.data.token);
                            self.addAuthorizationHttp();
                            userService.setUser(response.data.username);
                            userService.setRole(response.data.role);

                            if (self.debug) {
                                console.log(response.data);
                                console.log(response.data.username);
                                console.log(response.data.token);
                                console.log(response.data.role);
                            }
							$rootScope.$emit('userLogin');
                            defered.resolve(request);

                        },
                        function (error) {

                            request.status = error.status;
                            request.aswer = error.data;
                            if (self.debug) {
                                console.log(error.status);
                                console.log(error.data);
                            }
                            // clear les information stocker dans le service user 
                            userService.clearValues;

                            defered.reject(request);
                        });
                return defered.promise;
            },
            //requete qui donne les roles disponible pour cree un nouvelle utilisateur 
            getUserRoles: function () {
                // console.log(" call  get role");
                var defered = $q.defer();

                var request = {
                    'roles': [],
                    'answer': "",
                    'stateRequest': true,
                    'status': ""
                };

                let rolesUri = self.baseproxyUri + "roles";

                $http.get(rolesUri)
                    .then(
                        function (response) {
                            request.roles = response.data;
                            if (self.debug) {
                                console.log(request);
                            }
                            defered.resolve(request);

                        },
                        function (error) {
                            request.answer = error.data || 'Request failed';
                            request.status = error.status;
                            request.stateRequest = false;
                            if (self.debug) {
                                console.log(request);
                            }
                            defered.reject(request);

                        }
                    );
                return defered.promise;
            },

            // PART
            //TODO
            //showpart
            //DO
            createPart: function (header) {

                var defered = $q.defer();
                // valeur de resultas de retour 
                var request = {
                    'answer': false,
                    'status': ""
                };

                let createUriParts = self.baseproxyUri + "logcard/parts";

                if (self.debug) {
                    console.log(createUriParts);
                    console.log(header);
                }

                factory.createItem(header, createUriParts)
                    .then(function (reponse) {
                        defered.resolve(reponse);
                    }, function (error) {
                        defered.reject(error);
                    });
				$rootScope.$emit('chain-modification');
                return defered.promise;
            },

            getOnePart: function (idPart) {
                var defered = $q.defer();

                var request = {
                    'part': [],
                    'aswer': "",
                    'stateRequest': true,
                    'status': ""
                };

                let showPartUri = self.baseproxyUri + "logcard/parts/" + idPart;

                if (self.debug) {
                    console.log(showPartUri);
                }
                if (userService.getState()) {
                    $http.get(showPartUri)
                        .then(
                            function (response) {
                                request.assemblies = response.data;
                                request.status = response.status;
                                defered.resolve(request);
                                if (self.debug) {
                                    console.log(response.data);

                                }
                            },
                            function (error) {
                                request.answer = error.data || 'Request failed';
                                request.stateRequest = false;
                                defered.reject(request);
                            }
                        );
                }

                return defered.promise;
            },
			
			getParts: function () {
                var defered = $q.defer();

                var request = {
                    'parts': [],
                    'aswer': "",
                    'stateRequest': true,
                    'status': ""
                };

                let showPartsUri = self.baseproxyUri + "logcard/parts";

                if (self.debug) {
                    console.log(showPartsUri);
                }
                if (userService.getState()) {
                    $http.get(showPartsUri)
                        .then(
                            function (response) {
                                request.assemblies = response.data;
                                request.status = response.status;
                                defered.resolve(request);
                                if (self.debug) {
                                    console.log(response.data);

                                }
                            },
                            function (error) {
                                request.answer = error.data || 'Request failed';
                                request.stateRequest = false;
                                defered.reject(request);
                            }
                        );
                }

                return defered.promise;
            },

            scrappPart: function (partId) {

                var defered = $q.defer();
                // construction requete 
                let scrapParttUri = self.baseproxyUri + "logcard/parts/scrapp";
                if (self.debug) {
                    console.log(scrapParttUri);
                }
                factory.scrappTarget(partId, scrapParttUri)
                    .then(function (reponse) {
                        defered.resolve(reponse);
                    }, function (error) {
                        defered.reject(error);
                    })

                return defered.promise;

            },

            transfertPartOwnership: function (ownerName, idPart) {

                var defered = $q.defer();
                // construction requete 
                let transfertOwnershipPartUri = self.baseproxyUri + "logcard/parts/OwnerTransfer/" + idPart;

                if (self.debug) {
                    console.log(transfertOwnershipPartUri);
                }
                factory.transfertTargetOwnerShip(ownerName, transfertOwnershipPartUri)
                    .then(function (reponse) {
                        defered.resolve(reponse);
                    }, function (error) {
                        defered.reject(error);
                    })
                return defered.promise;
            },

            transfertPartResponsible: function (responsibleName, idPart) {
                var defered = $q.defer();
                // valeur de resultas de retour 
                var request = {
                    'answer': false,
                    'status': ""
                };
                // construction requete 
                let transfertResponsiblePartUri = self.baseproxyUri + "logcard/parts/RespoTransfer/" + idPart;

                if (self.debug) {
                    console.log(transfertResponsiblePartUri);
                }
                // construction data 
                var data = {
                    "responsible": responsibleName
                };

                $http.put(transfertResponsiblePartUri, data)
                    .then(function (reponse) {
                        if (self.debug)
                            console.log(reponse);
                        request.answer = reponse.data;
                        request.status = reponse.status;
						$rootScope.$emit('chain-modification');
                        defered.resolve(request);

                    }, function (error) {
                        request.status = error.status;
                        defered.reject(request);
                    })
                return defered.promise;
            },

            getListPartWithoutAssembly: function () {
                if (self.debug)
                    console.log(" Call getListPartWithoutAssembly");
                var defered = $q.defer();

                let showListPartWithoutAssembly = self.baseproxyUri + "logcard/partsNoAssembly/";

                factory.getPartsWitoutTargetContainer(showListPartWithoutAssembly)
                    .then(function (reponse) {
                        defered.resolve(reponse);
                    }, function (error) {
                        defered.reject(error);
                    })


                return defered.promise;
            },

            getListPartWithoutAirCraft: function () {

                if (self.debug)
                    console.log(" Call getListPartWithoutAssembly");

                var defered = $q.defer();

                let showListPartWithoutAssembly = self.baseproxyUri + "logcard/partsNoAircraft/";

                factory.getPartsWitoutTargetContainer(showListPartWithoutAssembly)
                    .then(function (reponse) {
                        defered.resolve(reponse);
                    }, function (error) {
                        defered.reject(error);
                    })

                return defered.promise;
            },
            //LOGPART
            // perform Ativites
            addLogOnPart: function (idPart, modType, description) {

                var defered = $q.defer();
                // valeur de resultas de retour 
                var request = {
                    'answer': false,
                    'status': ""
                };
                // construction requete 
                let PerformActsUri = self.baseproxyUri + "logcard/parts/PerformActs/" + idPart;

                if (self.debug) {
                    console.log(PerformActsUri);
                }

                // construction header 
                var header = {
                    "modType": modType,
                    "description": description
                };
                if (userService.getState()) {
                    $http.put(PerformActsUri, header)
                        .then(function (reponse) {
                            if (self.debug) {
                                console.log(reponse);
                            }
                            request.answer = reponse.data;
                            request.status = reponse.status;
							$rootScope.$emit('chain-modification');
                            defered.resolve(request);

                        }, function (error) {
                            request.status = error.status;
                            defered.reject(request);
                        })
                } else {
                    request.answer = "your are not connected ";
                    request.status = 403;
                    defered.reject(request);
                }

                return defered.promise;

            },
            // getlogType
            getListModificationType: function () {
                if (self.debug)
                    console.log(" CALL getListModificationType");
                var defered = $q.defer();

                var request = {
                    'ModificationTypes': [],
                    'aswer': "",
                    'stateRequest': true,
                    'status': ""
                };

                //self.addAuthorizationHttp();
                //contruction url rest 
                let showListModificationTypeUri = self.baseproxyUri + "logcard/List/modifications";

                if (userService.getState()) {
                    $http.get(showListModificationTypeUri)
                        .then(
                            function (response) {
                                request.ModificationTypes = response.data;
                                request.status = response.status;
                                defered.resolve(request);
                                if (self.debug) {
                                    console.log(response.data);
                                }
                            },
                            function (error) {
                                request.answer = error.data || 'Request failed';
                                request.stateRequest = false;
                                defered.reject(request);
                            }
                        );
                }
                return defered.promise;
            },

            //AIRCRAFT
            //TODO
            // Show Aircrafts
            //DO  

            // create new Aircraft 
            createAircraft: function (sn, an, name) {
                var defered = $q.defer();
                // valeur de resultas de retour 
                var request = {
                    'answer': false,
                    'status': ""
                };

                let createUriAirCraft = self.baseproxyUri + "logcard/aircrafts";
                // construction header 
                var header = {
                    "an": an,
                    "sn": sn,
                    "componentName": name
                };

                if (self.debug) {
                    console.log(createUriAirCraft);
                    console.log(header);
                }


                factory.createItem(header, createUriAirCraft)
                    .then(function (reponse) {
						$rootScope.$emit('chain-modification');
                        defered.resolve(reponse);
                    }, function (error) {
                        defered.reject(error);
                    })

                return defered.promise;


            },
            // add assemblies on aicraft
            addAssemblyOnAirCraft: function (aircraftId, assembyId) {
                var defered = $q.defer();
                // valeur de resultas de retour 
                var request = {
                    'answer': false,
                    'status': ""
                };
                // construction requete 
                let AddAssemblyOnAircraftUri = self.baseproxyUri + "logcard/aircrafts/add/assembly/" + aircraftId;

                if (self.debug)
                    console.log(AddAssemblyOnAircraftUri);

                // construction header 
                var header = {
                    "idassembly": assembyId
                };

                $http.put(AddAssemblyOnAircraftUri, header)
                    .then(function (reponse) {
                        if (self.debug)
                            console.log(reponse);
                        request.answer = reponse.data;
                        request.status = reponse.status;
						$rootScope.$emit('chain-modification');
                        defered.resolve(request);

                    }, function (error) {
                        request.status = error.status;
                        defered.reject(request);
                    })

                return defered.promise;

            },
            //add part on Aircrafts
            addPartOnAirCraft: function (idAircraft, idPart) {
                var defered = $q.defer();

                // construction requete 
                let addPartOnAirCraftUri = self.baseproxyUri + "logcard/aircrafts/add/part/" + idAircraft;

                if (self.debug)
                    console.log(addPartOnAirCraftUri);

                factory.addPartToContainer(idPart, addPartOnAirCraftUri)
                    .then(function (reponse) {
						$rootScope.$emit('chain-modification');
                        defered.resolve(reponse);
                    }, function (error) {
                        defered.reject(error);
                    })

                return defered.promise;

            },
            //
            getAircraftHistoric: function (idAirCraft) {
                var defered = $q.defer();

                var request = {
                    'aircraft': [],
                    'aswer': "",
                    'stateRequest': true,
                    'status': ""
                };

                //self.addAuthorizationHttp();

                let getAirCraftHistoryUri = self.baseproxyUri + "logcard/aircrafts/historic/" + idAirCraft;

                if (self.debug)
                    console.log(getAirCraftHistoryUri);

                if (userService.getState()) {
                    $http.get(getAirCraftHistoryUri)
                        .then(
                            function (response) {
                                request.aircraft = response.data;
                                request.status = response.status;
                                defered.resolve(request);
                                if (self.debug) {
                                    console.log(response.data);
                                }
                            },
                            function (error) {
                                request.answer = error.data || 'Request failed';
                                request.stateRequest = false;
                                defered.reject(request);
                            }
                        );
                }
                return defered.promise;
            },
            // recupere la liste des parts sur un aircraft
            getAircraftlistParts: function (idAircraft) {
                var defered = $q.defer();
                // construction requete 

                let getAircraftPartsListUri = self.baseproxyUri + "logcard/aircrafts/listing/parts/" + idAircraft;

                if (self.debug)
                    console.log(getAircraftPartsListUri);

                factory.getAircraftList(getAircraftPartsListUri)
                    .then(function (reponse) {
                        defered.resolve(reponse);
                    }, function (error) {
                        defered.reject(error);
                    })
                return defered.promise;
            },
            // recupere la liste des assemblies sur un aircraft 
            getAirCraftListAssemby: function (idAircraft) {
                var defered = $q.defer();
                // construction requete 

                let getAircraftAssembliesListUri = self.baseproxyUri + "logcard/aircrafts/listing/assemblies/" + idAircraft;

                if (self.debug)
                    console.log(getAircraftAssembliesListUri);

                factory.getAircraftList(getAircraftAssembliesListUri)
                    .then(function (reponse) {
                        defered.resolve(reponse);
                    }, function (error) {
                        defered.reject(error);
                    })
                return defered.promise;
            },
            // getaircraftlist generique 
            getAircraftList: function (restRequest) {
                var defered = $q.defer();
                // constrution de l'object de reponse 
                var request = {
                    'list': [],
                    'aswer': "",
                    'stateRequest': true,
                    'status': ""
                };

                //self.addAuthorizationHttp();

                //http request 

                $http.get(restRequest)
                    .then(
                        function (response) {
                            request.list = response.data;
                            request.status = response.status;
                            defered.resolve(request);
                            if (self.debug) {
                                console.log(response.data);

                            }
                        },
                        function (error) {
                            request.answer = error.data || 'Request failed';
                            request.stateRequest = false;
                            defered.reject(request);
                        }
                    );

                return defered.promise;

            },
            // remove part to aircraft
            removePartToAirCraft: function (idAirCraft, idPart) {

                var defered = $q.defer();
                // construction requete 
                let removePartOnAircraftUri = self.baseproxyUri + "logcard/aircrafts/remove/parts/" + idAirCraft;
                if (self.debug)
                    console.log(removePartOnAircraftUri);

                factory.removePartToContainer(idPart, removePartOnAircraftUri)
                    .then(function (reponse) {
                        defered.resolve(reponse);
                    }, function (error) {
                        defered.reject(error);
                    })

                return defered.promise;

            },
            //remove Assembly to aicraft
            removeAssemblyToAicraft: function (idAirCraft, idAssembly) {
                console.log(" call removeAssemblyToAicraft");
                var defered = $q.defer();
                // construction requete 
                let removeAssemblyOnAircraftUri = self.baseproxyUri + "logcard/aircrafts/remove/assembly/" + idAirCraft;
                if (self.debug)
                    console.log(removeAssemblyOnAircraftUri);

                factory.removeAssemblyToContainer(idAssembly, removeAssemblyOnAircraftUri)
                    .then(function (reponse) {
                        defered.resolve(reponse);
                    }, function (error) {
                        defered.reject(error);
                    });

                return defered.promise;

            },
            //replace Assembly On aircraft 
            replaceAssemblyOnAircraft: function (idOldAssembly, idNewAssembly, AirCraftId) {
                var defered = $q.defer();
                // construction requete 
                let replaceAssemblyOnAircraftUri = self.baseproxyUri + "logcard/aircrafts/replace/assembly/" + AirCraftId;


                if (self.debug)
                    console.log(replaceAssemblyOnAircraftUri);

                factory.replaceAssemblyOnContainer(idOldAssembly, idNewAssembly, replaceAssemblyOnAircraftUri)
                    .then(function (reponse) {
                        defered.resolve(reponse);
                    }, function (error) {
                        defered.reject(error);
                    });
                return defered.promise;


            },
            //replace Part On aircraft
            replacePartOnAircraft: function (idOldPart, idNewPart, AirCraftId) {
                var defered = $q.defer();
                // construction requete 
                let replacePartOnAircraftUri = self.baseproxyUri + "logcard/aircrafts/replace/part/" + AirCraftId;

                if (self.debug)
                    console.log(replacePartOnAircraftUri);

                factory.replacePartOnContainer(idOldPart, idNewPart, replacePartOnAircraftUri)
                    .then(function (reponse) {
                        defered.resolve(reponse);
                    }, function (error) {
                        defered.reject(error);
                    });
                return defered.promise;


            },
            // scrapp aircraft
            scrappAirCraft: function (AirCraftId) {

                var defered = $q.defer();
                // construction requete 
                let scrappAircraftUri = self.baseproxyUri + "logcard/aircrafts/scrapp";
                if (self.debug)
                    console.log(scrappAircraftUri);

                factory.scrappTarget(AirCraftId, scrappAircraftUri)
                    .then(function (reponse) {
                        defered.resolve(reponse);
                    }, function (error) {
                        defered.reject(error);
                    });

                return defered.promise;

            },
            // transfert ownerShip
            transfertAirCraftOwnerShip: function (ownerName, idAircraft) {
                var defered = $q.defer();
                // construction requete 
                let transfertOwnershipAircraftUri = self.baseproxyUri + "logcard/aircrafts/transfer/" + idAircraft;

                if (self.debug)
                    console.log(transfertOwnershipAircraftUri);

                factory.transfertTargetOwnerShip(ownerName, transfertOwnershipAircraftUri)
                    .then(function (reponse) {
                        defered.resolve(reponse);
                    }, function (error) {
                        defered.reject(error);
                    });
                return defered.promise;

            },
            //
            transfertAirCraftResponsible: function (responsibleName, idAirCraft) {
                var defered = $q.defer();
                // construction requete 
                let transfertResposiblepAirCraftUri = self.baseproxyUri + "logcard/AirCrafts/transferRespo/" + idAirCraft;

                if (self.debug)
                    console.log(transfertResposiblepAirCraftUri);

                factory.transfertTargetResponsible(responsibleName, transfertResposiblepAirCraftUri)
                    .then(function (reponse) {
                        defered.resolve(reponse);
                    }, function (error) {
                        defered.reject(error);
                    });

                return defered.promise;

            },


            //ASSEMBLY
            //TODO
            //DO  

            // add part on assemblies
            addPartOnAssembly: function (idAssembly, idPart) {

                var defered = $q.defer();

                // construction requete 
                let addPartOnAssemblyUri = self.baseproxyUri + "logcard/assemblies/add/" + idAssembly;

                if (self.debug)
                    console.log(addPartOnAssemblyUri);
                factory.addPartToContainer(idPart, addPartOnAssemblyUri)
                    .then(function (reponse) {
                        defered.resolve(reponse);
                    }, function (error) {
                        defered.reject(error);
                    })
                return defered.promise;
            },
            createAssembly: function (sn, an, name) {

                var defered = $q.defer();
                // valeur de resultas de retour 
                var request = {
                    'answer': false,
                    'status': ""
                };

                let createUriAssembly = self.baseproxyUri + "logcard/assemblies";
                // construction header 
                var header = {
                    "an": an,
                    "sn": sn,
                    "componentName": name
                };

                if (self.debug) {
                    console.log(createUriAssembly);
                    console.log(header);
                }

                factory.createItem(header, createUriAssembly)
                    .then(function (reponse) {
                        defered.resolve(reponse);
                    }, function (error) {
                        defered.reject(error);
                    });

                return defered.promise;

            },
            // show ASSEMBLIES
            getAssemblies: function () {
                var defered = $q.defer();

                var request = {
                    'assemblies': [],
                    'aswer': "",
                    'stateRequest': true,
                    'status': ""
                };

                //self.addAuthorizationHttp();
                let showAssembliesUri = self.baseproxyUri + "logcard/assemblies";
                if (self.debug)
                    console.log(showAssembliesUri);

                if (userService.getState()) {
                    $http.get(showAssembliesUri)
                        .then(
                            function (response) {
                                request.assemblies = response.data;
                                request.status = response.status;
                                defered.resolve(request);
                                if (self.debug) {
                                    console.log(response.data);

                                }
                            },
                            function (error) {
                                request.answer = error.data || 'Request failed';
                                request.stateRequest = false;
                                defered.reject(request);
                            }
                        );
                }

                return defered.promise;
            },
            // remove part on assembly
            removePartToAssembly: function (idAssembly, idPart) {
                var defered = $q.defer();
                // construction requete 
                let removePartOnAssemblyUri = self.baseproxyUri + "logcard/assemblies/remove/" + idAssembly;
                if (self.debug) {
                    console.log(removePartOnAssemblyUri);
                }

                factory.removePartToContainer(idPart, removePartOnAssemblyUri)
                    .then(function (reponse) {
                        defered.resolve(reponse);
                    }, function (error) {
                        defered.reject(error);
                    });

                return defered.promise;
            },
            //remplace Part On assembly
            replacePartOnAssembly: function (idOldPart, idNewPart, AssemblyId) {
                var defered = $q.defer();
                // construction requete 
                let replacePartOnAssemblyUri = self.baseproxyUri + "logcard/assemblies/replace/" + AssemblyId;


                if (self.debug) {
                    console.log(replacePartOnAssemblyUri);
                }

                factory.replacePartOnContainer(idOldPart, idNewPart, replacePartOnAssemblyUri)
                    .then(function (reponse) {
                        defered.resolve(reponse);
                    }, function (error) {
                        defered.reject(error);
                    });
                return defered.promise;
            },
            // scrapp assembly
            scrappAssembly: function (idAssembly) {

                var defered = $q.defer();
                // construction requete 
                let scrappAssemblytUri = self.baseproxyUri + "logcard/assemblies/scrapp";
                if (self.debug) {
                    console.log(scrappAssemblytUri);
                }

                factory.scrappTarget(idAssembly, scrappAssemblytUri)
                    .then(function (reponse) {
                        defered.resolve(reponse);
                    }, function (error) {
                        defered.reject(error);
                    });

                return defered.promise;

            },
            // transfert Ownership
            transfertAssemblyOwnerShip: function (ownerName, idAssembly) {
                var defered = $q.defer();
                // construction requete 
                let transfertOwnershipAssemblyUri = self.baseproxyUri + "logcard/assemblies/transfer/" + idAssembly;

                if (self.debug) {
                    console.log(transfertOwnershipAssemblyUri);
                }

                factory.transfertTargetOwnerShip(ownerName, transfertOwnershipAssemblyUri)
                    .then(function (reponse) {
                        defered.resolve(reponse);
                    }, function (error) {
                        defered.reject(error);
                    });

                return defered.promise;
            },
            //
            transfertAssemblyResponsible: function (responsibleName, idAssembly) {
                var defered = $q.defer();
                // construction requete 
                let transfertResposiblepAssemblyUri = self.baseproxyUri + "logcard/assemblies/transferRespo/" + idAssembly;

                if (self.debug) {
                    console.log(transfertResposiblepAssemblyUri);
                }

                factory.transfertTargetResponsible(responsibleName, transfertResposiblepAssemblyUri)
                    .then(function (reponse) {
                        defered.resolve(reponse);
                    }, function (error) {
                        defered.reject(error);
                    });

                return defered.promise;


            },
            //
            getListAssemblyWithoutAircraft: function () {
                if (self.debug) {
                    console.log(" CALL getListAssemblyWithoutAircraft");
                }
                var defered = $q.defer();

                var request = {
                    'assemblies': [],
                    'aswer': "",
                    'stateRequest': true,
                    'status': ""
                };

                //self.addAuthorizationHttp();
                //contruction url rest 
                let showListAssemblyWithoutAicraft = self.baseproxyUri + "logcard/assembliesNoAircraft";

                if (userService.getState()) {
                    $http.get(showListAssemblyWithoutAicraft)
                        .then(
                            function (response) {
                                request.assemblies = response.data;
                                request.status = response.status;
                                defered.resolve(request);
                                if (self.debug) {
                                    console.log(response.data);
                                }
                            },
                            function (error) {
                                request.answer = error.data || 'Request failed';
                                request.stateRequest = false;
                                defered.reject(request);
                            }
                        );
                }
                return defered.promise;
            },

            // ALL
            createItem: function (header, restRequest) {

                var defered = $q.defer();

                var request = {
                    'answer': false,
                    'status': ""
                };
                // test si l'utilisateur est connecte 
                if (userService.getState()) {
                    $http.post(restRequest, header)
                        .then(function (reponse) {
                                if (self.debug) {
                                    console.log(reponse);
                                }
                                request.answer = reponse.data;
                                request.status = reponse.status;
								$rootScope.$emit('chain-modification');
                                defered.resolve(request);

                            },
                            function (error) {

                                request.status = error.status;
                                request.answer = error.data;

                                if (self.debug) {
                                    console.log(error.status);
                                }

                                defered.reject(request);
                            });
                } else {
                    request.answer = "your are not connected ";
                    request.status = 403;
                    defered.reject(request);
                }

                return defered.promise;

            },
            // pemert d Ajouter un iteme a un autre essemble 
            addItemToContainer: function (itemData, restRequest) {
                var defered = $q.defer();
                // valeur de resultas de retour 
                var request = {
                    'answer': false,
                    'status': ""
                };
                // ajoute token autorisation 
                //self.addAuthorizationHttp();

                $http.put(restRequest, itemData)
                    .then(function (reponse) {
                        if (self.debug) {
                            console.log(reponse);
                        }
                        request.answer = reponse.data;
                        request.status = reponse.status;
						$rootScope.$emit('chain-modification');
                        defered.resolve(request);

                    }, function (error) {
                        request.status = error.status;
                        defered.reject(request);
                    });

                return defered.promise;

            },
            // permet de faire les ajout de part dans les autre ensemble 
            addPartToContainer: function (idTarget, restRequest) {
                var defered = $q.defer();

                // construction data 
                var Partdata = {
                    "idpart": idTarget
                };

                factory.addItemToContainer(Partdata, restRequest)
                    .then(function (reponse) {
                        defered.resolve(reponse);
                    }, function (error) {
                        defered.reject(error);
                    });

                return defered.promise;
            },
            //permet de faire les ajout de assembly dans les autre ensemble 
            addAssemblyToContainer: function (idTarget, restRequest) {
                var defered = $q.defer();

                // construction data 
                var Assemblydata = {
                    "idassembly": idTarget
                };

                factory.addItemToContainer(Assemblydata, restRequest)
                    .then(function (reponse) {
                        defered.resolve(reponse);
                    }, function (error) {
                        defered.reject(error);
                    });

                return defered.promise;
            },
            // scrapp function 
            scrappTarget: function (idTarget, restRequest) {
                var request = {
                    'answer': false,
                    'status': ""
                };
                var header = {
                    "id": idTarget
                };

                var defered = $q.defer();

                $http.put(restRequest, header)
                    .then(function (reponse) {
                        if (self.debug) {
                            console.log(reponse);
                        }
                        request.answer = reponse.data;
                        request.status = reponse.status;
						$rootScope.$emit('chain-modification');
                        defered.resolve(request);

                    }, function (error) {
                        request.status = error.status;
                        defered.reject(request);
                    });

                return defered.promise;

            },
            // remove item to container
            removeItemToContainer: function (data, restRequest) {
                var defered = $q.defer();
                // valeur de resultas de retour 
                var request = {
                    'answer': false,
                    'status': ""
                };

                $http.put(restRequest, data)
                    .then(function (reponse) {
                        if (self.debug) {
                            console.log(reponse);
                        }
                        request.answer = reponse.data;
                        request.status = reponse.status;
						$rootScope.$emit('chain-modification');
                        defered.resolve(request);

                    }, function (error) {
                        request.status = error.status;
                        defered.reject(request);
                    });

                return defered.promise;
            },

            removeAssemblyToContainer: function (idAssembly, restRequest) {
                console.log(" call removeAssemblyToContainer");
                var defered = $q.defer();

                // construction data 
                var dataAssembly = {
                    "idassembly": idAssembly
                };

                factory.removeItemToContainer(dataAssembly, restRequest)
                    .then(function (reponse) {
                        defered.resolve(reponse);

                    }, function (error) {
                        defered.reject(error);
                    });

                return defered.promise;

            },
            // remove Part
            removePartToContainer: function (idPart, restRequest) {
                if (self.debug) {
                    console.log(" call removePartToContainer");
                }
                var defered = $q.defer();

                // construction data 
                var dataPart = {
                    "idpart": idPart
                };

                factory.removeItemToContainer(dataPart, restRequest)
                    .then(function (reponse) {
                        defered.resolve(reponse);

                    }, function (error) {
                        defered.reject(error);
                    });

                return defered.promise;

            },
            // transferOwnerShip
            transfertTargetOwnerShip: function (ownerName, restRequest) {
                if (self.debug) {
                    console.log("transfertTargetOwnerShip ");
                    console.log("ownerName ");
                    console.log(ownerName);
                    console.log("restRequest ");
                    console.log(restRequest);
                }
                var defered = $q.defer();
                // valeur de resultas de retour 
                var request = {
                    'answer': false,
                    'status': ""
                };
                // construction data 
                var data = {
                    "owner": ownerName
                };

                $http.put(restRequest, data)
                    .then(function (reponse) {
                        if (self.debug) {
                            console.log(reponse);
                        }
                        request.answer = reponse.data;
                        request.status = reponse.status;
						$rootScope.$emit('chain-modification');
                        defered.resolve(request);

                    }, function (error) {
                        request.status = error.status;
                        defered.reject(request);
                    });
                return defered.promise;
            },
            //
            transfertTargetResponsible: function (responsibleName, restRequest) {
                if (self.debug) {
                    console.log("transfertTargetResponsible ");
                    console.log("responsibleName ");
                    console.log(responsibleName);
                    console.log("restRequest ");
                    console.log(restRequest);
                }
                var defered = $q.defer();
                // valeur de resultas de retour 
                var request = {
                    'answer': false,
                    'status': ""
                };
                // construction data 
                var data = {
                    "responsible": responsibleName
                };

                $http.put(restRequest, data)
                    .then(function (reponse) {
                        if (self.debug)
                            console.log(reponse);
                        request.answer = reponse.data;
                        request.status = reponse.status;
						$rootScope.$emit('chain-modification');
                        defered.resolve(request);

                    }, function (error) {
                        request.status = error.status;
                        defered.reject(request);
                    });

                return defered.promise;
            },
            // replace part on container (aircraft,assembly)
            replacePartOnContainer: function (idOldPart, idNewPart, restRequest) {

                var defered = $q.defer();
                // valeur de resultas de retour 
                var request = {
                    'answer': false,
                    'status': ""
                };
                // construction data 
                var data = {
                    "idpart": idOldPart,
                    "idpart1": idNewPart
                };

                $http.put(restRequest, data)
                    .then(function (reponse) {
                        if (self.debug) {
                            console.log(reponse);
                        }
                        request.answer = reponse.data;
                        request.status = reponse.status;
						$rootScope.$emit('chain-modification');
                        defered.resolve(request);

                    }, function (error) {
                        request.status = error.status;
                        defered.reject(request);
                    });

                return defered.promise;

            },
            // replace assembly on container (aircraft,assembly)
            replaceAssemblyOnContainer: function (idOldAssembly, idNewAssembly, restRequest) {
                var defered = $q.defer();

                // valeur de resultas de retour 
                var request = {
                    'answer': false,
                    'status': ""
                };
                // ajoute token autorisation 
                //self.addAuthorizationHttp();
                // construction data 
                var data = {
                    "idassembly": idOldAssembly,
                    "idassembly1": idNewAssembly
                };

                $http.put(restRequest, data)
                    .then(function (reponse) {
                        if (self.debug) {
                            console.log(reponse);
                        }
                        request.answer = reponse.data;
                        request.status = reponse.status;
						$rootScope.$emit('chain-modification');
                        defered.resolve(request);

                    }, function (error) {
                        request.status = error.status;
                        defered.reject(request);
                    });

                return defered.promise;

            },
            getPartsWitoutTargetContainer: function (restRequest) {
                if (self.debug)
                    console.log(" Call getPartsWitoutTargetContainer");

                var defered = $q.defer();

                var request = {
                    'parts': [],
                    'aswer': "",
                    'stateRequest': true,
                    'status': ""
                };

                if (userService.getState()) {
                    $http.get(restRequest)
                        .then(
                            function (response) {
                                request.parts = response.data;
                                request.status = response.status;
								$rootScope.$emit('chain-modification');
                                defered.resolve(request);
                                if (self.debug) {
                                    console.log(response.data);
                                }
                            },
                            function (error) {
                                request.answer = error.data || 'Request failed';
                                request.stateRequest = false;
                                request.status = error.status;
                                defered.reject(request);
                            }
                        );
                }
                return defered.promise;
            }
        }
        return factory;
    }

);
