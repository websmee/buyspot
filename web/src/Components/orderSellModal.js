function OrderSellModal(props) {
    return (
        <div id={props.id} className="menu menu-box-modal menu-box-detached">
            <div className="menu-title">
                <h1>
                    {props.assetName}
                    <span className="opacity-30 font-200" style={{marginLeft: "5px"}}>{props.assetTicker}</span>
                </h1>
                <p>{props.created}</p>
                <a href="#" className="close-menu"><i className="fa fa-times"></i></a>
            </div>
            <div className="divider divider-margins mb-1 mt-3"></div>
            <div className="content px-1">
                <p>
                    Current value: {props.amountTicker} <strong className={props.pnl[0] == "-" ? "color-red-light" : "color-sunny-light"}>{props.amount}</strong>
                    <br />
                    PNL: <strong className={props.pnl[0] == "-" ? "color-red-light" : "color-sunny-light"}>{props.pnl}</strong>
                    <br />
                    Will be sold automatically at {props.takeProfit} or {props.stopLoss}
                </p>
                <a href="#" className="close-menu btn btn-full btn-m bg-theme color-theme gradient-sunny rounded-sm text-uppercase font-800 mb-3">Sell Now</a>
            </div>
        </div>
    )
}

export default OrderSellModal