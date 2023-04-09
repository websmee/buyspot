import { useEffect, useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import { useNavigate } from "react-router-dom";

import { login } from "Store/reducer";
import stickymobile from "Utils/stickymobile"

function Login() {
    const dispatch = useDispatch();
    const navigate = useNavigate();
    const unauthorized = useSelector((state) => state.unauthorized);
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");

    useEffect(() => {
        if (unauthorized) {
            stickymobile.extendCard("login-card");
            stickymobile.hidePreloader();
        } else {
            navigate('/');
        }
    }, [unauthorized]);

    return (
        <div className="page-content pb-0">
            <div id="login-card" data-card-height="cover" className="card">
                <div className="card-center">
                    <div className="ps-5 pe-5">
                        <h1 className="text-center font-800 font-40 mb-1"><strong className="color-sunny-light">BUY</strong>SPOT</h1>
                        <p className="color-highlight text-center font-12">Let's get you logged in</p>

                        <div className="input-style no-borders has-icon validate-field">
                            <i className="fa fa-user"></i>
                            <input type="email" className="form-control validate-name" id="form1a" placeholder="Email" value={email} onChange={(e) => { setEmail(e.target.value) }}/>
                            <label htmlFor="form1a" className="color-blue-dark font-10 mt-1">Email</label>
                            <i className="fa fa-times disabled invalid color-red-dark"></i>
                            <i className="fa fa-check disabled valid color-green-dark"></i>
                            <em>(required)</em>
                        </div>

                        <div className="input-style no-borders has-icon validate-field mt-4">
                            <i className="fa fa-lock"></i>
                            <input type="password" className="form-control validate-password" id="form3a" placeholder="Password" value={password} onChange={(e) => { setPassword(e.target.value) }}/>
                            <label htmlFor="form3a" className="color-blue-dark font-10 mt-1">Password</label>
                            <i className="fa fa-times disabled invalid color-red-dark"></i>
                            <i className="fa fa-check disabled valid color-green-dark"></i>
                            <em>(required)</em>
                        </div>

                        <a
                            onClick={() => {
                                dispatch(login(email, password));
                            }}
                            className="back-button btn btn-full btn-m shadow-large rounded-sm text-uppercase font-700 bg-highlight">LOGIN</a>
                    </div>
                </div>
            </div>
        </div>
    )
}

export default Login