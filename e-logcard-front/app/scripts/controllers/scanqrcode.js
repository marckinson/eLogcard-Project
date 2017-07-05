'use strict';

/**
 * @ngdoc function
 * @name eLogcardFrontApp.controller:ScanqrcodeCtrl
 * @description
 * # ScanqrcodeCtrl
 * Controller of the eLogcardFrontApp
 */
angular.module('eLogcardFrontApp')
    .controller('ScanqrcodeCtrl', function ($location) {
        this.awesomeThings = [
      'HTML5 Boilerplate',
      'AngularJS',
      'Karma'
    ];

        var self = this;
        this.debug = true;
        this.Scan = false;
        this.QRcameras;
        this.QRcamera;

        this.closeCamera = function () {
            self.Scan = false;
            if (self.debug) {
                console.log("fin du scan ");
            }
            if (self.QRcameras.length > 0) {
                self.scanner.stop(self.QRcamera);
            } else {
                console.error('No cameras found.');
            }

        };


        Instascan.Camera.getCameras().then(function (cameras) {
            if (cameras.length > 0) {
                self.QRcameras = cameras;
                self.QRcamera = cameras[0];
            } else {
                console.error('No cameras found.');
            }
        }).catch(function (e) {
            console.error(e);
        });




        this.scanner = new Instascan.Scanner({
            video: document.getElementById('preview')
        });

        this.scanner.addListener('active', function (content) {
            if (self.debug) {
                console.log("camera passe active");
            }
        });
        this.scanner.addListener('scan', function (content) {
            if (self.debug) {
                console.log(content);
            }
            var move = true;
            move = confirm("Are you sure you want go on " + content + " ?");

            if (move == true) {
                // self.closeCamera();
                $location.path(content);



            }
        });
        this.scanner.addListener('inactive', function (content) {
            if (self.debug) {
                console.log("camera passe inactive");
            }
        });

        this.start = function () {
            self.Scan = true;
            if (self.debug) {
                console.log("fin du scan ");
            }
            if (self.QRcameras.length > 0) {
                self.scanner.start(self.QRcamera)
            } else {
                console.error('No cameras found.');
            }
        };

        this.stop = function () {
            self.closeCamera();
        }


    });
