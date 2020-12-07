function responseAlert(status, msg) {
	var statusType = status.toLowerCase();
	if (status == "Success") {
		Swal.fire({
			type: statusType,
			title: status,
			html: msg,
			padding: '3em',
			showConfirmButton: false,
			timer: 1500
		});
		return true;
	} else {
		Swal.fire({
			type: statusType,
			title: status,
			text: msg
		});
	}
}

//input field validation alert
function validationAlert(data) {
	$.each(data, function (key, value) {
		$('input[name="' + key + '"').addClass('is-invalid');
		$('select[name="' + key + '"').addClass('is-invalid');
		$('textarea[name="' + key + '"').addClass('is-invalid');
		$('input[name="' + key + '"').parents('.input-container').find('.validation_error').html('<p class="mt-3 text-danger"><i class="fa fa-warning"></i>' + ' ' + value + '</p>');
		$('select[name="' + key + '"').parents('.input-container').find('.validation_error').html('<p class="mt-3 text-danger"><i class="fa fa-warning"></i>' + ' ' + value + '</p>');
		$('textarea[name="' + key + '"').parents('.input-container').find('.validation_error').html('<p class="mt-3 text-danger"><i class="fa fa-warning"></i>' + ' ' + value + '</p>');
	});
}

//html sucsess message
function successAlert(msg) {
	$('.sucess_message').html('<p class="text-success" style="margin-bottom:0px"><i class="fa fa-check"></i>' + ' ' + msg + '</p>');
	$('.sucess_message').delay(100).fadeIn('slow');
	$('.sucess_message').delay(300).fadeOut('slow');
}

//input field validation alert using tooltip 
function tooltipValidation(data) {
	$.each(data, function (key, value) {
		$('input[name="' + key + '"').addClass('is-invalid');
		$('input[name="' + key + '"').attr('data-original-title', value);
		$('input[name="' + key + '"').tooltip({
			template: '<div class="tooltip tooltip-validation" role="tooltip"><div class="arrow"></div><div class="tooltip-inner"></div></div>'
		});
	});
}

function ajaxErrorAlert(error){
	Swal.fire({
		type: 'error',
		title: 'Error',
		text: 'Server Error Occurred',
		footer: '<a href="#" id="showDetailError">View Detail Error</a>'
	});
	$('#showDetailError').on('click', function () {
		showDetailError(error.responseText);
	});
}

//Show server detail error on a new page after an ajx error
function showDetailError(errorDetail){
    var x=window.open();
    x.document.open();
    x.document.write(errorDetail);
}
