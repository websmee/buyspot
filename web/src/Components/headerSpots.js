import { useSelector } from "react-redux";

function HeaderSpots() {
    const currentSpotsTotal = useSelector((state) => state.currentSpotsTotal);
    
    return (
        <span className="header-title">
            <span style={{marginRight:"5px"}}>Spots</span>
            <span className="badge rounded-xl bg-black">{currentSpotsTotal}</span>
        </span>
    );
}

export default HeaderSpots;