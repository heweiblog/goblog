<!DOCTYPE html>
<html>

<head>

  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">

  <title>{{config "String" "app.cname" "博客登录"}}</title>

  <link href="/static/layui/css/layui.css" rel="stylesheet" media="all"/>
  <script src="/static/layui/lay/dest/layui.all.js"></script>

  <script type="text/javascript">
    
  </script>

  <style type="text/css">

  </style>

</head>

<body>


  
<div style="margin: auto;width: 500px;height: 500px;">

<fieldset class="layui-elem-field layui-field-title" style="margin-top: 50px;">
  <legend>博客 登录</legend>
</fieldset>  
<form class="layui-form  layui-form-pane" action="/login" method="POST">
  <div class="layui-form-item">
    <label class="layui-form-label">用户名</label>
    <div class="layui-input-block">
      <input type="text" name="UserName" lay-verify="required" autocomplete="off" placeholder="请输入用户名" class="layui-input">
    </div>
  </div>
  <div class="layui-form-item">
    <label class="layui-form-label">密  码</label>
    <div class="layui-input-block">
      <input type="password" name="PassWord" lay-verify="required" placeholder="请输入密码" autocomplete="off" class="layui-input">
    </div>
  </div>

  <div class="checkbox">
  	<label>
	  <input type="checkbox" name="AutoLogin">自动登录
	</label>
  </div>

  <div class="layui-form-item">
    <div class="layui-input-block">
      <button class="layui-btn" lay-submit="" lay-filter="demo1">登录</button>
    </div>
  </div>
</form>
</div>

</body>

</html>
