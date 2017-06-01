'use strict';

/**
 * @ngdoc function
 * @name eLogcardFrontApp.controller:showAssemblysCtrl
 * @description
 * # showAssemblysCtrl
 * Controller of the eLogcardFrontApp
 */
app.controller('showAssemblysCtrl', function ($http, $location, userService) {
    this.debug = false;
    this.answer;
    this.status;
    var self = this;
    /*
        this.assemblies = [{
            "an": "a",
            "sn": "k",
            "id_assembly": "3de0a160-46a1-11e7-956e-cd3b1eedcf08",
            "owner": "sora",
            "parts": null,
            "logs": [{
                "log_type": "CREATE",
                "vDate": "2017/06/01 10:06:39",
                "owner": "sora",
                "responsible": "",
                "modType": "",
                "description": ""
           }]
        }];*/

    let showPartsUri = "/blockchain/logcard/assemblies";
    if (userService.getState()) {

        $http.get(showPartsUri)
            .then(
                function (response) {
                    self.assemblies = response.data;
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
});
