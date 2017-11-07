'use strict';

describe('Service: chainStatService', function () {

  // load the service's module
  beforeEach(module('eLogcardFrontApp'));

  // instantiate service
  var chainStatService;
  beforeEach(inject(function (_chainStatService_) {
    chainStatService = _chainStatService_;
  }));

  it('should do something', function () {
    expect(!!chainStatService).toBe(true);
  });

});
