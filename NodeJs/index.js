function palindrome(str){
    var re = /[^A-Za-z0-9]/g;
    str = str.toLowerCase().replace(re, '');
    var len = str.length;

    // Kamu hanya perlu mengecek setengah dari isi stringnya saja 
    // tidak perlu semuanya
    for(var i = 0; i < len/2; i++){
        if(str[i] !== str[len - 1 - i]){
            return false
        }
    }

    return true;
}

var str = "kupu-kupu";
console.log("Is " + str + " palindrome ? " + palindrome(str));
var str = "kasur ini rusak";
console.log("Is " + str + " palindrome ? " + palindrome(str));