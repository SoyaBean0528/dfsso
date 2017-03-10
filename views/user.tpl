<!-- box -->
<div class="box box-default">
  <!-- box-header -->
  <div class="box-header"> <h3 class="box-title">用户列表</h3> </div>
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
        </tr>
      </thead>
      <tbody>
      	{{range .UserList}}
      		<tr>
      			<td>{{.Id}}</td>
      			<td>{{.Username}}</td>
      			<td></td>
      			<td></td>
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
