function HeaderSpots(props) {
    return (
        <span className="header-title">
            <span style={{marginRight:"5px"}}>Spots</span>
            <span className="badge rounded-xl bg-black">{props.count}</span>
        </span>
    );
}

export default HeaderSpots;