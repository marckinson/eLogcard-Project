'use strict';

/**
 * @ngdoc function
 * @name eLogcardFrontApp.controller:GenerateqrcodeCtrl
 * @description
 * # GenerateqrcodeCtrl
 * Controller of the eLogcardFrontApp
 */
angular.module('eLogcardFrontApp')
    .controller('GenerateqrcodeCtrl', function ($routeParams) {
        this.awesomeThings = [
      'HTML5 Boilerplate',
      'AngularJS',
      'Karma'
    ];
        var self = this;
        this.debug = false;
        this.itemId = $routeParams.itemid;
        this.itemType = $routeParams.itemtype;

        var typeNumber = 4;
        var errorCorrectionLevel = 'L';
        var qr = qrcode(typeNumber, errorCorrectionLevel);
        qr.addData('showparts/'+ this.itemId);
        qr.make();
        document.getElementById('placeHolder').innerHTML = qr.createImgTag(10, 10);

    });
