/**
 * @param {number[]} nums
 * @param {number} original
 * @return {number}
 */
var findFinalValue = function(nums, original) {
    let x=original;
    let set=new Set(nums);
    while(set.has(x)){
        x=x*2
    }
    return x;
};