import { useDispatch } from "react-redux";

import ReactTimeAgo from "react-time-ago"

import { sellOrder } from "Store/reducer";

function OrderSellModal(props) {
    const dispatch = useDispatch();

    return (
        <div id={props.id} className="menu menu-box-modal menu-box-detached">
            <div className="menu-title">
                <h1>
                    {props.order.toAssetName}
                    <span className="opacity-30 font-200" style={{marginLeft: "5px"}}>{props.order.toSymbol}</span>
                </h1>
                <p><ReactTimeAgo date={Date.parse(props.order.created)} locale="en-US" /></p>
                <a className="close-menu"><i className="fa fa-times"></i></a>
            </div>
            <div className="divider divider-margins mb-1 mt-3"></div>
            <div className="content px-1">
                <p>
                    Current value: {props.order.amountSymbol} <strong className={props.order.pnl < 0 ? "color-red-light" : "color-sunny-light"}>{props.order.amountInBalanceSymbol}</strong>
                    <br />
                    PNL: <strong className={props.order.pnl < 0 ? "color-red-light" : "color-sunny-light"}>{props.order.pnl > 0 && "+"}{props.order.pnl}%</strong>
                    <br />
                    Will be sold automatically at {props.order.takeProfit > 0 && "+"}{props.order.takeProfit}% or {props.order.stopLoss}%
                </p>
                <a
                    className="close-menu btn btn-full btn-m bg-theme color-theme gradient-sunny rounded-sm text-uppercase font-800 mb-3"
                    onClick={() => {
                        dispatch(sellOrder(props.order.id));
                    }}
                >Sell Now</a>
            </div>
        </div>
    )
}

export default OrderSellModal