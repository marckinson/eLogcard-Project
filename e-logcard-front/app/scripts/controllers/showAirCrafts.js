'use strict';

/**
 * @ngdoc function
 * @name eLogcardFrontApp.controller:showAirCraftsCtrl
 * @description
 * # showAirCraftsCtrl
 * Controller of the eLogcardFrontApp
 */
app.controller('showAirCraftsCtrl', function ($http, $location, userService, eLogcardService) {
    this.debug = false;
    this.answer;
    this.status;
    var self = this;
    this.showId = false;
    this.deletedAirCrafts = {};

    // gestion evenement  pour consulter les log d'une aircraft
    this.doClickShowLogs = function (idAircraft) {

        let showLogsUri = "/showlogs/" + 'aircrafts' + "/" + idAircraft;

        $location.path(showLogsUri);
        if (self.debug) {
            console.log(idAircraft);
            console.log(showLogsUri)
        }
    }

    // gestion evenement  pour consulter les log d'une aircraft
    this.doClickTransfertOwnerShip = function (idAircraft) {

        let transferUri = "/transfer/" + 'aircraft/' + idAircraft;
        if (self.debug)
            console.log(transferUri);

        $location.path(transferUri);
    }

    // gestion evenement  pour consulter les log d'une aircraft
    this.doClickShowParts = function (idAircraft) {

        let showPartsUri = "/showaircraftlistingpart/" + idAircraft
        // let showPartsUri = "/showpartlist/" + 'aircrafts' + "/" + idAircraft;

        $location.path(showPartsUri);
        if (self.debug) {
            console.log(idAircraft);
            console.log(showPartsUri)
        }
    }

    // gestion evenement  pour consulter les log d'une aircraft
    this.doClickAddPart = function (idAircraft) {

        let attachPartsUri = "/attachpart/" + 'aircraft' + "/" + idAircraft;
        $location.path(attachPartsUri);

        if (self.debug) {
            console.log(idAircraft);
            console.log(attachPartsUri)
        }
    }

    // gestion evenement  pour scrapp une part
    this.doClickScrap = function (idAircraft) {
        let confirmScrapp = confirm("Are you sure you want to scrapp this Aircraft?");
        if (confirmScrapp == true) {


            if (self.debug)
                console.log("call doClickScrap");
            eLogcardService.scrappAirCraft(idAircraft)
                .then(function (reponse) {
                    self.deletedAirCrafts[idAircraft] = true;

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

    // execution de la requete pour rammene tout les Aicraft 
    let showAirCraftsUri = "/blockchain/logcard/aircrafts";
    if (userService.getState()) {

        $http.get(showAirCraftsUri)
            .then(
                function (response) {
                    self.airCrafts = response.data;
                    self.answer = response.data;
                    self.status = response.status;
                    if (self.debug) {
                        console.log(response.data);

                    }
                },
                function (response) {
                    self.answer = response.data || 'Request failed';
                }
            );
    }

});
