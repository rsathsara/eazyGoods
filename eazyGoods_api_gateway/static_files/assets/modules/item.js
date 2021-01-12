function itemForm() {
    //Select2 function for searchable dropdowns 
    $('.searchable-dropdown').select2({});

    //Onload events
    newItemNo();
    loadUnitDropdown();
    loadItemCat1Dropdown();
    loadItemCat2Dropdown();

    //load list
    $("#searchItemList").click(function () {
        itemList();
        $('#itemListModal').modal('show');
    });

    // Load Details
    $('#itemListTable').on('click', '.selectItem', function () {
        var cs = $(this).closest('tr');
        var index = cs.find('td:eq(0)').text()
        $('#itemListModal').modal('hide');
        itemDetails(index);
        $('#saveItem').hide();
        $('#editItem').show();
    });

    //Save
    $('#saveItem').on('click', function () {
        saveItem();
    });

    //Edit
    $('#editItem').on('click', function () {
        editItem();
    });

    //New
    $('#newItem').on('click', function () {
        clearFormData('#itemForm');
        clearFormErrorMsg('#itemForm');
        newItemNo();
        loadUnitDropdown();
        loadItemCat1Dropdown();
        loadItemCat2Dropdown();
        $('#editItem').hide();
        $('#saveItem').show();
    });
}

// **********************************Functions start from here*****************************************

function newItemNo() {
    data = request_handler({ url: getDefaultGateway() + "main/" + "newNumbers/FG", method: "GET", data: {} });
    if (data.status) {
        body = data.body
        if (body.status == "Success") {
            $('input[name="docNo"]').val(body.msg);
        } else {
            responseAlert(body.status, body.msg);
        }
    }
}

function loadUnitDropdown() {
    data = request_handler({ url: getDefaultGateway() + "main/" + "units", method: "GET", data: {} });
    if (data.status) {
        resData = (data.body.unitList == null) ? ([]) : (data.body.unitList)
        $('select[name="unit"]').find('option').remove().end();
        $('select[name="unit"]').append("<option selected disabled hidden style='display: none' value=''></option>");
        for (var i = 0; i < resData.length; i++) {
            $('select[name="unit"]').append("<option value='" + resData[i].id + "'>" + resData[i].description + "</option>");
        }
    }
}

function loadItemCat1Dropdown() {
    data = request_handler({ url: getDefaultGateway() + "main/" + "itemCat1", method: "GET", data: {} });
    if (data.status) {
        resData = (data.body.itemCat1List == null) ? ([]) : (data.body.itemCat1List)
        $('select[name="category1"]').find('option').remove().end();
        $('select[name="category1"]').append("<option selected disabled hidden style='display: none' value=''></option>");
        for (var i = 0; i < resData.length; i++) {
            $('select[name="category1"]').append("<option value='" + resData[i].id + "'>" +
                resData[i].description + "</option>");
        }
    }
}

function loadItemCat2Dropdown() {
    data = request_handler({ url: getDefaultGateway() + "main/" + "itemCat2", method: "GET", data: {} });
    if (data.status) {
        resData = (data.body.itemCat2List == null) ? ([]) : (data.body.itemCat2List)
        $('select[name="category2"]').find('option').remove().end();
        $('select[name="category2"]').append("<option selected disabled hidden style='display: none' value=''></option>");
        for (var i = 0; i < resData.length; i++) {
            $('select[name="category2"]').append("<option value='" + resData[i].id + "'>" +
                resData[i].description + "</option>");
        }
    }
}

function itemList() {
    data = request_handler({ url: getDefaultGateway() + "main/" + "items", method: "GET", data: {} });
    if (data.status) {
        list = data.body.itemList;
        var row = '<thead><tr>' +
            '<th class="th-sm" style="display:none;">Index</th>' +
            '<th class="th-sm">Code</th>' +
            '<th class="th-sm">Description</th>' +
            '</tr></thead><tbody>';
        for (var i = 0; i < list.length; i++) {
            row += '<tr class="selectItem"><td style="display:none;">' + list[i].id + '</td>' +
                '</td><td>' + list[i].code +
                '</td><td>' + list[i].description +
                '</td></tr>';
        }
        row += '</tbody>';

        $("#itemListTable").html(row);
        searchTables("#itemListTable", "#itemListTableSearch");
    }
}

function itemDetails(id) {
    data = request_handler({ url: getDefaultGateway() + "main/" + "items/" + id, method: "GET", data: {} });
    if (data.status) {
        body = data.body;
        if (body.responseMsg.status == "Success") {
            itemDetail = body.itemDetails;
            $('input[name="docId"]').val(itemDetail.id);
            $('input[name="docNo"]').val(itemDetail.code);
            $('input[name="itemDescription"]').val(itemDetail.description);
            $('#unit').val(itemDetail.unit.id).trigger('change');
            $('#category1').val(itemDetail.itemCat1Id).trigger('change');
            $('#category2').val(itemDetail.itemCat2Id).trigger('change');
            $('input[name="salePrice"]').val(itemDetail.salePrice);
        } else {
            responseAlert(body.responseMsg.status, body.responseMsg.msg);
        }
    }
}

function itemFormData(data) {
    switch (data.action) {
        case "sendData":
            var sendData = {
                id: parseInt($('input[name="docId"]').val()),
                code: $('input[name="docNo"]').val(),
                description: $('input[name="itemDescription"]').val(),
                unitId: parseInt($('select[name="unit"]').val()),
                itemCat1Id: parseInt($('select[name="category1"]').val()),
                itemCat2Id: parseInt($('select[name="category2"]').val()),
                salePrice: parseFloat($('input[name="salePrice"]').val()),
            }
            return sendData;
        default:
            return [];
    }
}

function saveItem() {
    var sendData = itemFormData({ action: "sendData"});
    data = request_handler({ url: getDefaultGateway() + "main/" + "items", method: "POST", data: JSON.stringify(sendData) });
    if (data.status) {
        body = data.body
        responseAlert(body.responseMsg.status, body.responseMsg.msg);
        if (body.responseMsg.status == "Success") {
            $('#newItem').trigger("click");
        }
    }
}

function editItem() {
    var id = $('input[name="docId"]').val();
    var sendData = itemFormData({ action: "sendData" });
    data = request_handler({ url: getDefaultGateway() + "main/" + "items/" + id, method: "PUT", data: JSON.stringify(sendData) });
    if (data.status) {
        body = data.body
        responseAlert(body.responseMsg.status, body.responseMsg.msg);
        if (body.responseMsg.status == "Success") {
            $('#newItem').trigger("click");
        }
    }
}
