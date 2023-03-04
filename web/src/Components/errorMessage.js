import { useDispatch, useSelector } from 'react-redux';

import { clearErrorMessage } from 'Store/reducer';

function ErrorMessage() {
    const dispatch = useDispatch();
    const errorMessage = useSelector((state) => state.errorMessage);

    return (
        <>
            {errorMessage && <div className="ms-3 me-3 mb-4 alert alert-small rounded shadow-xl bg-red-dark" role="alert">
                <span className="rounded-start rounded-start"><i className="fa fa-times"></i></span>
                <strong>{errorMessage}</strong>
                <button type="button" className="close color-white opacity-60 font-16" onClick={() => { dispatch(clearErrorMessage()); }}>×</button>
            </div>}
        </>
    );
}

export default ErrorMessage;