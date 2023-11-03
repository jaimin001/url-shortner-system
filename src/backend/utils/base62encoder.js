//utils/base62encoder.js

const alphabet = '0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz';

function generateShortUrl() {
    const randomNumber = Math.floor(Math.random() * 1000000);
    return encode(randomNumber);
}


function encode(number) {
    let base62 = '';
    do {
        base62 = alphabet[number % 62] + base62;
        number = Math.floor(number / 62);
    } while (number > 0)

    return base62;
}

function decodeShortUrl(string) {
    let number = 0;
    for (let i = 0; i < string.length; i++) {
        const char = string[i];
        number = number * 62 + alphabet.indexOf(char);
    }

    return number;
}

module.exports = { 
    generateShortUrl, decodeShortUrl
};