'use strict';

/**
 * @ngdoc function
 * @name eLogcardFrontApp.controller:addlogsCtrl
 * @description
 * # addlogsCtrl
 * Controller of the eLogcardFrontApp
 */
app.controller('addlogsCtrl', function ($location, $routeParams, eLogcardService) {
    this.awesomeThings = [
      'HTML5 Boilerplate',
      'AngularJS',
      'Karma'
    ];
    this.debug = false;
    var self = this;
    var faillureRequest = false;
    var status;
    this.data = {};
    this.itemId = $routeParams.itemid;
    this.itemType = $routeParams.itemtype;
    // jeux d essai
    this.modTypeSelected = "SB";
    this.description
    this.modTypes = [];


    this.doClickPerformActivitie = function (form) {
        if (form.$valid) {
            eLogcardService.addLogOnPart(self.itemId, self.modTypeSelected, self.description).then(
                function (response) {
                    self.answer = response.data;
                    self.status = response.status;
                    if (self.debug) {
                        console.log(response);
                        console.log(response.status);
                        console.log(response.data);
                    }
                    $location.path('/showpartlog/' + self.itemId);
                },
                function (error) {
                    self.answer = error.data;
                    self.faillureRequest = true;
                });
        }
    }

    // apelle function get list modification type 
    eLogcardService.getListModificationType().then(function (modTypeRequest) {
            self.modTypes = modTypeRequest.ModificationTypes;
            self.answer = modTypeRequest.aswer;

            // affecte la  valeur par defaut O SB 
            self.modTypeSelected = self.modTypes[0].value;
            if (self.debug) {
                console.log(modTypeRequest.ModificationTypes)
                console.log(modTypeRequest.aswer)
                console.log(modTypeRequest.status)
            }
        },
        function (error) {
            // permet d afficher que le requet role a echoue 
            self.faillureRolesRequest = true;
            if (self.debug) {
                console.log(modTypeRequest.aswer)
                console.log(modTypeRequest.status)
            }
        });


});
