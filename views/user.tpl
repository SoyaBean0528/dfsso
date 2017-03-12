<div class="modal modal-primary fade" id="createUserModal" tabindex="-1" role="dialog" aria-labelledby="create-user-label" aria-hidden="true">
  <div class="modal-dialog" style="width:380px;margin:10% auto">
    <div class="modal-content">
      <div class="modal-header">
        <button type="button" class="close" data-dismiss="modal" aria-label="Close">
          <span aria-hidden="true">&times;</span></button>
        <h4 class="modal-title" id="create-user-label">
          添加用户
          {{if .Msg}}
            <small id="addUserErrMsg" style="color:#f56954">{{.Msg}}</small>
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
          <div class="form-group">
            <label for="password" class="col-sm-3 control-label">密码</label>
            <div class="col-sm-8">
              <input type="text" class="form-control" name="password" id="password" placeholder="密码">
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
<!-- box -->
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
    <table id="userList" class="table table-bordered table-striped">
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
      	{{range .UserList}}
      		<tr>
      			<td>{{.Id}}</td>
      			<td>{{.Username}}</td>
      			<td></td>
      			<td>{{.Creator}}</td>
            <td>
              {{if ne .Id $.AdminID }}
                <div class="btn-group-xs">
                  <button type="button" class="btn btn-xs btn-info col-xs-2 col-xs-push-0">
                    <i class="fa fa-edit"></i>
                  </button>
                  <button type="button" class="btn btn-xs btn-danger col-xs-2 col-xs-push-1">
                    <i class="fa fa-trash"></i>
                  </button>
                </div>
              {{end}}
            </td>
      		</tr>
      	{{end}}
      	<!--
        <tr>
          <td>Trident</td>
          <td>Internet
            Explorer 4.0
          </td>
          <td>Win 95+</td>
          <td> 4</td>
          <td>X</td>
        </tr>
        -->
      </tbody>
      <tfoot>
        <tr>
          <th>ID</th>
          <th>用户名</th>
          <th>用户组</th>
          <th>创建者</th>
        </tr>
      </tfoot>
    </table>
  </div>
  <!-- /.box-body -->
</div>
<!-- /.box -->
