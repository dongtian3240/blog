
var Label = {
	globalUserId:0,
	newLabelValidator:'',
	 init:function(){
		Label.globalUserId = $('#_globalUserId').val();
		Label.ctrl.initEvent();
		Label.ctrl.initDialog();
		Label.ctrl.initValidator();
	},
	ctrl:{
		
		initValidator:function(){
			
			$('#newLabelform').bootstrapValidator({
            message: '不能为空',
            feedbackIcons: {
                valid: 'glyphicon glyphicon-ok',
                invalid: 'glyphicon glyphicon-remove',
                validating: 'glyphicon glyphicon-refresh'
            },
            fields: {
                labelName: {
                    message: '分类名称不能为空!',
                    validators: {
                        notEmpty: {
                            message: '分类名称不能为空'
                        },
                        stringLength: {
                            min: 3,
                            max: 15,
                            message: '用户必须大于3个字符小于15个字符'
                        }
                        /*remote: {
                            url: 'remote.php',
                            message: 'The username is not available'
                        },*/
                        
                    }
                }
            }
        });
		
		
			$('#editLabelform').bootstrapValidator({
            message: '不能为空',
            feedbackIcons: {
                valid: 'glyphicon glyphicon-ok',
                invalid: 'glyphicon glyphicon-remove',
                validating: 'glyphicon glyphicon-refresh'
            },
            fields: {
                labelName: {
                    message: '分类名称不能为空!',
                    validators: {
                        notEmpty: {
                            message: '分类名称不能为空'
                        },
                        stringLength: {
                            min: 3,
                            max: 15,
                            message: '用户必须大于3个字符小于15个字符'
                        }
                        /*remote: {
                            url: 'remote.php',
                            message: 'The username is not available'
                        },*/
                        
                    }
                }
            }
        });
			
		},
		initDialog:function(){
			
			$('#newLabelwrap').dialog({
				    title:"新建分类",
				   backdrop: "static",
				autoOpen:false,
				onClose: function() { 
				$(this).dialog("close"); 
				  } ,
             buttons: [
            {
                text: "关闭"
              , 'class': "btn-primary"
              , click: function() {
                  $(this).dialog("close");
                }
            },
            {
                text: "保存"
              , 'class': "btn-info"
              , click: function() {
                  /*your login handler*/
 					var bootstrapValidator = $('#newLabelform').data('bootstrapValidator')
                   bootstrapValidator.validate();
				 var sec =  bootstrapValidator.isValid();
                  if(sec == true){
					Label.ctrl.saveLabel();
				  }
                }
            }]
			});
			
			$('#editLabelwrap').dialog({
				    title:"修改分类",
				   backdrop: "static",
				autoOpen:false,
				onClose: function() { 
				$(this).dialog("close"); 
				  } ,
             buttons: [
            {
                text: "关闭"
              , 'class': "btn-primary"
              , click: function() {
                  $(this).dialog("close");
                }
            },
            {
                text: "保存"
              , 'class': "btn-info"
              , click: function() {
                  /*your login handler*/
 				var bootstrapValidator = $('#editLabelform').data('bootstrapValidator')
                   bootstrapValidator.validate();
				 var sec =  bootstrapValidator.isValid();
                  if(sec == true){
					 Label.ctrl.updateLabel();
				  }
                 
                }
            }]
			});
		},
		initEvent:function(){
			$('#btn_newLabel').on('click',function(){
				$('#newLabelwrap').dialog('open');
			});
			$('.btn_editLabel').on('click',function(){
				
				var va = $(this).attr("va");
				
					Label.ctrl.findLabelById(va);
				
			});
			
			$('.btn_deleteLabel').on('click',function(){
				var va = $(this).attr("va");
				layer.confirm('确认删除此分类吗？', {icon: 4,time: 1000}, function(index){
				    					layer.close(index);
				    					 Label.ctrl.deleteLabel(va);
				    				});
				
			});
		},
		
		updateLabel:function(){
			var formDat = $('#editLabelform').serialize();
			layer.load(1);
			$.ajax({
						url:"/admin/"+Label.globalUserId+"/label/update",
						cache:false,
						data:formDat,
						dataType:'json',
						type:'POST',
						error:function(XMLHttpRequest, textStatus, errorThrown){
							//关闭进度
							layer.closeAll('loading');
							layer.msg('提交失败,请重试!',{icon: 8});
						},
						success:function(data, textStatus, jqXHR) {
							//关闭进度
							layer.closeAll('loading');
							if(data.Success == true) {
								$('#editLabelwrap').dialog('close');
								 layer.msg(data.Message,{icon: 1,time: 1000}, function(){
										window.location.reload();
									});
							} else {
								layer.msg(data.Message,{icon: 8});
							}
							
						}
					});
			
			
			
		},
		findLabelById:function(id){
			
			var formDat = {"id":id};
			layer.load(1);
			$.ajax({
						url:"/admin/"+Label.globalUserId+"/label/findLabelById",
						cache:false,
						data:formDat,
						dataType:'json',
						type:'POST',
						error:function(XMLHttpRequest, textStatus, errorThrown){
							//关闭进度
							layer.closeAll('loading');
							layer.msg('提交失败,请重试!',{icon: 8});
						},
						success:function(data, textStatus, jqXHR) {
							//关闭进度
							layer.closeAll('loading');
							if(data.Success == true) {
								
								$('#editlabelName').val(data.Data.LabelName)
								$('#labelId').val(data.Data.LabelId);
								$('#editLabelwrap').dialog('open');
								 
							} else {
								layer.msg(data.Message,{icon: 8});
							}
							
						}
					});
		},
		saveLabel:function(){
			
			var formDat = $('#newLabelform').serialize();
			layer.load(1);
			$.ajax({
						url:"/admin/"+Label.globalUserId+"/label/new",
						cache:false,
						data:formDat,
						dataType:'json',
						type:'POST',
						error:function(XMLHttpRequest, textStatus, errorThrown){
							//关闭进度
							layer.closeAll('loading');
							layer.msg('提交失败,请重试!',{icon: 8});
						},
						success:function(data, textStatus, jqXHR) {
							//关闭进度
							layer.closeAll('loading');
							if(data.Success == true) {
								$('#newLabelwrap').dialog('close');
								 layer.msg(data.Message,{icon: 1,time: 1000}, function(){
										window.location.reload();
									});
							} else {
								layer.msg(data.Message,{icon: 8});
							}
							
						}
					});
			
		},
		deleteLabel:function(id){
			
			var formDat = {"id":id};
			layer.load(1);
			$.ajax({
						url:"/admin/"+Label.globalUserId+"/label/delete",
						cache:false,
						data:formDat,
						dataType:'json',
						type:'POST',
						error:function(XMLHttpRequest, textStatus, errorThrown){
							//关闭进度
							layer.closeAll('loading');
							layer.msg('提交失败,请重试!',{icon: 8});
						},
						success:function(data, textStatus, jqXHR) {
							//关闭进度
							layer.closeAll('loading');
							if(data.Success == true) {
								
								 layer.msg(data.Message,{icon: 1,time: 1000}, function(){
										window.location.reload();
									});
							} else {
								layer.msg(data.Message,{icon: 8});
							}
							
						}
					});
			
		}
	}
};

$(function(){
	
	Label.init();
});