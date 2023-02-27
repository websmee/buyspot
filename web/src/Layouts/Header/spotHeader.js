import HeaderSpots from 'Components/headerSpots';
import HeaderBalance from 'Components/headerBalance';

function SpotHeader() {
    return (
        <div className="header header-fixed header-logo-left">
            <HeaderSpots count="10"/>
            <HeaderBalance currency="USDT" amount="1234.56"/>
        </div>
    );
}

export default SpotHeader;