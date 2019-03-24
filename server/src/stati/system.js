var side = angular.module("side", []);

side.controller("ctrl", function($scope, $http) {

   $scope.saved = true
    
   $scope.loaded = true

   $http({
      method: "GET",
      url: "http://localhost:8080/folder/{all}"
    }).then(
      (response) => {
        $scope.files = response.data
        console.log("ok")},
      () => {console.log("error")}
    )

   $scope.save = () => {
      $http({ method : "POST", url : "http://localhost:8080/folder/{" + $scope.name + "}"})
      $http({
         method: "GET",
         url: "http://localhost:8080/folder/{all}"
       }).then(
         (response) => {
           $scope.files = response.data
           console.log("ok")},
         () => {console.log("error")}
       )   
      $scope.saved = false
   }

   $scope.load = (m) => {
      $http({ method : "GET", url : "http://localhost:8080/folder/{" + m + "}"})
      $http({
         method: "GET",
         url: "http://localhost:8080/folder/{all}"
       }).then(
         (response) => {
           $scope.files = response.data
           console.log("ok")},
         () => {console.log("error")}
       )
   
      $scope.loaded = false
   }

   $scope.delete = (m) => {
      $http({ method : "DELETE", url : "http://localhost:8080/folder/{" + m + "}"})
      $http({
         method: "GET",
         url: "http://localhost:8080/folder/{all}"
       }).then(
         (response) => {
           $scope.files = response.data
           console.log("ok")},
         () => {console.log("error")}
       )   
   }
     
})