'use strict';

/**
 * @ngdoc function
 * @name eLogcardFrontApp.controller:showPartsCtrl
 * @description
 * # showPartsCtrl
 * Controller of the eLogcardFrontApp
 */
app.controller('showPartsCtrl', function ($http, $location, userService) {
    this.debug = false;
    this.answer;
    this.status;
    var self = this;

    /* this.Parts = [{
         "pn": "Wffieng",
         "sn": "1024",
         "id": "x02048",
         "partName": "Wing",
         "type": "defence",
         "responsible": "sora",
         "owner": "florent",
         "helicopter": "tigre",
         "assembly": "3667 "
          }];*/

    let showPartsUri = "/blockchain/logcard/parts";
    if (userService.getState()) {

        $http.get(showPartsUri)
            .then(
                function (response) {
                    self.Parts = response.data;
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
    }

});
