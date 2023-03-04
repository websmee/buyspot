import { useEffect } from 'react';

import Order from "Components/order"
import OrderSellModal from "Components/orderSellModal"
import OrdersHeader from 'Layouts/ordersHeader';

function Orders() {
    useEffect(() => {
        window.bindAll();

        return () => {
            window.unbindAll();
        };
    }, []);

    return (
        <>
            <OrdersHeader />

            <div className="page-content header-clear-medium">
                <Order modalId="sell-modal-1" assetName="Bitcoin" assetTicker="BTC" created="1 hour ago" amountTicker="USDT" amount="123.45" pnl="+2.34%" />
                <Order modalId="sell-modal-2" assetName="Ethereum" assetTicker="ETH" created="24 hours ago" amountTicker="USDT" amount="111.11" pnl="-1.23%" />
            </div>

            <OrderSellModal id="sell-modal-1" assetName="Bitcoin" assetTicker="BTC" created="1 hour ago" amountTicker="USDT" amount="123.45" pnl="+2.34%" takeProfit="+3%" stopLoss="-1%" />
            <OrderSellModal id="sell-modal-2" assetName="Ethereum" assetTicker="ETH" created="24 hours ago" amountTicker="USDT" amount="111.11" pnl="-1.23%" takeProfit="+4%" stopLoss="-2%" />
        </>
    )
}

export default Orders