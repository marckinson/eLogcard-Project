'use strict';

/**
 * @ngdoc function
 * @name eLogcardFrontApp.controller:headerCtrl
 * @description
 * # HeaderCtrl
 * Controller of the eLogcardFrontApp
 */

app.controller('headerCtrl', function ($location, userService) {
    this.loginUser = userService.getUser();
    this.connected = userService.getState();

    this.doClicklogOut = function () {
        alert("clik on logout");
        console.log("clik on logout");
    }

});
