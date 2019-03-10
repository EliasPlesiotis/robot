var side = angular.module("side", []);

side.controller("ctrl", function($scope, $http) {

    $scope.saved = true
    
    $scope.loaded = true


    $scope.save = () => {
       $http({ method : "POST", url : "http://localhost:8080/folder/{" + $scope.name + "}"})
       $scope.saved = false
    }

    $scope.load = (m) => {
        $http({ method : "GET", url : "http://localhost:8080/folder/{" + m + "}"})
        $scope.loaded = false
     }
     
})