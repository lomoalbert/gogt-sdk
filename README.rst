Gt Golang SDK
===============
极验验证的非官方Golang SDK目前提供基于Beego框架的DEMO。

本项目是面向服务器端的，具体使用可以参考极验的 `文档 <http://www.geetest.com/install/sections/idx-server-sdk.html>`_ ,客户端相关开发请参考极验的 `前端文档。 <http://www.geetest.com/install/>`_.

开发环境
----------------

 - Golang (推荐6.2以上版本）
 - Beego框架

快速开始
---------------

下面使用示例代码的均以Beego框架为例。

1. 获取代码

从 `Github <https://github.com/lomoalbert/gtsdk>`__ 上Clone代码:

.. code-block:: bash

    $ go get https://github.com/lomoalbert/gtsdk.git

2. 初始化验证


在调用GeetestLib前请自行在app.conf设定公钥和私钥,用户id为可选项，默认为空字符串：

.. code-block :: conf

  PrivateKey = "你的私钥"
  CaptchaID = "你的公钥"

根据自己的私钥初始化验证

.. code-block :: golang

    func (ctl *RegisterController)Get() {
        userID := ""
        gt := gtsdk.GeetestLib(PrivateKey, CaptchaID)
        status := gt.PreProcess(userID)
        ctl.SetSession(gtsdk.GT_STATUS_SESSION_KEY, status)
        ctl.SetSession("user_id", userID)
        responseStr := gt.GetResponseStr()
        ctl.Ctx.WriteString(responseStr)
    }

4. 二次验证

.. code-block :: golang

    func (ctl *ValidateController)Post() {
        var result bool
        var respstr string
        gt := gtsdk.GeetestLib(PrivateKey, CaptchaID)
        challenge := ctl.GetString(gtsdk.FN_CHALLENGE)
        validate := ctl.GetString(gtsdk.FN_VALIDATE)
        seccode := ctl.GetString(gtsdk.FN_SECCODE)
        status := ctl.GetSession(gtsdk.GT_STATUS_SESSION_KEY).(int)
        userID := ctl.GetSession("user_id").(string)
        if status == 0 {
            result = gt.FailbackValidate(challenge, validate, seccode)
        } else {
            result = gt.SuccessValidate(challenge, validate, seccode, userID)
        }
        if result {
            respstr = "success"
        } else {
            respstr = "fail"
        }
        ctl.Ctx.WriteString(respstr)
    }


运行demo
---------------------

1. Beego demo运行：进入demo文件夹，运行：

.. code-block:: bash

    $ bee run

在浏览器中访问http://localhost:8080即可看到Demo界面

发布日志
-----------------
+ 3.2.0

 - 参照 `gt-python-sdk 3.2.0 <https://github.com/GeeTeam/gt-python-sdk/>`_实现极验接口
