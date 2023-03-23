import { useDispatch, useSelector } from "react-redux";
import { getOrders } from "Store/reducer";
import { useEffect } from "react";

import Order from "Components/order"
import OrderSellModal from "Components/orderSellModal"
import OrdersHeader from 'Layouts/ordersHeader';

function Orders() {
    const dispatch = useDispatch();
    const orders = useSelector((state) => state.orders);

    useEffect(() => {
        dispatch(getOrders());
    }, [dispatch]);

    return (
        <>
            <OrdersHeader />

            <div className="page-content header-clear-medium">
                {orders.map((order, i) =>
                    <Order key={i} modalId={"sell-modal-" + i} order={order} />
                )}
            </div>

            {orders.map((order, i) =>
                <OrderSellModal key={i} id={"sell-modal-" + i} order={order} />
            )}
        </>
    )
}

export default Orders