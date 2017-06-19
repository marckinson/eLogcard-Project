'use strict';

/**
 * @ngdoc function
 * @name eLogcardFrontApp.controller:headerCtrl
 * @description
 * # HeaderCtrl
 * Controller of the eLogcardFrontApp
 */

app.controller('headerCtrl', function ($location, userService) {
    var self = this;
    this.openParts = false;
    this.openAirCraft = false;
    this.openAssembly = false;

    this.loginUser = userService.getUser();
    this.connected = userService.getState();
    this.userRole = userService.getRole();

    this.doClicklogOut = function () {
        userService.disconnectUser();
        $location.path('/home');

    };

    this.doClickOpenParts = function () {
        self.openAirCraft = false;
        self.openAssembly = false;
        self.openParts = !self.openParts;
    };

    this.doClickOpenAssembly = function () {
        self.openParts = false;
        self.openAirCraft = false;
        self.openAssembly = !self.openAssembly;
    };

    this.doClickOpenAirCraf = function () {
        self.openParts = false;
        self.openAssembly = false;
        self.openAirCraft = !self.openAirCraft;
    };
});
