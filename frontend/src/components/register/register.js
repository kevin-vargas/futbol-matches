import React, {useRef} from 'react'
import { useNavigate } from "react-router-dom";

const Register = (props) => {

    const navigation = useNavigate();

    const userRef = {};
    userRef.name = useRef();
    userRef.lastname = useRef();
    userRef.username = useRef();
    userRef.password = useRef();
    userRef.email = useRef();
    userRef.phone = useRef();

    const handleCancel = () => {
        navigation("/");
    }

    return(
        <section className="container">
            <form className="login-form rounded row">

                <div className="col-md-6">
                    <label className="form-label" htmlFor="firstName">First name</label>
                    <input type="email" id="firstName" name="firstName" className="form-control"
                               ref={userRef.name}/>
                </div>

                <div className="col-md-6">
                    <label className="form-label" htmlFor="form2Example1">Last name</label>
                    <input type="email" id="lastName" name="lastName" className="form-control"
                           ref={userRef.lastname}/>
                </div>

                <div className="col-md-6">
                    <label className="form-label" htmlFor="firstName">User name</label>
                    <input type="email" id="userName" name="userName" className="form-control"
                           ref={userRef.username}/>
                </div>

                <div className="col-md-6">
                    <label className="form-label" htmlFor="phone">Phone</label>
                    <input type="email" id="phone" name="phone" className="form-control"
                           ref={userRef.phone}/>
                </div>

                <div className="col-md-12">
                    <label className="form-label" htmlFor="email">Email</label>
                    <input type="email" id="email" name="email" className="form-control"
                           ref={userRef.email}/>
                </div>

                <div className="col-md-12">
                    <label className="form-label" htmlFor="form2Example2">Password</label>
                    <input type="password" id="form2Example2" className="form-control"
                           ref={userRef.password} />
                </div>
                <div className="col-md-12">
                    <br />
                </div>
                <div className="col-md-6">
                    <button to="/" type="button" className="btn btn-danger" onClick={ e => handleCancel }>Cancel</button>
                </div>
                <div className="col-md-6">
                    <button type="button" className="btn btn-success" onClick={ e => props.handleRegistration(userRef) }>Save</button>
                </div>
            </form>
        </section>
    )
}

export default Register;
