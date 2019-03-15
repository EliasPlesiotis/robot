var side = angular.module("side", []);

side.controller("ctrl", function($scope, $http) {

    $http({
        method: "POST",
        url: "http://192.168.1.25:5000/connect"
    })

    $scope.move = (m) => {
        $http({ 
            method : "POST", 
            url : "http://192.168.1.25:5000/controller/"+ m +"+"+ $scope.speed
        })
    }
    
    $scope.stop = () => {
        $http({
            method: "POST",
            url: "http://192.168.1.25:5000/controller/stop+"+ $scope.speed
        })
    }

});
