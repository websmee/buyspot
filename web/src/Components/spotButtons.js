import { useEffect } from "react";

import stickymobile from "Utils/stickymobile";

function SpotButtons(props) {
    useEffect(() => {
        const menuOpenListener = stickymobile.getMenuOpenListener(props.buyModalId);
        const menuCloseListener = stickymobile.getMenuCloseListener();
        stickymobile.bindMenu(props.buyModalId, menuOpenListener, menuCloseListener);

        return () => {
            stickymobile.unbindMenu(props.buyModalId, menuOpenListener, menuCloseListener);
        }
    }, [])

    return (
        <div className="content mb-0">
            <div className="row mb-0">
                <div className="col-6 pe-1">
                    <a href="#" data-menu={props.buyModalId}
                    className="card-style d-block bg-theme gradient-sunny py-3 mx-0">
                        {props.activeOrdersCount > 0 && <span className="ps-3 pt-3 mt-n1 font-10 opacity-50 position-absolute">{props.activeOrdersCount} active orders for {props.assetTicker}</span>}
                        <span href="#" className="color-theme font-800 font-13 text-uppercase px-3">
                            <i className="fa fa-check pt-2 pe-3 float-end"></i>
                            Buy
                        </span>
                    </a>
                </div>
                <div className="col-6 ps-1">
                    <a href="#" data-menu="menu-transaction-request"
                    className="card-style d-block bg-theme gradient-dark py-3 mx-0">
                        <span className="ps-3 pt-3 mt-n1 font-10 opacity-50 position-absolute">spot {props.currentSpot} out of {props.spotCount}</span>
                        <span href="#" className="color-theme font-800 font-13 text-uppercase px-3"><i
                                className="fa fa-arrow-right pt-2 pe-3 float-end"></i>Next</span>
                    </a>
                </div>
            </div>
        </div>
    )
}

export default SpotButtons