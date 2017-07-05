'use strict';

/**
 * @ngdoc function
 * @name eLogcardFrontApp.controller:LoginCtrl
 * @description
 * # LoginCtrl
 * Controller of the eLogcardFrontApp
 */
app.controller("loginCtrl", function ($location, eLogcardService) {
    this.awesomeThings = [
      'HTML5 Boilerplate',
      'AngularJS',
      'Karma'
    ];
    // debug mode
    this.debug = false;
    //variable permetant affichage du volet login ou sign up 
    var self = this;
    this.newUser = false;
    this.role = "";
    this.roles;
    this.faillureRequest = false;
    this.faillureRolesRequest = false;
    this.userName;
    this.password;
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

    // recuperation des roles utilisateur via aux service Elocard
    eLogcardService.getUserRoles().then(function (rolesRequest) {

            self.roles = rolesRequest.roles;
            self.answer = rolesRequest.aswer;
            // affecte la  valeur par defaut 2 Supplier 
            self.role = self.roles[2].value;
        },
        function (error) {
            // permet d afficher que le requet role a echoue 
            self.faillureRolesRequest = true;
        });

    this.doClickCreateUser = function (form) {
        if (self.passwordVerify == self.password) {
            if (form.$valid) {

                eLogcardService.subscribe(self.userName, self.password, self.role)
                    .then(
                        function (response) {
                            self.answer = response.data;
                            self.status = response.status;
                            $location.path('/showparts');
                        },
                        function (error) {
                            self.answer = error.data;
                            self.status = error.status;
                            self.faillureRequest = true;
                        });
            }
        }
    }

    this.doClickConnectUser = function (form) {
        if (form.$valid) {
            eLogcardService.login(self.userName, self.password)
                .then(function (response) {

                        self.answer = response.data;
                        self.status = response.status;
                        if (self.debug)
                            console.log("connexion succes ");

                        $location.path('/showparts');

                    },
                    function (error) {
                        self.answer = error.data || 'Request failed';
                        self.status = error.status;
                        self.faillureRequest = true;

                    })
        }
    }
});
