import HeaderAmountInOrders from 'Components/headerAmountInOrders';
import HeaderBalance from 'Components/headerBalance';

function OrdersHeader() {
    return (
        <div className="header header-fixed header-logo-left">
            <HeaderAmountInOrders currency="USDT" amount="123.45"/>
            <HeaderBalance currency="USDT" amount="1234.56"/>
        </div>
    );
}

export default OrdersHeader;