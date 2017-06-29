'use strict';

/**
 * @ngdoc function
 * @name eLogcardFrontApp.controller:ReplaceCtrl
 * @description
 * # ReplaceCtrl
 * Controller of the eLogcardFrontApp
 */
app.controller('ReplaceCtrl', function ($routeParams, $location, eLogcardService) {
    this.awesomeThings = [
      'HTML5 Boilerplate',
      'AngularJS',
      'Karma'
    ];

    this.self = this;
    this.debug = false;
    this.container = $routeParams.container;
    this.containerId = $routeParams.containerid;
    this.item = $routeParams.item;
    this.itemId = $routeParams.itemid;

    this.targetId = "";
    this.aswer;
    this.satus;
    this.faillureRequest = false;
    /* voir tranfert.js pour gere les multi funtion */
    this.crossRoadReplace = {
        aircraft: {
            part: {
                url: "/showaircraftlistingpart/" + self.containerId,
                call: eLogcardService.replacePartOnAircraft
            },
            assembly: {
                url: "/aircraft/showassemblies/" + self.containerId,
                call: eLogcardService.replaceAssemblyOnAircraft
            }
        },
        assembly: {
            part: {
                url: "/showpartlist/assembly/" + self.containerId,
                call: eLogcardService.replacePartOnAssembly
            }
        }
    }

    this.doClickRemplacePart = function (partId) {
        var targetExist = self.crossRoadReplace[self.container][self.item];
        if (self.debug) {
            console.log("id  part clicker ");
            console.log(self.itemId);

            console.log("id  part saisie  ");
            console.log(self.targetId);
            console.log("id  aircraft ");
            console.log(self.containerId);


            console.log("container");
            console.log(self.container);
            console.log("item");
            console.log(self.item);

            if (targetExist) {
                console.log("targetExist");
                console.log(targetExist);
                console.log("replaceFunction");
                console.log(targetExist.call);
            }
        }
        if (targetExist) {
            var replaceFunction = targetExist.call;

            replaceFunction(self.itemId, self.targetId, self.containerId).then(function (reponse) {
                if (self.debug) {
                    console.log("add " + self.item + " succes ");
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
                    let url = targetExist.url
                    if (self.debug)
                        console.log(targetExist.url);
                    $location.path(targetExist.url);
                }

            }, function (error) {
                self.faillureRequest = true;
                self.status = error.status;
                self.answer = error.data;
            })

        }

    }

});
