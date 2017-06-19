'use strict';

/**
 * @ngdoc function
 * @name eLogcardFrontApp.controller:ShowPartListingCtrl
 * @description
 * # ShowPartListingCtrl
 * Controller of the eLogcardFrontApp
 */
app.controller('ShowPartListingCtrl', function ($location, $http, $routeParams, userService, eLogcardService) {
    this.awesomeThings = [
      'HTML5 Boilerplate',
      'AngularJS',
      'Karma'
    ];
    var self = this;
    this.itemId = $routeParams.itemid;
    this.itemType = $routeParams.itemtype;
    this.debug = false;
    this.item;
    this.name;
    this.parts;

    this.typeNameIndex = {
        "assemblies": "assemblyName",
        "parts": "partName",
        "aircrafts": "aircraftName"
    }


    // gestion evenement  pour consulter les log d'une part
    this.doClickShowLog = function (partId) {
        $location.path("/showpartlog/" + partId);
    }

    //construction requetes http 

    //recuperation des information de l'assemblie

    let showPartlogUriWitoutParameter = "/blockchain/logcard/" + this.itemType + "/historic/";
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
                    self.name = self.item[self.typeNameIndex[self.itemType]];
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
    let showPartListUriWitoutParameter = "/blockchain/logcard/" + this.itemType + "/listing/";
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
