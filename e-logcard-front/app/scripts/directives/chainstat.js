'use strict';

/**
 * @ngdoc directive
 * @name eLogcardFrontApp.directive:chainstat
 * @description
 * # chainstat
 */
angular.module('eLogcardFrontApp')
  .directive('chainstat', ['chainStatService','userService','$rootScope','$timeout',function (chainStatService,userService,$rootScope,$timeout) {
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
			self.showBlock=false;
			self.blocks=[];
			 $timeout(refreshblocks, 10000);
		};
		
		let refreshblocks = function(){
			
			if(userService.state){
				chainStatService.findBlocksSize().then(function (response){
					let blockSize=response;	
					let existingSize=self.blocks.length;
					
					if(blockSize!=existingSize){
						
						let startIndex=existingSize!=0 && blockSize>existingSize?existingSize:0;
						let endIndex=blockSize;
						for(var i=startIndex;i<endIndex;i++){
							let block={id:i};
							self.blocks.push(block);
						}
					
					}	
					self.showBlock=true;	
				});
				$timeout(refreshblocks, 10000);
			}
			else{
				self.showBlock=false;
			}
			 
		};
		
		self.onBlockClick=function(block){
			self.selectedBlock=null;
			self.blockDetail=null;
			chainStatService.findBlockDetail(block.id).then(function (response){
				self.selectedBlock=response;
				self.selectedBlock.id=block.id;
			});
		};
		init();
      }
    };
  }]);
