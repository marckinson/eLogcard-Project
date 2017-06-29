'use strict';

/**
 * @ngdoc function
 * @name eLogcardFrontApp.controller:ShowLogsCtrl
 * @description
 * # ShowLogsCtrl
 * Controller of the eLogcardFrontApp
 */
app.controller('showLogsCtrl', function ($location, $http, $routeParams, userService) {
    this.awesomeThings = [
      'HTML5 Boilerplate',
      'AngularJS',
      'Karma'
    ];
    var self = this;
    this.debug = false;
    this.itemId = $routeParams.itemid;
    this.itemType = $routeParams.itemtype;
    this.item;
    this.name;
    this.crossRoad = {
        assembly: {
            url: "assemblies"
        },
        aircraft: {
            url: "aircrafts"
        }
    };

    // construction de la requete en fonction du type d object demande
    if (this.debug)
        console.log(this.crossRoad[this.itemType]);
    let showPartlogUriWitoutParameter = "/blockchain/logcard/" + this.crossRoad[this.itemType].url + "/historic/";
    if (this.debug)
        console.log(this.itemId);
    let showPartlogUriIdParameter = showPartlogUriWitoutParameter + this.itemId;
    if (this.debug)
        console.log(showPartlogUriIdParameter);

    if (userService.getState()) {
        $http.get(showPartlogUriIdParameter)
            .then(
                function (response) {
                    self.item = response.data;
                    self.status = response.status;
                    self.name = self.item.componentName;
                    //self.name = self.name + " " + self.item.partName;
                    if (self.debug) {
                        console.log(response);
                        console.log(response.status);
                        console.log(response.data);
                    }
                },
                function (response) {
                    self.answer = response.data || 'Request failed';
                    if (self.debug) {
                        console.log(response);
                    }
                }
            );
    }
});
