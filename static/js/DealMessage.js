function showWorldRoom() {
    $('#msgPageUl li').each(function () {
        let id = $(this).attr('id');

        if (id === "liWorldRoom") {
            $(this).addClass("boxUsing");
        } else {
            $(this).removeClass("boxUsing");
        }
    });
    $.ajax({
        url: "/getWorldRoomMsg",
        method: "post",
        success: function (response) {
            $("#unreadNumWorldRoom").text("");
            let msgFrame = $("#msgFrame");
            msgFrame.empty();

            let msgs = response["msgWorldRoom"];
            let len = msgs.length;
            let imgs = response["imgWorldRoom"];
            let names = response["nameWorldRoom"];

            let originalStr = msgs[len - 1].CreatedAt;
            let resArr = originalStr.split("T");
            let finalIndex = resArr[1].length - 4;
            let finalArr = resArr[0] + " " + resArr[1].substring(0, finalIndex);
            let msgHtmlMsg = msgs[len - 1].Message;
            if (msgHtmlMsg.length > 5) {
                msgHtmlMsg = msgs[len - 1].Message.substring(0, 5) + "...";
            }
            $("#timeWorldRoom").text(finalArr);
            $("#msgWorldRoom").text(msgHtmlMsg);
            for (let i = 0; i < len; i++) {
                if (msgs[i].SenderAccount === acc) {
                    let originalStr = msgs[i].CreatedAt;
                    let resArr = originalStr.split("T");
                    let finalIndex = resArr[1].length - 4;
                    let finalStr = names[i] + " | " + resArr[0] + " " + resArr[1].substring(0, finalIndex);
                    let msgTemp = msgs[i].Message;
                    let msg = emojiShowFunc(msgTemp);
                    let msgRightHtml = `
                       <li class="d-flex flex-row justify-content-end">
                            <div>
                                <p class="small p-2 me-3 mb-1 text-black rounded-3" style="background-color: rgb(169, 234, 122)">
                                    ${msg}
                                </p>
                                <p class="small me-3 mb-3 rounded-3 text-muted">
                                    ${finalStr}
                                </p>
                            </div>
                            <img src="${imgs[i]}" style="width: 45px;height: 45px;border-radius: 8px;user-select: none">
                        </li>
            `;
                    msgFrame.append(msgRightHtml);
                } else {
                    let originalStr = msgs[i].CreatedAt;
                    let resArr = originalStr.split("T");
                    let finalIndex = resArr[1].length - 4;
                    let finalStr = names[i] + " | " + resArr[0] + " " + resArr[1].substring(0, finalIndex);
                    let msgTemp = msgs[i].Message;
                    let msg = emojiShowFunc(msgTemp);
                    let msgLeftHtml = `
                        <li class="d-flex flex-row justify-content-start">
                            <img src="${imgs[i]}" style="width: 45px;height: 45px;border-radius: 8px;user-select: none">
                            <div>
                                <p class="small p-2 ms-3 mb-1 text-black rounded-3" style="background-color: lightblue;">
                                    ${msg}
                                </p>
                                <p class="small ms-3 mb-3 rounded-3 text-muted float-end">
                                    ${finalStr}
                                </p>
                            </div>
                        </li>
            `;
                    msgFrame.append(msgLeftHtml);
                }
            }
            document.getElementById("editFrame").style.display = "block";
            msgFrame.scrollTop(msgFrame.prop("scrollHeight"));
            msgFrame.attr("data-remark", "WorldRoom");
            $("#liWorldRoom").addClass("boxAlreadyShow");
        }
    })
}

function showMsg(receiver) {
    let temp = $("#li" + receiver).is('*');
    if (!temp) {
        showPage('msgPage');
        msgBtn.addClass('using');
        $('#friendBtn, #infoBtn').removeClass('using');
        $.ajax({
            url: "/getChatMsg",
            method: "post",
            data: {
                "sender": acc,
                "receiver": receiver,
            },
            success: function (response) {
                let msgFrame = $("#msgFrame");
                msgFrame.empty();
                let msgs = response["msgs"];
                let len = msgs.length;
                let originalStr = msgs[len - 1].CreatedAt;
                let resArr = originalStr.split("T");
                let finalIndex = resArr[1].length - 4;
                let finalArr = resArr[0] + " " + resArr[1].substring(0, finalIndex);
                let msgHtmlMsg = msgs[len - 1].Message;
                if (msgHtmlMsg.length > 5) {
                    msgHtmlMsg = msgs[len - 1].Message.substring(0, 5) + "...";
                }
                let boxSpe = "li" + receiver;
                let msgSpe = "msg" + receiver;
                let timeSpe = "time" + receiver;
                let unreadSpe = "unreadNum" + receiver;
                let msgHtml = `
<li id="${boxSpe}" class="p-2 boxAlreadyShow" style="border-radius: 10px" onclick="showMsg2(${receiver})">
    <div class="d-flex justify-content-between">
        <div class="d-flex flex-row">
            <div>
                <img src="${response["receiverImg"]}" alt="avatar" class="d-flex align-self-center me-3" style="width: 50px;height: 50px;border-radius: 8px;user-select: none">
            </div>
            <div class="pt-1">
                <p class="fw-bold mb-0">${response["receiverUsername"]}</p>
                <p id="${msgSpe}" class="text-muted">${msgHtmlMsg}</p>
            </div>
        </div>
        <div class="pt-1">
            <p class="small text-muted mb-1" id="${timeSpe}">${finalArr}</p>
            <span id="${unreadSpe}" class="badge bg-danger rounded-pill float-end"></span>
        </div>
    </div>
</li>
            `;
                $("#msgPageUl").append(msgHtml);
                for (let i = 0; i < len; i++) {
                    if (msgs[i].SenderAccount === acc && msgs[i].ReceiverAccount === receiver) {
                        let originalStr = msgs[i].CreatedAt;
                        let resArr = originalStr.split("T");
                        let finalIndex = resArr[1].length - 4;
                        let finalArr = response["senderUsername"] + " | " + resArr[0] + " " + resArr[1].substring(0, finalIndex);
                        let msgTemp = msgs[i].Message;
                        let msg = emojiShowFunc(msgTemp);
                        let msgRightHtml = `
                       <li class="d-flex flex-row justify-content-end">
                            <div>
                                <p class="small p-2 me-3 mb-1 text-black rounded-3" style="background-color: rgb(169, 234, 122)">
                                    ${msg}
                                </p>
                                <p class="small me-3 mb-3 rounded-3 text-muted">
                                    ${finalArr}
                                </p>
                            </div>
                            <img src="${response["senderImg"]}" style="width: 45px;height: 45px;border-radius: 8px;user-select: none">
                        </li>
            `;
                        msgFrame.append(msgRightHtml);
                    } else if (msgs[i].SenderAccount === receiver && msgs[i].ReceiverAccount === acc) {
                        let originalStr = msgs[i].CreatedAt;
                        let resArr = originalStr.split("T");
                        let finalIndex = resArr[1].length - 4;
                        let finalArr = response["receiverUsername"] + " | " + resArr[0] + " " + resArr[1].substring(0, finalIndex);
                        let msgTemp = msgs[i].Message;
                        let msg = emojiShowFunc(msgTemp);
                        let msgLeftHtml = `
                        <li class="d-flex flex-row justify-content-start">
                            <img src="${response["receiverImg"]}" style="width: 45px;height: 45px;border-radius: 8px;user-select: none">
                            <div>
                                <p class="small p-2 ms-3 mb-1 text-black rounded-3" style="background-color: lightblue;">
                                    ${msg}
                                </p>
                                <p class="small ms-3 mb-3 rounded-3 text-muted float-end">
                                    ${finalArr}
                                </p>
                            </div>
                        </li>
            `;
                        msgFrame.append(msgLeftHtml);
                    }
                }
                document.getElementById("editFrame").style.display = "block";
                msgFrame.scrollTop(msgFrame.prop("scrollHeight"));
                msgFrame.attr("data-remark", receiver);
                $('#msgPageUl li').each(function () {
                    let id = $(this).attr('id');

                    if (id === "li" + receiver) {
                        $(this).addClass("boxUsing");
                    } else {
                        $(this).removeClass("boxUsing");
                    }
                });
            }
        });
    } else {
        showPage('msgPage');
        msgBtn.addClass('using');
        $('#friendBtn, #groupBtn, #infoBtn').removeClass('using');
    }
}

function showMsg2(receiver) {
    let receiverStr = String(receiver);
    $('#msgPageUl li').each(function () {
        let id = $(this).attr('id');

        if (id === "li" + receiverStr) {
            $(this).addClass("boxUsing");
        } else {
            $(this).removeClass("boxUsing");
        }
    });
    $.ajax({
        url: "/getChatMsg",
        method: "post",
        data: {
            "sender": acc,
            "receiver": receiverStr,
        },
        success: function (response) {
            $("#unreadNum"+receiverStr).text("");
            let msgFrame = $("#msgFrame");
            msgFrame.empty();
            let msgs = response["msgs"];
            let len = msgs.length;
            let originalStr = msgs[len - 1].CreatedAt;
            let resArr = originalStr.split("T");
            let finalIndex = resArr[1].length - 4;
            let finalArr = resArr[0] + " " + resArr[1].substring(0, finalIndex);
            let msgHtmlMsg = msgs[len - 1].Message;
            if (msgHtmlMsg.length > 5) {
                msgHtmlMsg = msgs[len - 1].Message.substring(0, 5) + "...";
            }
            $("#time" + receiverStr).text(finalArr);
            $("#msg" + receiverStr).text(msgHtmlMsg);
            for (let i = 0; i < len; i++) {
                if (msgs[i].SenderAccount === acc && msgs[i].ReceiverAccount === receiverStr) {
                    let originalStr = msgs[i].CreatedAt;
                    let resArr = originalStr.split("T");
                    let finalIndex = resArr[1].length - 4;
                    let finalArr = response["senderUsername"] + " | " + resArr[0] + " " + resArr[1].substring(0, finalIndex);
                    let msgTemp = msgs[i].Message;
                    let msg = emojiShowFunc(msgTemp);
                    let msgRightHtml = `
                       <li class="d-flex flex-row justify-content-end">
                            <div>
                                <p class="small p-2 me-3 mb-1 text-black rounded-3" style="background-color: rgb(169, 234, 122)">
                                    ${msg}
                                </p>
                                <p class="small me-3 mb-3 rounded-3 text-muted">
                                    ${finalArr}
                                </p>
                            </div>
                            <img src="${response["senderImg"]}" style="width: 45px;height: 45px;border-radius: 8px;user-select: none">
                        </li>
            `;
                    msgFrame.append(msgRightHtml);
                } else if (msgs[i].SenderAccount === receiverStr && msgs[i].ReceiverAccount === acc) {
                    let originalStr = msgs[i].CreatedAt;
                    let resArr = originalStr.split("T");
                    let finalIndex = resArr[1].length - 4;
                    let finalArr = response["receiverUsername"] + " | " + resArr[0] + " " + resArr[1].substring(0, finalIndex);
                    let msgTemp = msgs[i].Message;
                    let msg = emojiShowFunc(msgTemp);
                    let msgLeftHtml = `
                        <li class="d-flex flex-row justify-content-start">
                            <img src="${response["receiverImg"]}" style="width: 45px;height: 45px;border-radius: 8px;user-select: none">
                            <div>
                                <p class="small p-2 ms-3 mb-1 text-black rounded-3" style="background-color: lightblue;">
                                    ${msg}
                                </p>
                                <p class="small ms-3 mb-3 rounded-3 text-muted float-end">
                                    ${finalArr}
                                </p>
                            </div>
                        </li>
            `;
                    msgFrame.append(msgLeftHtml);
                }
            }
            document.getElementById("editFrame").style.display = "block";
            msgFrame.scrollTop(msgFrame.prop("scrollHeight"));
            msgFrame.attr("data-remark", receiverStr);
        }
    });
}