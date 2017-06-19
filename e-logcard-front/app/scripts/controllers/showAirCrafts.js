'use strict';

/**
 * @ngdoc function
 * @name eLogcardFrontApp.controller:showAirCraftsCtrl
 * @description
 * # showAirCraftsCtrl
 * Controller of the eLogcardFrontApp
 */
app.controller('showAirCraftsCtrl', function ($http, $location, userService) {
    this.debug = false;
    this.answer;
    this.status;
    var self = this;
    this.showId = false;

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


    // gestion evenement  pour consulter les log d'une aircraft
    this.doClickShowLogs = function (id) {

        let showLogsUri = "/showlogs/" + 'aircrafts' + "/" + id;

        $location.path(showLogsUri);
        if (self.debug) {
            console.log(id);
            console.log(showLogsUri)
        }
    }

    // gestion evenement  pour consulter les log d'une aircraft
    this.doClickShowParts = function (id) {

        let showPartsUri = "/showaircraftlistingpart/" + id;

        $location.path(showPartsUri);
        if (self.debug) {
            console.log(id);
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


});
