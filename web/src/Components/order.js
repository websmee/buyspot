import { useEffect } from "react";

import ReactTimeAgo from "react-time-ago";

import stickymobile from "Utils/stickymobile";

function Order(props) {

    useEffect(() => {
        const menuOpenListener = stickymobile.getMenuOpenListener(props.modalId);
        const menuCloseListener = stickymobile.getMenuCloseListener();
        stickymobile.bindMenu(props.modalId, menuOpenListener, menuCloseListener);

        return () => {
            stickymobile.unbindMenu(props.modalId, menuOpenListener, menuCloseListener);
        }
    }, [])

    return (
        <div className="card card-style mb-3" data-menu={props.modalId}>
            <div className="content m-3">
                <div className="d-flex">
                    <div className="align-self-center">
                        <h5 className="mb-n2">
                            {props.order.toAssetName}
                            <span className="opacity-30 font-200" style={{marginLeft: "3px"}}>{props.order.toTicker}</span>
                        </h5>
                        <p className="mt-n1 mb-0 font-10"><ReactTimeAgo date={Date.parse(props.order.created)} locale="en-US" /></p>
                    </div>
                    <div className="align-self-center ms-auto ps-3">
                        <h5 className="mb-n2">
                            <span className="opacity-30 font-200" style={{marginRight: "3px"}}>{props.order.fromTicker}</span>
                            <a className={props.order.pnl < 0 ? "color-red-light" : "color-sunny-light"}>{props.order.amountInBalanceTicker}</a>
                        </h5>
                        <p className="mt-n1 mb-0 font-10" style={{textAlign: "right"}}>{props.order.pnl > 0 && "+"}{props.order.pnl}%</p>
                    </div>
                </div>
            </div>
        </div>
    );
}

export default Order;