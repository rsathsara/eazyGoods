function request_handler(request) {
    var response;
    $.ajax({
        url: request.url,
        type: request.method,
        async: false,
        dataType: "json",
        data: request.data,
        success: function (data) {
            if (data.status != 200){
                response = {status: false, body: ""};
                ajaxErrorAlert(data.body);
            } else{
                response = {status: true, body: data.body};
            }
        },
        error: function(error){
            ajaxErrorAlert(error);
            response = {status: false, body: ""};
        }
    });
    return response;
}

//Views
$(".viewProfile").click(function () { $("#pageContentLoader").load("/userProfile") });
$(".viewChangePswd").click(function () { $("#pageContentLoader").load("/changePassword") });
$(".viewGrnForm").click(function () { $("#pageContentLoader").load("/grnFormPage") });
$(".viewBillingForm").click(function () { $("#pageContentLoader").load("/billingFormPage") });



