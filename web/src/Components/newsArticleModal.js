import ReactTimeAgo from 'react-time-ago'

import numbers from 'Utils/numbers';

function NewsArticleModal(props) {
    return (
        <div id={props.id} className="menu menu-box-bottom menu-box-detached">
            <div className="menu-title">
                <h4 className="pt-3 ps-3 pe-5">{props.title}</h4>
                <span className="ps-3 color-theme font-11 opacity-50">
                    <i className="far fa-clock fa-fw pe-2"></i><ReactTimeAgo date={Date.parse(props.created)} locale="en-US" />
                    <i className="far fa-eye fa-fw px-3"></i>{numbers.pretty(props.views, 1)}
                </span>
                <a href="#" className="close-menu"><i className="fa fa-times"></i></a>
            </div>
            <div className="content">
                <div className="divider mb-4"></div>
                <p style={{ overflowY: "scroll", maxHeight: "230px" }}>
                    {props.children}
                </p>
                <div className="divider mb-4"></div>
                <a href="#" className="close-menu btn btn-full btn-m rounded-sm bg-highlight font-800 text-uppercase mb-4">Close Article</a>
            </div>
        </div>
    )
}

export default NewsArticleModal