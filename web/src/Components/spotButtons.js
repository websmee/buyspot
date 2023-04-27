import { useEffect } from "react";
import { useDispatch, useSelector } from "react-redux";
import { getSpotByIndex } from "Store/reducer";

import stickymobile from "Utils/stickymobile";

function SpotButtons(props) {
    const dispatch = useDispatch();
    const currentSpotsIndex = useSelector((state) => state.currentSpotsIndex);
    const currentSpotsTotal = useSelector((state) => state.currentSpotsTotal);
    const currentSpotsNext = useSelector((state) => state.currentSpotsNext);
    const currentSpotsPrev = useSelector((state) => state.currentSpotsPrev);

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
                    <a data-menu={props.buyModalId}
                        className="card-style d-block bg-theme gradient-sunny py-3 mx-0">
                        {props.activeOrdersCount > 0 && <span className="ps-3 pt-3 mt-n1 font-10 opacity-50 position-absolute">{props.activeOrdersCount} active orders for {props.assetSymbol}</span>}
                        <span className="color-theme font-800 font-13 text-uppercase px-3">
                            <i className="fa fa-check pt-2 pe-3 float-end"></i>
                            Buy
                        </span>
                    </a>
                </div>
                <div className="col-2 ps-1 pe-1">
                    <a data-menu="menu-transaction-request" onClick={() => { dispatch(getSpotByIndex(currentSpotsPrev)); }}
                        className="card-style d-block bg-theme gradient-dark py-3 mx-0">
                        <span className="color-theme font-800 font-13 text-uppercase"><i
                            className="fa fa-arrow-left pt-2 float-end" style={{paddingRight: "42%"}}></i>&nbsp;</span>
                    </a>
                </div>
                <div className="col-4 ps-1">
                    <a data-menu="menu-transaction-request" onClick={() => { dispatch(getSpotByIndex(currentSpotsNext)); }}
                        className="card-style d-block bg-theme gradient-dark py-3 mx-0">
                        <span className="ps-3 pt-3 mt-n1 font-10 opacity-50 position-absolute">{currentSpotsIndex} out of {currentSpotsTotal}</span>
                        <span className="color-theme font-800 font-13 text-uppercase px-3"><i
                            className="fa fa-arrow-right pt-2 pe-3 float-end"></i>Next</span>
                    </a>
                </div>
            </div>
        </div>
    )
}

export default SpotButtons