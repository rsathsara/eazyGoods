function grnForm() {
    //Select2 function for searchable dropdowns 
    $('.searchable-dropdown').select2({});

    // Initialize grn item array
    var grnItems = [];

    //Onload events
    docDatePicker();
    newGrnNo();
    grnItemTable(grnItems);
    loadSupplierDropdown();
    loadItemDropdown();

    //load Grn List list
    $("#searchGrnList").click(function () {
        grnList();
        $('#grnListModal').modal('show');
    });

    // Load Grn Details
    $('#grnListTable').on('click', '.selectGrn', function () {
        var cs = $(this).closest('tr');
        var index = cs.find('td:eq(0)').text()
        $('#grnListModal').modal('hide');
        grnItems = grnDetail(index);
        grnItemTable(grnItems);
        $('#saveGrn').hide();
        $('#cancelGrn, #editGrn').show();
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
        grnItems = addItemHandler({ itemArray: grnItems, action: "add" });
        grnItemTable(grnItems);
        calGrnValue(grnItems);
    });
    // Add Item -> Delete item
    $('#grnItemTable').on('click', '.deleteItem', function () {
        var cs = $(this).closest('tr');
        var deleteId = cs.find('td:eq(1)').text();
        confirmationMsg({ msg: 'Do you want to delete this item?' }, function (confirmed) {
            if (confirmed) {
                grnItems = grnItems.filter(function (i) { return i.id != deleteId });
            }
            grnItemTable(grnItems);
            calGrnValue(grnItems);
        });
    });
    // Add Item -> Clear form
    $('#clearItem').on('click', function () {
        clearDivElements('#addItemForm');
        loadItemDropdown();
    });
    // Add Item -> Edit Item
    $('#grnItemTable').on('click', '.editItem', function () {
        var cs = $(this).closest('tr');
        grnItems = addItemHandler({itemArray: grnItems, action: "edit", editRow: cs});
    });
    $('#grnItemTable').on('click', '.saveEditedItem', function () {
        var cs = $(this).closest('tr');
        grnItems = addItemHandler({itemArray: grnItems, action: "saveEdit", editRow: cs, editingElements: {qty: 5, price: 6, value: 7}});
        calGrnValue(grnItems);
    });
    $('#grnItemTable').on('input', '.input-qty, .input-price', function () {
        var cs = $(this).closest('tr');
        var price = cs.find('td:eq(6)').find('input').val();
        var qty = cs.find('td:eq(5)').find('input').val();
        cs.find('td:eq(7)').text(price * qty);
    });
    

    // Calculate grn value
    $('input[name="cash"], input[name="creditCarde"]').on('input', function () {
        calGrnValue(grnItems);
    });

    //Save Grn
    $('#saveGrn').on('click', function () {
        saveGrn(grnItems);
    });

    //Edit Grn
    $('#editGrn').on('click', function () {
        editGrn(grnItems);
    });

    //Cancel Grn
    $('#cancelGrn').on('click', function () {
        cancelGrn();
    });

    //New Grn
    $('#newGrn').on('click', function () {
        clearFormData('#grnForm');
        clearFormErrorMsg('#grnForm');
        docDatePicker();
        newGrnNo();
        grnItems.length = 0;
        grnItemTable(grnItems);
        loadItemDropdown();
        loadSupplierDropdown();
        $('#editGrn, #cancelGrn').hide();
        $('#saveGrn').show();
    });
}

// **********************************Functions start from here*****************************************

function newGrnNo() {
    data = request_handler({ url: getDefaultGateway() + "main/" + "newNumbers/GRN", method: "GET", data: {} });
    if (data.status) {
        body = data.body
        if (body.status == "Success") {
            $('input[name="docNo"]').val(body.msg);
        } else {
            responseAlert(body.status, body.msg);
        }
    }
}

function loadSupplierDropdown() {
    data = request_handler({ url: getDefaultGateway() + "main/" + "suppliers", method: "GET", data: {} });
    if (data.status) {
        suppliers = (data.body == null) ? ([]) : (data.body)
        $('select[name="grnFrom"]').find('option').remove().end();
        for (var i = 0; i < suppliers.length; i++) {
            $('select[name="grnFrom"]').append("<option value='" + suppliers[i].id + "'>" + suppliers[i].description + "</option>");
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

function grnItemTable(data) {
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
            '</td><td style="width: 12%;"><input name="itemQty['+ data[i].id +']" type="text" value="'+ data[i].qty +'" autocomplete="off" data-toggle="tooltip"' +
            'data-placement="bottom" data-original-title="" class="form-control input-validation-tooltip hideOnNew input-qty" readonly>'+
            '</td><td style="width: 12%;"><input name="itemPrice['+ data[i].id +']" type="text" value="'+ data[i].price +'" autocomplete="off" data-toggle="tooltip"' +
            'data-placement="bottom" data-original-title="" class="form-control input-validation-tooltip hideOnNew input-price" readonly>'+
            '</td><td style="width: 12%;">'+ data[i].value +
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
    $("#grnItemTable").html(table);
    $("#grnItemTable").DataTable({ "dom": 'rt', "scrollY": "175px", "scrollX": true, "destroy": true });
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

function grnList() {
    data = request_handler({ url: getDefaultGateway() + "main/" + "grns", method: "GET", data: {} });
    if (data.status) {
        list = data.body.grn;
        var row = '<thead><tr>' +
            '<th class="th-sm" style="display:none;">Index</th>' +
            '<th class="th-sm">GRN No.</th>' +
            '<th class="th-sm">Date</th>' +
            '</tr></thead><tbody>';
        for (var i = 0; i < list.length; i++) {
            row += '<tr class="selectGrn"><td style="display:none;">' + list[i].id + '</td>' +
                '</td><td>' + list[i].grnNo +
                '</td><td>' + list[i].date +
                '</td></tr>';
        }
        row += '</tbody>';

        $("#grnListTable").html(row);
        searchTables("#grnListTable", "#grnListTableSearch");
    }
}

function grnDetail(id) {
    var returnArray = [];
    data = request_handler({ url: getDefaultGateway() + "main/" + "grns/" + id, method: "GET", data: {} });
    if (data.status) {
        grn = data.body.grn[0];
        $('input[name="docNo"]').val(grn.grnNo);
        $('input[name="docDate"]').val(grn.date);
        $('select[name="grnFrom"]').val(grn.grnFromId);
        $('input[name="grnTotal"]').val(grn.grnTotal);
        $('input[name="balance"]').val(grn.grnTotal);
        $('input[name="docId"]').val(grn.id);
        returnArray = grn.grnDetails;
    }
    return returnArray;
}

function calItemValue() {
    qty = parseFloat($('input[name="itemQty"]').val());
    price = parseFloat($('input[name="itemPrice"]').val());
    $('input[name="itemValue"]').val(qty * price);
}

function calGrnValue(itemArray) {
    var grnTotal = 0;
    var cash = parseFloat($('input[name="cash"]').val());
    var creditCard = parseFloat($('input[name="creditCard"]').val());
    for (var i = 0; i < itemArray.length; i++) {
        grnTotal += parseFloat(itemArray[i].value)
    }
    $('input[name="grnTotal"]').val(grnTotal);
    $('input[name="balance"]').val(grnTotal - (cash + creditCard));
}

function addItemHandler(data) {
    var itemArray = data.itemArray;
    var values = grnFormData({action: "addItemData", itemArray: itemArray});
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
                    if (el.length){
                        var val = row.find('td:eq(' + editingElements[key] + ')').find('input').val();
                    } else{
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

function grnFormData(data){
    switch (data.action){
        case "sendData":
            var sendData = {
                id: parseInt($('input[name="docId"]').val()),
                date: $('input[name="docDate"]').val(),
                grnNo: $('input[name="docNo"]').val(),
                grnFromId: parseInt($('select[name="grnFrom"]').val()),
                grnTotal: parseFloat($('input[name="grnTotal"]').val()),
                // cashPaid: parseFloat($('input[name="cash"]').val()),
                // creditPaid: parseFloat($('input[name="creditCard]').val()),
                grnDetails: data.itemArray,
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

function saveGrn(grnItems) {
    var sendData = grnFormData({action: "sendData", itemArray: grnItems});
    data = request_handler({ url: getDefaultGateway() + "main/" + "grns", method: "POST", data: JSON.stringify(sendData) });
    if (data.status) {
        body = data.body
        responseAlert(body.responseMsg.status, body.responseMsg.msg);
        if (body.responseMsg.status == "Success") {
            $('#newGrn').trigger("click");
        }
    }
}

function editGrn(grnItems) {
    var id = $('input[name="docId"]').val();
    var sendData = grnFormData({action: "sendData", itemArray: grnItems});
    data = request_handler({ url: getDefaultGateway() + "main/" + "grns/" + id, method: "PUT", data: JSON.stringify(sendData) });
    if (data.status) {
        body = data.body
        responseAlert(body.responseMsg.status, body.responseMsg.msg);
        if (body.responseMsg.status == "Success") {
            $('#newGrn').trigger("click");
        }
    }
}

function cancelGrn(){

}