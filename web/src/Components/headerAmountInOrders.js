function HeaderAmountInOrders(props) {
    return (
        <span className="header-title">
            <i className="fa fa-arrows-rotate color-sunny-light"></i>
            <span className="color-sunny-light" style={{marginLeft:"5px"}}>
                <span className="opacity-30 font-200" style={{marginRight:"3px"}}>{props.currency}</span>
                {props.amount}
            </span>
        </span>
    );
}

export default HeaderAmountInOrders;