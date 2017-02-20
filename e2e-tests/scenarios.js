'use strict';

/* https://github.com/angular/protractor/blob/master/docs/toc.md */

describe('chanteyBase', function() {


  it('should automatically redirect to /songs when location hash/fragment is empty', function() {
    browser.get('index.html');
    expect(browser.getLocationAbsUrl()).toMatch("/songs");
  });


  describe('view1', function() {

    beforeEach(function() {
      browser.get('index.html#!/songs');
    });


    it('should render view1 when user navigates to /songs', function() {
      expect(element.all(by.css('[ng-view] p')).first().getText()).
        toMatch(/partial for view 1/);
    });

  });


  describe('view2', function() {

    beforeEach(function() {
      browser.get('index.html#!/collections');
    });


    it('should render view2 when user navigates to /collections', function() {
      expect(element.all(by.css('[ng-view] p')).first().getText()).
        toMatch(/partial for view 2/);
    });

  });
});
