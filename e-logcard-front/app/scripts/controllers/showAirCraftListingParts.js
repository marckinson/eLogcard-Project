'use strict';

/**
 * @ngdoc function
 * @name eLogcardFrontApp.controller:showAircraftListingPartsCtrl
 * @description
 * # showAircraftListingPartsCtrl
 * Controller of the eLogcardFrontApp
 */

app.controller('showAircraftListingPartsCtrl', function ($location, $http, $routeParams, userService, eLogcardService) {
    this.awesomeThings = [
      'HTML5 Boilerplate',
      'AngularJS',
      'Karma'
    ];
    var self = this;
    this.debug = true;
    this.itemId = $routeParams.itemid;
    this.item = {};
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
        let confirmRemove = confirm("Are you sure you want to remove this Part from this aircraft ?");
        if (confirmRemove == true) {
            // apelle le service ici 
            eLogcardService.removePartToAirCraft(self.itemId, partId)
                .then(
                    function (reponse) {
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

                    },
                    function (error) {
                        self.answer = error.data || 'Request failed';
                        if (self.debug) {
                            console.log(error);
                        }
                    }
                )
        }
    }
    if (userService.getState()) {
        // recuperation des information de l'aicraft courant 
        eLogcardService.getAircraftHistoric(this.itemId)
            .then(
                function (response) {
                    self.item = response.aircraft;
                    self.status = response.status;
                    self.name = self.item["componentName"];
                    if (self.debug) {
                        console.log("name:" + self.name);
                        console.log("status: " + response.status);
                        console.log("item: ");
                        console.log(self.item);
                    }
                },
                function (response) {
                    self.answer = response.status || 'Request failed';
                    if (self.debug) {
                        console.log(response);
                    }
                }
            );

        // recuperation liste de part 
        eLogcardService.getAircraftlistParts(this.itemId)
            .then(
                function (response) {
                    self.parts = response.list;
                    self.status = response.status;

                    if (self.debug) {
                        console.log(response);
                        console.log(response.status);
                        console.log(response.list);
                    }
                },
                function (response) {
                    self.answer = response.status || 'Request failed';
                    if (self.debug) {
                        console.log(response);
                    }
                }
            );
    }

});
