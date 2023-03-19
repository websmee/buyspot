import { useEffect } from "react";

import stickymobile from "Utils/stickymobile";

function Order(props) {
    useEffect(() => {
        const menuOpenListener = stickymobile.getMenuOpenListener(props.modalId);
        const menuCloseListener = stickymobile.getMenuCloseListener();
        stickymobile.bindMenu(props.modalId, menuOpenListener, menuCloseListener);
        stickymobile.bindEmptyLinks();

        return () => {
            stickymobile.unbindMenu(props.modalId, menuOpenListener, menuCloseListener);
            stickymobile.unbindEmptyLinks();
        }
    }, [])

    return (
        <div className="card card-style mb-3" data-menu={props.modalId}>
            <div className="content m-3">
                <div className="d-flex">
                    <div className="align-self-center">
                        <h5 className="mb-n2">
                            {props.assetName}
                            <span className="opacity-30 font-200" style={{marginLeft: "3px"}}>{props.assetTicker}</span>
                        </h5>
                        <p className="mt-n1 mb-0 font-10">{props.created}</p>
                    </div>
                    <div className="align-self-center ms-auto ps-3">
                        <h5 className="mb-n2">
                            <span className="opacity-30 font-200" style={{marginRight: "3px"}}>{props.amountTicker}</span>
                            <a href="#" className={props.pnl[0] == "-" ? "color-red-light" : "color-sunny-light"}>{props.amount}</a>
                        </h5>
                        <p className="mt-n1 mb-0 font-10" style={{textAlign: "right"}}>{props.pnl}</p>
                    </div>
                </div>
            </div>
        </div>
    );
}

export default Order;