$(function(){
	$.get("/user/connectedUser",function(data){
		$('#connectedUserName').text(data.user);
	});
});