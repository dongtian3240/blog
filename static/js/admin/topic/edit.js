
/**新建博客**/
var EditTopic = {
	     globalUserId:0,
		init:function(){
			EditTopic.globalUserId = $('#_globalUserId').val();
			EditTopic.ctrl.initEvent();
			EditTopic.ctrl.findDefaultLabel();
			EditTopic.ctrl.initValidator();
		},
		ctrl:{
			
			findDefaultLabel:function(){
				
				var id = $('#topicId').val();
				var formDat = {"topicId":id};
				$.ajax({
						url:"/admin/"+EditTopic.globalUserId+"/label/findLabelListByTopicId",
						cache:false,
						data:formDat,
						dataType:'json',
						type:'POST',
						error:function(XMLHttpRequest, textStatus, errorThrown){
							
						},
						success:function(data, textStatus, jqXHR) {
							
							if(data.Success == true) {
								
								$.each(data.Data,function(ind,va){
										
										$('#labelCheck_'+va.LabelId).attr("checked",true);
								});
								
							} else {
								
							}
							
						}
					});
				
			},
			initEvent:function(){
				
				$('#btn_blog_publish').on('click',function(){
					var bootstrapValidator = $('#editTopicForm').data('bootstrapValidator')
                   bootstrapValidator.validate();
				 var sec =  bootstrapValidator.isValid();
                  if(sec == true){
					var content = $('#Content').val();
					if($.trim(content) == "") {
						 layer.msg("博客内容不能为空!");
						 return false;
					} else{
						EditTopic.ctrl.newTopic();
					}
				  }
					return false;
				});
			},
			initValidator:function(){
				
					$('#editTopicForm').bootstrapValidator({
            message: '不能为空',
            feedbackIcons: {
                valid: 'glyphicon glyphicon-ok',
                invalid: 'glyphicon glyphicon-remove',
                validating: 'glyphicon glyphicon-refresh'
            },
            fields: {
                LabelId:{
                    message: '博客分类不能为空!',
                    validators: {
                        notEmpty: {
                            message: '请至少选择一个博客分类'
                        }                        
                    }
                },
                Title:{
	                message:'博客标题不能为空',
	                validators:{
		                notEmpty: {
                            message: '博客标题不能为空'
                        }
	                }
                 }
            	}
        	});
			},
			//新建博客
			newTopic:function(){
				
			var formDat = $('#editTopicForm').serialize();
			layer.load(1);
			$.ajax({
						url:"/admin/"+EditTopic.globalUserId+"/topic/edit",
						cache:false,
						data:formDat,
						dataType:'json',
						type:'POST',
						error:function(XMLHttpRequest, textStatus, errorThrown){
							//关闭进度
							layer.closeAll('loading');
							layer.msg('提交失败,请重试!',{icon: 8,time: 1000});
						},
						success:function(data, textStatus, jqXHR) {
							//关闭进度
							layer.closeAll('loading');
							if(data.Success == true) {
								$('#editLabelwrap').dialog('close');
								 layer.msg(data.Message,{icon: 1,time: 1000}, function(){
										window.location.href = "/admin/"+EditTopic.globalUserId+"/topic";
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
	
	EditTopic.init();
});