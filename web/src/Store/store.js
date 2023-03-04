import { configureStore, getDefaultMiddleware } from "@reduxjs/toolkit";
import reducer from "Store/reducer";
import api from "Middleware/api";

export default function store() {
    return configureStore({
        reducer,
        middleware: [...getDefaultMiddleware(), api],
    });
}