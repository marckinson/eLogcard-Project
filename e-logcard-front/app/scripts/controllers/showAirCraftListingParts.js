'use strict';

/**
 * @ngdoc function
 * @name eLogcardFrontApp.controller:showAircraftListingPartsCtrl
 * @description
 * # showAircraftListingPartsCtrl
 * Controller of the eLogcardFrontApp
 */
angular.module('eLogcardFrontApp')
    .controller('showAircraftListingPartsCtrl', function ($location, $http, $routeParams, userService, eLogcardService) {
        this.awesomeThings = [
      'HTML5 Boilerplate',
      'AngularJS',
      'Karma'
    ];
        var self = this;
        this.itemId = $routeParams.itemid;
        this.debug = false;
        this.item;
        this.name;
        this.parts;
        this.deletedParts = {};

        // gestion evenement  pour consulter les log d'une part
        this.doClickShowLog = function (partId) {
            $location.path("/showpartlog/" + partId);
        }

        // gestion evenement  pour remplacer une part par une autre 
        this.doClickReplacePart = function (partId) {
            $location.path("/replace/aircraft/" + self.itemId + "/part/" + partId);
        }


        // gestion evenement  pour consulter les log d'une part
        this.doClickRemovePart = function (partId) {
            // $location.path("/showpartlog/" + partId);
            let confirmRemove = confirm("Are you sure you want to remove this Part to the aircraft ?");
            if (confirmRemove == true) {
                // apelle le service ici 
                eLogcardService.removePartToAirCraft(self.itemId, partId)
                    .then(function (reponse) {
                        self.deletedParts[partId] = true;
                        if (self.debug) {
                            console.log("remove part succes ");
                            console.log(reponse);
                        }
                        self.faillureRequest = false;
                        self.answer = reponse.answer;
                        if (self.debug) {
                            console.log("self.answer");
                            console.log(self.answer);
                        }

                    })
            }
        }


        let showPartlogUriWitoutParameter = "/blockchain/logcard/aircrafts/historic/";
        if (this.debug)
            console.log(this.itemId);
        let showPartlogUriIdParameter = showPartlogUriWitoutParameter + this.itemId;
        if (this.debug)
            console.log(showPartlogUriIdParameter);

        if (userService.getState()) {
            $http.get(showPartlogUriIdParameter)
                .then(
                    function (response) {
                        self.item = response.data;
                        self.status = response.status;
                        self.name = self.item["aircraftName"];
                        if (self.debug) {
                            console.log("name:" + self.name);
                            console.log("status: " + response.status);
                            console.log("data: ");
                            console.log(response.data);
                            console.log("item: ");
                            console.log(self.item);
                        }
                    },
                    function (response) {
                        self.answer = response.data || 'Request failed';
                        if (self.debug) {
                            console.log(response);
                        }
                    }
                );
        }
        ///logcard/aircrafts/listing/parts/d60be5e0-4d17-11e7-9e01-b7c4d567b2dd
        // recuperation liste de part 

        let showPartListUriWitoutParameter = "/blockchain/logcard/aircrafts/listing/parts/";
        if (this.debug)
            console.log(this.itemId);

        let showPartListUriIdParameter = showPartListUriWitoutParameter + this.itemId;
        if (this.debug)
            console.log(showPartListUriIdParameter);

        if (userService.getState()) {
            $http.get(showPartListUriIdParameter)
                .then(
                    function (response) {
                        self.parts = response.data;
                        self.status = response.status;

                        if (self.debug) {
                            console.log(response);
                            console.log(response.status);
                            console.log(response.data);
                        }
                    },
                    function (response) {
                        self.answer = response.data || 'Request failed';
                        if (self.debug) {
                            console.log(response);
                        }
                    }
                );
        }


    });
