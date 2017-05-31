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
    // set default value 
    // pour evite de reecrire 
    this.partNumber = "eeef";
    this.SerialNumber = "gefff";
    this.partName = "Helice";
    this.type = "Defense";
    this.responsible = "florent";
    this.helicopter = "Tigre";
    this.assembly = "Assembly75";


    this.doClickCreateParts = function (form) {
        if (form.$valid) {

            let createUri = "/blockchain/logcard/parts";
            var data = {
                "pn": self.partNumber,
                "sn": self.SerialNumber,
                "partName": self.partName,
                "type": self.type,
                "responsible": self.responsible,
                "helicopter": self.helicopter,
                "assembly": self.assembly
            };
            $http.post(createUri, data)
                .then(
                    function (response) {
                        self.answer = response.data;
                        self.status = response.status;
                        //console.log(self.status)
                        // if (self.status = 200)
                        $location.path('/showparts');
                    },
                    function (response) {
                        self.answer = response.data || 'Request failed';
                        self.faillureRequest = true;
                        self.status = response.status;
                        //console.log(self.status);
                    }
                );
        }
    }
});
