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
    this.debug = false;
    this.faillureRequest = false;

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

    //URI model
    ///logcard/parts/RespoTransfer/c080b720-4a8b-11e7-9968-836ae361e7eb
    ///logcard/parts/OwnerTransfer/c080b720-4a8b-11e7-9968-836ae361e7eb

    this.doClickSendTransfert = function (form) {
        switch (self.transferType) {
            case self.listTypeTransfert[0]:
                // owner
                self.data = {
                    "owner": self.transferTarget
                };
                if (self.debug)
                    console.log("data owner ")

                break;

            case self.listTypeTransfert[1]:
                // resposible
                self.data = {
                    "responsible": self.transferTarget
                };
                if (self.debug)
                    console.log("data responsible ")


                break;
        }
        if (self.debug)
            console.log(self.data);

        if (form.$valid) {
            if (self.debug) {
                console.log(self.transferTarget);
                console.log(self.itemType);
                console.log(self.transferType);
                console.log(self.itemId);
            }

            let transferUri = "/blockchain/logcard/" + self.itemType + "/" + self.typeTransfers[self.itemType][self.transferType].request + "/" + self.itemId;
            if (self.debug)
                console.log(transferUri);

            if (userService.getState()) {
                $http.put(transferUri, self.data)
                    .then(
                        function (response) {
                            self.answer = response.data;
                            self.status = response.status;
                            if (self.debug) {
                                console.log(response);
                                console.log(response.status);
                                console.log(response.data);
                            }
                            if (self.ownerMode == true)
                                $location.path('/showparts');
                            else {
                                $location.path('/showpartlog/' + self.itemId);
                            }
                        },
                        function (response) {
                            self.answer = response.data || 'Request failed';
                            self.faillureRequest = true;
                        }
                    );
            }
        }
    }

});
