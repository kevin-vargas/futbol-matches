import React from "react"
import { NavLink } from "react-router-dom";

const Login = (props) => {

    return (
        <section className="container login">
            <form className="login-form rounded">
                <div className="col-md-12">
                    <label className="form-label" htmlFor="form2Example1">Username</label>
                    <input type="email" id="form2Example1" className="form-control" ref={props.userName} />
                </div>
                <br />
                <div className="col-md-12">
                    <label className="form-label" htmlFor="form2Example2">Password</label>
                    <input type="password" id="form2Example2" className="form-control" ref={props.password}/>
                </div>
                <br />
                <div className="col-md-12">
                    <button type="button" className="btn btn-primary"
                    onClick={ props.handleLogin }>Login</button>
                </div>
                <br />
                <div className="col-md-12">
                    <NavLink to="/register">Don't have a user? Sign up!</NavLink>
                </div>
            </form>
        </section>)
}

export default Login;
