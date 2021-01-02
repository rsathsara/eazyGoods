const gateway = "eazyGoods_api/";
function getDefaultGateway(){
    return gateway
}

function request_handler(request) {
    var response;
    $.ajax({
        url: request.url,
        type: request.method,
        async: false,
        dataType: "json",
        data: request.data,
        success: function (data) {
            var body = JSON.parse(data.body)
            if (data.status != 200){
                response = {status: false, body: ""};
                ajaxErrorAlert(body);
            } else{
                response = {status: true, body: body};
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
$(".viewReportForm").click(function () { $("#pageContentLoader").load("/reportPage") });



