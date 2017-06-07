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

    /*

        this.airCrafts = [{
            "an": "CL4P7R4P",
            "sn": "hyperion",
            "id_aircraft": "2fb0dea0-46a8-11e7-956e-cd3b1eedcf08",
            "owner": "sora",
            "parts": null,
            "logs": [{
                "log_type": "CREATE",
                "vDate": "2017/06/01 10:56:22",
                "owner": "sora",
                "responsible": "",
                "modType": "",
                "description": ""
           }]
        }];*/

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

    // gestion evenement pour changer de volet login ou sign up 
    this.doClickShowLog = function (event) {
        $location.path("/showpartlog/:" + event);
        // console.log(event);
    }

});
