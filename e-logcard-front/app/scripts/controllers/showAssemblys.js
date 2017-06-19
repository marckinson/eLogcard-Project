'use strict';

/**
 * @ngdoc function
 * @name eLogcardFrontApp.controller:showAssemblysCtrl
 * @description
 * # showAssemblysCtrl
 * Controller of the eLogcardFrontApp
 */
app.controller('showAssemblysCtrl', function ($http, $location, userService, eLogcardService) {
    this.debug = false;
    this.answer;
    this.status;
    var self = this;
    this.showId = false;
    /*
        this.assemblies = [{
            "an": "a",
            "sn": "k",
            "id_assembly": "3de0a160-46a1-11e7-956e-cd3b1eedcf08",
            "owner": "sora",
            "parts": [{
                "pn": "Wffieng",
                "sn": "1024",
                "id": "x02048",
                "partName": "Wing",
                "type": "defence",
                "responsible": "sora",
                "owner": "florent",
                "helicopter": "tigre",
                "assembly": "3667 "
                          }],
            "logs": [{
                "log_type": "CREATE",
                "vDate": "2017/06/01 10:06:39",
                "owner": "sora",
                "responsible": "",
                "modType": "",
                "description": ""
                       }]
                    }];*/

    eLogcardService.getAssemblies().then(function (assembliesRequest) {
        self.assemblies = assembliesRequest.assemblies;
        self.status = assembliesRequest.satus;
    }, function (error) {
        self.answer = error.data || 'Request failed';
        self.status = error.status;
    });

    // gestion evenement  pour consulter les log d'une assembly
    this.doClickShowLogs = function (idAssembly) {

        let showLogsUri = "/showlogs/" + 'assemblies' + "/" + idAssembly;

        $location.path(showLogsUri);
        if (self.debug) {
            console.log(idAssembly);
            console.log(showLogsUri)
        }
    }

    // gestion evenement  pour consulter les log d'une assembly
    this.doClickShowParts = function (idAssembly) {

        let showPartsUri = "/showpartlist/" + 'assemblies' + "/" + idAssembly;

        $location.path(showPartsUri);
        if (self.debug) {
            console.log(idAssembly);
            console.log(showPartsUri)
        }
    }

    // gestion evenement  pour consulter les log d'une assembly
    this.doClickAddPart = function (idAssembly) {

        let attachPartsUri = "/attachpart/" + 'assembly' + "/" + idAssembly;
        $location.path(attachPartsUri);

        if (self.debug) {
            console.log(idAssembly);
            console.log(attachPartsUri)
        }
    }

});
