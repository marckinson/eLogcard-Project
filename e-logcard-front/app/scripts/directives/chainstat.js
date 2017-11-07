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
			self.blocks=[];
			$rootScope.$on('userLogin',function(){
				refreshblocks();
			});
			
			$rootScope.$on('chain-modification',function(){
				refreshblocks();
			});
			
		};
		
		let refreshblocks = function(){
			
			if(userService.state){
				chainStatService.findBlocksSize().then(function (response){
					let blockSize=response;		
					for(var i=Math.max(0,self.blocks.length-1);i<blockSize;i++){
						let block={id:i};
						self.blocks.push(block);
					}
				});
				
			}
		};
		
		self.onBlockClick=function(block){
			self.selectedBlock=null;
			self.blockDetail=null;
			chainStatService.findBlockDetail(block.id).then(function (response){
				self.selectedBlock=response;
				self.selectedBlock.id=block.id;
				//$('#myModal').modal('show');
			});
		};
		init();
      }
    };
  }]);
