import { useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';

import { getCurrentBalance } from 'Store/reducer';

function HeaderBalance(props) {
    const dispatch = useDispatch();
    const balance = useSelector((state) => state.balance);

    useEffect(() => {
        dispatch(getCurrentBalance());
    }, [dispatch]);

    return (
        <span className="header-title position-relative float-right pe-3 me-3">
            <i className="fa fa-wallet color-sunny-light"></i>
            <span className="color-sunny-light" style={{marginLeft:"5px"}}>
                <span className="opacity-30 font-200" style={{marginRight:"3px"}}>{balance.symbol}</span>
                {balance.amount}
            </span>
        </span>
    );
}

export default HeaderBalance;