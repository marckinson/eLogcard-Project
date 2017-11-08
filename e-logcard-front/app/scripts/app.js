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
	'angular-loading-bar', 'ngAnimate'
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
            .when('/showparts/:itemid', {
                templateUrl: 'views/showOnePart.html',
                controller: 'showOnePartCtrl',
                controllerAs: 'showOnePartCtrl'
            })
			.when('/showparts', {
                templateUrl: 'views/showParts.html',
                controller: 'showPartsCtrl',
                controllerAs: 'showPartsCtrl'
            })
            // utiliser pour les logs a voir plus tard si on ne peux pas 
            // mutualiser avec showlogs 
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
			 .when('/showassemblies/:itemid', {
                templateUrl: 'views/showOneAssembly.html',
                controller: 'showOneAssemblyCtrl',
                controllerAs: 'showOneAssemblyCtrl'
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
			.when('/showaircrafts/:itemid', {
                templateUrl: 'views/showOneAircraft.html',
                controller: 'showOneAirCraftCtrl',
                controllerAs: 'showOneAirCraftCtrl'
            })
            .when('/showaircraftlistingpart/:itemid', {
                templateUrl: 'views/showAirCraftListingParts.html',
                controller: 'showAircraftListingPartsCtrl',
                controllerAs: 'showAircraftListingPartsCtrl'
            })
            .when('/aircraft/showassemblies/:itemid', {
                templateUrl: 'views/showAssemblys.html',
                controller: 'ShowAirCraftAssembliesCtrl',
                controllerAs: 'showAssemblysCtrl'
            })
            /*ALL*/
            .when('/transfer/:itemtype/:itemid', {
                templateUrl: 'views/transfer.html',
                controller: 'transferCtrl',
                controllerAs: 'transferCtrl'
            })
            .when('/showlogs/:itemtype/:itemid', {
                templateUrl: 'views/showLogs.html',
                controller: 'showLogsCtrl',
                controllerAs: 'showLogsCtrl'
            })
            .when('/addlog/:itemtype/:itemid', {
                templateUrl: 'views/addLogs.html',
                controller: 'addlogsCtrl',
                controllerAs: 'addlogsCtrl'
            })
            .when('/showpartlist/:itemtype/:itemid', {
                templateUrl: 'views/showPartListing.html',
                controller: 'ShowPartListingCtrl',
                controllerAs: 'ShowPartListingCtrl'
            })
            .when('/attachpart/:itemtype/:itemid', {
                templateUrl: 'views/attachPart.html',
                controller: 'AttachpartCtrl',
                controllerAs: 'attachCtrl'
            })
            .when('/attachAssembly/:itemtype/:itemid', {
                templateUrl: 'views/attachPart.html',
                controller: 'AttachassemblyCtrl',
                controllerAs: 'attachCtrl'
            })
            .when('/replace/:container/:containerid/:item/:itemid', {
                templateUrl: 'views/replace.html',
                controller: 'ReplaceCtrl',
                controllerAs: 'replaceCtrl'
            })
            .when('/scanQrCode', {
                templateUrl: 'views/scanqrcode.html',
                controller: 'ScanqrcodeCtrl',
                controllerAs: 'scanQrcode'
            })
            .when('/generateqrcode/:itemtype/:itemid', {
                templateUrl: 'views/generateqrcode.html',
                controller: 'GenerateqrcodeCtrl',
                controllerAs: 'GenerateQR'
            })
            .otherwise({
                redirectTo: '/home'
            });
    });
