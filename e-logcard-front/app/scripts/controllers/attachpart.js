'use strict';

/**
 * @ngdoc function
 * @name eLogcardFrontApp.controller:AttachpartCtrl
 * @description
 * # AttachpartCtrl
 * Controller of the eLogcardFrontApp
 */
app.controller('AttachpartCtrl', function ($routeParams, $location, eLogcardService) {
    this.awesomeThings = [
      'HTML5 Boilerplate',
      'AngularJS',
      'Karma'
    ];

    self = this;
    this.debug = false;
    this.itemId = $routeParams.itemid;
    this.itemType = $routeParams.itemtype;
    this.targetId = "";
    this.aswer;
    this.satus;
    this.faillureRequest = false;

    this.crossRoad = {
        assembly: {
            url: "/showpartlist/assembly/",
            call: eLogcardService.addPartOnAssembly
        },
        aircraft: {
            url: "/showaircraftlistingpart/",
            call: eLogcardService.addPartOnAirCraft
        }
    };

    this.doClickAttach = function (partId) {
        if (self.debug) {
            console.log("call doClickAttachPart");
            console.log(self.itemType);
            console.log(self.crossRoad[self.itemType]);

        }

        let attachFunction = self.crossRoad[self.itemType].call;

        attachFunction(self.itemId, self.targetId)
            .then(function (reponse) {
                if (self.debug) {
                    console.log("add part succes ");
                    console.log(reponse);
                }
                self.faillureRequest = false;
                self.answer = reponse.answer;
                if (self.debug) {
                    console.log("self.answer");
                    console.log(self.answer);
                }

                if (reponse.data == false)
                    self.faillureRequest = true;

                if (self.answer == true) {

                    // redirection to show part on assemblies or arcraift or assemnlie on aircraft 
                    if (self.debug)
                        console.log(self.crossRoad[self.itemType].url);
                    $location.path(self.crossRoad[self.itemType].url + self.itemId);
                }

            }, function (error) {
                self.faillureRequest = true;
                self.status = error.status;
                self.answer = error.data;
            })
    }

});
