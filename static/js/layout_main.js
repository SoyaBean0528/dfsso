$(function() {  
    $("[data-toggle='popover']").popover({  
        html : true,    
        title: title(),   
        placement: "bottom", 
        delay:{show:0, hide:0},  
        content: function() {  
          return content();    
        }   
    });  
}); 

function title() {
    return '<div style="height:20px;line-height:20px">' +
           '<b style="color:black;">请选择头像</b>' +
           '<button type="button" class="btn btn-xs btn-default" style="float:right" onclick="">' +
                '上传' +
           '</button>' +
           '</div>';
}

function content() {
    var data = $('<table class="table no-border">'+
                '<tr>' +
                    '<td>' +
                        '<img src="../static/img/avatar/avatar1.png" class="img-circle" alt="User Image" style="height:50px;border:2px solid;border-color:transparent;border-color:rgba(0, 0, 0, 0.3);" >' +
                    '</td>' +
                    '<td>' +
                        '<img src="../static/img/avatar/avatar2.png" class="img-circle" alt="User Image" style="height:50px;border:2px solid;border-color:transparent;border-color:rgba(0, 0, 0, 0.3);" >' +
                    '</td>' +
                    '<td>' +
                        '<img src="../static/img/avatar/avatar3.png" class="img-circle" alt="User Image" style="height:50px;border:2px solid;border-color:transparent;border-color:rgba(0, 0, 0, 0.3);" >' +
                    '</td>' +
                '</tr>' +
                '<tr>' +
                    '<td>' +
                        '<img src="../static/img/avatar/avatar4.png" class="img-circle" alt="User Image" style="height:50px;border:2px solid;border-color:transparent;border-color:rgba(0, 0, 0, 0.3);" >' +
                    '</td>' +
                    '<td>' +
                        '<img src="../static/img/avatar/avatar5.png" class="img-circle" alt="User Image" style="height:50px;border:2px solid;border-color:transparent;border-color:rgba(0, 0, 0, 0.3);" >' +
                    '</td>' +
                '</tr>' +
                '</table>');
    
    return data;
}

function test() {
    alert('关注成功');
}