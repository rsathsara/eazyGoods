function billingForm() {
    //Select2 function for searchable dropdowns 
    $('.searchable-dropdown').select2({});

    // Initialize bill item array
    var billItems = [];

    //Onload events
    docDatePicker();
    newBillNo();
    billItemTable(billItems);
    loadCustomerDropdown();
    loadItemDropdown();
    // keyControl();

    // $(function(){
    //     $('#billingForm input[name="billTo"]').keypress(function (e) {
    //         console.log("bill :"+e.keyCode);
    //         if (e.keyCode == 9) {
    //             $('input[name="item"]').focus();
    //         }
    //     });
    //     $('#billingForm input[name="item"]').keypress(function (e) {
    //         console.log("item :"+e.keyCode);
    //         if (e.keyCode == 9) {
    //             $('input[name="itemQty"]').focus();
    //         }
    //     });
    // });

    //Add Item -> Select Item
    $("#item").on('change', function () {
        itemDetails($(this).val());
        $('input[name="itemQty"]').focus();
    });
    // Add Item -> Cal Item Value
    $('input[name="itemQty"], input[name="itemPrice"]').on('input', function () {
        calItemValue();
    });
    //Add Item -> Add item
    $('#addItem').on('click', function () {
        clearFormErrorMsg();
        billItems = addItemHandler({ itemArray: billItems, action: "add" });
        billItemTable(billItems);
        calBillValue(billItems);
    });
    //Add Item -> Delete item
    $('#billItemTable').on('click', '.deleteItem', function () {
        var cs = $(this).closest('tr');
        var deleteId = cs.find('td:eq(1)').text();
        confirmationMsg({msg:'Do you want to delete this item?'}, function (confirmed) {
            if(confirmed){
                billItems = billItems.filter(function (i) { return i.id != deleteId });
            }
            billItemTable(billItems);
            calBillValue(billItems);
        });
    });
    //Add Item -> Clear form
    $('#clearItem').on('click', function () {
        clearDivElements('#addItemForm');
        loadItemDropdown();
    });

    // Call bill value
    $('input[name="cash"], input[name="creditCarde"]').on('input', function () {
        calBillValue(billItems);
    });

    //Load Bill details


    //Save Bill
    $('#saveBill').on('click', function () {
        saveBill(billItems);
    });

    //Edit Bill
    $('#editBill').on('click', function () {
        editBill();
    });

    //New Bill
    $('#newBill').on('click', function () {
        clearFormData('#billingForm');
        clearFormErrorMsg('#billingForm');
        $("#date").datepicker("destroy");
        docDatePicker();
        newBillNo();
        billItems.length = 0;
        billItemTable(billItems);
        loadItemDropdown();
        loadCustomerDropdown();
        $('#editBill, #cancelBill').hide();
        $('#saveBill').show();
    });
}

function loadItemDropdown() {
    data = request_handler({ url: getDefaultGateway() + "main/" + "items", method: "GET", data: {} });
    if (data.status) {
        items = data.body.itemList
        $('select[name="item"]').find('option').remove().end();
        $('select[name="item"]').append("<option selected disabled hidden style='display: none' value=''></option>");
        for (var i = 0; i < items.length; i++) {
            $('select[name="item"]').append("<option value='" + items[i].id + "'>" +
                items[i].description + " (" + items[i].code + ")" + "</option>");
        }
    }
}

function itemDetails(id) {
    data = request_handler({ url: getDefaultGateway() + "main/" + "items/" + id, method: "GET", data: {} });
    if (data.status) {
        body = data.body
        if (body.responseMsg.status == "Success") {
            itemDetail = body.itemDetails;
            $('input[name="unitDes"]').val(itemDetail.unit.description);
            $('input[name="itemSOH"]').val(itemDetail.soh);
            $('input[name="itemDes"]').val(itemDetail.description);
            $('input[name="itemIndex"]').val(itemDetail.id);
            $('input[name="itemCode"]').val(itemDetail.code);
            $('input[name="unitId"]').val(itemDetail.unit.id);
        } else {
            responseAlert(body.responseMsg.status, body.responseMsg.msg);
        }
    }
}

function loadCustomerDropdown() {
    data = request_handler({ url: getDefaultGateway() + "main/" + "customers", method: "GET", data: {} });
    if (data.status) {
        customers = data.body
        $('select[name="billTo"]').find('option').remove().end();
        for (var i = 0; i < customers.length; i++) {
            $('select[name="billTo"]').append("<option value='" + customers[i].id + "'>" + customers[i].description + "</option>");
        }
    }
}

function calItemValue() {
    qty = parseFloat($('input[name="itemQty"]').val());
    price = parseFloat($('input[name="itemPrice"]').val());
    $('input[name="itemValue"]').val(qty * price);
}

function calBillValue(itemArray) {
    var billTotal = 0;
    var cash = parseFloat($('input[name="cash"]').val());
    var creditCard = parseFloat($('input[name="creditCard"]').val());
    for (var i = 0; i < itemArray.length; i++) {
        billTotal +=  parseFloat(itemArray[i].value)
    }
    $('input[name="billTotal"]').val(billTotal);
    $('input[name="balance"]').val(billTotal - (cash + creditCard));
}

function newBillNo() {
    data = request_handler({ url: getDefaultGateway() + "main/" + "newNumbers/INV", method: "GET", data: {} });
    if (data.status) {
        body = data.body
        if (body.status == "Success") {
            $('input[name="docNo"]').val(body.msg);
        } else {
            responseAlert(body.status, body.msg);
        }
    }
}

function addItemHandler(data) {
    var itemArray = data.itemArray;
    var values = {
        recNo: itemArray.length + 1,
        id: parseInt($('input[name="itemIndex"]').val()),
        code: $('input[name="itemCode"]').val(),
        description: $('input[name="itemDes"]').val(),
        unitId: parseInt($('input[name="unitId"]').val()),
        unitDescription: $('input[name="unitDes"]').val(),
        qty: parseFloat($('input[name="itemQty"]').val()),
        price: parseFloat($('input[name="itemPrice"]').val()),
        value: parseFloat($('input[name="itemValue"]').val()),
    };
    switch (data.action) {
        case "add":
            if (itemArray == undefined || Object.keys(itemArray).length === 0) {
                itemArray.push(values);
                shortPopUp({type: "success", title: "Success", msg: "Item Added Successfully."});
                clearDivElements('#addItemForm');
                loadItemDropdown();
            } else {
                var duplicate = itemArray.filter(function (i) { return i.id == values.id });
                if (Object.keys(duplicate).length === 0) {
                    itemArray.push(values);
                    shortPopUp({type: "success", title: "Success", msg: "Item Added Successfully."});
                    clearDivElements('#addItemForm');
                    loadItemDropdown();
                } else {
                    shortPopUp({type: "warning", title: "Warning", msg: "This Item already added."});
                }
            }
            return itemArray;
        case "edit":

            return itemArray;
        default:
            responseAlert("Warning", "Unrecognize Action.");
            return itemArray;
    }
}

function saveBill(itemArray) {
    var sendData = {
        id: parseInt($('input[name="docId"]').val()),
        date: $('input[name="docDate"]').val(),
        billNo: $('input[name="docNo"]').val(),
        billToId: parseInt($('select[name="billTo"]').val()),
        billDetails: itemArray,
        billTotal: parseFloat($('input[name="billTotal"]').val()),
        cashPaid: parseFloat($('input[name="cash"]').val()),
        creditPaid: parseFloat($('input[name="creditCard]').val()),
    }
    data = request_handler({ url: getDefaultGateway() + "main/" + "bills", method: "POST", data: JSON.stringify(sendData)});
    if (data.status) {
        body = data.body
        responseAlert(body.responseMsg.status, body.responseMsg.msg);
        if (body.responseMsg.status == "Success"){
            $('#newBill').trigger("click");
        }
    }
}

function editBill() {

}

function populateBillingForm(data) {

}

function billItemTable(data) {
    var table = '<thead><tr>' +
        '<th style="display:none;">recordId</th>' +
        '<th style="display:none;">itemIndex</th>' +
        '<th>Item Code</th>' +
        '<th>Item Description</th>' +
        '<th>Unit</th>' +
        '<th>Qty</th>' +
        '<th>Unit Price</th>' +
        '<th>Value</th>' +
        '<th>Actions</th>' +
        '</tr></thead><tbody>';
    for (var i in data) {
        table += '<tr>' +
            '<td style="display:none;">' + data[i].recNo +
            '<td style="display:none;">' + data[i].id +
            '</td><td>' + data[i].code +
            '</td><td>' + data[i].description +
            '</td><td>' + data[i].unitDescription +
            '</td><td>' + data[i].qty +
            '</td><td>' + data[i].price +
            '</td><td>' + data[i].value +
            '</td><td>' +
            '<a class="m-r-15 deleteItem" data-toggle="tooltip" title="Delete">' +
            '<i style="padding-left: 10px;color:#f00;" class="fas fa-trash-alt"></i></a>' +
            '</td></tr>';
    }
    table += '</tbody>';
    $("#billItemTable").html(table);
    $("#billItemTable").DataTable({ "dom": 'lrtip', "lengthChange": false, "aaSorting": [], pageLength: 5, "destroy": true });
}

function keyControl(){
    // $('input[name="billTo"]').on('change', function () {
    //     $('input[name="item"]').focus();
    // });
    // $('input[name="item"]').on('change', function () {
    //     $('input[name="itemQty"]').focus();
    // });
    
}