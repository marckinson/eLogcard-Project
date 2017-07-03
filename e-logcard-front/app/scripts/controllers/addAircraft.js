'use strict';

/**
 * @ngdoc function
 * @name eLogcardFrontApp.controller:addAircraftCtrl
 * @description
 * # addAircraftCtrl
 * Controller of the eLogcardFrontApp
 */
app.controller('addAircraftCtrl', function ($location, eLogcardService) {
    this.awesomeThings = [
      'HTML5 Boilerplate',
      'AngularJS',
      'Karma'
    ];
    var self = this;
    this.debug = false;
    this.faillureRequest = false;
    this.status;
    // set default value 
    // pour evite de reecrire 
    this.airCraftNumber = "1";
    this.SerialNumber = "1";
    this.name = "Aircraft1";

    this.doClickCreateParts = function (form) {
        if (form.$valid) {

            eLogcardService.createAircraft(self.SerialNumber, self.airCraftNumber, self.name).then(
                function (response) {
                    self.answer = response.answer;
                    self.status = response.status;
                    if (self.debug)
                        console.log(self.status)

                    $location.path('/showaircrafts');
                },
                function (error) {

                    self.answer = error.answer;
                    self.faillureRequest = true;
                    self.status = error.status;
                    if (self.debug)
                        console.log(self.status);

                })
        }
    }
});
