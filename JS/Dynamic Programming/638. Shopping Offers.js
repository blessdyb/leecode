/**
 * @param {number[]} price
 * @param {number[][]} special
 * @param {number[]} needs
 * @return {number}
 */
var shoppingOffers = function(price, special, needs) {
    const sum = needs.reduce((acc, item) => acc + item, 0);
    if(sum === 0) {
        return 0; 
    }
    let minPrice = Infinity;
    for(let i = 0; i < special.length; i++) {
        let isSpecialOfferValid = true;
        const newNeeds = [];
        for(let j = 0; j < needs.length; j++) {
            if (needs[j] < special[i][j]) {
                isSpecialOfferValid = false;
            } else {            
                newNeeds.push(needs[j] - special[i][j]);
            }
        }
        if(isSpecialOfferValid) {
            minPrice = Math.min(minPrice, special[i][needs.length] + shoppingOffers(price, special, newNeeds));
        }
    }
    minPrice = Math.min(minPrice, needs.reduce((acc, item, index) => acc + item * price[index], 0));
    return minPrice;
};
