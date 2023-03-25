export default {
    convert: (amount, fromTicker, toTicker, pricesByTickers) => {
        if (!pricesByTickers[fromTicker] || !pricesByTickers[toTicker]) return 0;

        const paidAmount = amount * pricesByTickers[fromTicker];
        const convertedAmount = paidAmount * pricesByTickers[toTicker];

        return Math.round(convertedAmount * 2) / 2;
    },

    calculatePNL: (fromAmount, fromTicker, toAmount, toTicker, pricesByTickers) => {
        if (!pricesByTickers[fromTicker] || !pricesByTickers[toTicker]) return 0;

        const paidAmount = fromAmount * pricesByTickers[fromTicker];
        const currentAmount = toAmount * pricesByTickers[toTicker];
        const diff = currentAmount - paidAmount;

        return Math.round(diff / currentAmount * 100 * 2) / 2;
    },
}