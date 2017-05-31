'use strict';

/**
 * @ngdoc function
 * @name eLogcardFrontApp.controller:LoginCtrl
 * @description
 * # LoginCtrl
 * Controller of the eLogcardFrontApp
 */
app.controller("loginCtrl", function ($http, $location, userService) {
    this.awesomeThings = [
      'HTML5 Boilerplate',
      'AngularJS',
      'Karma'
    ];
    //variable permetant affichage du volet login ou sign up 
    var self = this;
    this.newUser = false;
    this.role = "";
    this.roles;
    this.faillureRequest = false;
    this.faillureRolesRequest = false;
    this.userName;
    this.password;
    this.role;
    this.passwordVerify;

    // gestion evenement pour changer de volet login ou sign up 
    this.doClickLogin = function (event) {
        self.newUser = false;
        self.faillureRequest = false;
    }

    this.doClickSignUp = function (event) {
        self.newUser = true;
        self.faillureRequest = false;
    }
    let rolesUri = "/blockchain/roles";

    $http.get(rolesUri)
        .then(
            function (response) {
                self.roles = response.data;
                self.role = response.data[0].value;
            },
            function (response) {
                self.answer = response.data || 'Request failed';
                self.faillureRolesRequest = true;

            }
        );

    this.doClickCreateUser = function (form) {
        if (self.passwordVerify == self.password) {
            if (form.$valid) {
                let registrationUri = "/blockchain/registration";
                var data = {
                    "username": self.userName,
                    "password": self.password,
                    "role": self.role
                };

                $http.post(registrationUri, data)
                    .then(
                        function (response) {
                            self.answer = response.data;
                            self.status = response.status;
                            userService.setState(true);
                            userService.setToken(response.data);
                            userService.setUser(self.userName);
                            userService.setRole(self.role);

                            $location.path('/showparts');
                        },
                        function (response) {
                            self.answer = response.data || 'Request failed';
                            self.status = response.status;
                            userService.clearValues;
                            self.faillureRequest = true;
                        }
                    );
            }
        }
    }


    this.doClickConnectUser = function (form) {
        if (form.$valid) {

            let loginUri = "/blockchain/login";

            var data = {
                "username": self.userName,
                "password": self.password
            };

            $http.post(loginUri, data)
                .then(
                    function (response) {
                        self.answer = response.data;
                        self.status = response.status;
                        userService.setState(true);
                        userService.setToken(response.data.token);
                        userService.setUser(response.data.username);
                        userService.setRole(response.data.role);

                        $location.path('/showparts');
                    },
                    function (response) {
                        self.answer = response.data || 'Request failed';
                        self.status = response.status;
                        userService.clearValues;
                        self.faillureRequest = true;
                    }
                );
        }
    }
});
