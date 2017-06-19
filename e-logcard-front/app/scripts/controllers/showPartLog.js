'use strict';

/**
 * @ngdoc function
 * @name eLogcardFrontApp.controller:showPartlogCtrl
 * @description
 * # showPartlogCtrl
 * Controller of the eLogcardFrontApp
 */
app.controller('showPartlogCtrl', function ($http, $routeParams, userService) {
    this.awesomeThings = [
      'HTML5 Boilerplate',
      'AngularJS',
      'Karma'
    ];
    var self = this;
    this.partId = $routeParams.partId;
    this.debug = false;
    /*
        this.part = {
            "pn": "Wffieng",
            "sn": "1024",
            "id": "x02048",
            "partName": "Wing",
            "type": "defence",
            "responsible": "lucas",
            "owner": "florent",
            "helicopter": "tigre",
            "assembly": "3667 ",
            "logs": [{
                "log_type": "CREATE",
                "vDate": "2017/05/31 11:14:25",
                "owner": "lucas",
                "responsible": "cedric",
                "modType": "type",
                "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Suspendisse commodo nisi tortor, in blandit nunc vestibulum nec. Donec non ex accumsan, lobortis magna ut, sollicitudin purus. Suspendisse id turpis magna. Proin nec lacus tellus. Cras dui turpis, consequat eget risus a, scelerisque aliquet erat. In hac habitasse platea dictumst. Praesent ornare laoreet egestas. Duis consectetur suscipit felis non cursus. Sed elit risus, iaculis facilisis risus vitae, dictum varius tellus. Suspendisse potenti. Suspendisse potenti. Proin congue nec nulla ut faucibus. Vestibulum feugiat massa a risus bibendum, eget porta libero commodo. "
                   }]
        };*/




    let showPartlogUriWitoutParameter = "/blockchain/logcard/parts/historic/";
    if (this.debug)
        console.log(this.partId);
    let showPartlogUriIdParameter = showPartlogUriWitoutParameter + this.partId;
    if (this.debug)
        console.log(showPartlogUriIdParameter);
    if (userService.getState()) {
        $http.get(showPartlogUriIdParameter)
            .then(
                function (response) {
                    self.part = response.data;
                    self.status = response.status;
                    if (self.debug) {
                        console.log(response);
                        console.log(response.status);
                        console.log(response.data);
                    }
                },
                function (response) {
                    self.answer = response.data || 'Request failed';
                }
            );
    }


});
