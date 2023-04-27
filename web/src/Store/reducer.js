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
        unauthorized: false,
        errorMessage: "",
        currentPrices: {
            quote: "",
            pricesBySymbols: {
                "": 0,
            },
        },
        balance: {
            amount: 0,
            symbol: "",
        },
        currentSpotsIndex: 0,
        currentSpotsNext: 0,
        currentSpotsPrev: 0,
        currentSpotsTotal: 0,
        spot: {
            asset: {
                name: "",
                symbol: "",
                description: "",
                activeOrders: 0,
            },
            priceForecast: 0,
            chartsDataByQuotes: {
                "USDT": {
                    times: [],
                    prices: [],
                    forecast: [],
                    actual: [],
                    volumes: [],
                }
            },
            news: [
                // {
                //     sentiment: NEWS_ARTICLE_SENTIMENT.NEUTRAL,
                //     title: "",
                //     content: "",
                //     created: "2023-03-04T15:16:34.2960596+06:00",
                //     url: "https://...",
                //     imgURL: "https://...",
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
            isPorfitable: false,
        },
        orders: [
            // {
            //     id: "test123",
            //     fromAmount: 0,
            //     fromSymbol: "",
            //     toAmount: 0,
            //     toSymbol: "USDT",
            //     toAssetName: "",
            //     takeProfit: 0,
            //     stopLoss: 0,
            //     created: "2023-03-04T15:16:34.2960596+06:00",
            // }
        ],
    },

    reducers: {
        loginRequested: (state, action) => {
        },

        loginSuccess: (state, action) => {
            sessionStorage.setItem("jwt", action.payload)
            state.unauthorized = false;
        },

        loginRequestFailed: (state, action) => {
            handleFail(state, action);
        },

        currentPricesRequested: (state, action) => {
        },

        currentPricesReceived: (state, action) => {
            state.currentPrices = action.payload;
            state.errorMessage = "";
        },

        currentPricesRequestFailed: (state, action) => {
            handleFail(state, action);
        },

        currentBalanceRequested: (state, action) => {
        },

        currentBalanceReceived: (state, action) => {
            state.balance = action.payload;
            state.errorMessage = "";
        },

        currentBalanceRequestFailed: (state, action) => {
            handleFail(state, action);
        },

        currentSpotsDataRequested: (state, action) => {
        },

        currentSpotsDataReceived: (state, action) => {
            state.currentSpotsTotal = action.payload.currentSpotsTotal;

            if (action.payload.currentSpotsTotal == 0) {
                state.currentSpotsIndex = 0;
                state.currentSpotsNext = 0;
                state.currentSpotsPrev = 0;
            }

            if (action.payload.currentSpotsTotal > 0 && state.currentSpotsIndex == 0) {
                state.currentSpotsIndex = 1;
                state.currentSpotsNext = 2;
                state.currentSpotsPrev = action.payload.currentSpotsTotal;
            }

            if (action.payload.currentSpotsTotal < state.currentSpotsNext) {
                state.currentSpotsNext = 1;
                state.currentSpotsPrev = action.payload.currentSpotsTotal;
            };

            if (action.payload.spot) {
                state.spot = action.payload.spot;
            }

            stickymobile.hidePreloader();
            state.errorMessage = "";
        },

        currentSpotsDataRequestFailed: (state, action) => {
            state.currentSpotsTotal = 0;
            state.currentSpotsIndex = 0;
            state.currentSpotsNext = 0;
            state.currentSpotsPrev = 0;
            handleFail(state, action);
        },

        spotRequested: (state, action) => {
            stickymobile.showPreloader();
        },

        spotReceived: (state, action) => {
            state.currentSpotsIndex = action.payload.index;
            state.currentSpotsNext = action.payload.index < state.currentSpotsTotal ? action.payload.index + 1 : 1;
            state.currentSpotsPrev = action.payload.index > 1 ? action.payload.index - 1 : state.currentSpotsTotal;
            state.spot = action.payload;
            stickymobile.hidePreloader();
            state.errorMessage = "";
        },

        spotRequestFailed: (state, action) => {
            handleFail(state, action);
        },

        buySpotRequested: (state, action) => {
        },

        buySpotRequestSucceded: (state, action) => {
            state.spot.activeOrders++;
            state.balance = action.payload.updatedBalance;
            state.errorMessage = "";
        },

        buySpotRequestFailed: (state, action) => {
            handleFail(state, action);
        },

        sellOrderRequested: (state, action) => {
        },

        sellOrderRequestSucceded: (state, action) => {
            state.balance = action.payload.updatedBalance;
            state.orders = state.orders.filter(order => order.id != action.payload.orderID)
            state.errorMessage = "";
        },

        sellOrderRequestFailed: (state, action) => {
            handleFail(state, action);
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
            handleFail(state, action);
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
            order.fromSymbol,
            order.toAmount,
            order.toSymbol,
            state.currentPrices.pricesBySymbols,
        );
        order.amountInBalanceSymbol = converter.convert(
            order.toAmount,
            order.toSymbol,
            state.balance.symbol,
            state.currentPrices.pricesBySymbols,
        );
    });
};

const handleFail = (state, action) => {
    if (action.payload.status == 401) state.unauthorized = true;
    state.errorMessage = action.payload.message;
    stickymobile.hidePreloader();
}

export default slice.reducer;

const {
    currentPricesRequested, currentPricesReceived, currentPricesRequestFailed,
    currentBalanceRequested, currentBalanceReceived, currentBalanceRequestFailed,
    currentSpotsDataRequested, currentSpotsDataReceived, currentSpotsDataRequestFailed,
    spotRequested, spotReceived, spotRequestFailed,
    buySpotRequested, buySpotRequestSucceded, buySpotRequestFailed,
    sellOrderRequested, sellOrderRequestSucceded, sellOrderRequestFailed,
    ordersRequested, ordersReceived, ordersRequestFailed,
    loginRequested, loginSuccess, loginRequestFailed,
    clearErrorMessageRequested,
    updateOrdersDataRequested,
} = slice.actions;

export const login = (email, password) => (dispatch) => {
    return dispatch(
        apiCallBegan({
            url: "/api/v1/login",
            method: "post",
            data: { email, password },
            onStart: loginRequested.type,
            onSuccess: loginSuccess.type,
            onError: loginRequestFailed.type,
        })
    );
};

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

export const buySpot = (amount, symbol, takeProfit, stopLoss) => (dispatch) => {
    return dispatch(
        apiCallBegan({
            url: "/api/v1/spots/buy",
            method: "post",
            data: { amount, symbol, takeProfit, stopLoss },
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

export const sellOrder = (orderID) => (dispatch) => {
    return dispatch(
        apiCallBegan({
            url: "/api/v1/orders/" + orderID + "/sell",
            method: "post",
            onStart: sellOrderRequested.type,
            onSuccess: sellOrderRequestSucceded.type,
            onError: sellOrderRequestFailed.type,
        })
    );
};

export const clearErrorMessage = () => (dispatch) => {
    return dispatch({ type: clearErrorMessageRequested.type });
};

export const updateOrdersData = () => (dispatch) => {
    return dispatch({ type: updateOrdersDataRequested.type });
};