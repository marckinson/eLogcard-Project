'use strict';

/**
 * @ngdoc function
 * @name fiveAppApp.controller:headerCtrl
 * @description
 * # HeaderCtrl
 * Controller of the fiveAppApp
 */

app.controller("headerCtrl", function ($scope,
    userService
) {
    $scope.logUser = userService.getUser();

});
