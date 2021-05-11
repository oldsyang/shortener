/*
 * 作者:helang
 * 邮箱:helang.love@qq.com
 * jQuery插件库:https://www.jq22.com/mem395541
 * CSDN博客:https://blog.csdn.net/u013350495
 * 微信公众号:web-7258
 */

/* 搜索 */
var helangSearch = {
    /* 元素集 */
    els: {},
    /* 搜索类型序号 */
    searchIndex: 0,
    /* 火热的搜索列表 */
    hot: {
        /* 颜色 */
        color: ['#ff2c00', '#ff5a00', '#ff8105', '#fd9a15', '#dfad1c', '#6bc211', '#3cc71e', '#3cbe85', '#51b2ef', '#53b0ff'],
        /* 列表 */
        list: [
            'helang.love@qq.com',
            'qq:1846492969',
            'web前端梦之蓝',
            '公众号：web-7258',
            'jQuery插件库',
            'CSDN-IT博客',
            'jQuery之美-CSDN博客专栏',
            'jq22.com',
            'csdn.net',
            'mydarling.gitee.io'
        ]
    },
    /* 初始化 */
    init: function () {
        var _this = this;
        this.els = {
            pickerBtn: $(".picker"),
            pickerList: $(".picker-list"),
            logo: $(".logo"),
            hotList: $(".hot-list"),
            input: $("#search-input"),
            button: $(".search")
        };

        var hotListEle = this.els.hotList

        /* 设置热门搜索列表 */
        // this.els.hotList.html(function () {
        //     var str = '';
        //     $.each(_this.hot.list, function (index, item) {
        //         str += '<a href="https://www.baidu.com/s?ie=utf8&oe=utf8&tn=98010089_dg&ch=11&wd=' + item + '" target="_blank">'
        //             + '<div class="number" style="color: ' + _this.hot.color[index] + '">' + (index + 1) + '</div>'
        //             + '<div>' + item + '</div>'
        //             + '</a>';
        //     });
        //     return str;
        // });

        /* 注册事件 */
        /* 搜索类别选择按钮 */
        this.els.pickerBtn.click(function () {
            if (_this.els.pickerList.is(':hidden')) {
                setTimeout(function () {
                    _this.els.pickerList.show();
                }, 100);
            }
        });
        /* 搜索类别选择列表 */
        this.els.pickerList.on("click", ">li", function () {
            _this.els.logo.css("background-image", ('url(img/' + $(this).data("logo") + ')'));
            _this.searchIndex = $(this).index();
            _this.els.pickerBtn.html($(this).html())
        });
        // /* 搜索 输入框 点击*/
        // this.els.input.click(function () {
        //     if (!$(this).val()) {
        //         setTimeout(function () {
        //             _this.els.hotList.show();
        //         }, 100);
        //     }
        // });
        /* 搜索 输入框 输入*/
        this.els.input.on("input", function () {
            if ($(this).val()) {
                _this.els.hotList.hide();
            }
        });
        /* 搜索按钮 */
        this.els.button.click(function () {
            $.ajax({
                url: "/api/encode/",
                type: "POST",
                contentType: 'application/json;charset=UTF-8',
                data: JSON.stringify({
                    "url": $("input[name='content']").val()
                }),
                success: function (data) {
                    if (data.success === "true") {
                        var urlStr = '';
                        $.each([data.response.url], function (index, item) {
                            urlStr += '<a style="color:#333;float:right;margin-bottom:10px" href="' + item + '" target="_blank">'
                                + '<div style="display:inline;">' + '预览' + '</div>' + '</a>';
                        });

                        var str = '<div style="padding:20px; width:300px">' + data.response.url + '<br><br><br>' + urlStr + '</div>'
                        layer.open({
                            type: 1,
                            skin: 'layui-layer-demo', //样式类名
                            closeBtn: 0, //不显示关闭按钮
                            anim: 2,
                            shadeClose: true, //开启遮罩关闭
                            content: str
                        });

                        // hotListEle.html(function () {
                        //     var str = '';
                        //     $.each([data.response.url], function (index, item) {
                        //         str += '<a href="' + item + '" target="_blank">'
                        //             + '<div class="number" style="color: ' + _this.hot.color[index] + '">' + (index + 1) + '</div>'
                        //             + '<div>' + item + '</div>'
                        //             + '</a>';
                        //     });
                        //     return str;
                        // });
                        // hotListEle.show();
                    } else {
                        alert("生成失败");
                    }
                    $("#a3").attr("value", data);
                }
            })
        });
        /* 文档 */
        // $(document).click(function () {
        //     _this.els.pickerList.hide();
        //     _this.els.hotList.hide();
        // });
        /* 搜索按钮 */
    }
};