// init
$(function () {
	// trigger userList datatable
 	$("#userList").DataTable();
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
