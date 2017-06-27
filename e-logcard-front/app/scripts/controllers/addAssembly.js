'use strict';

/**
 * @ngdoc function
 * @name eLogcardFrontApp.controller:addAssemblyCtrl
 * @description
 * # addAssemblyCtrl
 * Controller of the eLogcardFrontApp
 */
app.controller('addAssemblyCtrl', function ($location, $http, $route, userService) {
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
    this.assemblyNumber = "EC";
    this.SerialNumber = "145";
    this.name = "turbine";


    this.doClickCreateAssembly = function (form) {
        if (self.debug)
            console.log("doClickCreateAssembly");
        if (form.$valid) {

            let createUriAssembly = "/blockchain/logcard/assemblies";
            var data = {
                "an": self.assemblyNumber,
                "sn": self.SerialNumber,
                "assemblyName": self.name
            };
            if (self.debug) {
                console.log(data);
                console.log(userService.getUser());
                console.log(userService.getRole());
            }

            $http.post(createUriAssembly, data)
                .then(
                    function (response) {
                        self.answer = response.data;
                        self.status = response.status;
                        if (self.debug) {

                            console.log(self.answer);
                            console.log(self.status);
                        }

                        $location.path('/showassemblies');
                    },
                    function (response) {
                        self.answer = response.data || 'Request failed';
                        self.faillureRequest = true;
                        self.status = response.status;
                        if (self.debug) {
                            console.log(self.status);
                        }
                    }
                );
        }
    }
});
