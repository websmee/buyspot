function NewsArticle(props) {
    return (
        <div className="card card-style mb-2" data-menu={props.modalId}>
            <a href="#" className="content ms-2 my-0">
                <div className="d-flex">
                    <div className="align-self-center">
                        <span className="icon icon-s rounded-xl me-3"><i className={"fa fa-light fa-regular font-28 " + props.sentimentIconClass}></i></span>
                    </div>
                    <div className="align-self-center w-100">
                        <h5 className="font-15 pt-2">{props.children}</h5>
                        <span className="color-theme font-11 opacity-50">
                            <i className="far fa-clock fa-fw pe-2"></i>{props.created}
                            <i className="far fa-eye fa-fw px-3"></i>{props.views}
                        </span>
                    </div>
                </div>
            </a>
        </div>
    );
}

export default NewsArticle;