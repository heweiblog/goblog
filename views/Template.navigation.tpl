{{define "navigation"}}
<nav class="navbar navbar-default navbar-fixed-top">
	 <div class="container">
    <!--导航区域开始-->
		<a class="navbar-brand" href="/">我的博客</a>
        <ul class="nav navbar-nav navbar-left">
            <li {{if .IsHome}}class="active"{{end}}><div class="nav-h1"></div><a href="/">首页</a></li>
            <li {{if .IsCategory}}class="active"{{end}}><div class="nav-h2"></div><a href="/category">分类</a></li>
            <li {{if .IsTopic}}class="active"{{end}}><div class="nav-h3"></div><a href="/topic">文章</a></li>
            <li {{if .IsArchtive}}class="active"{{end}}><div class="nav-h4"></div><a href="/archtive">归档</a></li>
            <li {{if .IsTag}}class="active"{{end}}><div class="nav-h5"></div><a href="/tag">标签</a></li>
            <li {{if .IsAbout}}class="active"{{end}}><div class="nav-h6"></div><a href="/about">关于</a></li>
            <li><div class="nav-h6"></div><a href="/404">公益404</a></li>
        </ul>
    <!--导航区域开始-->
		<div class="pull-right">
			<ul class="nav navbar-nav">
				{{if .IsLogin}}
				<li><a href="/login?exist=true">退出登录</a></li>
				{{else}}
				<li><a href="/login">管理员登录</a></li>
				{{end}}	
			</ul>
		</div>
	</div>
</nav>
{{end}}
