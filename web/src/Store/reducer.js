import { createSlice } from "@reduxjs/toolkit";

import { apiCallBegan } from "Store/api";
import converter from "Utils/converter";

import stickymobile from "Utils/stickymobile";

export const NEWS_ARTICLE_SENTIMENT = {
    NEUTRAL: "NEUTRAL",
    POSITIVE: "POSITIVE",
    NEGATIVE: "NEGATIVE",
};

const slice = createSlice({
    name: "reducer",
    initialState: {
        loading: false,
        errorMessage: "",
        currentPrices: {
            inTicker: "",
            pricesByTickers: {
                "": 0,
            },
        },
        balance: {
            amount: 0,
            ticker: "",
        },
        currentSpotsIndex: 1,
        currentSpotsTotal: 0,
        spot: {
            asset: {
                name: "",
                ticker: "",
                description: "",
                activeOrders: 0,
            },
            priceForecast: 0,
            chartsData: {
                times: [],
                prices: [],
                forecast: [],
                volumes: [],
            },
            news: [
                // {
                //     sentiment: NEWS_ARTICLE_SENTIMENT.NEUTRAL,
                //     title: "",
                //     content: "",
                //     created: "2023-03-04T15:16:34.2960596+06:00",
                //     views: 0,
                // },
            ],
            buyOrderSettings: {
                amount: 0,
                takeProfit: 0,
                takeProfitOptions: [
                    // {value: 0, text: ""},
                ],
                stopLoss: 0,
                stopLossOptions: [
                    // {value: 0, text: ""},
                ],
            },
        },
        orders: [
            // {
            //     id: "test123",
            //     fromAmount: 0,
            //     fromTicker: "",
            //     toAmount: 0,
            //     toTicker: "USDT",
            //     toAssetName: "",
            //     takeProfit: 0,
            //     stopLoss: 0,
            //     created: "2023-03-04T15:16:34.2960596+06:00",
            // }
        ],
    },

    reducers: {
        currentPricesRequested: (state, action) => {
        },

        currentPricesReceived: (state, action) => {
            state.currentPrices = action.payload;
            state.errorMessage = "";
        },

        currentPricesRequestFailed: (state, action) => {
            state.errorMessage = action.payload;
        },

        currentBalanceRequested: (state, action) => {
        },

        currentBalanceReceived: (state, action) => {
            state.balance = action.payload;
            state.errorMessage = "";
        },

        currentBalanceRequestFailed: (state, action) => {
            state.errorMessage = action.payload;
        },

        currentSpotsDataRequested: (state, action) => {
        },

        currentSpotsDataReceived: (state, action) => {
            state.currentSpotsTotal = action.payload.currentSpotsTotal;
            state.currentSpotsIndex = action.payload.currentSpotsTotal > 0 ? 1 : 0;
            state.errorMessage = "";
        },

        currentSpotsDataRequestFailed: (state, action) => {
            state.currentSpotsTotal = 0;
            state.currentSpotsIndex = 0;
            state.errorMessage = action.payload;
        },

        spotRequested: (state, action) => {
            stickymobile.showPreloader();
        },

        spotReceived: (state, action) => {
            state.spot = action.payload;
            stickymobile.hidePreloader();
            state.errorMessage = "";
        },

        spotRequestFailed: (state, action) => {
            state.errorMessage = action.payload;
            stickymobile.hidePreloader();
        },

        buySpotRequested: (state, action) => {
        },

        buySpotRequestSucceded: (state, action) => {
            state.spot.activeOrders++;
            state.balance = action.payload.updatedBalance;
            state.errorMessage = "";
        },

        buySpotRequestFailed: (state, action) => {
            state.errorMessage = action.payload;
        },

        ordersRequested: (state, action) => {
            stickymobile.showPreloader();
        },

        ordersReceived: (state, action) => {
            state.orders = action.payload;
            updateOrders(state);
            stickymobile.hidePreloader();
            state.errorMessage = "";
        },

        ordersRequestFailed: (state, action) => {
            state.errorMessage = action.payload;
            stickymobile.hidePreloader();
        },

        clearErrorMessageRequested: (state, action) => {
            state.errorMessage = "";
        },

        updateOrdersDataRequested: (state, action) => {
            updateOrders(state);
        },
    },
});

const updateOrders = state => {
    state.orders.forEach(order => {
        order.pnl = converter.calculatePNL(
            order.fromAmount,
            order.fromTicker,
            order.toAmount,
            order.toTicker,
            state.currentPrices.pricesByTickers,
        );
        order.amountInBalanceTicker = converter.convert(
            order.toAmount,
            order.toTicker,
            state.balance.ticker,
            state.currentPrices.pricesByTickers,
        );
    });
};

export default slice.reducer;

const {
    currentPricesRequested, currentPricesReceived, currentPricesRequestFailed,
    currentBalanceRequested, currentBalanceReceived, currentBalanceRequestFailed,
    currentSpotsDataRequested, currentSpotsDataReceived, currentSpotsDataRequestFailed,
    spotRequested, spotReceived, spotRequestFailed,
    buySpotRequested, buySpotRequestSucceded, buySpotRequestFailed,
    ordersRequested, ordersReceived, ordersRequestFailed,
    clearErrorMessageRequested,
    updateOrdersDataRequested,
} = slice.actions;

export const getCurrentPrices = () => (dispatch) => {
    return dispatch(
        apiCallBegan({
            url: "/api/v1/prices/current",
            method: "get",
            onStart: currentPricesRequested.type,
            onSuccess: currentPricesReceived.type,
            onError: currentPricesRequestFailed.type,
        })
    );
};

export const getCurrentBalance = () => (dispatch) => {
    return dispatch(
        apiCallBegan({
            url: "/api/v1/balances/current",
            method: "get",
            onStart: currentBalanceRequested.type,
            onSuccess: currentBalanceReceived.type,
            onError: currentBalanceRequestFailed.type,
        })
    );
};

export const getCurrentSpotsData = () => (dispatch) => {
    return dispatch(
        apiCallBegan({
            url: "/api/v1/spots/data",
            method: "get",
            onStart: currentSpotsDataRequested.type,
            onSuccess: currentSpotsDataReceived.type,
            onError: currentSpotsDataRequestFailed.type,
        })
    );
};

export const getSpotByIndex = (index) => (dispatch) => {
    return dispatch(
        apiCallBegan({
            url: "/api/v1/spots/" + index,
            method: "get",
            onStart: spotRequested.type,
            onSuccess: spotReceived.type,
            onError: spotRequestFailed.type,
        })
    );
};

export const buySpot = (amount, ticker, takeProfit, stopLoss) => (dispatch) => {
    return dispatch(
        apiCallBegan({
            url: "/api/v1/spots/buy",
            method: "post",
            data: { amount, ticker, takeProfit, stopLoss },
            onStart: buySpotRequested.type,
            onSuccess: buySpotRequestSucceded.type,
            onError: buySpotRequestFailed.type,
        })
    );
};

export const getOrders = () => (dispatch) => {
    return dispatch(
        apiCallBegan({
            url: "/api/v1/orders",
            method: "get",
            onStart: ordersRequested.type,
            onSuccess: ordersReceived.type,
            onError: ordersRequestFailed.type,
        })
    );
};

export const clearErrorMessage = () => (dispatch) => {
    return dispatch({ type: clearErrorMessageRequested.type });
};

export const updateOrdersData = () => (dispatch) => {
    return dispatch({ type: updateOrdersDataRequested.type });
};