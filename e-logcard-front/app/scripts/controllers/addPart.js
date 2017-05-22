'use strict';

/**
 * @ngdoc function
 * @name fiveAppApp.controller:addPartCtrl
 * @description
 * # addPartCtrl
 * Controller of the fiveAppApp
 */
app.controller('addPartCtrl', function ($scope, $location, $http, userService) {
    this.awesomeThings = [
      'HTML5 Boilerplate',
      'AngularJS',
      'Karma'
    ];

    $scope.partNumber = "ohi";
    $scope.SerialNumber = "o";
    $scope.partName = "Helice";
    $scope.type = "Defense";
    $scope.responsible = "cedric";
    $scope.helicopter = "Tigre";
    $scope.assembly = "Assembly75";


    $scope.doClickCreateParts = function (form) {
        if (form.$valid) {

            let createUri = "/blockchain/logcard/parts";
            var data = {
                "pn": $scope.partNumber,
                "sn": $scope.SerialNumber,
                "partName": $scope.partName,
                "type": $scope.type = "Defense",
                "resposible": $scope.responsible,
                "helicopter": $scope.helicopter,
                "assembly": $scope.assembly
            };
            var config = {
                headers: {
                    'Content-Type': 'application/x-www-form-urlencoded;charset=utf-8;',
                    'value': "Bearer " + userService.getToken()
                }
            }

            $http.post(createUri, data, config)
                .then(
                    function (response) {
                        $scope.answer = response.data;
                        $scope.status = response.status;
                        userService.setState(true);
                        userService.setToken(response.data);
                        userService.setUser($scope.userName);
                    },
                    function (response) {
                        $scope.answer = response.data || 'Request failed';
                        $scope.status = response.status;
                    }
                );
            $location.path('/home');
        }
    }
});
