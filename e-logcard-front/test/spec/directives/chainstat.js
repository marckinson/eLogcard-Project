'use strict';

describe('Directive: chainstat', function () {

  // load the directive's module
  beforeEach(module('eLogcardFrontApp'));

  var element,
    scope;

  beforeEach(inject(function ($rootScope) {
    scope = $rootScope.$new();
  }));

  it('should make hidden element visible', inject(function ($compile) {
    element = angular.element('<chainstat></chainstat>');
    element = $compile(element)(scope);
    expect(element.text()).toBe('this is the chainstat directive');
  }));
});
