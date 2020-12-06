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
$(".viewGrnList").click(function () { $("#pageContentLoader").load("/grnListPage") });
$(".viewBillingList").click(function () { $("#pageContentLoader").load("/billingListPage") });

