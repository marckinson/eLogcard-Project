'use strict';

/**
 * @ngdoc function
 * @name eLogcardFrontApp.controller:showOneAirCraftCtrl
 * @description
 * # showOneAirCraftCtrl
 * Controller of the eLogcardFrontApp
 */
app.controller('showOneAirCraftCtrl', function ($http,$routeParams, $location, userService, eLogcardService) {
    this.debug = false;
    var self = this;
	self.itemid = $routeParams.itemid;
    this.showId = false;
    this.deletedAirCrafts = {};
    // gestion evenement  pour consulter les log d'une aircraft
    this.doClickShowLogs = function (idAircraft) {

        let showLogsUri = "/showlogs/" + 'aircraft' + "/" + idAircraft;

        $location.path(showLogsUri);
        if (self.debug) {
            console.log(idAircraft);
            console.log(showLogsUri)
        }
    }
    // gestion evenement  pour consulter les log d'une aircraft
    this.doClickTransfertOwnerShip = function (idAircraft) {

        let transferUri = "/transfer/" + 'aircraft/' + idAircraft;
        if (self.debug)
            console.log(transferUri);

        $location.path(transferUri);
    }

    // gestion evenement  pour consulter les Assemblies d'une aircraft
    this.doClickAssemblies = function (idAircraft) {

        let airCraftAssembliesUri = '/aircraft/showassemblies/' + idAircraft;
        if (self.debug)
            console.log(airCraftAssembliesUri);

        $location.path(airCraftAssembliesUri);
    }

    // gestion evenement  pour consulter les log d'une aircraft
    this.doClickShowParts = function (idAircraft) {

        let showPartsUri = "/showaircraftlistingpart/" + idAircraft

        $location.path(showPartsUri);
        if (self.debug) {
            console.log(idAircraft);
            console.log(showPartsUri)
        }
    }
    // gestion evenement  ajoute une part sur aircraft
    this.doClickAddPart = function (idAircraft) {

        let attachPartsUri = "/attachpart/" + 'aircraft' + "/" + idAircraft;
        $location.path(attachPartsUri);

        if (self.debug) {
            console.log(idAircraft);
            console.log(attachPartsUri)
        }
    }
    // gestion evenement ajouter une assembly on part 
    this.doClickAddAssembly = function (idAircraft) {

        let attachAssemblyUri = "/attachAssembly/" + 'aircraft' + "/" + idAircraft;
        $location.path(attachAssemblyUri);

        if (self.debug) {
            console.log(idAircraft);
            console.log(attachAssemblyUri)
        }
    }
    // gestion evenement  pour scrapp une part
    this.doClickScrap = function (idAircraft) {
        let confirmScrapp = confirm("Are you sure you want to scrap this Aircraft?");
        if (confirmScrapp == true) {

            if (self.debug)
                console.log("call doClickScrap");

            eLogcardService.scrappAirCraft(idAircraft)
                .then(function (reponse) {
                    self.deletedAirCrafts[idAircraft] = true;

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
    // execution de la requete pour rammene tout les Aicraft 
    let getAirCraftHistoryUri = "/blockchain/logcard/aircrafts" + self.itemid;
  //   if (userService.getState()) {

        $http.get(getAirCraftHistoryUri)
            .then(
                function (response) {
					var parts=typeof response.data==="Array"?response.data:[response.data];

                    self.airCrafts = response.data;
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
//   }

});
