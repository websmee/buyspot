import { useEffect, useState } from 'react';
import { useDispatch } from 'react-redux';

import { buySpot } from 'Store/reducer';

function SpotBuyModal(props) {
    const dispatch = useDispatch();

    const [orderAmount, setOrderAmount] = useState(0);
    const [orderTakeProfit, setOrderTakeProfit] = useState(0);
    const [orderStopLoss, setOrderStopLoss] = useState(0);

    useEffect(() => {
        setOrderAmount(props.amount);
        setOrderTakeProfit(props.takeProfit);
        setOrderStopLoss(props.stopLoss);
    }, [props.amount, props.takeProfit, props.stopLoss])

    return (
        <div id={props.id} className="menu menu-box-modal menu-box-detached">
            <div className="menu-title"><h1>Buy {props.assetName}</h1>
                <p className="color-highlight">Convert {props.balanceSymbol} to {props.assetSymbol}</p><a className="close-menu"><i className="fa fa-times"></i></a>
            </div>
            <div className="divider divider-margins mb-1 mt-3"></div>
            <div className="content px-1">
                <div className="input-style input-style-always-active validate-field no-borders no-icon">
                    <input type="number" className="form-control validate-number" id="f3ab" value={orderAmount} onChange={(e) => { setOrderAmount(e.target.value) }} />
                    <label htmlFor="f3ab" className="color-theme opacity-30 text-uppercase font-700 font-10 mt-1">Amount in {props.balanceSymbol}</label>
                    <i className="fa fa-times disabled invalid color-red-dark"></i>
                    <i className="fa fa-check disabled valid color-green-dark"></i>
                    <em>(required)</em>
                </div>
                <div className="input-style input-style-always-active validate-field no-borders no-icon">
                    <input type="number" className="form-control validate-number" id="f1" value={orderTakeProfit} onChange={(e) => { setOrderTakeProfit(e.target.value) }} />
                    <label htmlFor="f1" className="color-theme opacity-30 text-uppercase font-700 font-10 mt-1">Take Profit At</label>
                    <i className="fa fa-times disabled invalid color-red-dark"></i>
                    <i className="fa fa-check disabled valid color-green-dark"></i>
                    <em></em>
                </div>
                <div className="input-style input-style-always-active no-borders no-icon">
                    <input type="number" className="form-control validate-number" id="f1a" value={orderStopLoss} onChange={(e) => { setOrderStopLoss(e.target.value) }} />
                    <label htmlFor="f1a" className="color-theme opacity-30 text-uppercase font-700 font-10 mt-1">Stop Loss At</label>
                    <i className="fa fa-times disabled invalid color-red-dark"></i>
                    <i className="fa fa-check disabled valid color-green-dark"></i>
                    <em></em>
                </div>
                <a
                    className="close-menu btn btn-full btn-m bg-theme color-theme gradient-sunny rounded-sm text-uppercase font-800 mb-3"
                    onClick={() => {
                        dispatch(buySpot(
                            props.spotID,
                            parseFloat(orderAmount),
                            props.assetSymbol,
                            parseFloat(orderTakeProfit),
                            parseFloat(orderStopLoss),
                        ));
                    }}
                >Create Order</a>
            </div>
        </div>
    )
}

export default SpotBuyModal