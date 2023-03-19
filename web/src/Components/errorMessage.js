import { useDispatch, useSelector } from 'react-redux';

import { clearErrorMessage } from 'Store/reducer';

function ErrorMessage() {
    const dispatch = useDispatch();
    const errorMessage = useSelector((state) => state.errorMessage);

    return (
        <>
            {errorMessage && <div className="ms-3 me-3 mb-4 alert alert-small shadow-xl bg-red-dark" role="alert" style={{borderRadius: "15px"}}>
                <span style={{borderRadius: "15px 0 0 15px", left: "0", top: "0", bottom: "0"}}><i className="fa fa-times"></i></span>
                <strong>{errorMessage}</strong>
                <button type="button" className="close color-white opacity-60 font-16" onClick={() => { dispatch(clearErrorMessage()); }}>×</button>
            </div>}
        </>
    );
}

export default ErrorMessage;