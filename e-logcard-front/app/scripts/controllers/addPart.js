'use strict';

/**
 * @ngdoc function
 * @name eLogcardFrontApp.controller:addPartCtrl
 * @description
 * # addPartCtrl
 * Controller of the eLogcardFrontApp
 */
app.controller('addPartCtrl', function ($location, eLogcardService) {
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
    this.partNumber
    this.SerialNumber
    this.partName
    this.type = "Defense";

    this.doClickCreateParts = function (form) {

        if (form.$valid) {
            // recupertation information a envoyer 
            var header = {
                "pn": self.partNumber,
                "sn": self.SerialNumber,
                "partName": self.partName,
                "type": self.type,
                "helicopter": "",
                "assembly": ""
            };

            eLogcardService.createPart(header).then(
                function (response) {
                    self.answer = response.answer;
                    self.status = response.status;
                    if (self.debug) {
                        console.log(self.status)
                    }
                    $location.path('/showparts');
                },
                function (error) {

                    self.answer = error.answer;
                    self.faillureRequest = true;
                    self.status = error.status;
                    if (self.debug) {
                        console.log(self.status);
                    }

                });
        }
    }
});
