var crypto = require('crypto');

var name = 'braitsch';


GenSignature(1615365252 , 20032, "2345245", "rand", "zOOlS80Xz0BHng7l")


function GenSignature (unixSec , appid , echo , rand , token) {
    basic = unixSec + "|" + echo + "|" + appid + "|" + token
    console.log(basic)

    var hash = crypto.createHash('md5').update(Buffer.from(basic), 'binary').digest();
    console.log(hash); // 9b74c9897bac770ffc029102a200c5de



    h2 = crypto.createHash('md5').update(hash, 'binary').update(rand).digest('hex');


    console.log(h2);
}