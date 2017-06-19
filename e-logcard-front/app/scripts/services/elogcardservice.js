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
            // scrap
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
            //AIRCRAFT
            //Show Aircrafts
            // scrap parts 
            // scrap assembli
            //add part on Aircrafts
            addPartOnAirCraft: function (idAircraft, idPart) {
                var defered = $q.defer();

                // construction requete 
                let addPartOnAirCraftUri = self.baseproxyUri + "logcard/aircrafts/add/part/" + idAircraft;

                if (self.debug)
                    console.log(addPartOnAirCraftUri);

                factory.addTargetToSource(idAircraft, idPart, addPartOnAirCraftUri)
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


            //ASSEMBLY
            // scrat parts 
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
            // add part on assemblies
            addPartOnAssembly: function (idAssembly, idPart) {

                var defered = $q.defer();

                // construction requete 
                let addPartOnAssemblyUri = self.baseproxyUri + "logcard/assemblies/add/" + idAssembly;

                if (self.debug)
                    console.log(addPartOnAssemblyUri);
                factory.addTargetToSource(idAssembly, idPart, addPartOnAssemblyUri)
                    .then(function (reponse) {
                        defered.resolve(reponse);
                    }, function (error) {
                        defered.reject(error);
                    })
                return defered.promise;
            },
            // scrapp assembly
            scrappAssembly: function (AssemblyId) {

                var defered = $q.defer();
                // construction requete 
                let scrappAssemblytUri = self.baseproxyUri + "logcard/assemblies/scrapp";
                if (self.debug)
                    console.log(scrappAssemblytUri);

                factory.scrappTarget(AssemblyId, scrappAssemblytUri)
                    .then(function (reponse) {
                        defered.resolve(reponse);
                    }, function (error) {
                        defered.reject(error);
                    })

                return defered.promise;

            },

            // all 
            // permet de faire les ajout de parts dans les autre ensemble 
            addTargetToSource: function (idSource, idTarget, restRequest) {
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

            }
        }
        return factory;

    }

);
