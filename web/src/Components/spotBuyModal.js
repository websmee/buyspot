function SpotBuyModal(props) {
    return (
        <div id={props.id} className="menu menu-box-modal menu-box-detached">
            <div className="menu-title"><h1>Buy {props.assetName}</h1>
                <p className="color-highlight">Convert {props.balanceTicker} to {props.assetTicker}</p><a href="#" className="close-menu"><i className="fa fa-times"></i></a>
            </div>
            <div className="divider divider-margins mb-1 mt-3"></div>
            <div className="content px-1">
                <div className="input-style input-style-always-active validate-field no-borders no-icon">
                    <input type="number" className="form-control validate-number" id="f3ab" placeholder="100" />
                    <label htmlFor="f3ab" className="color-theme opacity-30 text-uppercase font-700 font-10 mt-1">Amount in {props.balanceTicker}</label>
                    <i className="fa fa-times disabled invalid color-red-dark"></i>
                    <i className="fa fa-check disabled valid color-green-dark"></i>
                    <em>(required)</em>
                </div>
                <div className="input-style input-style-always-active no-borders no-icon">
                    <label htmlFor="f1" className="color-theme opacity-30 text-uppercase font-700 font-10 mt-1">Take Profit
                        At</label>
                    <select id="f1">
                        <option value="default" selected>+3%</option>
                        <option value="1">+2%</option>
                        <option value="2">+1%</option>
                    </select>
                    <span><i className="fa fa-chevron-down"></i></span>
                    <i className="fa fa-check disabled valid color-green-dark"></i>
                    <em></em>
                </div>
                <div className="input-style input-style-always-active no-borders no-icon">
                    <label htmlFor="f1a" className="color-theme opacity-30 text-uppercase font-700 font-10 mt-1">Stop Loss
                        At</label>
                    <select id="f1a">
                        <option value="default" selected>-1%</option>
                        <option value="1">-2%</option>
                        <option value="2">-3%</option>
                    </select>
                    <span><i className="fa fa-chevron-down"></i></span>
                    <i className="fa fa-check disabled valid color-green-dark"></i>
                    <em></em>
                </div>
                <a href="#" className="close-menu btn btn-full btn-m bg-theme color-theme gradient-sunny rounded-sm text-uppercase font-800 mb-3">Create Order</a>
            </div>
        </div>
    )
}

export default SpotBuyModal