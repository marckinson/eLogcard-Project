'use strict';

/**
 * @ngdoc function
 * @name eLogcardFrontApp.controller:addAssemblyCtrl
 * @description
 * # addAssemblyCtrl
 * Controller of the eLogcardFrontApp
 */
app.controller('addAssemblyCtrl', function ($location, eLogcardService) {
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
    this.assemblyNumber = "1";
    this.SerialNumber = "1";
    this.name = "Assembly1";

    this.doClickCreateAssembly = function (form) {
        if (self.debug)
            console.log("doClickCreateAssembly");
        if (form.$valid) {

            eLogcardService.createAssembly(self.SerialNumber, self.assemblyNumber, self.name).then(
                function (response) {
                    self.answer = response.answer;
                    self.status = response.status;
                    if (self.debug)
                        console.log(self.status)

                    $location.path('/showassemblies');
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
