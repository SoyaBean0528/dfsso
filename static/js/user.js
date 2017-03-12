// init
$(function () {
	// trigger userList datatable
 	$("#userList").DataTable({
 		"aoColumnDefs": [ { "bSortable": false, "aTargets": [ 4 ] }]
 	});
 	// check msg
 	if(document.getElementById("addUserErrMsg")){
 		$("#createUserModal").removeClass('fade')
 		$('#createUserModal').modal('show')
 	}
});

// submit addUserForm
function submitAddUserForm() {
  	document.getElementById("addUserForm").submit();
}
