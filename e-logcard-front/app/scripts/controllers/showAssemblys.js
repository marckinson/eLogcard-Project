'use strict';

/**
 * @ngdoc function
 * @name eLogcardFrontApp.controller:showAssemblysCtrl
 * @description
 * # showAssemblysCtrl
 * Controller of the eLogcardFrontApp
 */
app.controller('showAssemblysCtrl', function ($http, $location, userService, eLogcardService) {
    this.debug = false;
    this.answer;
    this.status;
    var self = this;
    this.showId = true;
    this.deletedAssemblies = {};
    this.message = "MyAssemblies"
    this.aircraftMode = false;
/*
    this.assemblies = [{
        "an": "a",
        "sn": "k",
        "id_assembly": "3de0a160-46a1-11e7-956e-cd3b1eedcf08",
        "owner": "sora",
        "componentName": "moteur ",
        "responsible": "sora",
        "parts": [{
            "pn": "Wffieng",
            "sn": "1024",
            "id": "x02048",
            "partName": "Wing",
            "type": "defence",
            "responsible": "sora",
            "owner": "florent",
            "helicopter": "tigre",
            "assembly": "3667 "
                                          }],
        "logs": [{
            "log_type": "CREATE",
            "vDate": "2017/06/01 10:06:39",
            "owner": "sora",
            "responsible": "",
            "modType": "",
            "description": ""
                                       }]
                                    }];
*/
    // EVENT
    // gestion evenement  pour consulter les log d'une assembly
    this.doClickShowLogs = function (idAssembly) {

        let showLogsUri = "/showlogs/" + 'assembly' + "/" + idAssembly;
        if (self.debug) {
            console.log(idAssembly);
            console.log(showLogsUri)
        }
        $location.path(showLogsUri);
    }
    // gestion evenement envoi ver la vue le transfer d'une part 
    this.doClickTransfertOwnerShip = function (idAssembly) {

        let transferUri = "/transfer/" + 'assembly/' + idAssembly;
        if (self.debug)
            console.log(transferUri);

        $location.path(transferUri);

    }

    // gestion evenement  pour consulter les log d'une assembly
    this.doClickShowParts = function (idAssembly) {

        let showPartsUri = "/showpartlist/" + 'assembly' + "/" + idAssembly;
        if (self.debug) {
            console.log(idAssembly);
            console.log(showPartsUri)
        }
        $location.path(showPartsUri);
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


    // gestion evenement  pour scrapp une part
    this.doClickScrap = function (idAssembly) {
        let confirmScrapp = confirm("Are you sure you want to scrap this Assembly?");
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
    // requete de recuperation des assemblies 
    eLogcardService.getAssemblies().then(function (response) {
        self.assemblies = response.assemblies;
        self.status = response.satus;
    }, function (error) {
        self.answer = error.data || 'Request failed';
        self.status = error.status;
    });

});
