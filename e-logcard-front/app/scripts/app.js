'use strict';

/**
 * @ngdoc overview
 * @name eLogcardFrontApp
 * @description
 * # eLogcardFrontApp
 * Main module of the application.
 */
var app = angular
    .module('eLogcardFrontApp', [
    'ngResource',
    'ngRoute',
  ])
    .config(function ($routeProvider) {
        $routeProvider
            .when('/home', {
                templateUrl: 'views/loginBS.html',
                controller: 'loginCtrl',
                controllerAs: 'loginCtrl'
            })
            /* PARTS */
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
            /* ASSEMBLIES */
            .when('/addassembly', {
                templateUrl: 'views/addAssembly.html',
                controller: 'addAssemblyCtrl',
                controllerAs: 'addAssemblyCtrl'
            })
            .when('/showassemblies', {
                templateUrl: 'views/showAssemblys.html',
                controller: 'showAssemblysCtrl',
                controllerAs: 'showAssemblysCtrl'
            })

            /* AIRCRAFTS */
            .when('/addaircraft', {
                templateUrl: 'views/addAircraft.html',
                controller: 'addAircraftCtrl',
                controllerAs: 'addAircraftCtrl'
            })
            .when('/showaircrafts', {
                templateUrl: 'views/showAircrafts.html',
                controller: 'showAirCraftsCtrl',
                controllerAs: 'showAirCraftsCtrl'
            })
            /*ALL*/
            .when('/transfer/:itemtype/:itemid', {
                templateUrl: 'views/transfer.html',
                controller: 'transferCtrl',
                controllerAs: 'transferCtrl'
            })

            .when('/showLogs/:itemtype/:itemid', {
                templateUrl: 'views/showLogs.html',
                controller: 'showLogsCtrl',
                controllerAs: 'showLogsCtrl'
            })
            .otherwise({
                redirectTo: '/home'
            });
    });
