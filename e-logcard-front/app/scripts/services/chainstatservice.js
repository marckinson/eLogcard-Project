'use strict';

/**
 * @ngdoc service
 * @name eLogcardFrontApp.chainStatService
 * @description
 * # chainStatService
 * Service in the eLogcardFrontApp.
 */
angular.module('eLogcardFrontApp')
  .service('chainStatService', ['$http','$q',function ($http,$q) {
		var self=this;
		var self = this;
        self.baseproxyUri = "/blockchain/stats/";
		
		self.findBlocksSize=function(){
			var defered = $q.defer();
			$http.get(self.baseproxyUri+'blocks').then(function (response) {
				// The then function here is an opportunity to modify the response
				console.log(response);
				// The return value gets picked up by the then in the controller.
				return defered.resolve(response.data.height);
			}); 
			return defered.promise;
		}
  }]);
