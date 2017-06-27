'use strict';

/**
 * @ngdoc function
 * @name eLogcardFrontApp.controller:ShowaircraftassembliesCtrl
 * @description
 * # ShowaircraftassembliesCtrl
 * Controller of the eLogcardFrontApp
 */
app.
controller('ShowAirCraftAssembliesCtrl', function ($routeParams, $location, eLogcardService, userService) {
    this.awesomeThings = [
      'HTML5 Boilerplate',
      'AngularJS',
      'Karma'
    ];
    this.debug = false;
    this.answer;
    this.status;
    var self = this;
    this.showId = false;
    this.deletedAssemblies = {};
    this.itemId = $routeParams.itemid;
    this.message = "Show Aircraft's (" +this.itemId + ") assemblies  "
    this.aircraftMode = true;

    this.doClickShowLogs = function (idAssembly) {

        let showLogsUri = "/showlogs/" + 'assembly' + "/" + idAssembly;

        $location.path(showLogsUri);
        if (self.debug) {
            console.log(idAssembly);
            console.log(showLogsUri)
        }
    }

    this.doClickTransfertOwnerShip = function (idAssembly) {


        let transferUri = "/transfer/" + 'assembly/' + idAssembly;
        if (self.debug)
            console.log(transferUri);

        $location.path(transferUri);

    }
    // gestion evenement  pour consulter les log d'une assembly
    this.doClickShowParts = function (idAssembly) {

        let showPartsUri = "/showpartlist/" + 'assembly' + "/" + idAssembly;

        $location.path(showPartsUri);
        if (self.debug) {
            console.log(idAssembly);
            console.log(showPartsUri)
        }
    }

    // gestion evenement  pour consulter les log d'une assembly
    this.doClickAddPart = function (idAssembly) {

        let attachPartsUri = "/attachpart/" + 'assembly' + "/" + idAssembly;
        $location.path(attachPartsUri);

        if (self.debug) {
            console.log(idAssembly);
            console.log(attachPartsUri)
        }
    }

    this.doClickReplace = function (idAssembly) {

        $location.path("/replace/aircraft/" + self.itemId + "/assembly/" + idAssembly);
    }



    this.doClickRemove = function (idAssembly) {

        let confirmRemove = confirm("Are you sure you want to remove this assembly from this aircraft ?");
        if (confirmRemove == true) {
            eLogcardService.removeAssemblyToAicraft(self.itemId, idAssembly)
                .then(function (reponse) {
                    self.deletedAssemblies[idAssembly] = true;
                    if (self.debug) {
                        console.log("remove part succes ");
                        console.log(reponse);
                    }
                    self.faillureRequest = false;
                    self.answer = reponse.answer;
                    if (self.debug) {
                        console.log("self.answer");
                        console.log(self.answer);
                    }
                })
        }
    }

    // gestion evenement  pour scrapp une part
    this.doClickScrap = function (idAssembly) {
        let confirmScrapp = confirm("Are you sure you want to scrapp this Assembly?");
        if (confirmScrapp == true) {

            if (self.debug)
                console.log("call doClickScrap");
            eLogcardService.scrappAssembly(idAssembly)
                .then(function (reponse) {
                    self.deletedAssemblies[idAssembly] = true;

                    if (self.debug) {
                        console.log("scrapp part succes ");
                        console.log(reponse);


                    }
                    self.faillureRequest = false;
                    self.answer = reponse.answer;
                    if (self.debug) {
                        console.log("self.answer");
                        console.log(self.answer);
                    }
                })
        }
    }


    if (userService.getState()) {
        // recuperation liste de assemblies 
        eLogcardService.getAirCraftListAssemby(this.itemId)
            .then(
                function (response) {
                    self.assemblies = response.list;
                    self.status = response.status;

                    if (self.debug) {
                        console.log(response);
                        console.log(response.status);
                        console.log(response.list);
                    }
                },
                function (response) {
                    self.answer = response.status || 'Request failed';
                    if (self.debug) {
                        console.log(response);
                    }
                }
            );
    }

});
