const gateway = "eazyGoods_api/";
const services = {mainService: "main/"};

function get_item_list(){
    var response = request_handler({url: gateway + services.mainService + "itemList", method: "GET" , data: {}});
    return response;
}

function getBillNo(){

}

function saveBill(){

}

function editBill(){

}

