'use strict';

/**
 * @ngdoc function
 * @name eLogcardFrontApp.controller:transferCtrl
 * @description
 * # transferCtrl
 * Controller of the eLogcardFrontApp
 */
app.controller('transferCtrl', function ($location, $http, $routeParams, userService) {
    this.awesomeThings = [
      'HTML5 Boilerplate',
      'AngularJS',
      'Karma'
    ];

    var self = this;
    this.ownerMode = true;
    this.itemId = $routeParams.itemid;
    this.itemType = $routeParams.itemtype;
    this.transferTarget;
    this.answer;
    this.data = {};

    this.listTypeTransfert = ["Owner", "Responsible"];

    this.typeTransfers = {
        "parts": {
            "Owner": {
                "label": "Owner",
                "value": "owner",
                "request": "OwnerTransfer"
            },
            "Responsible": {
                "label": "Responsible",
                "value": "responsible",
                "request": "RespoTransfer"
            }
        }
    };

    this.requestTransfer = [{
        "Parts": [{
            "Onwer": "OwnerTransfer",
            "Resposible": "RespoTransfer"
        }]
    }];

    this.transferType = this.listTypeTransfert[0];


    // console.log("test: " + self.transfertType);


    this.doClickSendTransfert = function (form) {
        switch (self.transferType) {
            case self.listTypeTransfert[0]:
                // owner
                self.data = {
                    "owner": self.transferTarget
                };
                console.log("data owner ")

                break;

            case self.listTypeTransfert[1]:
                // resposible
                self.data = {
                    "responsible": self.transferTarget
                };
                console.log("data responsible ")


                break;
        }
        console.log(self.data);

        if (form.$valid) {
            console.log(self.transferTarget);
            console.log(self.itemType);
            console.log(self.transferType);
            console.log(self.itemId);

            let transferUri = "/blockchain/logcard/" + self.itemType + "/" + self.typeTransfers[self.itemType][self.transferType].request + "/" + self.itemId;
            console.log(transferUri);

            if (userService.getState()) {
                $http.put(transferUri, self.data)
                    .then(
                        function (response) {
                            self.answer = response.data;
                            self.status = response.status;
                            console.log(response);
                            console.log(response.status);
                            console.log(response.data);
                        },
                        function (response) {
                            self.answer = response.data || 'Request failed';
                        }
                    );

            }
        }
    }


    ///logcard/parts/RespoTransfer/c080b720-4a8b-11e7-9968-836ae361e7eb
    ///logcard/parts/OwnerTransfer/c080b720-4a8b-11e7-9968-836ae361e7eb
    //  let showPartlogUriWitoutParameter = "logcard/parts/RespoTransfer/c080b720-4a8b-11e7-9968-836ae361e7eb";
    // let transferUri = "/blockchain/logcard/" + self.itemType + "/" + /self.typeTransfers[self.itemType][self.transferType].request + "/" + self.itemId;
    // console.log(transferUri);
    // ne fonctionne pas pour l instan 
    // ng-click="transferCtrl.doClickRadioOwner()" 
    // let showPartlogUriIdParameter = showPartlogUriWitoutParameter + this.partId;
    /* console.log(transferUri);

        if (userService.getState()) {
            $http.put(transferUri)
                .then(
                    function (response) {
                        self.part = response.data;
                        self.status = response.status;
                        console.log(response);
                        console.log(response.status);
                        console.log(response.data);
                    },
                    function (response) {
                        self.answer = response.data || 'Request failed';
                    }
                );
        }
    };*/

});
