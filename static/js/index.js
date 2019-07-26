$(function () {
    $(".nav3").click(function () {
        $.ajax({
            type: 'get',
            url: '/Urlinfo',
            data: {
                'urlId': 3,
            },
            success: function (data, textStatus) {
                if (textStatus == 'success') {
                    var data_string = JSON.stringify(data);
                    var body = jQuery.parseJSON(data_string);
                    $("#method").val(body.method);
                    $("#url").val(body.url);

                    console.log(body)
                }
            }
        });
    })
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