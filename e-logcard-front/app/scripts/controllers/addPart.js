'use strict';

/**
 * @ngdoc function
 * @name eLogcardFrontApp.controller:addPartCtrl
 * @description
 * # addPartCtrl
 * Controller of the eLogcardFrontApp
 */
app.controller('addPartCtrl', function ($location, $http, $route, userService) {
    this.awesomeThings = [
      'HTML5 Boilerplate',
      'AngularJS',
      'Karma'
    ];
    var self = this;
    var faillureRequest = false;
    var status;
    this.debug = false;
    // set default value 
    // pour evite de reecrire 
    this.partNumber = "MSP430FG";
    this.SerialNumber = "IQZ342";
    this.partName = "Helice";
    this.type = "Defense";
    this.responsible = "AHSupplier";
    this.helicopter = "";
    this.assembly = "";


    this.doClickCreateParts = function (form) {
      
        if (form.$valid) {
           

            let createUriPart = "/blockchain/logcard/parts";
            var data = {
                "pn": self.partNumber,
                "sn": self.SerialNumber,
                "partName": self.partName,
                "type": self.type,
                "responsible": self.responsible,
                "helicopter": self.helicopter,
                "assembly": self.assembly
            };
            $http.post(createUriPart, data)
                .then(
                    function (response) {
                        self.answer = response.data;
                        self.status = response.status;
                        if (self.debug) {
                            console.log(self.status);
                            console.log(self.answer);
                        }
                        $location.path('/showparts');
                    },
                    function (response) {
                        self.answer = response.data || 'Request failed';
                        self.faillureRequest = true;
                        self.status = response.status;
                        if (self.debug)
                            console.log(self.status);
                    }
                );
        }
    }
});
