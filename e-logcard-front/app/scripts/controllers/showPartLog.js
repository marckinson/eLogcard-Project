'use strict';

/**
 * @ngdoc function
 * @name eLogcardFrontApp.controller:showPartlogCtrl
 * @description
 * # addPartCtrl
 * Controller of the eLogcardFrontApp
 */
app.controller('showPartlogCtrl', function ($location, $http, $routeParams, userService) {
    this.awesomeThings = [
      'HTML5 Boilerplate',
      'AngularJS',
      'Karma'
    ];
    var self = this;
    this.partId = "8a9ccee0-45e1-11e7-acc6-f3c56d6be63d"; // $routeParams.partId;
    /*this.name = "test";
    this.part = {
        "pn": "Wffieng",
        "sn": "1024",
        "id": "x02048",
        "partName": "Wing",
        "type": "defence",
        "responsible": "sora",
        "owner": "florent",
        "helicopter": "tigre",
        "assembly": "3667 ",
        "logs": [{
            "log_type": "CREATE",
            "vDate": "2017/05/31 11:14:25",
            "owner": "lucas",
            "responsible": "cedric",
            "modType": "",
            "description": ""
           }]
    };*/


    //let showPartlogUriWitoutParameter = "/blockchain/logcard/parts/historic?id=";
    var showPartlogUriWitoutParameter = "/blockchain/logcard/parts/historic?id=";
    // ne fonctionne pas pour l instan 
    //let showPartlogUriIdParameter = showPartlogUriWitoutParameter + this.partId;
    var showPartlogUriIdParameter = showPartlogUriWitoutParameter + this.partId;
    console.log(showPartlogUriIdParameter);
    $http.get(showPartlogUriIdParameter)
        .then(
            function (response) {
                self.part = response.data;
                self.status = response.status;
                console.log(response);
                console.log(response.status);
            },
            function (response) {
                self.answer = response.data || 'Request failed';
            }
        );
    //}



});
