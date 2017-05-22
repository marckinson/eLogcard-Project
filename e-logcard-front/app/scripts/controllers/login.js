'use strict';

/**
 * @ngdoc function
 * @name fiveAppApp.controller:LoginCtrl
 * @description
 * # LoginCtrl
 * Controller of the fiveAppApp
 */
app.controller("loginCtrl", function ($scope, $http, $location, userService) {
    this.awesomeThings = [
      'HTML5 Boilerplate',
      'AngularJS',
      'Karma'
    ];
    // init valeur par defaut de la combo box ;
    //variable permetant affichage du volet login ou sign up 
    $scope.newUser = false;
    // gestion evenement pour changer de volet login ou sign up 
    $scope.doClickLogin = function (event) {
        $scope.newUser = false;
    }

    $scope.doClickSignUp = function (event) {
        $scope.newUser = true;
    }
    // init valeur par defaut du role 
    $scope.role = "Supplier";

    let rolesUri = "/blockchain/roles";
    // doit etre reactive de que la fonction role 
    // est corrige
    /*$http.get(rolesUri)
        .then(
            function (response) {
                $scope.roles = response.data;
            },
            function (response) {
                $scope.answer = response.data || 'Request failed';
    */
    $scope.roles = [
        {
            label: 'Auditor authority',
            value: 'Auditor_authority'
                    },
        {
            label: 'AH admin',
            value: 'AH_admin',
                    },
        {
            label: 'Supplier',
            value: 'Supplier'
                    },
        {
            label: 'Manufacturer',
            value: 'Manufacturer'
                    },
        {
            label: 'Customer',
            value: 'Customer'
                    },
        {
            label: 'Maintenance user',
            value: 'Maintenance_user'
                    }
                ];
    /*
        }
    );*/

    $scope.doClickCreateUser = function (form) {
        if (form.$valid) {
            let registrationUri = "/blockchain/registration";

            var data = {
                "username": $scope.userName,
                "password": $scope.password,
                "role": $scope.role
            };

            $http.post(registrationUri, data)
                .then(
                    function (response) {
                        $scope.answer = response.data;
                        $scope.status = response.status;
                        userService.setState(true);
                        userService.setToken(response.data);
                        console.log("user :" + $scope.userName);
                        userService.setUser($scope.userName);
                        console.log("role :" + $scope.role);
                    },
                    function (response) {
                        $scope.answer = response.data || 'Request failed';
                        $scope.status = response.status;
                        userService.clearValues;
                    }
                );
            $location.path('/home');
        }
    }

    $scope.doClickConnectUser = function (form) {
        if (form.$valid) {

            let loginUri = "/blockchain/login";

            var data = {
                "username": $scope.userName,
                "password": $scope.password
            };

            $http.post(loginUri, data)
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
                        userService.clearValues;
                    }
                );
            $location.path('/home');
        }
    }
});
