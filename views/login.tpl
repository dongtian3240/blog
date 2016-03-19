<!DOCTYPE html>
<html lang="en">

<head>

    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <meta name="description" content="">
    <meta name="author" content="">

    <title>请登录</title>

    <!-- Bootstrap Core CSS -->
    <link href="/static/sbadmin/bower_components/bootstrap/dist/css/bootstrap.min.css" rel="stylesheet">

    <!-- MetisMenu CSS -->
    <link href="/static/sbadmin/bower_components/metisMenu/dist/metisMenu.min.css" rel="stylesheet">

    <!-- Custom CSS -->
    <link href="/static/sbadmin/css/sb-admin-2.css" rel="stylesheet">

    <!-- Custom Fonts -->
    <link href="/static/sbadmin/bower_components/font-awesome/css/font-awesome.min.css" rel="stylesheet" type="text/css">
	
		<!--validator-->
 <link href="/static/validator/css/bootstrapValidator.css" rel="stylesheet" type="text/css">
    <!-- HTML5 Shim and Respond.js IE8 support of HTML5 elements and media queries -->
    <!-- WARNING: Respond.js doesn't work if you view the page via file:// -->
    <!--[if lt IE 9]>
       <script src="/static/bootstrap/js/html5shiv.min.js"></script>
      <script src="/static/bootstrap/js/1.4.2/respond.min.js"></script>
    <![endif]-->

</head>

<body>

    <div class="container">
        <div class="row">
            <div class="col-md-4 col-md-offset-4">
                <div class="login-panel panel panel-default">
                    <div class="panel-heading">
                        <h3 class="panel-title">请登录</h3>
                    </div>
                    <div class="panel-body">
                        <form id="loginForm" action="/login" method="POST" role="form">
                            <fieldset>
                                <div class="form-group">
                                    <input class="form-control" value="admin" placeholder="账号" name="userName" type="text" autofocus>
                                </div>
                                <div class="form-group">
                                    <input class="form-control" value="123456" placeholder="密码" name="password" type="password" value="">
                                </div>
                                <div class="checkbox">
                                    <label>
                                        <input  checked name="remember" type="checkbox" value="true" >记住我
                                    </label>
                                </div>
                                <!-- Change this to a button or input when using this as a form -->
                                <button type="submit"  class="btn btn-lg btn-success btn-block">登录</button>
                            </fieldset>
                        </form>
                    </div>
                </div>
            </div>
        </div>
    </div>


<br><br>
<hr>
<br>
<footer class="footer ">
      <div class="container">
        <div class="row footer-top">
          <div class="col-sm-6 col-lg-6">
            <h4>
              <img src="https://ss1.baidu.com/6ONXsjip0QIZ8tyhnq/it/u=619180273,1394682420&fm=58">
            </h4>
          </div>
          <div class="col-sm-6  col-lg-5 col-lg-offset-1">
            <div class="row about">
              <div class="col-xs-3">
                <h4>关于</h4>
                <ul class="list-unstyled">
                  <li><a href="/about/">关于我们</a></li>
                  <li><a href="/ad/">广告合作</a></li>
                  <li><a href="/links/">友情链接</a></li>
                  <li><a href="/hr/">招聘</a></li>
                </ul>
              </div>
              <div class="col-xs-3">
                <h4>联系方式</h4>
                <ul class="list-unstyled">
                  <li><a target="_blank" title="Bootstrap中文网官方微博" href="http://weibo.com/bootcss">新浪微博</a></li>
                  <li><a href="mailto:admin@bootcss.com">电子邮件</a></li>
                </ul>
              </div>
              <div class="col-xs-3">
                <h4>旗下网站</h4>
                <ul class="list-unstyled">
                  <li><a target="_blank" href="http://www.golaravel.com/">Laravel中文网</a></li>
                  <li><a target="_blank" href="http://www.ghostchina.com/">Ghost中国</a></li>
                </ul>
              </div>
              <div class="col-xs-3">
                <h4>赞助商</h4>
                <ul class="list-unstyled">
                  <li><a target="_blank" href="http://www.ucloud.cn/">UCloud</a></li>
                  <li><a target="_blank" href="https://www.upyun.com">又拍云</a></li>
                </ul>
              </div>
            </div>
    
          </div>
        </div>
        <hr>
        <div class="row footer-bottom">
          <ul class="list-inline text-center">
            <li><a target="_blank" href="http://www.miibeian.gov.cn/">京ICP备11008151号</a></li><li>京公网安备11010802014853</li>
          </ul>
        </div>
      </div>
    </footer>
    <!-- jQuery -->
    <script src="/static/sbadmin/bower_components/jquery/dist/jquery.min.js"></script>

    <!-- Bootstrap Core JavaScript -->
    <script src="/static/sbadmin/bower_components/bootstrap/dist/js/bootstrap.min.js"></script>
    <script src="/static/validator/js/bootstrapValidator.js"></script>
    <script src="/static/validator/js/language/zh_CN.js"></script>
    <!-- Metis Menu Plugin JavaScript -->
    <script src="/static/sbadmin/bower_components/metisMenu/dist/metisMenu.min.js"></script>

<!--layer-->
 <script src="/static/layer/layer.js"></script>
    <!-- Custom Theme JavaScript -->
    <script src="/static/sbadmin/js/sb-admin-2.js"></script>
	
	<script type="text/javascript" src="/static/js/login.js"></script>

  </body>
</html>