export default {
    convert: (amount, fromSymbol, toSymbol, pricesBySymbols) => {
        if (!pricesBySymbols[fromSymbol] || !pricesBySymbols[toSymbol]) return 0;

        const paidAmount = amount * pricesBySymbols[fromSymbol];
        const convertedAmount = paidAmount * pricesBySymbols[toSymbol];

        return Math.round(convertedAmount * 100) / 100;
    },

    calculatePNL: (fromAmount, fromSymbol, toAmount, toSymbol, pricesBySymbols) => {
        if (!pricesBySymbols[fromSymbol] || !pricesBySymbols[toSymbol]) return 0;

        const paidAmount = fromAmount * pricesBySymbols[fromSymbol];
        const currentAmount = toAmount * pricesBySymbols[toSymbol];
        const diff = currentAmount - paidAmount;

        return Math.round(diff / currentAmount * 100 * 100) / 100;
    },
}