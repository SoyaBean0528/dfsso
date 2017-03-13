// init
$(function () {
	// trigger userList datatable
 	$("#userListTable").DataTable({
 		"aoColumnDefs": [ { "bSortable": false, "aTargets": [ 4 ] }],
 	});
 	// check add msg
 	if(document.getElementById("addUserErrMsg")){
 		$("#createUserModal").removeClass('fade');
 		$("#createUserModal").modal('show');
        $('#createUserModal').on('hidden.bs.modal', function () {
            if (!$("#createUserModal").hasClass("fade")) {
                $("#createUserModal").addClass('fade');
                $('#addUserErrMsg').remove();
            } 
        })
 	}
    // check del msg 
    if(document.getElementById("delUserErrMsg")) {
        $("#delUserModal").removeClass('fade');
        $("#delUserModal").modal('show');
        $('#delUserModal').on('hidden.bs.modal', function () {
            if (!$("#delUserModal").hasClass("fade")) {
                $("#delUserModal").addClass('fade');
                $('#delUserErrMsg').remove();
            } 
        })
    }
    // check update msg
    if(document.getElementById("updateUserErrMsg")){
        $("#editUserModal").removeClass('fade');
        $("#editUserModal").modal('show');
        $('#editUserModal').on('hidden.bs.modal', function () {
            if (!$("#editUserModal").hasClass("fade")) {
                $("#editUserModal").addClass('fade');
                $('#updateUserErrMsg').remove();
            } 
        })
    }
});

// submit addUserForm
function submitAddUserForm() {
  	document.getElementById("addUserForm").submit();
}

// show del user modal
function showDelUserModal(obj) {  
    // get username from userListTable by id
    var id = $(obj).attr("id");  
    var row = parseInt($(obj).attr("row")) + 1;
    var username = document.getElementById("userListTable").rows[row].cells[1].innerText;
    // set modal val
    var content = "确定删除用户" + username + "吗？";
    document.getElementById("delUserFormId").value = id;
    document.getElementById("delUserModalOK").value = id; 
    document.getElementById("delUserContent").innerHTML = content;
    $('#delUserModal').modal('show');  
}  

// del user
function delUser(obj) {
    var id = $(obj).attr("value")
    document.getElementById("delUserForm").submit();
}

// show edit user modal
function showEditUserModal(obj) {  
    // get userdata from userListTable by id
    var id = $(obj).attr("id");  
    var row = parseInt($(obj).attr("row")) + 1;
    var username = document.getElementById("userListTable").rows[row].cells[1].innerText;
    // set modal val
    document.getElementById("editUserFormUserId").value = id
    document.getElementById("editUserFormUserName").value = username;
    $('#editUserModal').modal('show');  
}  

// submit eidtUserForm
function submitEditUserForm() {
    document.getElementById("editUserForm").submit();
}