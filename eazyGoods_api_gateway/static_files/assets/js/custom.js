// Clear form error messages
function clearFormErrorMsg(formId) {
	$(formId + " .is-invalid").removeClass("is-invalid");
	$(formId + " .validation_error").html("");
	$(formId + " .sucess_message").html("");
	$(formId + ' .input-validation-tooltip').attr('data-original-title', "");
}

//Clear Form Data
function clearFormData(formId) {
	$(formId).trigger("reset");
	$(formId + " .hideOnNew").val('');
	$(formId + ' select').find('.inactive_select_option').remove();
	$(formId + " select option:selected").removeAttr("selected");//remove selected atribute in select elements
	// $(formId + ' select').append("<option selected disabled hidden style='display: none' value=''></option>"); //Add a hidden option to select element
}

function docDatePicker() {
	$(function () {
		function cb(selectedTime) { $('#date').val(selectedTime.format('YYYY-MM-DD')); }
		$('#date').daterangepicker({ singleDatePicker: true, showDropdowns: true, autoUpdateInput: true, locale: { format: 'YYYY-MM-DD' } }, cb);
	});
}

function keyboardControl(data, e) {
	var returnVal = false;
	var codes = [{ id: 9, name: "tab" }];
	var code = e.keyCode || e.which;
	var result = codes.filter(function (codes) { return codes.name == data.key });
	if (code === result.id) {
		e.preventDefault();
		returnVal = true;
	} else {
		returnVal = false;
	}
	return returnVal;
}

function focusElements(element) {
	$(element).focus();
	$(window).scrollTop($(element).position().top)
}

function confirmationMsg(data, callback) {
	Swal.fire({
		title: "Are you sure?",
		text: data.msg,
		type: 'warning',
		showCancelButton: true,
		confirmButtonColor: '#d33',
		cancelButtonColor: '#3085d6',
		cancelButtonText: 'No',
		confirmButtonText: "Yes, Delete it!"
	}).then((confirmed) => {
		if (confirmed.value) {
			callback(true);
		} else {
			callback(false);
		}
	});
}

function clearDivElements(divId) {
	$(divId).children().find('input, select').each(function () {
		$(this).val('');
	});
}

function shortPopUp(data) {
	var placementFrom = "top";
	var placementAlign = "right";
	var state = data.type;
	var content = {};

	content.message = data.msg;
	content.title = data.title;
	if (state == "warning") {
		content.icon = 'fas fa-exclamation-triangle';
	} else if (state == "success"){
		content.icon = 'fa fa-check';
	} else{
		content.icon = 'none';
	}
	$.notify(content, {
		type: state,
		placement: {
			from: placementFrom,
			align: placementAlign
		},
		time: 1000,
		delay: 1000,
	});
}

function searchTables(tableId, tableSearchId) {
    var otherEntryTable = $(tableId).DataTable({ "dom": 'lrtip', "lengthChange": false, "aaSorting": [], pageLength: 5, "destroy": true });
    $(tableSearchId).keyup(function () {
        otherEntryTable.search($(this).val()).draw();
    });
}

function Binding(b) {
    _this = this
    this.elementBindings = []
    this.value = b.object[b.property]
    this.valueGetter = function(){
        return _this.value;
    }
    this.valueSetter = function(val){
        _this.value = val
        for (var i = 0; i < _this.elementBindings.length; i++) {
            var binding=_this.elementBindings[i]
            binding.element[binding.attribute] = val
        }
    }
    this.addBinding = function(element, attribute, event){
        var binding = {
            element: element,
            attribute: attribute
        }
        if (event){
            element.addEventListener(event, function(event){
                _this.valueSetter(element[attribute]);
            })
            binding.event = event
        }       
        this.elementBindings.push(binding)
        element[attribute] = _this.value
        return _this
    }

    Object.defineProperty(b.object, b.property, {
        get: this.valueGetter,
        set: this.valueSetter
    }); 

	b.object[b.property] = this.value;
	console.log(this.value);
}

// var formData = {itemQty:34, itemPrice:50}
// new Binding({object: formData, property: "itemQty"}).addBinding($('input[name="itemQty"]'), "value", "keyup");
// new Binding({object: formData, property: "itemPrice"}).addBinding($('input[name="itemPrice"]'), "value", "keyup");

function validator(data){
	
}

function defaultValidation(){
	
}
