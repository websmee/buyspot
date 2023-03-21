import HeaderSpots from 'Components/headerSpots';
import HeaderBalance from 'Components/headerBalance';

function SpotHeader() {
    return (
        <div className="header header-fixed header-logo-left">
            <HeaderSpots />
            <HeaderBalance />
        </div>
    );
}

export default SpotHeader;