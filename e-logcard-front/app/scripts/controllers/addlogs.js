'use strict';

/**
 * @ngdoc function
 * @name eLogcardFrontApp.controller:addlogsCtrl
 * @description
 * # addlogsCtrl
 * Controller of the eLogcardFrontApp
 */
angular.module('eLogcardFrontApp')
    .controller('addlogsCtrl', function ($location, $http, $routeParams, userService) {
        this.awesomeThings = [
      'HTML5 Boilerplate',
      'AngularJS',
      'Karma'
    ];
        var self = this;
        var faillureRequest = false;
        var status;
        this.data = {};
        this.itemId = $routeParams.itemid;
        this.itemType = $routeParams.itemtype;
        this.debug = false;
        this.modType = "SB";
        this.description = "changement de l'helice é@";

        // jeux de donnée exemple 
        /*{
	   "modType":"SB",
	   "description": "changement de l'helice"
        }
        */


        this.doClickPerformActivitie = function (form) {
            self.data = {
                "modType": self.modType,
                "description": self.description
            };
            if (self.debug)
                console.log(self.data);

            if (form.$valid) {
                if (self.debug) {
                    console.log("type: " + self.itemType);
                    console.log("id: " + self.itemId);
                }

                // exemple requete de base 
                // /logcard/parts/PerformActs/c1458970-4c1a-11e7-a9ae-998c688c5600
                let PerformActsUri = "/blockchain/logcard/parts/PerformActs/" + self.itemId;
                if (self.debug)
                    console.log(PerformActsUri);

                if (userService.getState()) {
                    $http.put(PerformActsUri, self.data)
                        .then(
                            function (response) {
                                self.answer = response.data;
                                self.status = response.status;
                                if (self.debug) {
                                    console.log(response);
                                    console.log(response.status);
                                    console.log(response.data);
                                }
                                $location.path('/showpartlog/' + self.itemId);
                            },
                            function (response) {
                                self.answer = response.data || 'Request failed';
                                self.faillureRequest = true;
                            }
                        );

                } else {
                    self.faillureRequest = true;
                    self.answer = " no user is connected "
                }
            }
        }
    });