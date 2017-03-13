<!-- add user modal -->
<div class="modal modal-primary fade" id="createUserModal" tabindex="-1" role="dialog" aria-labelledby="create-user-label" aria-hidden="true">
  <div class="modal-dialog" style="width:380px;margin:10% auto">
    <div class="modal-content">
      <div class="modal-header">
        <button type="button" class="close" data-dismiss="modal" aria-label="Close">
          <span aria-hidden="true">&times;</span></button>
        <h4 class="modal-title" id="create-user-label">
          添加用户
          {{if .AddUserMsg}}
            <small id="addUserErrMsg" style="color:#f56954">{{.AddUserMsg}}</small>
          {{end}}
        </h4>
      </div>
      <div class="modal-body">
        <form class="form-horizontal" id="addUserForm" action="addUser" method="post">
          <div class="form-group">
            <label for="username" class="col-sm-3 control-label">用户名</label>
            <div class="col-sm-8">
              <input type="text" class="form-control" name="username" id="username" placeholder="用户名">
            </div>
          </div>
        </form>
        <!-- ./form -->
      </div>
      <div class="modal-footer">
        <button type="button" class="btn btn-outline pull-left" data-dismiss="modal">关闭</button>
        <button type="button" class="btn btn-outline" onclick="submitAddUserForm();">添加</button>
      </div>
    </div>
    <!-- /.modal-content -->
  </div>
  <!-- /.modal-dialog -->
</div>

<!-- del user modal -->
<div class="modal modal-primary fade" id="delUserModal" tabindex="-1" role="dialog" aria-labelledby="create-user-label" aria-hidden="true">
  <div class="modal-dialog" style="width:380px;margin:10% auto">
    <div class="modal-content">
      <div class="modal-header">
        <button type="button" class="close" data-dismiss="modal" aria-label="Close">
          <span aria-hidden="true">&times;</span></button>
        <h4 class="modal-title" id="create-user-label">
          删除用户
        </h4>
      </div>
      <div class="modal-body">
        <form id="delUserForm" type="hidden" action="delUser" method="post">
          <input type="hidden" name="userid" id="delUserFormId"></form> 
        </form>
          
        {{if .DelUserMsg}}
          <p id="delUserContent"></p>
          <p id="delUserErrMsg" style="color:#f56954">{{.DelUserMsg}}</p>
        {{else}}
          <p></p>
          <p id="delUserContent"></p>
        {{end}}
      </div>
      <div class="modal-footer">
        <button type="button" class="btn btn-outline pull-left" data-dismiss="modal">取消</button>
        <button id="delUserModalOK" type="button" class="btn btn-outline" onclick="delUser(this);">
          删除
        </button>
      </div>
    </div>
    <!-- /.modal-content -->
  </div>
  <!-- /.modal-dialog -->
</div>

<!-- edit user modal -->
<div class="modal modal-primary fade" id="editUserModal" tabindex="-1" role="dialog" aria-labelledby="create-user-label" aria-hidden="true">
  <div class="modal-dialog" style="width:380px;margin:10% auto">
    <div class="modal-content">
      <div class="modal-header">
        <button type="button" class="close" data-dismiss="modal" aria-label="Close">
          <span aria-hidden="true">&times;</span></button>
        <h4 class="modal-title" id="create-user-label">
          编辑用户
          {{if .UpdateUserMsg}}
            <small id="updateUserErrMsg" style="color:#f56954">{{.UpdateUserMsg}}</small>
          {{end}}
        </h4>
      </div>
      <div class="modal-body">
        <form class="form-horizontal" id="editUserForm" action="editUser" method="post">
          <input id="editUserFormUserId" type="hidden" name="userid" id="userid">
          <div class="form-group">
            <label for="username" class="col-sm-3 control-label">用户名</label>
            <div class="col-sm-8">
              <input id="editUserFormUserName" type="text" readonly="readonly" class="form-control" name="username" id="username" placeholder="用户名">
          </div>
          </div> 
          <div class="form-group">
            <label for="password" class="col-sm-3 control-label">密码</label>
            <div class="col-sm-8">
              <input id="editUserFormPassword" value="******" type="text" class="form-control" name="password" id="password" placeholder="密码">
            </div>
          </div>
        </form>
        <!-- ./form -->
      </div>
      <div class="modal-footer">
        <button type="button" class="btn btn-outline pull-left" data-dismiss="modal">取消</button>
        <button id="delUserModalOK" type="button" class="btn btn-outline" onclick="submitEditUserForm();">
          保存
        </button>
      </div>
    </div>
    <!-- /.modal-content -->
  </div>
  <!-- /.modal-dialog -->
</div>

<!-- user list -->
<div class="box box-default">
  <!-- box-header -->
  <div class="box-header"> 
    <h3 class="box-title">用户列表</h3>
    <div class="box-tools pull-right">
      <button type="button" class="btn btn-box-tool" data-toggle="modal" data-target="#createUserModal">
        <i class="fa fa-plus"></i>
      </button>
    </div>
  </div>
  <!-- /.box-header -->
  <!-- box-body -->
  <div class="box-body">
    <table id="userListTable" class="table table-bordered table-striped">
      <thead>
        <tr>
          <th>ID</th>
          <th>用户名</th>
          <th>用户组</th>
          <th>创建者</th>
          <th>操作</th>
        </tr>
      </thead>
      <tbody>
      	{{range $index, $user := .UserList}}
      		<tr>
      			<td>{{$user.Id}}</td>
      			<td>{{$user.Username}}</td>
      			<td></td>
      			<td>{{$user.Creator}}</td>
            <td>
              {{if ne $user.Id $.AdminID}}
                <button id="{{$user.Id}}" row="{{$index}}" type="button" class="btn btn-xs btn-info" style="width:30px" onclick="showEditUserModal(this);">
                  <i class="fa fa-edit"></i>
                </button>
                {{if ne $user.Id $.UserData.Id}}
                  <button id="{{$user.Id}}" row="{{$index}}" type="button" class="btn btn-xs btn-danger" style="width:30px" onclick="showDelUserModal(this);">
                    <i class="fa fa-trash"></i>
                  </button>
                {{end}}
              {{end}}
            </td>
      		</tr>
      	{{end}}
      </tbody>
      <tfoot>
        <tr>
          <th>ID</th>
          <th>用户名</th>
          <th>用户组</th>
          <th>创建者</th>
          <th>操作</th>
        </tr>
      </tfoot>
    </table>
  </div>
  <!-- /.box-body -->
</div>
<!-- /.box -->
