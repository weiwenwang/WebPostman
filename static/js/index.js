$(function () {
    $(".nav1").click(function () {
        $(this).children("span").toggleClass("glyphicon glyphicon-triangle-bottom");
        $(this).children("span").toggleClass("glyphicon glyphicon-triangle-right");
        var a = $(this).siblings("ul");

        a.slideToggle(200);
        console.log(a)
    })
    $(".nav2").click(function () {
        $(this).children("span").toggleClass("glyphicon glyphicon-triangle-bottom");
        $(this).children("span").toggleClass("glyphicon glyphicon-triangle-right");

        var a = $(this).siblings("ul");

        a.slideToggle(200);
        console.log(a)
    })
});

function sendSearch() {
    $("#nav-search-content").css('display', 'block');
    $("#nav-performance-content").css('display', 'none');
}

function sendPreformance() {
    $.ajax({
        type: 'get',
        url: '/perform',
        data: {},
        success: function (data, textStatus) {
            if (textStatus == 'success') {
                // 基于准备好的dom，初始化echarts实例
                var myChart = echarts.init(document.getElementById('main'));

                // 指定图表的配置项和数据
                var option = {
                    title: {
                        text: 'ECharts 入门示例'
                    },
                    tooltip: {},
                    legend: {
                        data: ['销量']
                    },
                    xAxis: {
                        data: ["衬衫", "羊毛衫", "雪纺衫", "裤子", "高跟鞋", "袜子"]
                    },
                    yAxis: {},
                    series: [{
                        name: '销量',
                        type: 'bar',
                        data: data.VALUE,
                    }]
                };

                // 使用刚指定的配置项和数据显示图表。
                myChart.setOption(option);
                console.log(data.VALUE);
            } else {
            }
        }
    })


    $("#nav-search-content").css('display', 'none');
    $("#nav-performance-content").css('display', 'block');
}

function sendClick() {
    var select_host = $("#select_host").val();
    var key = $("#key").val();

    $.ajax({
        type: 'get',
        url: '/receive',
        data: {
            'key': key,
            'select_host': select_host,
        },
        success: function (data, textStatus) {
            if (textStatus == 'success') {
                var data_string = JSON.stringify(data);
                var body = jQuery.parseJSON(data_string);
                console.log(body.value);
                $("#json-renderer").empty();

                var ttl = body.TTL + "s";
                if (body.TTL == -1) {
                    ttl = 'unset ttl(-1)'
                }

                var type = 'string';
                switch (body.TYPE) {
                    case 1:
                        type = 'string';
                        break;
                    case 2:
                        type = 'list';
                        break;
                    case 3:
                        type = 'hash';
                        break;
                    case 4:
                        type = 'set';
                        break;
                    case 5:
                        type = 'zset';
                        break;
                    default:
                        type = 'string';
                }

                var header_content =
                    '<div class="row">' +
                    '<div class="col-md-8"> TTL: ' + ttl + ' </div>' +
                    '<div class="col-md-8"> TYPE: ' + type + '</div>' +
                    '<div class="col-md-8"> VALUE: <br>' + body.VALUE + '</div>' +
                    '</div>';
                $("#json-renderer").append(header_content);
            } else {
            }
        }
    });
}
