'use strict';

/**
 * @ngdoc overview
 * @name eLogcardFrontApp
 * @description
 * # eLogcardFrontApp
 *
 * Main module of the application.
 */
var app = angular
    .module('eLogcardFrontApp', [
    'ngResource',
    'ngRoute'
  ])
    .config(function ($routeProvider) {
        $routeProvider
            .when('/home', {
                templateUrl: 'views/loginBS.html',
                controller: 'loginCtrl',
                controllerAs: 'login'
            })
            .when('/addpart', {
                templateUrl: 'views/addParts.html',
                controller: 'addPartCtrl',
                controllerAs: 'addPart'
            })
            .otherwise({
                redirectTo: '/home'
            });
    });

/* service de gestion des connexion utilisateur */
app.service('userService', function () {

    this.user = '';
    this.role = '';
    this.token = '';
    this.state = false;

    this.getUser = function () {
        return this.user;
    }

    this.getRole = function () {
        return this.role;
    }

    this.getToken = function () {
        return this.token;
    }

    this.getState = function () {
        return this.state;
    }

    this.setUser = function (user) {
        this.user = user;
    }

    this.setState = function (state) {
        this.state = state;
    }

    this.setRole = function (role) {
        this.role = role;
    }

    this.setToken = function (token) {
        this.token = token;
    }

});
