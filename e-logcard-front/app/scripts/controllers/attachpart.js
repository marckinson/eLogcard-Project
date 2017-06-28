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

    var self = this;
    this.debug = false;
    this.itemId = $routeParams.itemid;
    this.itemType = $routeParams.itemtype;
    this.targetId = "";
    this.aswer;
    this.satus;
    this.faillureRequest = false;
    // tableau de items(part) disponible 
    this.items;
    this.defautcombo = "Selected Part";

    this.crossRoad = {
        assembly: {
            url: "/showpartlist/assembly/",
            callAttach: eLogcardService.addPartOnAssembly,
            callgetItem: eLogcardService.getListPartWithoutAssembly
        },
        aircraft: {
            url: "/showaircraftlistingpart/",
            callAttach: eLogcardService.addPartOnAirCraft,
            callgetItem: eLogcardService.getListPartWithoutAirCraft

        }
    };

    this.doClickAttach = function (form) {
        if (self.debug) {
            console.log("call doClickAttachPart");
            console.log(self.itemType);
            console.log(self.crossRoad[self.itemType]);
            console.log(form.$valid);
        }

        if (form.$valid) {
            if (self.debug) {
                console.log("call doClickAttachPart");
                console.log(self.itemType);
                console.log(self.crossRoad[self.itemType]);

            }


            let attachFunction = self.crossRoad[self.itemType].callAttach;
            if (self.debug) {
                console.log("attachFunction");
                console.log(attachFunction);

            }
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
    }
    //charge la liste de part disponible 

    let getItemFunction = self.crossRoad[self.itemType].callgetItem;

    if (this.debug) {
        console.log("getItemFunction");
        console.log(getItemFunction);

    }
    getItemFunction()
        .then(function (reponse) {
            self.items = reponse.parts;
            if (self.debug) {
                console.log(self.items)
                console.log(reponse.status)
            }
        }, function (error) {
            if (self.debug) {
                console.log(error.aswer)
                console.log(error.status)
            }

        })

});
