'use strict';

/**
 * @ngdoc service
 * @name eLogcardFrontApp.eLogcardService
 * @description
 * # eLogcardService
 * Service in the eLogcardFrontApp.
 */
app.service('eLogcardService', function ($http, $q, userService) {

        var self = this;
        this.debug = false;
        this.baseproxyUri = "/blockchain/";

        this.addAuthorizationHttp = function () {
            $http.defaults.headers.common.Authorization = 'Bearer ' + userService.getToken();

        };

        var factory = {
            //TEST
            gettest: function () {
                console.log("call service eLogcardService verifiez si bien execute ");
            },

            // USER
            //TODO
            //login
            //subscribe
            //DO
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
                            //console.log(request);
                            defered.resolve(request);

                        },
                        function (error) {
                            request.answer = error.data || 'Request failed';
                            request.status = error.status;
                            request.stateRequest = false;
                            defered.reject(request);
                            // console.log(request);
                        }
                    );
                return defered.promise;
            },

            // PART
            //TODO
            //showpart
            //DO
            // scrapp part 
            scrappPart: function (partId) {

                var defered = $q.defer();
                // construction requete 
                let scrapParttUri = self.baseproxyUri + "logcard/parts/scrapp";
                if (self.debug)
                    console.log(scrapParttUri);

                factory.scrappTarget(partId, scrapParttUri)
                    .then(function (reponse) {
                        defered.resolve(reponse);
                    }, function (error) {
                        defered.reject(error);
                    })

                return defered.promise;

            },
            // tranfert part ownership
            transfertPartOwnership: function (ownerName, idPart) {

                var defered = $q.defer();
                // construction requete 
                let transfertOwnershipPartUri = self.baseproxyUri + "logcard/parts/OwnerTransfer/" + idPart;

                if (self.debug)
                    console.log(transfertOwnershipPartUri);

                factory.transfertTargetOwnerShip(ownerName, transfertOwnershipPartUri)
                    .then(function (reponse) {
                        defered.resolve(reponse);
                    }, function (error) {
                        defered.reject(error);
                    })
                return defered.promise;
            },
            // transfert part resposible
            transfertPartResponsible: function (responsibleName, idPart) {
                var defered = $q.defer();
                // valeur de resultas de retour 
                var request = {
                    'answer': false,
                    'status': ""
                };
                // construction requete 
                let transfertResponsiblePartUri = self.baseproxyUri + "logcard/parts/RespoTransfer/" + idPart;

                if (self.debug)
                    console.log(transfertResponsiblePartUri);

                // ajoute token autorisation 
                self.addAuthorizationHttp();

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
                        defered.resolve(request);

                    }, function (error) {
                        request.status = error.status;
                        defered.reject(request);
                    })

                return defered.promise;

            },

            //AIRCRAFT
            //TODO
            // Show Aircrafts
            //DO  
            //add part on Aircrafts
            addPartOnAirCraft: function (idAircraft, idPart) {
                var defered = $q.defer();

                // construction requete 
                let addPartOnAirCraftUri = self.baseproxyUri + "logcard/aircrafts/add/part/" + idAircraft;

                if (self.debug)
                    console.log(addPartOnAirCraftUri);

                factory.addPartToSource(idAircraft, idPart, addPartOnAirCraftUri)
                    .then(function (reponse) {
                        defered.resolve(reponse);
                    }, function (error) {
                        defered.reject(error);
                    })

                return defered.promise;

            },
            // remove assembly to aircraft
            removePartToAirCraft: function (idAirCraft, idPart) {
                //   /logcard/aircrafts/remove/assembly

                // remove part form aircraft
                //   /logcard/aircrafts/transfer/ec5101c0-4ed2-11e7-ae79-c5bbeb60c360
                // data :owner

                var defered = $q.defer();
                // construction requete 
                let removePartOnAircraftUri = self.baseproxyUri + "logcard/aircrafts/remove/" + idAirCraft;
                if (self.debug)
                    console.log(removePartOnAircraftUri);

                factory.removePartToSource(idAirCraft, idPart, removePartOnAircraftUri)
                    .then(function (reponse) {
                        defered.resolve(reponse);
                    }, function (error) {
                        defered.reject(error);
                    })

                return defered.promise;

            },
            //replace Assembly On aircraft 
            replaceAssemblyOnAircraft: function (idOldAssembly, idNewAssembly, AirCraftId) {
                var defered = $q.defer();
                // construction requete 
                let replaceAssemblyOnAircraftUri = self.baseproxyUri + "logcard/aircrafts/replace/assembly/" + idAircraft;

                if (self.debug)
                    console.log(replaceAssemblyOnAircraftUri);

                factory.replacePartOnContainer(idOldAssembly, idNewAssembly, replaceAssemblyOnAircraftUri)
                    .then(function (reponse) {
                        defered.resolve(reponse);
                    }, function (error) {
                        defered.reject(error);
                    })
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
                    })
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
                    })

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
                    })
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
                factory.addPartToSource(idAssembly, idPart, addPartOnAssemblyUri)
                    .then(function (reponse) {
                        defered.resolve(reponse);
                    }, function (error) {
                        defered.reject(error);
                    })
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

                self.addAuthorizationHttp();
                let showPartsUri = self.baseproxyUri + "logcard/assemblies";

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
            // remove part on assembly
            removePartToAssembly: function (idAssembly, idPart) {
                // remove part to assembly
                // /logcard/assemblies/remove/7e082d30-4b88-11e7-bce3-ed8077462ef1
                var defered = $q.defer();
                // construction requete 
                let removePartOnAssemblyUri = self.baseproxyUri + "logcard/assemblies/remove/" + idAssembly;
                if (self.debug)
                    console.log(removePartOnAssemblyUri);

                factory.removePartToSource(idAssembly, idPart, removePartOnAssemblyUri)
                    .then(function (reponse) {
                        defered.resolve(reponse);
                    }, function (error) {
                        defered.reject(error);
                    })

                return defered.promise;
            },
            //remplace Part On assembly
            replacePartOnAssembly: function (idOldPart, idNewPart, AssemblyId) {
                var defered = $q.defer();
                // construction requete 
                let replacePartOnAssemblyUri = self.baseproxyUri + "logcard/assemblies/replace/" + AssemblyId;


                if (self.debug)
                    console.log(replacePartOnAssemblyUri);

                factory.replacePartOnContainer(idOldPart, idNewPart, replacePartOnAssemblyUri)
                    .then(function (reponse) {
                        defered.resolve(reponse);
                    }, function (error) {
                        defered.reject(error);
                    })
                return defered.promise;
            },
            // scrapp assembly
            scrappAssembly: function (idAssembly) {

                var defered = $q.defer();
                // construction requete 
                let scrappAssemblytUri = self.baseproxyUri + "logcard/assemblies/scrapp";
                if (self.debug)
                    console.log(scrappAssemblytUri);

                factory.scrappTarget(idAssembly, scrappAssemblytUri)
                    .then(function (reponse) {
                        defered.resolve(reponse);
                    }, function (error) {
                        defered.reject(error);
                    })

                return defered.promise;

            },
            // transfert Ownership
            transfertAssemblyOwnerShip: function (ownerName, idAssembly) {
                var defered = $q.defer();
                // construction requete 
                let transfertOwnershipAssemblyUri = self.baseproxyUri + "logcard/assemblies/transfer/" + idAssembly;

                if (self.debug)
                    console.log(transfertOwnershipAssemblyUri);

                factory.transfertTargetOwnerShip(ownerName, transfertOwnershipAssemblyUri)
                    .then(function (reponse) {
                        defered.resolve(reponse);
                    }, function (error) {
                        defered.reject(error);
                    })

                return defered.promise;


            },

            // ALL
            // permet de faire les ajout de parts dans les autre ensemble 
            addPartToSource: function (idSource, idTarget, restRequest) {
                var defered = $q.defer();
                // valeur de resultas de retour 
                var request = {
                    'answer': false,
                    'status': ""
                };
                // ajoute token autorisation 
                self.addAuthorizationHttp();
                // construction data 
                var data = {
                    "idpart": idTarget
                };

                $http.put(restRequest, data)
                    .then(function (reponse) {
                        if (self.debug)
                            console.log(reponse);
                        request.answer = reponse.data;
                        request.status = reponse.status;
                        defered.resolve(request);

                    }, function (error) {
                        request.status = error.status;
                        defered.reject(request);
                    })

                return defered.promise;
            },
            // scrapp function 
            scrappTarget: function (idTarget, restRequest) {
                var request = {
                    'answer': false,
                    'status': ""
                };
                var data = {
                    "id": idTarget
                };


                var defered = $q.defer();
                // construction requete 

                // ajoute token autorisation 
                self.addAuthorizationHttp();
                // requette http put 
                $http.put(restRequest, data)
                    .then(function (reponse) {
                        if (self.debug)
                            console.log(reponse);
                        request.answer = reponse.data;
                        request.status = reponse.status;
                        defered.resolve(request);

                    }, function (error) {
                        request.status = error.status;
                        defered.reject(request);
                    })

                return defered.promise;

            },
            // remove Part
            removePartToSource: function (idSource, idPart, restRequest) {
                var defered = $q.defer();
                // valeur de resultas de retour 
                var request = {
                    'answer': false,
                    'status': ""
                };
                // ajoute token autorisation 
                self.addAuthorizationHttp();
                // construction data 
                var data = {
                    "idpart": idPart
                };

                $http.put(restRequest, data)
                    .then(function (reponse) {
                        if (self.debug)
                            console.log(reponse);
                        request.answer = reponse.data;
                        request.status = reponse.status;
                        defered.resolve(request);

                    }, function (error) {
                        request.status = error.status;
                        defered.reject(request);
                    })

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
                // ajoute token autorisation 
                self.addAuthorizationHttp();
                // construction data 
                var data = {
                    "owner": ownerName
                };

                $http.put(restRequest, data)
                    .then(function (reponse) {
                        if (self.debug)
                            console.log(reponse);
                        request.answer = reponse.data;
                        request.status = reponse.status;
                        defered.resolve(request);

                    }, function (error) {
                        request.status = error.status;
                        defered.reject(request);
                    })

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
                // ajoute token autorisation 
                self.addAuthorizationHttp();
                // construction data 
                var data = {
                    "idpart": idOldPart,
                    "idpart1": idNewPart
                };
                $http.put(restRequest, data)
                    .then(function (reponse) {
                        if (self.debug)
                            console.log(reponse);
                        request.answer = reponse.data;
                        request.status = reponse.status;
                        defered.resolve(request);

                    }, function (error) {
                        request.status = error.status;
                        defered.reject(request);
                    })

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
                self.addAuthorizationHttp();
                // construction data 
                var data = {
                    "idassembly": idOldAssembly,
                    "idassembly1": idNewAssembly
                };

                $http.put(restRequest, data)
                    .then(function (reponse) {
                        if (self.debug)
                            console.log(reponse);
                        request.answer = reponse.data;
                        request.status = reponse.status;
                        defered.resolve(request);

                    }, function (error) {
                        request.status = error.status;
                        defered.reject(request);
                    })

                return defered.promise;

            }
        }
        return factory;

    }

);
