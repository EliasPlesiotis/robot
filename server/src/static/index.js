var side = angular.module("side", []);

side.controller("ctrl", function($scope, $http) {
  $http({
    method: "GET",
    url: "http://localhost:8080/files"
  }).then(
    (response) => {
      $scope.moves = response.data
      console.log("ok")},
    () => {console.log("error")}
  )

  $http({
    method: "POST",
    url: "http://127.0.0.1:5000/connect"
  })

  $scope.move = (m) => {
    if($scope.moves.length > 15) {
      return
    }

    if($scope.duration > 10 || $scope.duration < 1) {
      return
    }

    $http({ method : "POST", url : "http://localhost:8080/command/{0}+{"+ m +"}+{"+ $scope.duration +"}+{90}"})
    $http({
      method: "GET",
      url: "http://localhost:8080/files"
    }).then((response) => {
      $scope.moves = response.data;
      console.log("ok")},
    () => {console.log("error")}
    )
  }

  $scope.delete = (Id) => {
    $http({
      method: "DELETE",
      url: "http://localhost:8080/command/{"+ Id +"}+{left}+{4}+{90}"
    })

    $http({
      method: "GET",
      url: "http://localhost:8080/files"
    }).then((response) => {
      $scope.moves = response.data;console.log("ok")}, 
    () => {console.log("error")}
    )
  }

  $scope.start = () => {
    $http({
      method: "POST",
      url: "http://127.0.0.1:5000/start"
    })
  }

  $scope.stop = () => {
    $http({
      method: "POST",
      url: "http://127.0.0.1:5000/stop"
    })
  }

});
