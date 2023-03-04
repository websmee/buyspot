import HeaderSpots from 'Components/headerSpots';
import HeaderBalance from 'Components/headerBalance';

function SpotHeader(props) {
    return (
        <div className="header header-fixed header-logo-left">
            <HeaderSpots count={props.spotsCount} />
            <HeaderBalance />
        </div>
    );
}

export default SpotHeader;