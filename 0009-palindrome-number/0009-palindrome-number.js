/**
 * @param {number} x
 * @return {boolean}
 */
var isPalindrome = function(x) {
    let str=String(x);
    let strr=str.split("").reverse().join("");
    if(str===strr){
        return true;
    }else{ return false; }
};