'use strict';

/**
 * @ngdoc function
 * @name eLogcardFrontApp.controller:showPartsCtrl
 * @description
 * # showPartsCtrl
 * Controller of the eLogcardFrontApp
 */
app.controller('showPartsCtrl', function ($http, $location, userService, eLogcardService) {
    this.debug = false;
    var self = this;
    this.showId = true;
    this.deletedParts = {};
    // EVENT
    // gestion evenement envoi ver la vue le transfer d'une part 
    this.doClickTransfer = function (partId) {

        let transferUri = "/transfer/" + 'part/' + partId;
        if (self.debug)
            console.log(transferUri);

        $location.path(transferUri);
    }
    // gestion evenement envoi ver la vue d'ajout de log 
    this.doClickAddLog = function (partId) {

        if (self.debug)
            console.log("click doClickAddLog");

        let addLogUri = "/addlog/" + 'part/' + partId;
        if (self.debug)
            console.log(addLogUri);
        $location.path(addLogUri);
    }
    // gestion evenement  pour consulter les log d'une part
    this.doClickShowLog = function (partId) {
        $location.path("/showpartlog/" + partId);
    }
    // gestion evenement  pour generaion qr code  les log d'une part
    this.doClickShowQrcode = function (partId) {
        $location.path("/generateqrcode/part/" + partId);
    }
    // gestion evenement  pour scrapp une part
    this.doClickScrap = function (partId) {
        let confirmScrapp = confirm("Are you sure you want to scrap this Part?");
        if (confirmScrapp == true) {
            if (self.debug)
                console.log("call doClickScrap");
            eLogcardService.scrappPart(partId)
                .then(function (reponse) {
                    self.deletedParts[partId] = true;
                    if (self.debug) {
                        console.log("scrapp part succes ");
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
    if (self.debug)
        console.log("request block");
    // REQUEST
    // requete de recuperation des des parts generale 
    let showPartsUri = "/blockchain/logcard/parts";
	
   //  if (userService.getState()) {
        if (self.debug)
            console.log("request si connecter");

        $http.get(showPartsUri)
            .then(
                function (response) {
					var parts=typeof response.data==="Array"?response.data:[response.data];
                    self.Parts = response.data;
                    self.answer = response.data;
                    self.status = response.status;
                    if (self.debug) {
                        console.log(response.data);
                        console.log(response.status);
                    }
                },
                function (error) {
                    self.answer = error.data || 'Request failed';
                    if (self.debug) {
                        console.log(error.data);
                        console.log(error.status);
                    }
                }
            );
  //   }

});