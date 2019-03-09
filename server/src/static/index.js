var side = angular.module("side", []);

side.controller("ctrl", function($scope, $http) {
  $http({
    method: "GET",
    url: "http://localhost:8080/"
  }).then(
    (response) => {
      $scope.moves = response.data
      console.log("ok")},
    () => {console.log("error")}
  )

  $scope.move = (m) => {
    if($scope.moves.length > 15) {
      return
    }

    if($scope.duration > 10 || $scope.duration < 1) {
      return
    }

    $http({ method : "POST", url : "http://localhost:8080/command/{0}+{"+ m +"}+{"+ $scope.duration +"}"})
    $http({
      method: "GET",
      url: "http://localhost:8080/"
    }).then((response) => {
      $scope.moves = response.data;
      console.log("ok")},
    () => {console.log("error")}
    )
  }

  $scope.delete = (Id) => {
    $http({
      method: "DELETE",
      url: "http://localhost:8080/command/{"+ Id +"}+{left}+{4}"
    })

    $http({
      method: "GET",
      url: "http://localhost:8080/"
    }).then((response) => {
      $scope.moves = response.data;console.log("ok")}, 
    () => {console.log("error")}
    )
  }

  $scope.start = () => {
    $http({
      method: "POST",
      url: "http://localhost:8080/command/{0}+{}+{0}"
    })
  }

  $scope.stop = () => {
    $http({
      method: "POST",
      url: "http://localhost:8080/command/{1}+{forward}+{-4}"
    })
  }

});
