$(function () {
    $(".nav3").click(function () {
        $.ajax({
            type: 'get',
            url: '/urlinfo',
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
    alert("sendclick")
    var method = $("#method").val();
    var url = $("#url").val();

    $.ajax({
        type: 'get',
        url: '/receive',
        data: {
            'url': url,
            'method': method,
            // 'inputbrandname': 'brand',
            // 'brandlistdiv': 'brandList',
            // 'brid': 'brid'
        },
        success: function (data, textStatus) {
            if (textStatus == 'success') {
                var data_string = JSON.stringify(data);
                var body = jQuery.parseJSON(data_string);
                var header_arr = jQuery.parseJSON(body.header);
                var options = {
                    collapsed: false,
                    withQuotes: true
                };
                $('#json-renderer').jsonViewer(jQuery.parseJSON(data.body), options);
                // $("#pre").html(body.body);
                // console.log(body_arrr);
                console.log(header_arr);
                $("#response-content").empty("");
                $.each(header_arr, function (i, val) {
                    var header_content =
                        '<div class="row" style="margin-left: 15px;">' +
                        '<div class="col-md-4"> <b>' + i + '</b></div>' +
                        '<div class="col-md-8"> ' + val + '</div>' +
                        '</div>';
                    $("#response-content").append(header_content);

                    console.log(i);
                    console.log(val);
                });


            }
            else {
            }
        }
    });


}