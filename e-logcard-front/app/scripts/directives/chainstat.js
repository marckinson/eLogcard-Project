'use strict';

/**
 * @ngdoc directive
 * @name eLogcardFrontApp.directive:chainstat
 * @description
 * # chainstat
 */
angular.module('eLogcardFrontApp')
  .directive('chainstat', ['chainStatService','userService','$rootScope',function (chainStatService,userService,$rootScope) {
    return {
      templateUrl: 'views/template/chainstat.html',
      restrict: 'AE',
      link: function postLink(scope, element, attrs) {
		let self=scope;
        let init=function(){
			$rootScope.$on('userLogin',function(){
				refreshblocks();
			});
			
			$rootScope.$on('chain-modification',function(){
				refreshblocks();
			});
			
		};
		
		let refreshblocks = function(){
			self.blocks=[];
			if(userService.state){
				chainStatService.findBlocksSize().then(function (response){
					let blockSize=response;		
					for(var i=0;i<blockSize;i++){
						let block={id:i};
						self.blocks[i]=block;
					}
				});
				
			}
		};
		
		init();
      }
    };
  }]);
