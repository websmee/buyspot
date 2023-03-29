import axios from "axios";

import * as actions from "Store/api";
import config from "Config/config";

const api =
    ({ dispatch }) =>
        (next) =>
            async (action) => {
                if (action.type !== actions.apiCallBegan.type) return next(action);

                const { url, method, data, onStart, onSuccess, onError, then } =
                    action.payload;

                if (onStart) dispatch({ type: onStart });

                next(action);

                try {
                    const response = await axios.request({
                        baseURL: config.API_BASE_URL,
                        url,
                        method,
                        data,
                    });
                    dispatch(actions.apiCallSuccess(response.data));
                    if (onSuccess) dispatch({ type: onSuccess, payload: response.data });
                    if (then) then(response.data);
                } catch (error) {
                    dispatch(actions.apiCallFailed(error.message));
                    if (onError) dispatch({ type: onError, payload: error.message });
                }
            };

export default api;