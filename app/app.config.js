
angular
  .module('chanteyBase')
  .config(['$locationProvider', '$routeProvider',
           function($locationProvider, $routeProvider) {
             $locationProvider.hashPrefix('!');
             
             $routeProvider
               .when('/songs', {
                 template: '<songs></songs>'
               })
               .when('/collections', {
                 template: '<collections></collections>'
               })
               .otherwise({redirectTo: '/songs'});
           }]);
