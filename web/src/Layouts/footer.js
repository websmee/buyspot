import classNames from 'classnames';
import { useState } from 'react';
import { Link, useLocation } from 'react-router-dom';

function Footer() {
    const { pathname } = useLocation();
    const [activeNav, setActiveNav] = useState(pathname);

    const links = [
        { path: "/", text: "Spots", icon: "fa fa-chart-line" },
        { path: "/orders", text: "Orders", icon: "fa fa-list" },
        { path: "/settings", text: "Settings", icon: "fa fa-cog" },
        { path: "/profile", text: "Profile", icon: "fa fa-user" },
    ];

    return (
        <div id="footer-bar" className="footer-bar-1">{links.map((l, i) => {
            return <Link key={i} to={l.path} className={classNames({ "active-nav": activeNav == l.path })} onClick={() => { setActiveNav(l.path) }}>
                <i className={l.icon}></i>
                <span>{l.text}</span>
            </Link>
        })}</div>
    );
}

export default Footer