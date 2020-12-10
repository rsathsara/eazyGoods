// Clear form error messages
function clearFormErrorMsg(formId){
	$(formId + " .is-invalid" ).removeClass("is-invalid");
	$(formId + " .validation_error" ).html("");
	$(formId + " .sucess_message" ).html("");
	$(formId + ' .input-validation-tooltip').attr('data-original-title', "");
}

//Clear Form Data
function clearFormData(formId){
	$(formId).trigger("reset");
    $(formId + " .hideOnNew").val('');
    // $(formId + " #saveBtnText").html('Save');
	$(formId + ' select').find('.inactive_select_option').remove();
	$(formId + " select option:selected").removeAttr("selected");//remove selected atribute in select elements
	$(formId + ' select').append("<option selected disabled hidden style='display: none' value=''></option>"); //Add a hidden option to select element
}

function docDatePicker() {
    $(function () {
        function cb(selectedTime) { $('#date').val(selectedTime.format('YYYY-MM-DD')); }
        $('#date').daterangepicker({ singleDatePicker: true, showDropdowns: true, autoUpdateInput: true, locale: { format: 'YYYY-MM-DD' } }, cb);
    });
}