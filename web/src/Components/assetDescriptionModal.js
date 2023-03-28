function AssetDescriptionModal(props) {
    return (
        <div id={props.id} className="menu menu-box-bottom menu-box-detached">
            <div className="menu-title">
                <h1 className="pt-3 ps-3 pe-5">{props.assetName}</h1>
                <h4 className="ps-3 font-700 text-uppercase mt-n2 font-12 opacity-30">{props.assetSymbol}</h4>
                <a className="close-menu"><i className="fa fa-times"></i></a>
            </div>
            <div className="content">
                <div className="divider mb-4"></div>
                <p style={{overflowY: "scroll",maxHeight: "230px"}}>
                    {props.children}
                </p>
                <div className="divider mb-4"></div>
                <a className="close-menu btn btn-full btn-m rounded-sm bg-highlight font-800 text-uppercase mb-4">Close
                    Description</a>
            </div>
        </div>
    )
}

export default AssetDescriptionModal