'use strict';

/**
 * @ngdoc function
 * @name eLogcardFrontApp.controller:addAircraftCtrl
 * @description
 * # addAircraftCtrl
 * Controller of the eLogcardFrontApp
 */
app.controller('addAircraftCtrl', function ($location, $http, $route, userService) {
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
    this.airCraftNumber = "H";
    this.SerialNumber = "160";


    this.doClickCreateParts = function (form) {
        if (form.$valid) {

            let createUriAirCraft = "/blockchain/logcard/aircrafts";
            var data = {
                "an": self.airCraftNumber,
                "sn": self.SerialNumber,
            };
            $http.post(createUriAirCraft, data)
                .then(
                    function (response) {
                        self.answer = response.data;
                        self.status = response.status;
                        if (self.debug)
                            console.log(self.status)
                        $location.path('/showaircrafts');
                    },
                    function (response) {
                        self.answer = response.data || 'Request failed';
                        self.faillureRequest = true;
                        self.status = response.status;
                        if (sefl.debug)
                            console.log(self.status);
                    }
                );
        }
    }
});
