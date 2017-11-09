'use strict';

/**
 * @ngdoc function
 * @name eLogcardFrontApp.controller:transferCtrl
 * @description
 * # transferCtrl
 * Controller of the eLogcardFrontApp
 */
app.controller('transferCtrl', function ($location, $routeParams, userService, eLogcardService) {
    this.awesomeThings = [
      'HTML5 Boilerplate',
      'AngularJS',
      'Karma'
    ];
    //transferTarget
    this.debug = false;
    var self = this;
    this.ownerMode = true;
    this.itemId = $routeParams.itemid;
    this.itemType = $routeParams.itemtype;
    this.UsertransferTarget;
    this.answer;
    this.status;
    this.data = {};
    this.faillureRequest = false;
    // utiliser pour genere les bouton radio
    this.listTypeTransfert = ["Owner", "Responsible"];
	let init=function(){
		eLogcardService.findAllUsers().then(function(response){
			self.users = response;
		});
		
	};
	init();
    
    // initialise le objet d Aiguillage

    this.crossRoadTransfers = {
        part: {
            Owner: {
                label: "Owner",
                url: "/showparts",
                call: eLogcardService.transfertPartOwnership
            },
            Responsible: {
                label: "Responsible",
                url: "/showpartlog/" + self.itemId,
                call: eLogcardService.transfertPartResponsible
            }
        },
        assembly: {
            Owner: {
                label: "Owner",
                url: "/showassemblies",
                call: eLogcardService.transfertAssemblyOwnerShip
            },
            Responsible: {
                label: "Responsible",
                url: "/showlogs/assembly/" + self.itemId,
                call: eLogcardService.transfertAssemblyResponsible
            }
        },
        aircraft: {
            Owner: {
                label: "Owner",
                url: "/showaircrafts",
                call: eLogcardService.transfertAirCraftOwnerShip
            },
            Responsible: {
                label: "Responsible",
                url: "/showlogs/aircraft/" + self.itemId,
                call: eLogcardService.transfertAirCraftResponsible
            }
        }
    };

    // initialise la valeur par defaut des bouton radio 

    this.radioTransferType = this.listTypeTransfert[0];
    if (this.debug) {
        console.log("itemtype");
        console.log(this.itemType);
    }

    this.doClickSendTransfert = function (form) {

        // on reseigne data en fonction du type de transfer 

        if (self.debug) {
            console.log(self.data);
        }

        if (form.$valid) {

            if (self.debug) {
                console.log('userTaget');
                console.log(self.UsertransferTarget);
                console.log('type de litem ');
                console.log(self.itemType);
                console.log('type de transfert ');
                console.log(self.radioTransferType);
                console.log('id item');
                console.log(self.itemId);
            }

            if (userService.getState()) {

                var targetExist = self.crossRoadTransfers[self.itemType][self.radioTransferType];

                if (self.debug) {
                    console.log("target existe")
                    console.log(targetExist);
                }

                if (targetExist) {

                    var transfertFunction = targetExist.call;

                    if (self.debug) {
                        console.log("TransfertFunction");
                        console.log(transfertFunction);
                    }

                    transfertFunction(self.UsertransferTarget, self.itemId)
                        .then(function (reponse) {
                            if (self.debug) {
                                console.log("transfert " + targetExist.label +
                                    " part succes ");
                                console.log(reponse);
                            }

                            $location.path(targetExist.url);
                            self.faillureRequest = false;
                            self.answer = reponse.answer;

                            if (self.debug) {
                                console.log("self.answer");
                                console.log(self.answer);
                            }
                        }, function (error) {
                            self.faillureRequest = true;
                            self.status = error.status;
                            self.answer = error.data;
                        });
                }
            }
        }
    }
});
