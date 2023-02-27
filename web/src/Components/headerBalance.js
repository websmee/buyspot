function HeaderBalance(props) {
    return (
        <span className="header-title position-relative float-right pe-3 me-3">
            <i className="fa fa-wallet color-sunny-light"></i>
            <span className="color-sunny-light" style={{marginLeft:"5px"}}>
                <span className="opacity-30 font-200" style={{marginRight:"3px"}}>{props.currency}</span>
                {props.amount}
            </span>
        </span>
    );
}

export default HeaderBalance;