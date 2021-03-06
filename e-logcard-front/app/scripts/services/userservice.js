'use strict';

/**
 * @ngdoc service
 * @name eLogcardFrontApp.userService
 * @description
 * # userService
 * Service in the eLogcardFrontApp.
 */
app.service('userService', ['$http', function ($http) {

    this.user = '';
    this.role = '';
    this.token = '';
    this.state = false;

    this.getUser = function () {
        return this.user;
    };

    this.getRole = function () {
        return this.role;
    };

    this.getToken = function () {
        return this.token;
    };

    this.getState = function () {
        return this.state;
    };

    this.setUser = function (user) {
        this.user = user;
    };

    this.setState = function (state) {
        this.state = state;
    };

    this.setRole = function (role) {
        this.role = role;
    };

    this.setToken = function (token) {
        this.token = token;
        // a conserver temporairement tant que les http sont present dans les controler 
        $http.defaults.headers.common.Authorization = 'Bearer ' + token;
    };

    this.disconnectUser = function () {
        this.user = '';
        this.role = '';
        this.token = '';
        this.state = false;
    };

}]);
