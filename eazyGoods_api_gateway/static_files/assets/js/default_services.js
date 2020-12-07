const gateway = "eazyGoods_api/";
const mainService = "main_service/";

function get_item_list(){
    var response = request_handler({url: gateway + mainService + "itemList", method: "GET" , data: {}});
    return response;
}

function getBillNo(){

}

function saveBill(){

}

function editBill(){

}

