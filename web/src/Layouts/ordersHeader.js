import HeaderAmountInOrders from 'Components/headerAmountInOrders';
import HeaderBalance from 'Components/headerBalance';
import {useSelector} from "react-redux";

function OrdersHeader() {
    const amountInOrders = useSelector((state) => state.amountInOrders);
    
    return (
        <div className="header header-fixed header-logo-left">
            <HeaderAmountInOrders currency="USDT" amount={amountInOrders}/>
            <HeaderBalance/>
        </div>
    );
}

export default OrdersHeader;