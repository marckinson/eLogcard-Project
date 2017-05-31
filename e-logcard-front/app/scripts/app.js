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
                controllerAs: 'loginCtrl'
            })
            .when('/addpart', {
                templateUrl: 'views/addParts.html',
                controller: 'addPartCtrl',
                controllerAs: 'addPartCtrl'
            })
            .when('/showparts', {
                templateUrl: 'views/showParts.html',
                controller: 'showPartsCtrl',
                controllerAs: 'showPartsCtrl'
            })
            .when('/showpartsdemo', {
                templateUrl: 'views/showPartsdemo.html',
                controller: 'showPartsCtrl',
                controllerAs: 'showPartsCtrl'
            })
            .when('/showpartlog', {
                templateUrl: 'views/showPartLog.html',
                controller: 'showPartlogCtrl',
                controllerAs: 'showPartlogCtrl'
            })
            .when('/showpartlog/:partId', {
                templateUrl: 'views/showPartLog.html',
                controller: 'showPartlogCtrl',
                controllerAs: 'showPartlogCtrl'
            })
            .when('/header', {
                templateUrl: 'includes/header.html',
                controller: 'headerCtrl',
                controllerAs: 'headerCtrl'
            })
            .otherwise({
                redirectTo: '/home'
            });
    });
