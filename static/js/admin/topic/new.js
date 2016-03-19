
/**新建博客**/
var NewTopic = {
	
		init:function(){
			NewTopic.ctrl.initEvent();
			NewTopic.ctrl.initValidator();
		},
		ctrl:{
			
			
			initEvent:function(){
				
				$('#btn_blog_publish').on('click',function(){
					var bootstrapValidator = $('#newTopicForm').data('bootstrapValidator')
                   bootstrapValidator.validate();
				 var sec =  bootstrapValidator.isValid();
                  if(sec == true){
					var content = $('#Content').val();
					if($.trim(content) == "") {
						 layer.msg("博客内容不能为空!");
						 return false;
					} else{
						NewTopic.ctrl.newTopic();
					}
				  }
					return false;
				});
			},
			initValidator:function(){
				
					$('#newTopicForm').bootstrapValidator({
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
				
			var formDat = $('#newTopicForm').serialize();
			layer.load(1);
			$.ajax({
						url:'/admin/topic/new',
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
										window.location.href = "/admin/topic";
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
	
	NewTopic.init();
});