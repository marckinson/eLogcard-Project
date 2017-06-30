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
    this.debug = true;
    var self = this;
    this.itemId = $routeParams.itemid;
    this.itemType = $routeParams.itemtype;
    this.item;
    this.name;
    this.parts;
    this.deletedParts = {};

    this.typeNameIndex = {
        assembly: "assemblies",
        aircraft: "aircrafts"
    }

    // gestion evenement  pour consulter les log d'une part
    this.doClickShowLog = function (partId) {
        $location.path("/showpartlog/" + partId);
    }

    // gestion evenement  pour remplacer une part par une autre 
    this.doClickReplacePart = function (partId) {
        $location.path("/replace/assembly/" + self.itemId + "/part/" + partId);
    }
    // gestion evenement pour remove une part de l assembly
    this.doClickRemovePart = function (partId) {

        let confirmRemove = confirm("Are you sure you want to remove this Part from this assembly ?");
        if (confirmRemove == true) {
            eLogcardService.removePartToAssembly(self.itemId, partId)
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
    //recuperation des information de l'assemblie
    //construction requetes http 
    let showPartlogUriWitoutParameter = "/blockchain/logcard/" + this.typeNameIndex[this.itemType] + "/historic/";
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
                    self.name = self.item["componentName"];
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
    // recuperation liste de part 
    let showPartListUriWitoutParameter = "/blockchain/logcard/" + this.typeNameIndex[this.itemType] + "/partslisting/";
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
