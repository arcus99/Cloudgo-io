<!DOCTYPE html>
<!-- saved from url=(0035)http://write.blog.csdn.net/mdeditor -->
<html manifest="cache.manifest" style="overflow: hidden;"><head><meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
    <script src="./README_files/hm.js.下载"></script><script type="text/javascript" src="./README_files/tingyun-rum.js.下载"></script>
    <title>CSDN-markdown编辑器</title>
    <meta name="viewport" content="user-scalable=no, width=device-width, initial-scale=1, maximum-scale=1">
    <meta name="mobile-web-app-capable" content="yes">
    <meta name="apple-mobile-web-app-capable" content="yes">
    <meta name="apple-mobile-web-app-status-bar-style" content="black">
    <meta name="msvalidate.01" content="5E47EE6F67B069C17E3CDD418351A612">   
  
    <script src="./README_files/jquery.min.js.下载"></script>
          
         
    <style type="text/css">
        body {
            margin: 0;
        }

        .csdn-toolbar {
            position: absolute !important;
            top: 0;
            left: 0;
            z-index: 1005;
            width: 100% !important;
            padding: 0 35px !important;
            transition: all .5s;
            box-sizing: border-box;
        }

            .csdn-toolbar .container {
                width: 100% !important;
            }

        #ie8 {
            display: none;
            display: block \9;
            width: 500px;
            padding: 50px 0;
            border: 1px solid #ddd;
            background: #eee;
            text-align: center;
            font-size: 24px;
            margin: 200px 0 0 -250px;
            position: absolute;
            left: 50%;
            z-index: -1;
        }
    </style>
    <script type="text/javascript">

        //var isExpert = '';

        // Use ?debug to serve original JavaScript files instead of minified
               
            window.baseDir = 'http://static.blog.csdn.net/public/res'

            if (!/(\?|&)debug($|&)/.test(location.search)) {
                window.baseDir += '-min';
            }
            window.require = {
                baseUrl: window.baseDir,
                deps: ['main'],
                urlArgs: 'v=1026'
            };
           document.domain = "csdn.net";
        
        
        //document.domain = "csdn.net";

        //var addPushCheckInterval = setInterval(function () {
        //    var modal = $(".modal-content:visible");
        //    var modaltitle = modal.find(".modal-title");
        //    if (modaltitle != undefined && modal.length > 0 && $(modaltitle).html() == "文章设置") {
        //        $("#csdn-at-blog-checkbox").hide();
        //        $(".modal-footer #csdn-at-blog-checkbox").parent().find("span").hide();
        //        $("#bole_help_bt").hide();
        //        clearInterval(addPushCheckInterval);
        //    }
        //}, 500);
                             
        setTimeout(function () {
            $("#csdnEditor").keyup(function () {
                writedisplay(500);

                setTimeout(function () {
                    $("#preview-contents a[target!='_blank']").attr("target", "_blank");
                }, 1000);
                setTimeout(function () {
                    $("#preview-contents a[target!='_blank']").attr("target", "_blank");
                }, 2000);
            });

            $(".btn.btn-danger.action-insert-link").click(function () {
                setTimeout(function () {
                    $("#preview-contents a[target!='_blank']").attr("target", "_blank");
                }, 1000);
                setTimeout(function () {
                    $("#preview-contents a[target!='_blank']").attr("target", "_blank");
                }, 2000);
            });

            $("#preview-contents a[target!='_blank']").attr("target", "_blank");
        }, 500);
        
        writedisplay(3000);
        writedisplay(5000);
        writedisplay(8000);
        function writedisplay(time)
        {
            setTimeout(function () {               
                $(".math[role='math']").each(function (index, value) { $(this).find("span").last().css("color", "#f9f9f9"); });
                $("#preview-contents a[target!='_blank']").attr("target", "_blank");
            }, time);
        }
        var isval = false;
        var addPushSelectInterval = setInterval(function () {
            var modal = $(".modal-content:visible");
            var modaltitle = modal.find(".modal-title");
            if (modaltitle != undefined && modal.length > 0 && $(modaltitle).html() == "文章设置") {                
                if (isval == false) {
                    setchannel();
                    isval = true;
                }
                clearInterval(addPushSelectInterval);
            }
        }, 500);
        function setchannel()
        {
           var selid=$("#input-blog-channel").val();
            var html = '<option value="0" selected="selected">请选择</option>';
             var json = "[{\"Id\":28,\"Name\":\"人工智能\",\"Alias\":\"ai\",\"OrderNumber\":1,\"ParentId\":0,\"Type\":3,\"Status\":0,\"Description\":\"\"},{\"Id\":1,\"Name\":\"移动开发\",\"Alias\":\"mobile\",\"OrderNumber\":2,\"ParentId\":0,\"Type\":3,\"Status\":0,\"Description\":\"iOS,Android,Windows Phone,webOS,apple,mobile,近期\\r\\n最受欢迎的博客，CSDN博客频道\"},{\"Id\":29,\"Name\":\"物联网\",\"Alias\":\"iot\",\"OrderNumber\":3,\"ParentId\":0,\"Type\":3,\"Status\":0,\"Description\":\"\"},{\"Id\":15,\"Name\":\"架构\",\"Alias\":\"enterprise\",\"OrderNumber\":4,\"ParentId\":0,\"Type\":3,\"Status\":0,\"Description\":\"Hibernate,Spring,Struts,框架,安全,企业应用，近期最\\r\\n受欢迎的博客，CSDN博客频道\"},{\"Id\":2,\"Name\":\"云计算/大数据\",\"Alias\":\"cloud\",\"OrderNumber\":5,\"ParentId\":0,\"Type\":3,\"Status\":0,\"Description\":\"云计算,云存储,云平台,Cloud computing,近期最受欢迎\\r\\n的博客，CSDN博客频道\"},{\"Id\":30,\"Name\":\"游戏开发\",\"Alias\":\"game\",\"OrderNumber\":6,\"ParentId\":0,\"Type\":3,\"Status\":0,\"Description\":\"\"},{\"Id\":12,\"Name\":\"运维\",\"Alias\":\"system\",\"OrderNumber\":8,\"ParentId\":0,\"Type\":3,\"Status\":0,\"Description\":\"Windows,Linux,Ubuntu,网络管理,网管,系统分析,MacOS\\r\\n，近期最受欢迎的博客，CSDN博客频道\"},{\"Id\":6,\"Name\":\"数据库\",\"Alias\":\"database\",\"OrderNumber\":9,\"ParentId\":0,\"Type\":3,\"Status\":0,\"Description\":\"Oracle,MySQL,NoSQL,SQL Server,PostgreSQL,MongoDB，\\r\\n近期最受欢迎的博客，CSDN博客频道\"},{\"Id\":14,\"Name\":\"前端\",\"Alias\":\"web\",\"OrderNumber\":10,\"ParentId\":0,\"Type\":3,\"Status\":0,\"Description\":\"JavaScript,CSS,jQuery,HTML5,firebug,近期最受欢迎的\\r\\n博客，CSDN博客频道\"},{\"Id\":16,\"Name\":\"编程语言\",\"Alias\":\"code\",\"OrderNumber\":12,\"ParentId\":0,\"Type\":3,\"Status\":0,\"Description\":\"Java,Ruby,Rails,PHP,Python,.net,C#,C,C++,JSP,ASP.net,ASP,JDK,编程,Scala，近期最\\r\\n受欢迎的博客，CSDN博客频道\"},{\"Id\":3,\"Name\":\"研发管理\",\"Alias\":\"software\",\"OrderNumber\":13,\"ParentId\":0,\"Type\":3,\"Status\":0,\"Description\":\"项目管理,软件测试,敏捷开发,Git,SVN，近期最受欢迎的\\r\\n博客，CSDN博客频道\"},{\"Id\":32,\"Name\":\"安全\",\"Alias\":\"safe\",\"OrderNumber\":14,\"ParentId\":0,\"Type\":3,\"Status\":0,\"Description\":\"\"},{\"Id\":33,\"Name\":\"程序人生\",\"Alias\":\"programlife\",\"OrderNumber\":15,\"ParentId\":0,\"Type\":3,\"Status\":0,\"Description\":\"\"},{\"Id\":34,\"Name\":\"区块链\",\"Alias\":\"block\",\"OrderNumber\":16,\"ParentId\":0,\"Type\":3,\"Status\":0,\"Description\":\"\"},{\"Id\":35,\"Name\":\"音视频开发\",\"Alias\":\"audiovideo\\n\",\"OrderNumber\":17,\"ParentId\":0,\"Type\":3,\"Status\":0,\"Description\":\"\"},{\"Id\":36,\"Name\":\"资讯\",\"Alias\":\"news\",\"OrderNumber\":18,\"ParentId\":0,\"Type\":3,\"Status\":0,\"Description\":\"\"},{\"Id\":37,\"Name\":\"计算机理论与基础\",\"Alias\":\"basics\",\"OrderNumber\":19,\"ParentId\":0,\"Type\":3,\"Status\":0,\"Description\":\"\"}]";
            $.each(eval(json), function (index, item) {
                var obj = item;
                html += '<option value="'+obj.Id+'">'+obj.Name +'</option>';
            });
            $("#input-blog-channel").html(html).val(selid);
        }

    </script>   
             

    <!--<script src="../public/res-min/require.js"></script>-->
      
    
    
    

    <link href="./README_files/index.css" media="screen" rel="stylesheet" type="text/css">
<script type="text/javascript" charset="utf-8" async="" data-requirecontext="_" data-requiremodule="main" src="./README_files/main.js.下载"></script><link type="text/css" rel="stylesheet" href="./README_files/default.css"><script type="text/x-mathjax-config;executed=true">MathJax.Hub.Config({
	skipStartupTypeset: true,
    "HTML-CSS": {
        preferredFont: "TeX",
        availableFonts: [
            "STIX",
            "TeX"
        ],
        linebreaks: {
            automatic: true
        },
        EqnChunk: 10,
        imageFont: null
    },
    tex2jax: { inlineMath: [["$","$"],["\\\\(","\\\\)"]], displayMath: [["$$","$$"],["\\[","\\]"]], processEscapes: true },
    TeX: $.extend({
        noUndefined: {
            attributes: {
                mathcolor: "red",
                mathbackground: "#FFEEEE",
                mathsize: "90%"
            }
        },
        Safe: {
            allow: {
                URLs: "safe",
                classes: "safe",
                cssIDs: "safe",
                styles: "safe",
                fontsize: "all"
            }
        }
    }, undefined),
    messageStyle: "none"
});
</script><script type="text/javascript" charset="utf-8" async="" data-requirecontext="_" data-requiremodule="mathjax" src="./README_files/MathJax.js.下载"></script><style type="text/css">@media (min-width: 0px) {
#wmd-input {
   font-size: 16px;
}
#preview-contents {
   font-size: 16px;
}
}@media (min-width: 600px) {
#wmd-input {
   font-size: 17px;
}
#preview-contents {
   font-size: 17px;
}
}@media (min-width: 1200px) {
#wmd-input {
   font-size: 18px;
}
#preview-contents {
   font-size: 18px;
}
}</style><style type="text/css">@keyframes working-indicator-bar0 {
    0% { opacity:0.25; }
    0.01% { opacity:0.25; }
    0.02% { opacity:1; }
    50.01% { opacity:0.25; }
    100% { opacity:0.25; }
}@-webkit-keyframes working-indicator-bar0 {
    0% { opacity:0.25; }
    0.01% { opacity:0.25; }
    0.02% { opacity:1; }
    50.01% { opacity:0.25; }
    100% { opacity:0.25; }
}@keyframes working-indicator-bar1 {
    0% { opacity:0.25; }
    20.01% { opacity:0.25; }
    20.02% { opacity:1; }
    70.01% { opacity:0.25; }
    100% { opacity:0.25; }
}@-webkit-keyframes working-indicator-bar1 {
    0% { opacity:0.25; }
    20.01% { opacity:0.25; }
    20.02% { opacity:1; }
    70.01% { opacity:0.25; }
    100% { opacity:0.25; }
}@keyframes working-indicator-bar2 {
    0% { opacity:0.25; }
    40.01% { opacity:0.25; }
    40.02% { opacity:1; }
    90.01% { opacity:0.25; }
    100% { opacity:0.25; }
}@-webkit-keyframes working-indicator-bar2 {
    0% { opacity:0.25; }
    40.01% { opacity:0.25; }
    40.02% { opacity:1; }
    90.01% { opacity:0.25; }
    100% { opacity:0.25; }
}</style><style type="text/css">.MathJax_Hover_Frame {border-radius: .25em; -webkit-border-radius: .25em; -moz-border-radius: .25em; -khtml-border-radius: .25em; box-shadow: 0px 0px 15px #83A; -webkit-box-shadow: 0px 0px 15px #83A; -moz-box-shadow: 0px 0px 15px #83A; -khtml-box-shadow: 0px 0px 15px #83A; border: 1px solid #A6D ! important; display: inline-block; position: absolute}
.MathJax_Hover_Arrow {position: absolute; width: 15px; height: 11px; cursor: pointer}
</style><style type="text/css">#MathJax_About {position: fixed; left: 50%; width: auto; text-align: center; border: 3px outset; padding: 1em 2em; background-color: #DDDDDD; color: black; cursor: default; font-family: message-box; font-size: 120%; font-style: normal; text-indent: 0; text-transform: none; line-height: normal; letter-spacing: normal; word-spacing: normal; word-wrap: normal; white-space: nowrap; float: none; z-index: 201; border-radius: 15px; -webkit-border-radius: 15px; -moz-border-radius: 15px; -khtml-border-radius: 15px; box-shadow: 0px 10px 20px #808080; -webkit-box-shadow: 0px 10px 20px #808080; -moz-box-shadow: 0px 10px 20px #808080; -khtml-box-shadow: 0px 10px 20px #808080; filter: progid:DXImageTransform.Microsoft.dropshadow(OffX=2, OffY=2, Color='gray', Positive='true')}
.MathJax_Menu {position: absolute; background-color: white; color: black; width: auto; padding: 2px; border: 1px solid #CCCCCC; margin: 0; cursor: default; font: menu; text-align: left; text-indent: 0; text-transform: none; line-height: normal; letter-spacing: normal; word-spacing: normal; word-wrap: normal; white-space: nowrap; float: none; z-index: 201; box-shadow: 0px 10px 20px #808080; -webkit-box-shadow: 0px 10px 20px #808080; -moz-box-shadow: 0px 10px 20px #808080; -khtml-box-shadow: 0px 10px 20px #808080; filter: progid:DXImageTransform.Microsoft.dropshadow(OffX=2, OffY=2, Color='gray', Positive='true')}
.MathJax_MenuItem {padding: 2px 2em; background: transparent}
.MathJax_MenuArrow {position: absolute; right: .5em; color: #666666}
.MathJax_MenuActive .MathJax_MenuArrow {color: white}
.MathJax_MenuArrow.RTL {left: .5em; right: auto}
.MathJax_MenuCheck {position: absolute; left: .7em}
.MathJax_MenuCheck.RTL {right: .7em; left: auto}
.MathJax_MenuRadioCheck {position: absolute; left: 1em}
.MathJax_MenuRadioCheck.RTL {right: 1em; left: auto}
.MathJax_MenuLabel {padding: 2px 2em 4px 1.33em; font-style: italic}
.MathJax_MenuRule {border-top: 1px solid #CCCCCC; margin: 4px 1px 0px}
.MathJax_MenuDisabled {color: GrayText}
.MathJax_MenuActive {background-color: Highlight; color: HighlightText}
.MathJax_Menu_Close {position: absolute; width: 31px; height: 31px; top: -15px; left: -15px}
</style><style type="text/css">#MathJax_Zoom {position: absolute; background-color: #F0F0F0; overflow: auto; display: block; z-index: 301; padding: .5em; border: 1px solid black; margin: 0; font-weight: normal; font-style: normal; text-align: left; text-indent: 0; text-transform: none; line-height: normal; letter-spacing: normal; word-spacing: normal; word-wrap: normal; white-space: nowrap; float: none; box-shadow: 5px 5px 15px #AAAAAA; -webkit-box-shadow: 5px 5px 15px #AAAAAA; -moz-box-shadow: 5px 5px 15px #AAAAAA; -khtml-box-shadow: 5px 5px 15px #AAAAAA; filter: progid:DXImageTransform.Microsoft.dropshadow(OffX=2, OffY=2, Color='gray', Positive='true')}
#MathJax_ZoomOverlay {position: absolute; left: 0; top: 0; z-index: 300; display: inline-block; width: 100%; height: 100%; border: 0; padding: 0; margin: 0; background-color: white; opacity: 0; filter: alpha(opacity=0)}
#MathJax_ZoomFrame {position: relative; display: inline-block; height: 0; width: 0}
#MathJax_ZoomEventTrap {position: absolute; left: 0; top: 0; z-index: 302; display: inline-block; border: 0; padding: 0; margin: 0; background-color: white; opacity: 0; filter: alpha(opacity=0)}
</style><style type="text/css">.MathJax_Preview {color: #888}
#MathJax_Message {position: fixed; left: 1em; bottom: 1.5em; background-color: #E6E6E6; border: 1px solid #959595; margin: 0px; padding: 2px 8px; z-index: 102; color: black; font-size: 80%; width: auto; white-space: nowrap}
#MathJax_MSIE_Frame {position: absolute; top: 0; left: 0; width: 0px; z-index: 101; border: 0px; margin: 0px; padding: 0px}
.MathJax_Error {color: #CC0000; font-style: italic}
</style></head>

<body style="" class=""><div id="MathJax_Message" style="display: none;"></div><div class="csdn-toolbar csdn-toolbar-skin-black ">        <div class="container row center-block ">          <div class="col-md-3 pull-left logo clearfix"><a href="http://www.csdn.net/?ref=toolbar" title="CSDN首页" target="_blank" class="icon"></a><a title="频道首页" href="http://write.blog.csdn.net/?ref=toolbar_logo" class="img blog-icon"></a></div>          <div class="pull-right login-wrap ">            <ul class="btns">              <li class="loginlink"><a href="https://passport.csdn.net/account/login?ref=toolbar" target="_top">登录&nbsp;</a>|<a target="_top" href="http://passport.csdn.net/account/mobileregister?ref=toolbar&amp;action=mobileRegister">&nbsp;注册</a></li>              <li class="search">                <div class="icon on-search-icon">                  <div class="wrap">                    <div class="curr-icon-wrap">                      <div class="curr-icon"></div>                    </div>                    <form action="http://so.csdn.net/search" id="toolbar_search" method="get" target="_blank">                      <input type="hidden" value="toolbar" name="ref" accesskey="2">                      <div class="border">                        <input placeholder="搜索" type="text" value="" name="q" accesskey="2"><span class="icon-enter-sm"></span>                      </div>                    </form>                  </div>                </div>              </li>              <li class="favor">                <div class="icon on-favor-icon">                  <div class="wrap">                    <div class="curr-icon-wrap">                      <div class="curr-icon"></div>                    </div>                    <div style="display:none;" class="favor-success"><span class="msg">收藏成功</span>                      <div class="btns"><span class="btn btn-primary ok">确定</span></div>                    </div>                    <div style="display:none;" class="favor-failed"><span class="icon-danger-lg"></span><span class="msg">收藏失败，请重新收藏</span>                      <div class="btns"><span class="btn btn-primary ok">确定</span></div>                    </div>                    <form role="form" class="form-horizontal favor-form">                      <div class="form-group">                        <div class="clearfix">                          <label for="input-title" class="col-sm-2 control-label"><span class="red_txt">*</span>标题</label>                          <div class="col-sm-10">                            <input id="inputTitle" type="text" placeholder="" class="title form-control">                          </div>                        </div>                        <div class="alert alert-danger"><strong></strong>标题不能为空</div>                      </div>                      <div class="form-group" style="display:none;">                        <label for="input-url" class="col-sm-2 control-label">网址</label>                        <div class="col-sm-10">                          <input id="input-url" type="text" placeholder="" class="url form-control">                        </div>                      </div>                      <div class="form-group">                        <label for="input-tag" class="col-sm-2 tag control-label">标签</label>                        <div class="col-sm-10">                          <input id="input-tag" type="text" class="form-control tag">                        </div>                      </div>                      <div class="form-group">                        <label for="input-description" class="description col-sm-2 control-label">位置</label>                        <div class="col-sm-10">                          <div class="my_lib_box">                            个人主页&nbsp;-&nbsp;<a href="http://my.csdn.net/" target="_blank">我的知识</a>                          </div>                          <div class="checkbox">                            <div class="pull-left">                              <label>                                <input type="checkbox" name="share" class="save_lib_map">同时保存至：                              </label>                            </div>                            <div class="pull-left">                              <div class="dropdown">                                <button id="toolbar_sele_map" type="button">                                  选择知识图谱                                  <i class="fa fa-chevron-down"></i>                                </button>                                <div class="top_arr"></div>                                <div class="outside">                                  <ul class="dropdown-menu" id="toolbar_Design_knowledge"></ul>                                </div>                              </div>                            </div>                            <div class="pull-left new_txt">                              <a href="http://lib.csdn.net/my/create/structure" target="_blank">新建？</a>                            </div>                          </div>                        </div>                      </div>                      <div class="form-group">                        <div class="col-sm-offset-2 col-sm-10 ft">                          <div class="col-sm-4 pull-left" style="display:none">                            <div class="checkbox">                              <label>                                <input type="checkbox" name="share" checked="checked" class="share">公开                              </label>                            </div>                          </div>                          <div class="col-sm-8 pull-right favor-btns">                            <button type="button" class="cancel btn btn-default">取消</button>                            <button type="submit" class="submit btn btn-primary">收藏</button>                          </div>                        </div>                      </div>                    </form>                  </div>                </div>              </li>              <li class="notify">                <div style="display:none" class="number"></div>                <div style="display:none" class="icon-hasnotes-sm"></div>                <div id="header_notice_num"></div>                <div class="icon on-notify-icon">                  <div class="wrap">                    <div class="curr-icon-wrap">                      <div class="curr-icon"></div>                    </div>                    <div id="note1" class="csdn_note">                      <div class="box"></div>                    </div>                  </div>                </div>              </li>              <li class="ugc">                <div class="icon on-ugc-icon">                  <div class="wrap clearfix">                    <div class="curr-icon-wrap">                      <div class="curr-icon"></div>                    </div>                    <dl>                      <dt><a href="http://geek.csdn.net/news/expert?ref=toolbar" target="_blank" class="p-news clearfix" style="display:none;"><em class="icon"></em><span>分享资讯</span></a></dt>                      <dt style="border: none;"><a href="http://u.download.csdn.net/upload?ref=toolbar" target="_blank" class="p-doc clearfix"><em class="icon"></em><span>传PPT/文档</span></a></dt>                      <dt><a href="http://bbs.csdn.net/topics/new?ref=toolbar" target="_blank" class="p-ask clearfix"><em class="icon"></em><span>提问题</span></a></dt>                      <dt><a href="http://write.blog.csdn.net/postedit?ref=toolbar" target="_blank" class="p-blog clearfix"><em class="icon"></em><span>写博客</span></a></dt>                      <dt><a href="http://gitbook.cn/new/gitchat/activity?utm_source=csdnblog1" target="_blank" class="gitchat clearfix"><em class="icon"></em><span>发布Chat</span></a></dt>                      <dt><a href="http://u.download.csdn.net/upload?ref=toolbar" target="_blank" class="p-src clearfix"><em class="icon"></em><span>传资源</span></a></dt>                      <dt><a href="https://code.csdn.net/projects/new?ref=toolbar" target="_blank" class="c-obj clearfix"><em class="icon"></em><span>创建项目</span></a></dt>                      <dt><a href="https://code.csdn.net/snippets/new?ref=toolbar" target="_blank" class="c-code clearfix"><em class="icon"></em><span>创建代码片</span></a></dt>                    </dl>                  </div>                </div>              </li>              <li class="profile">                <div class="icon on-profile-icon"><img src="./README_files/2_leprechaun_.jpg" class="curr-icon-img">                  <div class="wrap clearfix">                    <div class="curr-icon-wrap">                      <div class="curr-icon"></div>                    </div>                    <div class="bd">                      <dl class="clearfix">                        <dt class="pull-left img"><a target="_blank" href="http://my.csdn.net/?ref=toolbar" class="avatar"><img src="./README_files/2_leprechaun_.jpg"></a></dt>                        <dd class="info" style="border: none;"><a target="_blank" href="http://my.csdn.net/?ref=toolbar" class="nickname">leprechaun_</a><span class="dec"><a class="fill-dec" href="http://my.csdn.net/" target="_blank">编辑自我介绍，让更多人了解你<span class="write-icon"></span></a></span></dd>                      </dl>                    </div>                    <div class="ft clearfix"><a target="_blank" href="http://my.csdn.net/my/account/changepwd?ref=toolbar" class="pull-left"><span class="icon-cog"></span>帐号设置</a><a href="https://passport.csdn.net/account/logout?ref=toolbar" target="_top" class="pull-left" style="margin-left:132px; width:18px; height:27px; white-space:nowrap; overflow:hidden;"><span class="icon-signout"></span><span class="out">退出</span></a></div>                  </div>                </div>              </li>              <li class="apps">                <div id="chasnew123" class="hasnew" style="display: none;"></div>                <div id="cappsarea123" class="icon on-apps-icon">                  <div class="wrap clearfix">                    <div class="curr-icon-wrap">                      <div class="curr-icon"></div>                    </div>                  <div class="detail">                    <dl>                      <dt>                        <h5>社区</h5>                      </dt>                      <dd> <a href="http://blog.csdn.net/?ref=toolbar" target="_blank">博客</a></dd>                      <dd> <a href="http://bbs.csdn.net/?ref=toolbar" target="_blank">论坛</a></dd>                      <dd> <a href="http://download.csdn.net/?ref=toolbar" target="_blank">下载</a></dd>                      <dd> <a href="http://lib.csdn.net/?ref=toolbar" target="_blank">知识库</a></dd>                      <dd><a href="http://ask.csdn.net/?ref=toolbar" target="_blank">技术问答</a></dd>                      <dd><a href="http://geek.csdn.net/?ref=toolbar" target="_blank">极客头条</a></dd>                      <dd style="display:none"> <a href="http://hero.csdn.net/?ref=toolbar" target="_blank">英雄会</a></dd>                    </dl>                  </div>                  <div class="detail">                    <dl>                      <dt>                        <h5>服务</h5>                      </dt>                      <dd style="display:none"> <a href="http://job.csdn.net/?ref=toolbar" target="_blank">JOB<img src="./README_files/new.gif" style="display: none; margin-top: -26px; width: 23px;"></a></dd>                      <dd> <a href="http://edu.csdn.net/?ref=toolbar" target="_blank">学院<img src="./README_files/new.gif" style="display: none; margin-top: -26px; width: 23px;"></a></dd>                      <dd> <a href="https://code.csdn.net/?ref=toolbar" target="_blank">CODE</a></dd>                      <dd> <a href="http://huiyi.csdn.net/?ref=toolbar" target="_blank">活动</a></dd>                      <dd> <a href="http://www.csto.com/?ref=toolbar" target="_blank">CSTO</a></dd>                      <dd> <a href="http://mall.csdn.net/?ref=toolbar" target="_blank">C币兑换<img src="./README_files/new.gif" style="display: none; margin-top: -26px; width: 23px;"></a></dd>                    </dl>                  </div>                  <div class="detail last">                    <dl>                      <dt>                        <h5>俱乐部</h5>                      </dt>                      <dd> <a href="http://cto.csdn.net/?ref=toolbar" target="_blank">CTO俱乐部</a></dd>                      <dd> <a href="http://student.csdn.net/?ref=toolbar" target="_blank">高校俱乐部</a></dd>                    </dl>                  </div>                </div>              </div>            </li>            </ul>          </div>        </div>    </div>    
    <script type="text/javascript">      
        var ua = window.navigator.userAgent;
        if (ua.indexOf("MSIE") >= 1) {
            {
                location.replace("/mdeditor/tip");
            }
        }
    </script>

        <script src="./README_files/require.js.下载"></script>

    <script fixed="true" id="toolbar-tpl-scriptId" prod="blog" skin="black" src="./README_files/html.js.下载" type="text/javascript"></script>    

    <div class="working-indicator"><div class="hide"><div class="bar" style="animation: working-indicator-bar0 0.7s linear infinite;"></div><div class="bar" style="animation: working-indicator-bar1 0.7s linear infinite;"></div><div class="bar" style="animation: working-indicator-bar2 0.7s linear infinite;"></div></div></div>
    <div id="csdnEditor" class=" ltr" style="position: absolute; top: 0px; left: 0px; bottom: 0px; right: 0px; overflow: hidden;"><div class="layout-wrapper-l1" style="transform: translate(0px, 0px); width: 1600px; height: 769px;">
	<div class="layout-wrapper-l2" style="left: 0px; width: 1600px; height: 769px;">
		<div class="navbar navbar-default">
			<div class="navbar-inner">
				<div class="buttons-dropdown dropdown hide">
					<div class="nav">
						<button class="btn btn-success" data-toggle="dropdown" title="" data-original-title="Show buttons">
							<i class="icon-th-large"></i>
						</button>
						<div class="dropdown-menu">
						</div>
					</div>
				</div>
				
				
				<a class="btn btn-danger btn-lg btn-blog-publish" data-original-title="" title="">发表博客</a>
				<!--<a class="btn btn-danger btn-lg btn-blog-publish" data-original-title="" title="">发表博客

           <span class="send_bt_arrow" style="	border-top: 6px solid #fff;
 border-left: 6px solid transparent;
 border-right: 6px solid transparent;
 border-bottom: 0;
 vertical-align: -17px;
 margin-left: 0px;"></span>
					<ul class="send_bt_list" style="width: 120px;height: 78px;position: absolute;left: 0; bottom: -80px;padding: 0;margin: 0;background: #fff;color: #333333;
 font-size: 14px;
 font-weight: normal;z-index: 9999;list-style: none;display: none;">
						<li class="send_bt_list_first" style="cursor: pointer;height: 35px;line-height: 50px;">普通发表</li>
						<li class="send_bt_list_second" style="cursor: pointer;;height: 35px;line-height: 35px;">
							发表并@博乐
						</li>
					</ul>


				</a>-->
				<script>
					//				alert(1);
					//					$(".btn-blog-publish").hover(
					//
					//						function () {
					//							$(this).find("send_bt_list")[0].style.display="block";
					//						},
					//						function () {
					//							$(this).find("send_bt_list")[0].style.display="none";
					//						}
					//					);

				</script>
				
				
				<ul class="nav pull-right right-buttons">
					<li class="offline-status hide">
						<div class="text-danger">
							<i class="icon-attention-circled"></i>离线
						</div>
					</li>
					<li class="extension-buttons" style="display:none"><div class="btn-group"><a class="btn btn-success button-publish" title="" data-original-title="Update document publication"><i class="icon-upload"></i></a></div></li>
					<li>
						<a class="btn btn-success" href="javascript:void(0);" style="display:none" id="csdn-lock-editor-button" title="" data-original-title="锁定MarkDown编辑器">
							<i class="icon-lock-open"></i>
						</a>
					</li>
					<li>
						<div class="dropdown">
                            <span class="btn btn-success" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false" data-original-title="" title="">
                                博客管理
                                <i class="icon-arrow-down"></i>
                            </span>
							<ul class="dropdown-menu pull-right" role="menu">
								<li>
									<a target="_blank" href="http://write.blog.csdn.net/postedit?type=edit">切换到HTML编辑器</a>
								</li>
								<li>
									<a target="_blank" href="http://write.blog.csdn.net/postlist">文章管理</a>
								</li>
								<li>
									<a target="_blank" href="http://write.blog.csdn.net/category">类别管理</a>
								</li>
								<li>
									<a target="_blank" href="http://write.blog.csdn.net/feedback">评论管理</a>
								</li>
								<li>
									<a target="_blank" href="http://write.blog.csdn.net/configure">博客配置</a>
								</li>
								<li>
									<a target="_blank" href="http://write.blog.csdn.net/configure/column">博客栏目</a>
								</li>
								<li>
									<a target="_blank" href="http://write.blog.csdn.net/postlist/0/all/draft">草稿箱</a>
								</li>
								<li>
									<a target="_blank" href="http://write.blog.csdn.net/postlist/0/all/deleted">回收站</a>
								</li>                               
							</ul>
						</div>
					</li>
				</ul>
				<ul class="nav left-buttons">
					<li class="wmd-button-group1 btn-group"><li class="wmd-button btn btn-success" id="wmd-bold-button" title="" style="left: 0px;" data-original-title="加粗 &lt;strong&gt; Ctrl/Cmd+B"><span style="background-position: 0px 0px; display: none;"></span><i class="icon-bold"></i></li><li class="wmd-button btn btn-success" id="wmd-italic-button" title="" style="left: 0px;" data-original-title="斜体 &lt;em&gt; Ctrl/Cmd+I"><span style="background-position: -20px 0px; display: none;"></span><i class="icon-italic"></i></li><li class="wmd-button btn btn-success" id="wmd-link-button" title="" style="left: 0px;" data-original-title="Hyperlink &lt;a&gt; Ctrl/Cmd+L"><span style="background-position: -40px 0px; display: none;"></span><i class="icon-link"></i></li><li class="wmd-button btn btn-success" id="wmd-quote-button" title="" style="left: 0px;" data-original-title="引用 &lt;blockquote&gt; Ctrl/Cmd+Q"><span style="background-position: -60px 0px; display: none;"></span><i class="icon-quote"></i></li><li class="wmd-button btn btn-success" id="wmd-code-button" title="" style="left: 0px;" data-original-title="代码片 &lt;pre&gt;&lt;code&gt; Ctrl/Cmd+K"><span style="background-position: -80px 0px; display: none;"></span><i class="icon-code"></i></li><li class="wmd-button btn btn-success" id="wmd-image-button" title="" style="left: 0px;" data-original-title="图片 &lt;img&gt; Ctrl/Cmd+G"><span style="background-position: -100px 0px; display: none;"></span><i class="icon-picture"></i></li><li class="wmd-button btn btn-success" id="wmd-olist-button" title="" style="left: 0px;" data-original-title="有序列表 &lt;ol&gt; Ctrl/Cmd+O"><span style="background-position: -120px 0px; display: none;"></span><i class="icon-list-ordered"></i></li><li class="wmd-button btn btn-success" id="wmd-ulist-button" title="" style="left: 0px;" data-original-title="无序列表 &lt;ul&gt; Ctrl/Cmd+U"><span style="background-position: -140px 0px; display: none;"></span><i class="icon-list-unordered"></i></li><li class="wmd-button btn btn-success" id="wmd-heading-button" title="" style="left: 0px;" data-original-title="标题 &lt;h1&gt;/&lt;h2&gt; Ctrl/Cmd+H"><span style="background-position: -160px 0px; display: none;"></span><i class="icon-heading"></i></li><li class="wmd-button btn btn-success" id="wmd-hr-button" title="" style="left: 0px;" data-original-title="横线 &lt;hr&gt; Ctrl/Cmd+R"><span style="background-position: -180px 0px; display: none;"></span><i class="icon-line"></i></li><li class="wmd-button btn btn-success" id="wmd-undo-button" title="" style="left: 0px;" data-original-title="撤销 - Ctrl/Cmd+Z"><span style="background-position: -200px 0px; display: none;"></span><i class="icon-undo"></i></li><li class="wmd-button btn btn-success disabled" id="wmd-redo-button" title="" style="left: 0px;" data-original-title="重做 - Ctrl/Cmd+Y"><span style="background-position: -220px -20px; display: none;"></span><i class="icon-redo"></i></li></li>
				</ul><ul class="nav left-buttons">
					<li class="wmd-button-group2 btn-group">
						<a class="btn btn-success btn-blog-new" title="" data-original-title="写新文章"><i class="icon-pencil"></i></a>
						<a class="btn btn-success btn-blog-save changed" title="" data-original-title="保存到线上草稿箱"><i class="icon-disk"></i></a>
						<a class="btn btn-success btn-blog-abstract" title="" data-original-title="生成摘要"><i class="icon-abstract"></i></a>
						<a class="btn btn-success btn-blog-setting" title="" data-original-title="文章设置"><i class="icon-doc-setting"></i></a>
					</li>
				</ul><ul class="nav left-buttons">
					<li class="wmd-button-group3 btn-group">
						<a data-toggle="modal" data-target=".modal-import-url" class="btn btn-success btn-import-online" title="" data-original-title="从线上导入">
							<i class="icon-earth"></i>
						</a>
						<a data-toggle="modal" data-target=".modal-import-harddrive-markdown" class="btn btn-success btn-import" title="" data-original-title="从本机导入">
							<i class="icon-hdd"></i>
						</a>
						<a data-toggle="modal" data-target=".modal-export-doc" class="btn btn-success btn-export" title="" data-original-title="导出到本地">
							<i class="icon-upload"></i>
						</a>
					</li>
				</ul><ul class="nav left-buttons">
					<li class="wmd-button-group5 btn-group">
						<a class="btn btn-success btn-modal-helps" title="" data-original-title="Markdown语法帮助">
							<i class="icon-help"></i>
						</a>
					</li>
				</ul><ul class="nav pull-right title-container">
					<li><a class="btn btn-success file-title-navbar" title="" data-placement="bottom" style="width: 1360px;" data-original-title="点击修改标题">欢迎使用CSDN-markdown编辑器</a></li>
					<li>
						<div class="input-file-title-container"><input type="text" maxlength="100" class="col-sm-4 form-control input-file-title hide" placeholder="文章标题" style="width: 1360px;"></div>
					</li>
				</ul>
			</div>
		</div>
		<div class="layout-wrapper-l3" style="top: 170px; width: 1600px; height: 599px;">
			<pre id="wmd-input" class="form-control font-rich" style="width: 784px; height: 599px;"><div class="editor-content" contenteditable="true" style="padding-left: 35px; padding-right: 35px; padding-bottom: 539px;"><span id="wmd-input-section-1" class="wmd-input-section"></span><span id="wmd-input-section-37" class="wmd-input-section"><span class="token h1"><span class="token md md-hash">#</span> Cloudgo-io</span><span class="token lf">
</span><span class="token lf">
</span><span class="token p">这次的作业按照老师的教程[cloudio] (<span class="token url">http://blog.csdn.net/pmlpml/article/details/78539261)</span>  进行实现，需要添加改动的地方不多。</span><span class="token lf">
</span><span class="token lf">
</span><span class="token lf">
</span></span><span id="wmd-input-section-3" class="wmd-input-section"><span class="token h2"><span class="token md md-hash">##</span> 基本要求</span><span class="token lf">
</span><span class="token lf">
</span></span><span id="wmd-input-section-4" class="wmd-input-section"><span class="token h3"><span class="token md md-hash">###</span> 1、支持静态文件服务</span><span class="token lf">
</span><span class="token lf">
</span><span class="token p"><span class="token img"><span class="token md md-bang">!</span><span class="token md md-bracket-start">[</span><span class="token md md-alt">这里写图片描述</span><span class="token md md-bracket-end">]</span><span class="token md img-parens"><span class="token md md-paren-start">(</span><span class="token md md-src">http://img.blog.csdn.net/20171121202431629?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvbGVwcmVjaGF1bl8=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast</span><span class="token md md-paren-end">)</span></span></span></span><span class="token lf">
</span><span class="token lf">
</span></span><span id="wmd-input-section-110" class="wmd-input-section"><span class="token h3"><span class="token md md-hash">###</span> 2、支持简单的js访问</span><span class="token lf">
</span><span class="token lf">
</span><span class="token p"><span class="token img"><span class="token md md-bang">!</span><span class="token md md-bracket-start">[</span><span class="token md md-alt">这里写图片描述</span><span class="token md md-bracket-end">]</span><span class="token md img-parens"><span class="token md md-paren-start">(</span><span class="token md md-src">http://img.blog.csdn.net/20171121204509880?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvbGVwcmVjaGF1bl8=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast</span><span class="token md md-paren-end">)</span></span></span></span><span class="token lf">
</span><span class="token lf">
</span></span><span id="wmd-input-section-153" class="wmd-input-section"><span class="token h3"><span class="token md md-hash">###</span> 3、提交表单，并输出一个表格</span><span class="token lf">
</span><span class="token lf">
</span><span class="token p">注册（提交）表单界面</span><span class="token lf">
</span><span class="token lf">
</span><span class="token p"><span class="token img"><span class="token md md-bang">!</span><span class="token md md-bracket-start">[</span><span class="token md md-alt">这里写图片描述</span><span class="token md md-bracket-end">]</span><span class="token md img-parens"><span class="token md md-paren-start">(</span><span class="token md md-src">http://img.blog.csdn.net/20171121203031129?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvbGVwcmVjaGF1bl8=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast</span><span class="token md md-paren-end">)</span></span></span></span><span class="token lf">
</span><span class="token lf">
</span><span class="token p">输出表格</span><span class="token lf">
</span><span class="token lf">
</span><span class="token p"><span class="token img"><span class="token md md-bang">!</span><span class="token md md-bracket-start">[</span><span class="token md md-alt">这里写图片描述</span><span class="token md md-bracket-end">]</span><span class="token md img-parens"><span class="token md md-paren-start">(</span><span class="token md md-src">http://img.blog.csdn.net/20171121204550880?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvbGVwcmVjaGF1bl8=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast</span><span class="token md md-paren-end">)</span></span></span></span><span class="token lf">
</span><span class="token lf">
</span></span><span id="wmd-input-section-147" class="wmd-input-section"><span class="token h3"><span class="token md md-hash">###</span> 4、对 /unknown 给出开发中的提示，返回码 5xx</span><span class="token lf">
</span><span class="token lf">
</span><span class="token p"><span class="token img"><span class="token md md-bang">!</span><span class="token md md-bracket-start">[</span><span class="token md md-alt">这里写图片描</span><span class="token md md-bracket-end">]</span><span class="token md img-parens"><span class="token md md-paren-start">(</span><span class="token md md-src">http://img.blog.csdn.net/20171121203202240?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvbGVwcmVjaGF1bl8=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast</span><span class="token md md-paren-end">)</span></span></span></span></span><span class="token lf">
</span></div><div class="editor-margin" style="width: 35px;"></div></pre>
			<div class="preview-panel" style="transform: translate(784px, 0px); width: 816px; height: 599px;">
				<div class="layout-resizer layout-resizer-preview open" style="user-select: none; -webkit-user-drag: none; -webkit-tap-highlight-color: rgba(0, 0, 0, 0); touch-action: pan-y; height: 599px;"></div>
				<div class="layout-toggler layout-toggler-navbar btn btn-info open" title="头部收起/展开" style="height: 60px;"><i class="icon-th"></i>
				</div>
				<div class="layout-toggler layout-toggler-preview btn btn-info open" title="预览区显示/隐藏（可拖拽）" style="transform: translate(0px, 269.5px); height: 60px;"><i class="icon-none"></i></div>
				<div class="preview-container" style="left: 32px; width: 784px; height: 599px;">
					<div id="preview-contents" style="padding-left: 35px; padding-right: 35px; padding-bottom: 539px;">
						<div id="wmd-preview" class="preview-content"></div>
					<div id="wmd-preview-section-1" class="wmd-preview-section preview-content">

</div><div id="wmd-preview-section-37" class="wmd-preview-section preview-content">

<h1 id="cloudgo-io">Cloudgo-io</h1>

<p>这次的作业按照老师的教程[cloudio] (<a href="http://blog.csdn.net/pmlpml/article/details/78539261" target="_blank">http://blog.csdn.net/pmlpml/article/details/78539261</a>)  进行实现，需要添加改动的地方不多。</p></div><div id="wmd-preview-section-3" class="wmd-preview-section preview-content">

<h2 id="基本要求">基本要求</h2>

</div><div id="wmd-preview-section-4" class="wmd-preview-section preview-content">

<h3 id="1支持静态文件服务">1、支持静态文件服务</h3>

<p><img src="./README_files/20171121202431629" alt="这里写图片描述" title=""></p>

</div><div id="wmd-preview-section-110" class="wmd-preview-section preview-content">

<h3 id="2支持简单的js访问">2、支持简单的js访问</h3>

<p><img src="./README_files/20171121204509880" alt="这里写图片描述" title=""></p></div><div id="wmd-preview-section-153" class="wmd-preview-section preview-content">

<h3 id="3提交表单并输出一个表格">3、提交表单，并输出一个表格</h3>

<p>注册（提交）表单界面</p>

<p><img src="./README_files/20171121203031129" alt="这里写图片描述" title=""></p>

<p>输出表格</p>

<p><img src="./README_files/20171121204550880" alt="这里写图片描述" title=""></p></div><div id="wmd-preview-section-14" class="wmd-preview-section preview-content">

<h3 id="4对-unknown-给出开发中的提示返回码-5xx">4、对 /unknown 给出开发中的提示，返回码 5xx</h3>

<p><img src="./README_files/20171121203202240" alt="这里写图片描" title=""></p></div><div id="wmd-preview-section-footnotes" class="preview-content"></div></div>
				</div>
			</div>
		</div>
		<div class="extension-preview-buttons closed animate" style="user-select: none; -webkit-user-drag: none; -webkit-tap-highlight-color: rgba(0, 0, 0, 0); touch-action: pan-y; transform: translate(92px, -6px);">
			<div class="btn-group drag-me dropup" title="这里可以拖拽">
				<i class="icon-ellipsis-vert"></i>
			</div>
		<div class="btn-group dropup"></div><div class="btn-group dropup"><button class="btn btn-success dropdown-toggle" title="标题列表" data-toggle="dropdown">
    <i class="icon-list"></i>
</button>
<div class="dropdown-menu pull-right">
    <h4>标题列表</h4>
    <hr>
    <div class="table-of-contents"><div class="toc">
<ul>
<li><a href="http://write.blog.csdn.net/mdeditor#cloudgo-io">Cloudgo-io</a><ul>
<li><a href="http://write.blog.csdn.net/mdeditor#基本要求">基本要求</a><ul>
<li><a href="http://write.blog.csdn.net/mdeditor#1支持静态文件服务">1支持静态文件服务</a></li>
<li><a href="http://write.blog.csdn.net/mdeditor#2支持简单的js访问">2支持简单的js访问</a></li>
<li><a href="http://write.blog.csdn.net/mdeditor#3提交表单并输出一个表格">3提交表单并输出一个表格</a></li>
<li><a href="http://write.blog.csdn.net/mdeditor#4对-unknown-给出开发中的提示返回码-5xx">4对 unknown 给出开发中的提示返回码 5xx</a></li>
</ul>
</li>
</ul>
</li>
</ul>
</div>
</div>
    <hr>
</div>
</div><div class="btn-group dropup"><button class="btn btn-success dropdown-toggle stat-button" title="统计信息" data-toggle="dropdown">
	<i class="icon-chart-bar"></i>
	<span class="value">181</span>
</button>
<div class="dropdown-menu pull-right stat-button-dropdown">
	<h4>统计信息</h4>
    <hr>
	<div class="stat">
		<div>
			字数: <span class="value1">181</span>
		</div>
		<!--<div>-->
			<!--词（单词）: <span class="value2"></span>-->
		<!--</div>-->
		<div>
			段落: <span class="value3">9</span>
		</div>
	</div>
</div>
</div></div>
	<div class="find-replace" style="display: none;"><button type="button" class="close button-find-replace-dismiss">×</button>
<div class="form-inline">
    <div class="form-group">
        <label for="input-find-replace-search-for">搜索</label>
        <input class="form-control" id="input-find-replace-search-for" placeholder="关键词">
    </div>
    <div class="form-group">
        <label for="input-find-replace-replace-with">替换为</label>
        <input class="form-control" id="input-find-replace-replace-with" placeholder="替换文本">
    </div>
</div>
<div class="pull-right">
    <div class="help-block text-right">
        <span class="found-counter">0</span> 处
    </div>
    <div>
        <button type="button" class="btn btn-sm btn-primary search-button">查找</button>
        <button type="button" class="btn btn-sm btn-warning replace-button">替换</button>
        <button type="button" class="btn btn-sm btn-danger replace-all-button">全部替换</button>
    </div>
</div>
<div class="pull-left">
    <div class="checkbox">
        <label>
            <input type="checkbox" class="checkbox-case-sensitive"> 区分大小写
        </label>
    </div>
    <div class="checkbox">
        <label>
            <input type="checkbox" class="checkbox-regexp"> 正则表达式
        </label>
    </div>
</div>
</div></div>
	<div id="wmd-button-bar" class="hide"><ul id="wmd-button-row" class="wmd-button-row"><li class="wmd-spacer wmd-spacer1 btn btn-success" id="wmd-spacer1" style="left: 0px;"></li><li class="wmd-spacer wmd-spacer2 btn btn-success" id="wmd-spacer2" style="left: 0px;"></li><li class="wmd-spacer wmd-spacer3 btn btn-success" id="wmd-spacer3" style="left: 0px;"></li></ul></div>


	<div class="document-panel">

		<div class="panel-content">
			<div class="list-group document-list"><ul class="nav"><li><a href="http://write.blog.csdn.net/mdeditor#" class="list-group-item file active" data-file-index="file.xBphcdiVf9b7Jx4gzGXVDrnu">   <i class="icon-provider-csdnblog"></i>欢迎使用CSDN-markdown编辑器</a></li></ul></div>
			<div class="list-group document-list-filtered hide"><ul class="nav"><li><a href="http://write.blog.csdn.net/mdeditor#" class="list-group-item file active" data-file-index="file.xBphcdiVf9b7Jx4gzGXVDrnu">   <i class="icon-provider-csdnblog"></i>欢迎使用CSDN-markdown编辑器</a></li></ul></div>
		</div>
	</div>
</div>

<div class="modal fade modal-document-manager">
	<div class="modal-dialog">
		<div class="modal-content">
			<div class="modal-header">
				<h4 class="modal-title">文档管理</h4>
			</div>
			<div class="modal-body">
				<div></div>
				<ul class="nav nav-pills document-list">
					<li class="pull-right dropdown"><a data-toggle="dropdown"><i class="icon-check"></i> 选择操作 <b class="caret"></b></a>
						<ul class="dropdown-menu">
							<li><a class="action-select-all"><i class="icon-check"></i> 选择全部</a></li>
							<li><a class="action-unselect-all"><i class="icon-check-empty"></i> 取消选择</a></li>
							<li class="divider"></li>
							<li><a class="action-move-items"><i class="icon-forward"></i> 移到指定文件夹</a></li>
							<li><a class="action-delete-items"><i class="icon-trash"></i> 删除</a></li>
						</ul>
					</li>
					<li class="pull-right"><a class="action-create-folder"> <i class="icon-folder"></i> 新建文件夹</a></li>
					<li class="disabled"><a><i class="icon-file"></i> <span class="document-count"></span></a></li>
					<li class="disabled"><a><i class="icon-folder"></i> <span class="folder-count"></span></a></li>
				</ul>
				<div class="list-group document-list"></div>
				<p class="confirm-delete hide">下列文档将从你本地存储中删除：</p>

				<p class="choose-folder hide">请选择一个目标文件夹：</p>

				<div class="list-group selected-document-list hide"></div>
				<div class="list-group select-folder-list hide"></div>
			</div>
			<div class="modal-footer">
				<a class="btn btn-default confirm-delete choose-folder action-cancel hide">取消</a>
				<a class="btn btn-danger confirm-delete action-delete-items-confirm hide">确定</a>
				<a class="btn btn-danger document-list" data-dismiss="modal">关闭</a>
			</div>
		</div>
	</div>
</div>


<div class="modal fade modal-insert-link">
	<div class="modal-dialog">
		<div class="modal-content">
			<div class="modal-header">
				<button type="button" class="close" data-dismiss="modal" aria-hidden="true">×</button>
				<h4 class="modal-title">添加链接</h4>
			</div>
			<div class="modal-body">
				<p>链接书写格式为：链接地址 + 空格 + "链接提示"。</p>

				<div class="input-group">
					<span class="input-group-addon">
                        <i class="icon-link"></i>
                    </span>
					<input id="input-insert-link" type="text" class="col-sm-5 form-control" placeholder="http://example.com/ &quot;optional title&quot;">
				</div>
			</div>
			<div class="modal-footer">
				<a class="btn btn-default" data-dismiss="modal">取消</a>
				<a class="btn btn-danger action-insert-link" data-dismiss="modal">确定</a>
			</div>
		</div>
	</div>
</div>


<div class="modal fade modal-insert-image" aria-hidden="true" style="display: none;">
	<div class="modal-dialog">
		<div class="modal-content">
			<div class="modal-header">
				<button type="button" class="close" data-dismiss="modal" aria-hidden="true">×</button>
				<h4 class="modal-title">插入图片</h4>
			</div>
			<div class="modal-body">
				<ul class="nav nav-tabs" role="tablist">
					<li role="presentation" class=""><a href="http://write.blog.csdn.net/mdeditor#img_online" role="tab" data-toggle="tab">在线图片</a>
					</li>
					<li role="presentation" class="active"><a href="http://write.blog.csdn.net/mdeditor#img_upload" role="tab" data-toggle="tab">上传图片</a></li>
				</ul>

				<div class="tab-content">
					<div role="tabpanel" class="tab-pane" id="img_online">
						<p>图片书写格式为：图片地址 + 空格 + "图片提示"。</p>

						<div class="input-group">
							<span class="input-group-addon"><i class="icon-picture"></i></span>
							<input id="input-insert-image" type="text" class="col-sm-5 form-control" placeholder="http://example.com/image.jpg &quot;optional title&quot;">
						</div>
					</div>
					<div role="tabpanel" class="tab-pane active" id="img_upload">
						<!--<iframe src="http://ask.csdn.net/upload.html"></iframe>-->
						<iframe name="up_img" src="./README_files/UploadImgMarkdown.html"></iframe>
					</div>
				</div>
			</div>
			<div class="modal-footer">
				<a class="btn btn-default" data-dismiss="modal">取消</a> <a class="btn btn-danger action-insert-image" data-dismiss="modal">确定</a>
			</div>
		</div>
	</div>
</div>

<div class="modal fade modal-start-abstract modal-bg">
	<div class="modal-dialog">
		<div class="modal-content">
			<div class="modal-header">
				<button type="button" class="close" data-dismiss="modal" aria-hidden="true">×</button>
				<h4 class="modal-title">摘要 标签</h4>
			</div>
			<div class="modal-body">
				<div id="alert-nll" class="alert alert-danger fade"><i class="icon-warning"></i> 摘要内容不能为空！</div>
				<div class="form-group">
					<textarea class="form-control" rows="10" maxlength="210" id="input-blog-description" required="required"> </textarea>
				</div>
				<div class="form-group">
					<div id="blog_tagRecomm">正在获取推荐标签<i class="icon-spinner icon-spin"></i></div>
					<div class="div_tags clearfix">
						<div id="tags-con-blog" class="tags_con">
							<span class="tags_label"></span>
							<input type="text" maxlength="10">
						</div>
						<input type="hidden" id="tags-val-blog" class="form-control" value="">
					</div>
				</div>
			</div>
			<div class="modal-footer">
				<a class="btn btn-default" data-dismiss="modal">取消</a>
				<a class="btn btn-danger" id="csdn-tags-blog-button">确定</a>
			</div>
		</div>
	</div>
</div>


<div class="modal fade modal-doc-setting modal-bg">
	<div class="modal-dialog">
		<div class="modal-content">
			<div class="modal-header">
				<button type="button" class="close" data-dismiss="modal" aria-hidden="true">×</button>
				<h4 class="modal-title">文章设置</h4>
			</div>
			<div class="modal-body">
				<div class="form-horizontal">
					<div class="form-group">
						<label class="col-sm-2 control-label">文章类型</label>

						<div class="col-sm-5">
							<select class="form-control" id="input-blog-type">
								<option value="0" selected="selected">请选择</option>
								<option value="original">原创</option>
								<option value="repost">转载</option>
								<option value="translated">翻译</option>
							</select>
						</div>
					</div>
					<div class="form-group">
						<label class="col-sm-2 control-label">个人分类</label>

						<div class="col-sm-10">
							<div id="tags-con-categories" class="tags_con">
								<span class="tags_label"></span>
								<input type="text" maxlength="10">
							</div>
							<input type="hidden" class="form-control" id="tags-val-categories" value="">

							<div id="blog-categories-loadding"></div>
							<div id="blog-categories-box"></div>
						</div>
					</div>
					<div class="form-group">
						<label class="col-sm-2 control-label">文章分类</label>

						<div class="col-sm-5">
							<select class="form-control" name="radChl" id="input-blog-channel">
								<option value="0" selected="selected">请选择</option>
								<option value="1">移动开发</option>
								<option value="14">Web前端</option>
								<option value="15">架构设计</option>
								<option value="16">编程语言</option>
								<option value="17">互联网</option>
								<option value="6">数据库</option>
								<option value="12">系统运维</option>
								<option value="2">云计算</option>
								<option value="3">研发管理</option>
								<option value="7">综合</option>
							</select>
						</div>
						<div class="col-sm-3"><p class="form-control-static">分类到首页</p></div>
					</div>

                    <div class="form-group">
						<label class="col-sm-2 control-label" title="添加关联的文章url">更多文章</label>
						<div class="col-sm-10">
							<div class="tags_con">
                                  <input id="inputArticleMore" type="text" style="width:446px;" maxlength="100">                                  
							</div>	
                            	
                             <input type="button" id="submitMoreArticle" value="添加">   					
						</div>
                        <div class="col-smtags_con10">
                            <div id="articlePanle" style="margin-left:113px"></div>       
                        </div>
					</div>

                  <!-- <div class="form-group" id="panelCopyright" style="display:none" >
						<label class="col-sm-2 control-label">版权声明</label>
						<div class="col-sm-10">
							<div class="tags_con">							
								 <input type="text" maxlength="100" id="inputCopyright"  value="本文为博主原创文章，未经博主允许不得转载。"  style="width:446px" >
							</div>							
						</div>
					</div>-->

					<div class="form-group" style="display:none">
						<div class="col-sm-5 col-sm-offset-2">
							<label class="checkbox-inline">
								<input type="checkbox" id="input-blog-level">发布到CSDN博客首页
							</label>
						</div>
						<div class="col-sm-5"><p class="form-control-static">需积分&gt;100且文章类型非转载</p></div>
					</div>
				</div>
			</div>
			<div class="modal-footer">
                <a class="btn btn-default" data-dismiss="modal">取消</a>
				<a class="btn btn-danger" id="csdn-channel-blog-button" data-dismiss="modal">确定</a>
				<a class="btn btn-danger" id="csdn-post-blog-button" data-dismiss="modal">发布</a>
			</div>
		</div>
	</div>
</div>

<div class="modal fade modal-export-doc">
	<div class="modal-dialog">
		<div class="modal-content">
			<div class="modal-header">
				<button type="button" class="close" data-dismiss="modal" aria-hidden="true">×</button>
				<h4 class="modal-title">导出文档</h4>
			</div>
			<div class="modal-body">
				<ul class="nav">
					<li><a class="action-download-md"><i class="icon-file-markdown"></i>

						<p>Markdown</p></a></li>
					<li><a class="action-download-html"><i class="icon-file-html"></i>

						<p>仅内容HTML</p></a></li>
					<li><a class="action-download-template"><i class="icon-file-html-t"></i>

						<p>带模板HTML</p></a></li>
					<!--<li><a class="action-download-pdf"><i class="icon-file-pdf"></i> <p>PDF文档</p></a></li>-->
				</ul>
			</div>
		</div>
	</div>
</div>
<div class="modal-helps hide">
	<div class="modal-dialog">
		<div class="modal-content">
			<div class="modal-header">
				<button type="button" class="close" onclick="$(&#39;.modal-helps&#39;).addClass(&#39;hide&#39;)">×</button>
				<h4 class="modal-title">Markdown语法帮助</h4>
			</div>
			<div class="modal-body">
				<ol class="help-menu">
					<li class="active"><a>标题</a>

						<div>
							<h4>在文字写书写不同数量的#可以完成不同的标题，如下：</h4>

							<p># 一级标题<br>
								## 二级标题<br>
								### 三级标题<br>
								#### 四级标题<br>
								##### 五级标题<br>
								###### 六级标题</p>
							<h4>等号及减号也可以进行标题的书写，不过只能书写二级标题，并且需要写在文字的下面，减号及等号的数量不会影响标题的基数，如下：</h4>

							<p>二级标题<br>
								=========</p>

							<p>二级标题<br>
								---------</p>
						</div>
					</li>
					<li><a>列表</a>

						<div>
							<h4>无序列表的使用，在符号“-”后加空格使用。如下：</h4>

							<p> - 无序列表1<br>
								- 无序列表2<br>
								- 无序列表3</p>
							<h4>如果要控制列表的层级，则需要在符号“-”前使用空格。如下：</h4>

							<p> - 无序列表1<br>
								- 无序列表2<br>
								&nbsp;&nbsp;- 无序列表2.1<br>
								&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;- 列表内容<br>
								&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;- 列表内容</p>

							<p>&nbsp;</p>
							<h4>有序列表的使用，在数字及符号“.”后加空格几个，如下：</h4>

							<p> 1. 有序列表1<br>
								2. 有序列表2<br>
								3. 有序列表3<br>
							</p>
							<h4>有序列表如果要区分层级，也可以在数字前加空格。</h4>
						</div>
					</li>
					<li><a>引用</a>

						<div>
							<h4>引用的格式是使用符号“&gt;”后面书写文字，及可以使用引用。如下：</h4>

							<p> &gt;这个是引用<br> &gt; 是不是和电子邮件中的<br>
								&gt; 引用格式很像</p>
						</div>
					</li>
					<li><a>粗体与斜体</a>

						<div>
							<h4>
								粗体的使用是在需要加粗的文字前后各加两个“*”，而斜体的使用则是在需要斜体的文字前后各加一个“*”，如果要使用粗体和斜体，那么就是在需要操作的文字前后各加三个“*”。如下：</h4>

							<p> **这个是粗体**<br>
								*这个是斜体*<br>
								***这个是粗体加斜体***</p>
						</div>
					</li>
					<li><a>链接与图片</a>

						<div>
							<h4>在文中直接加链接，中括号中是需要添加链接的文字，圆括号中是需要添加的链接，如下：</h4>

							<p> [link text](http://example.com/ "optional title")</p>
							<h4>在引用中加链接，第一个中括号添加需要添加的文字，第二个中括号中是引用链接的id，之后在引用中，使用id加链接：如下：</h4>

							<p> [link text][id]<br>
								[id]: http://example.com/ "optional title here"</p>
							<h4>在文中直接引用链接，直接使用尖括号，把链接加入到尖括号中就可以实现，如下：</h4>

							<p>&lt;http://example.com/&gt; or &lt;address@example.com&gt;<br>
							</p>
							<h4>插入互联网上图片，格式如下：</h4>

							<p>![这里写图片描述](http://img3.douban.com/mpic/s1108264.jpg)<br>
								![这里写图片描述][jane-eyre-douban]<br>
								[jane-eyre-douban]: http://img3.douban.com/mpic/s1108264.jpg</p>
						</div>
					</li>
					<li><a>代码块</a>

						<div>
							<h4>用TAB键起始的段落，会被认为是代码块，如下：</h4>

							<p> &nbsp;&nbsp;&nbsp;&nbsp;&lt;php&gt;<br>
								&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;echo “hello world";<br>
								&nbsp;&nbsp;&nbsp;&nbsp;&lt;/php&gt;</p>
							<h4>如果在一个行内需要引用代码，只要用反引号`引起来就好，如下：</h4>

							<p> Use the <code>`printf()` </code>function.</p>
						</div>
					</li>
					<li><a>分割线与删除线</a>

						<div>
							<h4>可以在一行中用三个以上的星号、减号、底线来建立一个分隔线，同时需要在分隔线的上面空一行。如下：</h4>

							<p> ---<br>
								****<br>
								___</p>
							<h4>删除线的使用，在需要删除的文字前后各使用两个符合“~”，如下</h4>

							<p> ~~Mistaken text.~~</p>
						</div>
					</li>
					<li><a>代码块与语法高亮</a>

						<div>
							<h4>在需要高亮的代码块的前一行及后一行使用三个反引号“`”，同时第一行反引号后面表面代码块所使用的语言，如下：</h4>

							<p>```ruby<br>
								require 'redcarpet'<br>
								markdown = Redcarpet.new("Hello World!")<br>
								puts markdown.to_html<br>
								```</p>
						</div>
					</li>
					<li><a>表格</a>

						<div>
							<h4>可以使用冒号来定义表格的对齐方式，如下：</h4>

							<p> | Tables | Are | Cool |<br>
								| ------------- |:-------------:| -----:|<br>
								| col 3 is | right-aligned | $1600 |<br>
								| col 2 is | centered | $12 |<br>
								| zebra stripes | are neat | $1 |<br>
							</p>
						</div>
					</li>
					<li><a>LaTex数学公式</a>

						<div>
							<h4>使用MathJax渲染*LaTex* 数学公式，详见[<a href="http://math.stackexchange.com/" target="_blank">math.stackexchange.com</a>]
							</h4>
							<h4>行内公式，数学公式为：</h4>

							<p>$\Gamma(n) = (n-1)!\quad\forall n\in\mathbb N$<br>
							</p>
							<h4>块级公式：</h4>

							<p>$$ x = \dfrac{-b \pm \sqrt{b^2 - 4ac}}{2a} $$</p>

							<p>更多LaTex语法请参考 [<a href="http://meta.math.stackexchange.com/questions/5020/mathjax-basic-tutorial-and-quick-reference" target="_blank">这儿</a>]。</p>

						</div>
					</li>
					<li><a>UML 图</a>

						<div>
							<h4>可以渲染序列图：</h4>

							<p>```sequence<br>
								张三-&gt;李四: 嘿，小四儿, 写博客了没?<br>
								Note right of 李四: 李四愣了一下，说：<br>
								李四--&gt;张三: 忙得吐血，哪有时间写。<br>
								```</p>
							<h4>或者流程图：</h4>

							<p>```flow<br>
								st=&gt;start: 开始<br>
								e=&gt;end: 结束<br>
								op=&gt;operation: 我的操作<br>
								cond=&gt;condition: 确认？</p>

							<p>st-&gt;op-&gt;cond<br>
								cond(yes)-&gt;e<br>
								cond(no)-&gt;op<br>
								```</p>

							<p>- 关于 **序列图** 语法，参考 [<a href="http://bramp.github.io/js-sequence-diagrams/" target="_blank">这儿</a>]<br>
								- 关于 **流程图** 语法，参考 [<a href="http://adrai.github.io/flowchart.js/" target="_blank">这儿</a>]<br>
							</p>
						</div>
					</li>
				</ol>

			</div>
		</div>
	</div>
</div>

<div class="modal fade modal-alert">
	<div class="modal-dialog">
		<div class="modal-content">
			<div class="modal-header">
				<!--<button type="button" class="close" data-dismiss="modal" aria-hidden="true">&times;</button>-->
			</div>
			<div class="modal-body">
			</div>
			<div class="modal-footer text-center">
			</div>
		</div>
	</div>
</div>


<div class="modal fade modal-import-image">
	<div class="modal-dialog">
		<div class="modal-content">

			<div class="modal-header">
				<button type="button" class="close" data-dismiss="modal" aria-hidden="true">×</button>
				<h2 class="modal-title">Google+ image import</h2>
			</div>
			<div class="modal-body">
				<div class="form-horizontal">
					<div class="form-group">
						<div class="col-sm-7">
							<img>
						</div>
					</div>
					<div class="form-group">
						<label class="col-sm-4 control-label" for="input-import-image-title">Title (optional)</label>

						<div class="col-sm-7">
							<input type="text" id="input-import-image-title" placeholder="Image title" class="form-control">
						</div>
					</div>
					<div class="form-group">
						<label class="col-sm-4 control-label" for="input-import-image-size">Size limit (optional)</label>

						<div class="col-sm-7 form-inline">
							<input type="text" id="input-import-image-size" placeholder="0" class="col-sm-3 form-control"> px
						</div>
					</div>
				</div>
			</div>
			<div class="modal-footer">
				<a class="btn btn-default" data-dismiss="modal">取消</a>
				<a class="btn btn-danger action-import-image" data-dismiss="modal">确定</a>
			</div>
		</div>
	</div>
</div>


<div class="modal fade modal-remove-file-confirm">
	<div class="modal-dialog">
		<div class="modal-content">

			<div class="modal-header">
				<button type="button" class="close" data-dismiss="modal" aria-hidden="true">×</button>
				<h2 class="modal-title">Delete</h2>
			</div>
			<div class="modal-body">
				<p>
					Are you sure you want to delete "<span class="file-title">欢迎使用CSDN-markdown编辑器</span>"?
				</p>
				<blockquote>
					<p><b>Note:</b> It won't delete the file on synchronized locations.</p>
				</blockquote>
			</div>
			<div class="modal-footer">
				<a class="btn btn-default" data-dismiss="modal">取消</a>
				<a class="btn btn-danger action-remove-file" data-dismiss="modal">Delete</a>
			</div>
		</div>
	</div>
</div>


<div class="modal fade modal-import-url">
	<div class="modal-dialog">
		<div class="modal-content">
			<div class="modal-header">
				<button type="button" class="close" data-dismiss="modal" aria-hidden="true">×</button>
				<h4 class="modal-title">从线上导入</h4>
			</div>
			<div class="modal-body">
				<div class="alert alert-danger">
					<p><i class="icon-warning"></i> 注意：当前操作只要成功导入在线.md地址就会冲掉当前的文章，请确定是否保存或者发表后再进行导入操作。</p>

					<p>PS：IE11以下版本不支持导入功能。</p>
				</div>
				<p>请输入在线的markdown文件地址：</p>

				<div class="form-horizontal">
					<div class="form-group">
						<label class="col-sm-2 control-label" for="input-import-url">URL</label>

						<div class="col-sm-9">
							<input type="text" id="input-import-url" placeholder="http://code.csdn.net/test.md" class="form-control">
						</div>
					</div>
				</div>
			</div>
			<div class="modal-footer">
				<a class="btn btn-default" data-dismiss="modal">取消</a>
				<a data-dismiss="modal" class="btn btn-danger action-import-url">确定</a>
			</div>
		</div>
	</div>
</div>


<div class="modal fade modal-import-harddrive-markdown">
	<div class="modal-dialog">
		<div class="modal-content">
			<div class="modal-header">
				<button type="button" class="close" data-dismiss="modal" aria-hidden="true">×</button>
				<h4 class="modal-title">从本机导入</h4>
			</div>
			<div class="modal-body">
				<div class="alert alert-danger">
					<p><i class="icon-warning"></i> 注意：当前操作只要成功选择或者拖入.md文件就会冲掉当前的文章，请确定是否保存或者发表后再进行导入操作。</p>

					<p>PS：IE11以下版本不支持导入功能。</p>
				</div>
				<p>请选择你要导入的.md文档</p>

				<p>
					<input type="file" id="input-file-import-harddrive-markdown" multiple="" class="form-control">
				</p>

				<p>或者将.md文档拖到这里</p>

				<p id="dropzone-import-harddrive-markdown" class="drop-zone">拖拽文件到这里</p>
			</div>
			<div class="modal-footer">
				<a class="btn btn-danger" data-dismiss="modal">关闭</a>
			</div>
		</div>
	</div>
</div>


<div class="modal fade modal-import-harddrive-html">
	<div class="modal-dialog">
		<div class="modal-content">

			<div class="modal-header">
				<button type="button" class="close" data-dismiss="modal" aria-hidden="true">×</button>
				<h2 class="modal-title">Convert HTML to Markdown</h2>
			</div>
			<div class="modal-body">
				<p>Please select your HTML files here:</p>

				<p>
					<input type="file" id="input-file-import-harddrive-html" multiple="" class="form-control">
				</p>

				<p>Or drag and drop your HTML files here:</p>

				<p id="dropzone-import-harddrive-html" class="drop-zone">Drop files here</p>

				<p>Or insert your HTML code here:</p>
				<textarea id="input-convert-html" class="form-control"></textarea>
			</div>
			<div class="modal-footer">
				<a class="btn btn-default" data-dismiss="modal">取消</a>
				<a class="btn btn-danger action-convert-html" data-dismiss="modal">确定</a>
			</div>
		</div>
	</div>
</div>

<div class="modal fade modal-publish">
	<div class="modal-dialog">
		<div class="modal-content">
			<div class="modal-header">
				<button type="button" class="close" data-dismiss="modal" aria-hidden="true">×</button>
				<h2 class="modal-title">
					Publish on <span class="publish-provider-name"></span>
				</h2>
			</div>
			<div class="modal-body">
				<div class="form-horizontal">
					<div class="form-group modal-publish-ssh">
						<label class="col-sm-4 control-label" for="input-publish-ssh-host">Host</label>

						<div class="col-sm-7">
							<input type="text" id="input-publish-ssh-host" placeholder="hostname.or.ip" class="form-control"> <span class="help-block"> Host must be accessible publicly, unless you're hosting your own StackEdit instance.
							</span>
						</div>
					</div>
					<div class="form-group modal-publish-ssh">
						<label class="col-sm-4 control-label" for="input-publish-ssh-port">Port
							(optional)</label>

						<div class="col-sm-2">
							<input type="text" id="input-publish-ssh-port" placeholder="22" class="form-control">
						</div>
					</div>
					<div class="form-group modal-publish-ssh">
						<label class="col-sm-4 control-label" for="input-publish-ssh-username">Username</label>

						<div class="col-sm-7">
							<input type="text" id="input-publish-ssh-username" placeholder="username" class="form-control">
						</div>
					</div>
					<div class="form-group modal-publish-ssh">
						<label class="col-sm-4 control-label" for="input-publish-ssh-password">Password</label>

						<div class="col-sm-7">
							<input type="password" id="input-publish-ssh-password" placeholder="password" class="form-control">
						</div>
					</div>
					<div class="form-group modal-publish-github">
						<label class="col-sm-4 control-label" for="input-publish-github-repo">Repository</label>

						<div class="col-sm-7">
							<input type="text" id="input-publish-github-repo" placeholder="Repository name or URL" class="form-control">
						</div>
					</div>
					<div class="form-group modal-publish-github">
						<label class="col-sm-4 control-label" for="input-publish-github-branch">Branch</label>

						<div class="col-sm-7">
							<input type="text" id="input-publish-github-branch" placeholder="branch-name" class="form-control">
						</div>
					</div>
					<div class="form-group modal-publish-ssh modal-publish-github">
						<label class="col-sm-4 control-label" for="input-publish-file-path">File path</label>

						<div class="col-sm-7">
							<input type="text" id="input-publish-file-path" placeholder="path/to/file.md" class="form-control">
							<span class="help-block"> File path is composed of both
								folder and filename. </span>
						</div>
					</div>
					<div class="form-group modal-publish-gist">
						<label class="col-sm-4 control-label" for="input-publish-filename">Filename</label>

						<div class="col-sm-7">
							<input type="text" id="input-publish-filename" placeholder="filename" class="form-control">
						</div>
					</div>
					<div class="form-group modal-publish-gist">
						<label class="col-sm-4 control-label" for="input-publish-gist-id">Existing
							ID (optional)</label>

						<div class="col-sm-7">
							<input type="text" id="input-publish-gist-id" placeholder="GistID" class="form-control">
						</div>
					</div>
					<div class="form-group modal-publish-gist">
						<label class="col-sm-4 control-label" for="input-publish-gist-public">Public</label>

						<div class="col-sm-7">
							<div class="checkbox">
								<input type="checkbox" id="input-publish-gist-public" checked="checked">
							</div>
						</div>
					</div>
					<div class="form-group modal-publish-blogger modal-publish-bloggerpage">
						<label class="col-sm-4 control-label" for="input-publish-blogger-url">Blog URL</label>

						<div class="col-sm-7">
							<input type="text" id="input-publish-blogger-url" placeholder="http://exemple.blogger.com/" class="form-control">
						</div>
					</div>
					<div class="form-group modal-publish-tumblr">
						<label class="col-sm-4 control-label" for="input-publish-tumblr-hostname">Blog hostname</label>

						<div class="col-sm-7">
							<input type="text" id="input-publish-tumblr-hostname" placeholder="exemple.tumblr.com" class="form-control">
						</div>
					</div>
					<div class="form-group modal-publish-wordpress">
						<label class="col-sm-4 control-label" for="input-publish-tumblr-hostname">WordPress site</label>

						<div class="col-sm-7">
							<input type="text" id="input-publish-wordpress-site" placeholder="exemple.wordpress.com" class="form-control">
							<span class="help-block"> <a target="_blank" href="http://jetpack.me/">Jetpack plugin</a> is required for
								self-hosted sites.
							</span>
						</div>
					</div>
					<div class="form-group modal-publish-blogger modal-publish-tumblr modal-publish-wordpress">
						<label class="col-sm-4 control-label" for="input-publish-postid">Update
							existing post ID (optional)</label>

						<div class="col-sm-7">
							<input type="text" id="input-publish-postid" placeholder="PostID" class="form-control">
						</div>
					</div>
					<div class="form-group modal-publish-bloggerpage">
						<label class="col-sm-4 control-label" for="input-publish-pageid">Update
							existing page ID (optional)</label>

						<div class="col-sm-7">
							<input type="text" id="input-publish-pageid" placeholder="PageID" class="form-control">
						</div>
					</div>
					<div class="form-group modal-publish-dropbox">
						<label class="col-sm-4 control-label" for="input-publish-dropbox-path">File path</label>

						<div class="col-sm-7">
							<input type="text" id="input-publish-dropbox-path" placeholder="/path/to/My Document.html" class="form-control">
							<span class="help-block"> File path is composed of both
								folder and filename. </span>
						</div>
					</div>
					<div class="form-group modal-publish-gdrive">
						<label class="col-sm-4 control-label" for="input-publish-gdrive-fileid">File ID (optional)</label>

						<div class="col-sm-7">
							<input type="text" id="input-publish-gdrive-fileid" placeholder="FileID" class="form-control"> <span class="help-block">If no file ID is supplied, a new file
								will be created in your Google Drive root folder. You can move
								the file afterwards within Google Drive.</span>
						</div>
					</div>
					<div class="form-group modal-publish-gdrive">
						<label class="col-sm-4 control-label" for="input-publish-gdrive-filename">Force filename
							(optional)</label>

						<div class="col-sm-7">
							<input type="text" id="input-publish-gdrive-filename" placeholder="Filename" class="form-control"> <span class="help-block">If no file name is supplied, the
								document title will be used.</span>
						</div>
					</div>

					<div class="form-group">
						<label class="col-sm-4 control-label">Format</label>

						<div class="col-sm-7">
							<div class="radio">
								<label> <input type="radio" name="radio-publish-format" value="markdown"> Markdown
								</label>
							</div>
							<div class="radio">
								<label> <input type="radio" name="radio-publish-format" value="html"> HTML
								</label>
							</div>
							<div class="radio">
								<label> <input type="radio" name="radio-publish-format" value="template"> Template
								</label>
							</div>
						</div>
					</div>
					<div class="collapse publish-custom-template-collapse">
						<div class="form-group">
							<div class="col-sm-4"></div>
							<div class="col-sm-7">
								<div class="checkbox">
									<label> <input type="checkbox" id="checkbox-publish-custom-template"> Custom template
									</label> <a class="tooltip-template">(?)</a>
								</div>
							</div>
						</div>
						<div class="form-group">
							<div class="col-sm-4"></div>
							<div class="col-sm-7">
								<textarea class="form-control" id="textarea-publish-custom-template" disabled=""></textarea>
							</div>
						</div>
					</div>
				</div>
				<blockquote class="front-matter-info modal-publish-blogger modal-publish-tumblr modal-publish-wordpress">
					<p><b>Tip:</b> You can use a
						<a href="http://jekyllrb.com/docs/frontmatter/" target="_blank">YAML front matter</a> to specify the title and the tags/labels of your
						publication.</p>

					<p><b>Interpreted variables:</b> <code>title</code>, <code>tags</code>, <code>published</code>,
						<code>date</code>.</p>
				</blockquote>
				<blockquote class="front-matter-info modal-publish-bloggerpage">
					<p><b>Tip:</b> You can use a
						<a href="http://jekyllrb.com/docs/frontmatter/" target="_blank">YAML front matter</a> to specify the title of your page.</p>

					<p><b>Interpreted variables:</b> <code>title</code>.</p>
				</blockquote>
				<blockquote class="url-info modal-publish-bloggerpage">
					<p><b>About URL:</b> For newly created page , Blogger API will append a generated number to the url
						like <code>about-me-1234.html</code>, if you deeply care about your URL naming, you should first
						create the page on Blogger and then update them with StackEdit specifying the pageId when
						publishing.
					</p>

					<p><b>About page visibility:</b> Blogger API does not respect published status for pages.When
						publishing the page to Blogger, the page will be <strong>live</strong> but not added to the page
						listing. You should arrange the page listing from Blogger dashboard.
					</p>
				</blockquote>
			</div>
			<div class="modal-footer">
				<a class="btn btn-default" data-dismiss="modal">取消</a>
				<a data-dismiss="modal" class="btn btn-danger action-process-publish">确定</a>
			</div>
		</div>
	</div>
</div>


<div class="modal fade modal-manage-publish">
	<div class="modal-dialog">
		<div class="modal-content">

			<div class="modal-header">
				<button type="button" class="close" data-dismiss="modal" aria-hidden="true">×</button>
				<h2 class="modal-title">Publication</h2>
			</div>
			<div class="modal-body">
				<p>
					"<span class="file-title">欢迎使用CSDN-markdown编辑器</span>" is published on the following
					location(s):
				</p>

				<div class="publish-list"><div class="entry">
    <div class="input-group">
	<span class="input-group-addon" title="CsdnBlog">
		<i class="icon-provider-csdnblog"></i>
	</span> <input class="form-control" type="text" value="id:, htmlContent:&lt;h1 id=\cloudgo\&gt;Cloudgo&lt;/h1&gt;\n\n&lt;h2 id=\框架选择\&gt;框架选择&lt;/h2&gt;\n\n&lt;p&gt;这里所选择使用的框架是课程群里推荐的Martini。Martini 是一个非常新的 Go 语言的 Web 框架，使用 Go 的 net/http 接口开发，类似 Sinatra 或者 Flask 之类的框架，你可使用自己的 DB 层、会话管理和模板。&lt;/p&gt;\n\n&lt;p&gt;特性：&lt;/p&gt;\n\n&lt;p&gt;&lt;strong&gt;·&lt;/strong&gt;   使用非常简单 &lt;br&gt;\n&lt;strong&gt;·&lt;/strong&gt;   无侵入设计 &lt;br&gt;\n&lt;strong&gt;·&lt;/strong&gt;   可与其他 Go 的包配合工作 &lt;br&gt;\n&lt;strong&gt;·&lt;/strong&gt;   超棒的路径匹配和路由 &lt;br&gt;\n&lt;strong&gt;·&lt;/strong&gt;   模块化设计，可轻松添加工具 &lt;br&gt;\n&lt;strong&gt;·&lt;/strong&gt;   大量很好的处理器和中间件 &lt;br&gt;\n&lt;strong&gt;·&lt;/strong&gt;   很棒的开箱即用特性 &lt;br&gt;\n&lt;strong&gt;·&lt;/strong&gt;   完全兼容 http.HandlerFunc 接口&lt;/p&gt;\n\n&lt;h2 id=\框架安装\&gt;框架安装&lt;/h2&gt;\n\n&lt;p&gt;直接使用 &lt;br&gt;\n&lt;code&gt;go get github.com/go-martini/martini&lt;/code&gt;&lt;/p&gt;\n\n&lt;p&gt;即可。&lt;/p&gt;\n\n&lt;h2 id=\测试\&gt;测试&lt;/h2&gt;\n\n&lt;h4 id=\启动服务器\&gt;&lt;strong&gt;·&lt;/strong&gt;   启动服务器&lt;/h4&gt;\n\n&lt;p&gt;&lt;code&gt;go run main.go -p9090&lt;/code&gt; &lt;br&gt;\n然后在浏览器中（不关闭终端）打开 &lt;br&gt;\n&lt;code&gt;http://localhost:9090/hello/主机名&lt;/code&gt; &lt;br&gt;\n可以看到页面上出现service中的字符： &lt;br&gt;\n&lt;img src=\http://img.blog.csdn.net/20171114223654673?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvbGVwcmVjaGF1bl8=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast\ alt=\浏览器\ title=\\&gt; &lt;br&gt;\n此时终端中的监听信息如下 &lt;br&gt;\n&lt;img src=\http://img.blog.csdn.net/20171114223722445?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvbGVwcmVjaGF1bl8=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast\ alt=\这里写图片描述\ title=\\&gt;&lt;/p&gt;\n\n&lt;h4 id=\进行curl测试\&gt;&lt;strong&gt;·&lt;/strong&gt;   进行curl测试&lt;/h4&gt;\n\n&lt;p&gt;在不关闭当前终端的情况下开启另一个终端，并在该终端上使用curl命令 &lt;br&gt;\n&lt;code&gt;curl -v http://localhost:9090/hello/主机名&lt;/code&gt; &lt;br&gt;\n&lt;img src=\http://img.blog.csdn.net/20171114223949301?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvbGVwcmVjaGF1bl8=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast\ alt=\这里写图片描述\ title=\\&gt;&lt;/p&gt;\n\n&lt;h4 id=\进行apachebench测试\&gt;&lt;strong&gt;·&lt;/strong&gt;   进行ApacheBench测试&lt;/h4&gt;\n\n&lt;p&gt;使用命令 &lt;br&gt;\n&lt;code&gt;ab -n 1000 -c 100 http://localhost:9090/hello/主机名&lt;/code&gt; &lt;br&gt;\n测试结果如下 &lt;br&gt;\n&lt;img src=\http://img.blog.csdn.net/20171114224242549?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvbGVwcmVjaGF1bl8=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast\ alt=\这里写图片描述\ title=\\&gt; &lt;br&gt;\n&lt;img src=\http://img.blog.csdn.net/20171114224253357?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvbGVwcmVjaGF1bl8=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast\ alt=\这里写图片描述\ title=\\&gt; &lt;br&gt;\n从结果中我们可以得到，总共发送了1000个请求，失败0个，耗时0.207秒，其中50%的请求可以在19ms内被响应，90%的请求可以在25ms内被响应，最长的响应时间为35ms。&lt;/p&gt;, markdowndirectory:, type:0, channel:0, categories:, description:, tags:, level:0, status:2, articleedittype:1, articlemore:, format:markdown" disabled="">
    </div>
    <div class="text-right">
        <a class="btn btn-link btn-sm open-location" target="_blank" href="https://blog.csdn.net/">
            <i class="icon-link-ext-alt"></i> Open in CsdnBlog
        </a>
        <a class="btn btn-link btn-sm remove-button" data-publish-index="publish.cBC6qD1XNU1ifpMVEKejSQGQ"><i class="icon-trash"></i> Remove location</a>
    </div>
</div>
</div>
				<blockquote>
					<p><b>Note:</b> Removing a publish location will not delete the actual publication.</p>
				</blockquote>
			</div>
			<div class="modal-footer">
				<a class="btn btn-danger" data-dismiss="modal">Close</a>
			</div>
		</div>
	</div>
</div>

<div class="modal fade modal-non-unique">
	<div class="modal-dialog">
		<div class="modal-content">
			<div class="modal-header">
				<h4 class="modal-title">不好意思...</h4>
			</div>
			<div class="modal-body">
				<p>Markdown编辑器已经停止，因为另一个实例中相同的浏览器中运行。</p>
				<blockquote>
					<p>如果你想重新打开Markdown编辑器，请点击 “重新加载”</p>
				</blockquote>
			</div>
			<div class="modal-footer">
				<a href="javascript:window.location.reload();" class="btn btn-danger">重新加载</a>
			</div>
		</div>
	</div>
</div>


<div class="modal fade modal-redirect-confirm">
	<div class="modal-dialog">
		<div class="modal-content">

			<div class="modal-header">
				<h2 class="modal-title">Redirection</h2>
			</div>
			<div class="modal-body">
				<p class="redirect-msg"></p>
				<blockquote>
					<p>Please click <b>确定</b> to proceed.</p>
				</blockquote>
			</div>
			<div class="modal-footer">
				<a class="btn btn-default" data-dismiss="modal">取消</a>
				<a class="btn btn-danger action-redirect-confirm" data-dismiss="modal">确定</a>
			</div>
		</div>
	</div>
</div>


<div class="modal fade modal-app-reset">
	<div class="modal-dialog">
		<div class="modal-content">

			<div class="modal-header">
				<h2 class="modal-title">Reset application</h2>
			</div>
			<div class="modal-body">
				<p>This will delete all your local documents.</p>
				<blockquote><b>Are you sure?</b></blockquote>
			</div>
			<div class="modal-footer">
				<a class="btn btn-default" data-dismiss="modal">取消</a>
				<a class="btn btn-danger action-app-reset" data-dismiss="modal">确定</a>
			</div>
		</div>
	</div>
</div>


<div class="modal fade modal-import-docs-settings">
	<div class="modal-dialog">
		<div class="modal-content">

			<div class="modal-header">
				<h2 class="modal-title">Import documents and settings</h2>
			</div>
			<div class="modal-body">
				<p>This will delete all existing local documents.</p>
				<blockquote><b>Are you sure?</b></blockquote>
			</div>
			<div class="modal-footer">
				<a class="btn btn-default" data-dismiss="modal">取消</a>
				<a class="btn btn-danger action-import-docs-settings-confirm" data-dismiss="modal">确定</a>
			</div>
		</div>
	</div>
</div>


</div>
       
    <div id="loginform" style="display: none"></div>    
     






<ul class="complete_list"></ul><ul class="complete_list" style="display: none;"></ul></body><input style="width:1px;height:1px;border:none;margin:0;padding:0;" tabindex="-1"></html>