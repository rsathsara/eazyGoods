function request_handler(request) {
    var response;
    $.ajax({
        url: request.url,
        type: request.method,
        async: false,
        dataType: "json",
        data: request.data,
        success: function (data) {
            response = data;
        },
        error: function(error){
            ajaxErrorAlert(error);
        }
    });
    return response;
}

//Views
$(".viewProfile").click(function () { $("#pageContentLoader").load("/userProfile") });
$(".viewChangePswd").click(function () { $("#pageContentLoader").load("/changePassword") });
$(".viewGrnForm").click(function () { $("#pageContentLoader").load("/grnFormPage") });
$(".viewBillingForm").click(function () { $("#pageContentLoader").load("/billingFormPage") });



