showEmojis.click(function (event) {
    event.stopPropagation();
    emojiContainer.toggle();
    let buttonPosition = showEmojis.offset();
    emojiContainer.css({
        top: buttonPosition.top - emojiContainer.height() - 10,
        left: buttonPosition.left
    });
});

$(document).click(function (event) {
    if (!emojiContainer.is(event.target) && !showEmojis.is(event.target)) {
        emojiContainer.hide();
    }
});

emojiContainer.on('click', '.emoji', function () {
    let emoji = $(this).attr('alt');
    inputFrame.val(inputFrame.val() + emoji);
    inputFrame.focus();
});

function emojiShowFunc(str) {
    return str.replace(/\[([^\]]+)\]/g, (match, content) => {
        if (content === "脱单doge") {
            return "<img src='/static/img/emojis/1.png' style='margin-bottom: 3px' width='20px'>";
        } else if (content === "微笑") {
            return "<img src='/static/img/emojis/2.png' style='margin-bottom: 3px' width='20px'>";
        } else if (content === "口罩") {
            return "<img src='/static/img/emojis/3.png' style='margin-bottom: 3px' width='20px'>";
        } else if (content === "doge") {
            return "<img src='/static/img/emojis/4.png' style='margin-bottom: 3px' width='20px'>";
        } else if (content === "妙啊") {
            return "<img src='/static/img/emojis/5.png' style='margin-bottom: 3px' width='20px'>";
        } else if (content === "OK") {
            return "<img src='/static/img/emojis/6.png' style='margin-bottom: 3px' width='20px'>";
        } else if (content === "星星眼") {
            return "<img src='/static/img/emojis/7.png' style='margin-bottom: 3px' width='20px'>";
        } else if (content === "辣眼睛") {
            return "<img src='/static/img/emojis/8.png' style='margin-bottom: 3px' width='20px'>";
        } else if (content === "吃瓜") {
            return "<img src='/static/img/emojis/9.png' style='margin-bottom: 3px' width='20px'>";
        } else if (content === "滑稽") {
            return "<img src='/static/img/emojis/10.png' style='margin-bottom: 3px' width='20px'>";
        } else if (content === "呲牙") {
            return "<img src='/static/img/emojis/11.png' style='margin-bottom: 3px' width='20px'>";
        } else if (content === "打call") {
            return "<img src='/static/img/emojis/12.png' style='margin-bottom: 3px' width='20px'>";
        } else if (content === "歪嘴") {
            return "<img src='/static/img/emojis/13.png' style='margin-bottom: 3px' width='20px'>";
        } else if (content === "调皮") {
            return "<img src='/static/img/emojis/14.png' style='margin-bottom: 3px' width='20px'>";
        } else if (content === "嗑瓜子") {
            return "<img src='/static/img/emojis/15.png' style='margin-bottom: 3px' width='20px'>";
        } else if (content === "哭笑") {
            return "<img src='/static/img/emojis/16.png' style='margin-bottom: 3px' width='20px'>";
        } else if (content === "脸红") {
            return "<img src='/static/img/emojis/17.png' style='margin-bottom: 3px' width='20px'>";
        } else if (content === "给心心") {
            return "<img src='/static/img/emojis/18.png' style='margin-bottom: 3px' width='20px'>";
        } else if (content === "嘟嘟") {
            return "<img src='/static/img/emojis/19.png' style='margin-bottom: 3px' width='20px'>";
        } else if (content === "哦呼") {
            return "<img src='/static/img/emojis/20.png' style='margin-bottom: 3px' width='20px'>";
        } else if (content === "喜欢") {
            return "<img src='/static/img/emojis/21.png' style='margin-bottom: 3px' width='20px'>";
        } else if (content === "酸了") {
            return "<img src='/static/img/emojis/22.png' style='margin-bottom: 3px' width='20px'>";
        } else if (content === "嫌弃") {
            return "<img src='/static/img/emojis/23.png' style='margin-bottom: 3px' width='20px'>";
        } else if (content === "害羞") {
            return "<img src='/static/img/emojis/24.png' style='margin-bottom: 3px' width='20px'>";
        } else if (content === "大哭") {
            return "<img src='/static/img/emojis/25.png' style='margin-bottom: 3px' width='20px'>";
        } else if (content === "疑惑") {
            return "<img src='/static/img/emojis/26.png' style='margin-bottom: 3px' width='20px'>";
        } else if (content === "喜极而泣") {
            return "<img src='/static/img/emojis/27.png' style='margin-bottom: 3px' width='20px'>";
        } else if (content === "笑") {
            return "<img src='/static/img/emojis/28.png' style='margin-bottom: 3px' width='20px'>";
        } else if (content === "偷笑") {
            return "<img src='/static/img/emojis/29.png' style='margin-bottom: 3px' width='20px'>";
        } else if (content === "惊讶") {
            return "<img src='/static/img/emojis/30.png' style='margin-bottom: 3px' width='20px'>";
        } else if (content === "捂脸") {
            return "<img src='/static/img/emojis/31.png' style='margin-bottom: 3px' width='20px'>";
        } else if (content === "阴险") {
            return "<img src='/static/img/emojis/32.png' style='margin-bottom: 3px' width='20px'>";
        } else if (content === "呆") {
            return "<img src='/static/img/emojis/33.png' style='margin-bottom: 3px' width='20px'>";
        } else if (content === "抠鼻") {
            return "<img src='/static/img/emojis/34.png' style='margin-bottom: 3px' width='20px'>";
        } else if (content === "大笑") {
            return "<img src='/static/img/emojis/35.png' style='margin-bottom: 3px' width='20px'>";
        } else if (content === "惊喜") {
            return "<img src='/static/img/emojis/36.png' style='margin-bottom: 3px' width='20px'>";
        } else if (content === "无语") {
            return "<img src='/static/img/emojis/37.png' style='margin-bottom: 3px' width='20px'>";
        } else if (content === "点赞") {
            return "<img src='/static/img/emojis/38.png' style='margin-bottom: 3px' width='20px'>";
        } else if (content === "鼓掌") {
            return "<img src='/static/img/emojis/39.png' style='margin-bottom: 3px' width='20px'>";
        } else if (content === "尴尬") {
            return "<img src='/static/img/emojis/40.png' style='margin-bottom: 3px' width='20px'>";
        } else if (content === "灵魂出窍") {
            return "<img src='/static/img/emojis/41.png' style='margin-bottom: 3px' width='20px'>";
        } else if (content === "委屈") {
            return "<img src='/static/img/emojis/42.png' style='margin-bottom: 3px' width='20px'>";
        } else if (content === "傲娇") {
            return "<img src='/static/img/emojis/43.png' style='margin-bottom: 3px' width='20px'>";
        } else if (content === "疼") {
            return "<img src='/static/img/emojis/44.png' style='margin-bottom: 3px' width='20px'>";
        } else if (content === "冷") {
            return "<img src='/static/img/emojis/45.png' style='margin-bottom: 3px' width='20px'>";
        } else if (content === "热") {
            return "<img src='/static/img/emojis/46.png' style='margin-bottom: 3px' width='20px'>";
        } else if (content === "生病") {
            return "<img src='/static/img/emojis/47.png' style='margin-bottom: 3px' width='20px'>";
        } else if (content === "捂眼") {
            return "<img src='/static/img/emojis/48.png' style='margin-bottom: 3px' width='20px'>";
        } else if (content === "嘘声") {
            return "<img src='/static/img/emojis/49.png' style='margin-bottom: 3px' width='20px'>";
        } else if (content === "思考") {
            return "<img src='/static/img/emojis/50.png' style='margin-bottom: 3px' width='20px'>";
        } else if (content === "翻白眼") {
            return "<img src='/static/img/emojis/51.png' style='margin-bottom: 3px' width='20px'>";
        } else if (content === "哈欠") {
            return "<img src='/static/img/emojis/52.png' style='margin-bottom: 3px' width='20px'>";
        } else if (content === "奋斗") {
            return "<img src='/static/img/emojis/53.png' style='margin-bottom: 3px' width='20px'>";
        } else if (content === "墨镜") {
            return "<img src='/static/img/emojis/54.png' style='margin-bottom: 3px' width='20px'>";
        } else if (content === "难过") {
            return "<img src='/static/img/emojis/55.png' style='margin-bottom: 3px' width='20px'>";
        } else if (content === "撇嘴") {
            return "<img src='/static/img/emojis/56.png' style='margin-bottom: 3px' width='20px'>";
        } else if (content === "抓狂") {
            return "<img src='/static/img/emojis/57.png' style='margin-bottom: 3px' width='20px'>";
        } else if (content === "生气") {
            return "<img src='/static/img/emojis/58.png' style='margin-bottom: 3px' width='20px'>";
        } else if (content === "爱心") {
            return "<img src='/static/img/emojis/59.png' style='margin-bottom: 3px' width='20px'>";
        } else if (content === "胜利") {
            return "<img src='/static/img/emojis/60.png' style='margin-bottom: 3px' width='20px'>";
        } else if (content === "加油") {
            return "<img src='/static/img/emojis/61.png' style='margin-bottom: 3px' width='20px'>";
        } else if (content === "抱拳") {
            return "<img src='/static/img/emojis/62.png' style='margin-bottom: 3px' width='20px'>";
        } else if (content === "保佑") {
            return "<img src='/static/img/emojis/63.png' style='margin-bottom: 3px' width='20px'>";
        } else if (content === "厉害") {
            return "<img src='/static/img/emojis/64.png' style='margin-bottom: 3px' width='20px'>";
        } else if (content === "拥抱") {
            return "<img src='/static/img/emojis/65.png' style='margin-bottom: 3px' width='20px'>";
        } else if (content === "派蒙") {
            return "<img src='/static/img/emojis/66.png' style='margin-bottom: 3px' width='20px'>";
        } else {
            return match;
        }
    })
}