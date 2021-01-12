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

    //load Bill List list
    $("#searchBillList").click(function () {
        billList();
        $('#billListModal').modal('show');
    });

    // Load Bill Details
    $('#billListTable').on('click', '.selectBill', function () {
        var cs = $(this).closest('tr');
        var index = cs.find('td:eq(0)').text()
        $('#billListModal').modal('hide');
        billItems = billDetail(index);
        billItemTable(billItems);
        $('#saveBill').hide();
        $('#cancelBill, #editBill').show();
    });

    // Add Item -> Select Item
    $("#item").on('change', function () {
        itemDetails($(this).val());
        $('input[name="itemQty"]').focus();
    });
    // Add Item -> Cal Item Value
    $('input[name="itemQty"], input[name="itemPrice"]').on('input', function () {
        calItemValue();
    });
    // Add Item -> Add item
    $('#addItem').on('click', function () {
        clearFormErrorMsg();
        billItems = addItemHandler({ itemArray: billItems, action: "add" });
        billItemTable(billItems);
        calBillValue(billItems);
    });
    // Add Item -> Delete item
    $('#billItemTable').on('click', '.deleteItem', function () {
        var cs = $(this).closest('tr');
        var deleteId = cs.find('td:eq(1)').text();
        confirmationMsg({ msg: 'Do you want to delete this item?' }, function (confirmed) {
            if (confirmed) {
                billItems = billItems.filter(function (i) { return i.id != deleteId });
            }
            billItemTable(billItems);
            calBillValue(billItems);
        });
    });
    // Add Item -> Clear form
    $('#clearItem').on('click', function () {
        clearDivElements('#addItemForm');
        loadItemDropdown();
    });
    // Add Item -> Edit Item
    $('#billItemTable').on('click', '.editItem', function () {
        var cs = $(this).closest('tr');
        billItems = addItemHandler({ itemArray: billItems, action: "edit", editRow: cs });
    });
    $('#billItemTable').on('click', '.saveEditedItem', function () {
        var cs = $(this).closest('tr');
        billItems = addItemHandler({ itemArray: billItems, action: "saveEdit", editRow: cs, editingElements: { qty: 5, price: 6, value: 7 } });
        calBillValue(billItems);
    });
    $('#billItemTable').on('input', '.input-qty, .input-price', function () {
        var cs = $(this).closest('tr');
        var price = cs.find('td:eq(6)').find('input').val();
        var qty = cs.find('td:eq(5)').find('input').val();
        cs.find('td:eq(7)').text(price * qty);
    });


    // Calculate bill value
    $('input[name="cash"], input[name="creditCarde"]').on('input', function () {
        calBillValue(billItems);
    });

    //Save Bill
    $('#saveBill').on('click', function () {
        saveBill(billItems);
    });

    //Edit Bill
    $('#editBill').on('click', function () {
        editBill(billItems);
    });

    //Cancel Bill
    $('#cancelBill').on('click', function () {
        cancelBill();
    });

    //New Bill
    $('#newBill').on('click', function () {
        clearFormData('#billingForm');
        clearFormErrorMsg('#billingForm');
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

// **********************************Functions start from here*****************************************

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

function loadCustomerDropdown() {
    data = request_handler({ url: getDefaultGateway() + "main/" + "customers", method: "GET", data: {} });
    if (data.status) {
        customers = (data.body == null) ? ([]) : (data.body)
        $('select[name="billTo"]').find('option').remove().end();
        for (var i = 0; i < customers.length; i++) {
            $('select[name="billTo"]').append("<option value='" + customers[i].id + "'>" + customers[i].description + "</option>");
        }
    }
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
            '</td><td style="width: 12%;">' + data[i].code +
            '</td><td>' + data[i].description +
            '</td><td style="width: 10%;">' + data[i].unitDescription +
            '</td><td style="width: 12%;"><input name="itemQty[' + data[i].id + ']" type="text" value="' + data[i].qty + '" autocomplete="off" data-toggle="tooltip"' +
            'data-placement="bottom" data-original-title="" class="form-control input-validation-tooltip hideOnNew input-qty" readonly>' +
            '</td><td style="width: 12%;"><input name="itemPrice[' + data[i].id + ']" type="text" value="' + data[i].price + '" autocomplete="off" data-toggle="tooltip"' +
            'data-placement="bottom" data-original-title="" class="form-control input-validation-tooltip hideOnNew input-price" readonly>' +
            '</td><td style="width: 12%;">' + data[i].value +
            '</td><td style="width: 10%;">' +
            '<a class="m-r-15 deleteItem" data-toggle="tooltip" title="Delete">' +
            '<i style="padding-left: 10px;color:#f00;" class="fas fa-trash-alt"></i></a>' +
            '<a class="m-r-15 editItem" data-toggle="tooltip" title="Edit" data-original-title="Edit">' +
            '<i style="padding-left: 10px;color:#0028ff;" class="fa fa-edit"></i></a>' +
            '<a hidden class="m-r-15 saveEditedItem" data-toggle="tooltip" title="Save" data-original-title="Save">' +
            '<i style="padding-left: 10px;color:#0028ff;" class="fa fa-save"></i></a>' +
            '</td></tr>';
    }
    table += '</tbody>';
    $("#billItemTable").html(table);
    $("#billItemTable").DataTable({ "dom": 'rt', "scrollY": "175px", "scrollX": true, "destroy": true });
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
            $('input[name="itemPrice"]').val(itemDetail.salePrice);
            $('input[name="itemQty"]').val("");
            $('input[name="itemValue"]').val("");
        } else {
            responseAlert(body.responseMsg.status, body.responseMsg.msg);
        }
    }
}

function billList() {
    data = request_handler({ url: getDefaultGateway() + "main/" + "bills", method: "GET", data: {} });
    if (data.status) {
        list = data.body.bill;
        var row = '<thead><tr>' +
            '<th class="th-sm" style="display:none;">Index</th>' +
            '<th class="th-sm">Bill No.</th>' +
            '<th class="th-sm">Date</th>' +
            '</tr></thead><tbody>';
        for (var i = 0; i < list.length; i++) {
            row += '<tr class="selectBill"><td style="display:none;">' + list[i].id + '</td>' +
                '</td><td>' + list[i].billNo +
                '</td><td>' + list[i].date +
                '</td></tr>';
        }
        row += '</tbody>';

        $("#billListTable").html(row);
        searchTables("#billListTable", "#billListTableSearch");
    }
}

function billDetail(id) {
    var returnArray = [];
    data = request_handler({ url: getDefaultGateway() + "main/" + "bills/" + id, method: "GET", data: {} });
    if (data.status) {
        bill = data.body.bill[0];
        $('input[name="docNo"]').val(bill.billNo);
        $('input[name="docDate"]').val(bill.date);
        $('select[name="billTo"]').val(bill.billToId).trigger('change');
        $('input[name="billTotal"]').val(bill.billTotal);
        $('input[name="balance"]').val(bill.billTotal);
        $('input[name="docId"]').val(bill.id);
        returnArray = bill.billDetails;
    }
    return returnArray;
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
        billTotal += parseFloat(itemArray[i].value)
    }
    $('input[name="billTotal"]').val(billTotal);
    $('input[name="balance"]').val(billTotal - (cash + creditCard));
}

function addItemHandler(data) {
    var itemArray = data.itemArray;
    var values = billFormData({ action: "addItemData", itemArray: itemArray });
    switch (data.action) {
        case "add":
            if (itemArray == undefined || Object.keys(itemArray).length === 0) {
                itemArray.push(values);
                shortPopUp({ type: "success", title: "Success", msg: "Item Added Successfully." });
                clearDivElements('#addItemForm');
                loadItemDropdown();
            } else {
                var duplicate = itemArray.filter(function (i) { return i.id == values.id });
                if (Object.keys(duplicate).length === 0) {
                    itemArray.push(values);
                    shortPopUp({ type: "success", title: "Success", msg: "Item Added Successfully." });
                    clearDivElements('#addItemForm');
                    loadItemDropdown();
                } else {
                    shortPopUp({ type: "warning", title: "Warning", msg: "This Item already added." });
                }
            }
            return itemArray;
        case "edit":
            var row = data.editRow;
            row.find('input').removeAttr('readonly');
            row.find('.editItem').prop('hidden', true);
            row.find('.saveEditedItem').prop('hidden', false);
            return itemArray;
        case "saveEdit":
            var row = data.editRow;
            var editingElements = data.editingElements;
            row.find('input').prop('readonly', true);
            row.find('.editItem').prop('hidden', false);
            row.find('.saveEditedItem').prop('hidden', true);
            for (var key in editingElements) {
                if (editingElements.hasOwnProperty(key)) {
                    var id = row.find('td:eq(0)').text();
                    var el = row.find('td:eq(' + editingElements[key] + ')').find('input');
                    if (el.length) {
                        var val = row.find('td:eq(' + editingElements[key] + ')').find('input').val();
                    } else {
                        var val = row.find('td:eq(' + editingElements[key] + ')').text();
                    }
                    for (var i = 0; i < itemArray.length; i++) {
                        if (itemArray[i].recNo == id) {
                            itemArray[i][key] = parseFloat(val);
                        }
                    }
                }
            }
            return itemArray;
        default:
            responseAlert("Warning", "Unrecognize Action.");
            return itemArray;
    }
}

// var validatingElements = {
//     "date": { value: "", element: "", labal: "", warning: "", conditions: [{ name: "", msg: "", data: {} }] },
//     "billTo": { value: "", element: "", labal: "", warning: "", conditions: [{ name: "", msg: "", data: {} }] },
//     "itemId": { value: "", element: "", labal: "", warning: "", conditions: [{ name: "", msg: "", data: {} }] },
//     "billItemQty": { value: "", element: "", labal: "", warning: "", conditions: [{ name: "", msg: "", data: {} }] },
//     "billItemPrice": { value: "", element: "", labal: "", warning: "", conditions: [{ name: "", msg: "", data: {} }] },
// };

function billFormData(data) {
    switch (data.action) {
        case "sendData":
            var sendData = {
                id: parseInt($('input[name="docId"]').val()),
                date: $('input[name="docDate"]').val(),
                billNo: $('input[name="docNo"]').val(),
                billToId: parseInt($('select[name="billTo"]').val()),
                billTotal: parseFloat($('input[name="billTotal"]').val()),
                cashPaid: parseFloat($('input[name="cash"]').val()),
                creditPaid: parseFloat($('input[name="creditCard]').val()),
                billDetails: data.itemArray,
            }
            return sendData;
        case "addItemData":
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
            return values;
        default:
            return [];
    }
}

function saveBill(billItems) {
    var sendData = billFormData({ action: "sendData", itemArray: billItems });
    data = request_handler({ url: getDefaultGateway() + "main/" + "bills", method: "POST", data: JSON.stringify(sendData) });
    if (data.status) {
        body = data.body
        responseAlert(body.responseMsg.status, body.responseMsg.msg);
        if (body.responseMsg.status == "Success") {
            $('#newBill').trigger("click");
        }
    }
}

function editBill(billItems) {
    var id = $('input[name="docId"]').val();
    var sendData = billFormData({ action: "sendData", itemArray: billItems });
    data = request_handler({ url: getDefaultGateway() + "main/" + "bills/" + id, method: "PUT", data: JSON.stringify(sendData) });
    if (data.status) {
        body = data.body
        responseAlert(body.responseMsg.status, body.responseMsg.msg);
        if (body.responseMsg.status == "Success") {
            $('#newBill').trigger("click");
        }
    }
}

function cancelBill() {

}