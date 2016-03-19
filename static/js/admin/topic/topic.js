

var Topic = {
	 init:function(){
		Topic.ctrl.initEvent();
	},
	ctrl:{
		
		deleteTopic:function(id){
			var formDat = {"topicId":id};
			layer.load(1);
			$.ajax({
						url:'/admin/topic/deleteTopic',
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
			
		},
		initEvent:function(){
			$('#btn_newTopic').on('click',function(){
					
					window.location.href = "/admin/topic/new";
			});
			
			$('.btn_editTopic').on('click',function(){
				var id = $(this).attr("va")
				window.location.href = "/admin/topic/edit?id=" + id;
			});
			
			$('.btn_deleteTopic').on('click',function(){
				
				var va = $(this).attr("va");
				layer.confirm('确认删除此博客吗？', {icon: 4,time: 1000}, function(index){
				    					layer.close(index);
				    					 Topic.ctrl.deleteTopic(va);
				    				});
			});
		}
	}
};


$(function(){
	Topic.init();
});