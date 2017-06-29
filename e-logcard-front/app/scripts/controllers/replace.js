'use strict';

/**
 * @ngdoc function
 * @name eLogcardFrontApp.controller:ReplaceCtrl
 * @description
 * # ReplaceCtrl
 * Controller of the eLogcardFrontApp
 */
app.controller('ReplaceCtrl', function ($routeParams, $location, eLogcardService, userService) {
    this.awesomeThings = [
      'HTML5 Boilerplate',
      'AngularJS',
      'Karma'
    ];

    var self = this;
    this.debug = false;
    this.container = $routeParams.container;
    this.containerId = $routeParams.containerid;
    this.item = $routeParams.item;
    this.itemId = $routeParams.itemid;

    this.targetId = "";
    this.aswer;
    this.satus;
    this.faillureRequest = false;
    // use on combo box 
    this.items;
    this.showComboBox = true;

    this.indexCombobox = {
        part: {
            part: true,
            defautcombo: "Selected Part",
            labelCombobox: "Parts List"
        },
        assembly: {
            part: false,
            defautcombo: "Selected assembly",
            labelCombobox: "Assemblies list"
        }
    };

    this.crossRoadReplace = {
        aircraft: {
            part: {
                url: "/showaircraftlistingpart/" + self.containerId,
                call: eLogcardService.replacePartOnAircraft,
                callGetList: eLogcardService.getListPartWithoutAirCraft
            },
            assembly: {
                url: "/aircraft/showassemblies/" + self.containerId,
                call: eLogcardService.replaceAssemblyOnAircraft,
                callGetList: eLogcardService.getListAssemblyWithoutAircraft
            }
        },
        assembly: {
            part: {
                url: "/showpartlist/assembly/" + self.containerId,
                call: eLogcardService.replacePartOnAssembly,
                callGetList: eLogcardService.getListPartWithoutAssembly
            }
        }
    };
    this.defautcombo = this.indexCombobox[this.item].defautcombo;
    this.labelCombobox = this.indexCombobox[this.item].labelCombobox;
    // load List items 
    if (userService.getState()) {

        let getlistFunction = self.crossRoadReplace[self.container][self.item].callGetList;

        if (this.debug) {
            console.log("getlistFunction");
            console.log(getlistFunction);
        }

        getlistFunction()
            .then(function (reponse) {
                    if (self.indexCombobox[self.item].part == true)
                        self.items = reponse.parts;
                    else
                        self.items = reponse.assemblies;

                    // pour l afficher dans la vue attash 
                    let taille = self.items.length;
                    for (let i = 0; i < taille; i++) {
                        if (self.indexCombobox[self.item].part == true)
                            self.items[i].componentName = self.items[i].partName;
                        else
                            self.items[i].id = self.items[i].id_assembly;
                    }
                    if (self.debug) {
                        console.log(self.items)
                        console.log(reponse.status)
                    }
                },
                function (error) {
                    if (self.debug) {
                        console.log(error.aswer)
                        console.log(error.status)
                    }
                })
    }

    this.doClickRemplacePart = function (partId) {
        if (userService.getState()) {
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
    }
});
