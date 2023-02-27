import {Link} from 'react-router-dom';

function Footer() {
    return (
        <div id="footer-bar" className="footer-bar-1">
            <Link to="/" className="active-nav"><i className="fa fa-chart-line"></i><span>Spots</span></Link>
            <Link to="/orders"><i className="fa fa-list"></i><span>Orders</span></Link>
            <Link to="/settings"><i className="fa fa-cog"></i><span>Settings</span></Link>
            <Link to="/profile"><i className="fa fa-user"></i><span>Profile</span></Link>
        </div>
    );
}

export default Footer