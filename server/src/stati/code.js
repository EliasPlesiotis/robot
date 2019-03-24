var side = angular.module("side", []);

side.controller("ctrl", function($scope, $http) {

    $http({
        method: "POST",
        url: "http://192.168.1.25:5000/connect"
    })

    $scope.run = () => {
        $http({
            method: "GET",
            url: "http://192.168.1.25:5000/code"
        })
    }
});