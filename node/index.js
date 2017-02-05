let cheerio = require('cheerio')
let fs = require('fs')
let eval = require('node-eval')

fs.readFile('../data.jsonp', 'utf8', function (err, text) {
    var res = eval(
            'function jsonp423(html){return html;};' + text
            ) ;

    date_line = []
    doc = []

    var moment = require('moment-timezone');
    var hoge = moment().tz("Asia/Tokyo");
    var year = hoge.year();
    var month = hoge.month() + 1;
    var date = hoge.date();
    var hour = hoge.hour();
    var second = hoge.second();

    date_line.push(year)
    date_line.push(month)
    date_line.push(date)



    let $ = cheerio.load(res)
    //$('.item4line1').find('dl').find('dd.detail')
    $('.item4line1').find('dl')
    .each(function(i, element){
        var ch0 = element.children;
        for(var val0 in ch0) {
            //DATAID
            var dataid = ch0[val0].parent.attribs['data-id'];
            var cv0 = ch0[val0];
            var ch1 = cv0.children;
            if (cv0.type == 'tag' && cv0.name == 'dd') {
                for (var var1 in ch1) {
                    var cv1 = ch1[var1];
                    var ch2 = cv1.children;
                    if (cv1.type == 'tag' && cv1.name == 'a') {
                        //DETAIL
                        var detail = ch2[0].data.trim();
                    }

                    if (cv1.type == 'tag' && cv1.name == 'div') {
                        for (var var2 in ch2) {
                            var cv2 = ch2[var2];
                            var ch3 = cv2.children;
                            if (cv2.type = 'tag' && cv2.name == 'div') {
                                //PRICE
                                var price = "¥" + ch3[1].children[0].data.trim();
                            }
                        }
                    }
                }
            }
        }
    var result = []
    result.push(...date_line);
    result.push(dataid);
    result.push(price);
    result.push(detail);
    result.push("\n");
    var joined_result =  result.join("\t")
    doc.push(joined_result)
    });
    var doc_result = doc.join("\n")


});

