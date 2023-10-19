$("#settings").click(() => {
    $("#settingsModal").modal('show');
})

$("#updateName").click(() => {
    $.ajax({
        url: "/updateName",
        method: "post",
        data: {
            "account": acc,
            "username": $("#newName").val(),
        },
        success: function (response) {
            if (response["status"] === "1") {
                $("#newName").val("");
                $.message({
                    message: '修改成功！',
                    type: 'success'
                });
                location.reload(true);
            } else {
                $.message({
                    message: '该昵称已经被他人使用了，请换一个吧。',
                    type: 'error'
                });
            }
        }
    });
})

$("#updateImg").click(() => {
    let file = $("#file")[0].files[0];
    let form_data = new FormData();
    form_data.append("file", file);
    form_data.append("account", acc);
    $.ajax({
        url: "/updateImg",
        type: "POST",
        data: form_data,
        dataType: "json",
        contentType: false,
        processData: false,
        success: function (data) {
            $.message({
                message: '修改成功！',
                type: 'success'
            });
            location.reload(true);
        },
    });
})