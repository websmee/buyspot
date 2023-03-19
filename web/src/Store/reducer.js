import { createSlice } from "@reduxjs/toolkit";

import { apiCallBegan } from "Store/api";

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
        balance: {
            amount: 0,
            ticker: "",
        },
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
            currentSpotsIndex: 0,
            currentSpotsTotal: 0,
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
    },

    reducers: {
        currentBalanceRequested: (state, action) => {
        },

        currentBalanceReceived: (state, action) => {
            state.balance = action.payload;
        },

        currentBalanceRequestFailed: (state, action) => {
            state.errorMessage = action.payload;
        },

        nextSpotRequested: (state, action) => {
            stickymobile.showPreloader();
        },

        nextSpotReceived: (state, action) => {
            state.spot = action.payload;
            stickymobile.hidePreloader();
        },

        nextSpotRequestFailed: (state, action) => {
            state.errorMessage = action.payload;
            stickymobile.hidePreloader();
        },

        buySpotRequested: (state, action) => {
        },

        buySpotRequestSucceded: (state, action) => {
            state.spot.asset.activeOrders++;
            state.balance = action.payload.updatedBalance;
        },

        buySpotRequestFailed: (state, action) => {
            state.errorMessage = action.payload;
        },

        clearErrorMessageRequested: (state, action) => {
            state.errorMessage = "";
        },
    },
});

export default slice.reducer;

const {
    currentBalanceRequested, currentBalanceReceived, currentBalanceRequestFailed,
    nextSpotRequested, nextSpotReceived, nextSpotRequestFailed,
    buySpotRequested, buySpotRequestSucceded, buySpotRequestFailed,
    clearErrorMessageRequested,
} = slice.actions;

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

export const getNextSpot = () => (dispatch) => {
    return dispatch(
        apiCallBegan({
            url: "/api/v1/spots/next",
            method: "get",
            onStart: nextSpotRequested.type,
            onSuccess: nextSpotReceived.type,
            onError: nextSpotRequestFailed.type,
        })
    );
};

export const buySpot = (assetTicker, balanceTicker, amount, takeProfit, stopLoss) => (dispatch) => {
    return dispatch(
        apiCallBegan({
            url: "/api/v1/spots/buy",
            method: "post",
            data: { assetTicker, balanceTicker, amount, takeProfit, stopLoss },
            onStart: buySpotRequested.type,
            onSuccess: buySpotRequestSucceded.type,
            onError: buySpotRequestFailed.type,
        })
    );
};

export const clearErrorMessage = () => (dispatch) => {
    return dispatch({type: clearErrorMessageRequested.type});
};