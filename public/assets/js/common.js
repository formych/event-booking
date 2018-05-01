$("body").on("click", "#sign_button", function(event) {
	event.preventDefault();
	let formData = new FormData(document.getElementById('signup-form'));
    $.ajax({
        url : document.getElementById('signup-form').getAttribute("action"), 
        type : "POST", 
        data: formData,
        processData : false, 
        contentType : false,
        success : function(data) { 
            if(data.code === 200){
            	if (data.token !== undefined) {
                    console.log(data)
                    localStorage["token"] = data.token
                    localStorage["name"] = data.name
 					window.location.href='/index.html'
 					return
            	}
            	window.location.href='/login.html';
            } else{
                alert(data.msg);
            }
        }, 
        error : function(data) { 
            console.log(data);
        } 
    })
})