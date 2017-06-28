'use strict';

/**
 * @ngdoc function
 * @name eLogcardFrontApp.controller:AttachassemblyCtrl
 * @description
 * # AttachassemblyCtrl
 * Controller of the eLogcardFrontApp
 */
app.controller('AttachassemblyCtrl', function ($routeParams, $location, eLogcardService) {
    this.awesomeThings = [
      'HTML5 Boilerplate',
      'AngularJS',
      'Karma'
    ];
    'use strict';

    var self = this;
    this.debug = false;
    this.itemId = $routeParams.itemid;
    this.itemType = $routeParams.itemtype;
    this.targetId = "";
    this.aswer;
    this.satus;
    this.faillureRequest = false;
    // tableau de items(assembly) disponible 
    this.items;
    this.defautcombo = "Selected assembly";
    this.showComboBox = false;


    this.doClickAttach = function (partId) {
        if (self.debug) {
            console.log("call doClickAttachPart");
            console.log(self.itemType);
        }


        let attachFunction = eLogcardService.addAssemblyOnAirCraft

        attachFunction(self.itemId, self.targetId)
            .then(function (reponse) {
                if (self.debug) {
                    console.log("add assembly");
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
                    let url = "aircraft/showassemblies/" + self.itemId
                    if (self.debug)
                        console.log(url);
                    $location.path(url);
                }


            }, function (error) {
                self.faillureRequest = true;
                self.status = error.status;
                self.answer = error.data;
            })
    }


    //charge la liste de assemblies disponible 
    /*
    eLogcardService.getListPartWithoutAirCraft()
        .then(function (reponse) {
            self.items = reponse.assemblies;
            if (self.debug) {
                console.log(self.items)
                console.log(reponse.status)
            }
        }, function (error) {
            if (debug) {
                console.log(error.aswer)
                console.log(error.status)
            }

        })*/

});
