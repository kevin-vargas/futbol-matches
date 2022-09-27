import React from 'react';
import {useNavigate} from "react-router-dom";

const Profile = (props) => {
    const navigation = useNavigate();

    console.log("PROFILE: ", props.user);
    const user = props.user;

    const handleSave = (user) => {
        console.log("SAVE: ", user);
        navigation("/principal");
    }

    const handleCancel = () => {
        console.log("CANCELLL");
        navigation("/principal");
    }

    return (
        <section className="container">
            <form className="login-form rounded row">

                <div className="col-md-6">
                    <label className="form-label" htmlFor="firstName">First name</label>
                    <input type="email" id="firstName" name="firstName" className="form-control"
                           defaultValue={user.name}/>
                </div>

                <div className="col-md-6">
                    <label className="form-label" htmlFor="form2Example1">Last name</label>
                    <input type="email" id="lastName" name="lastName" className="form-control"
                           defaultValue={user.lastname}/>
                </div>

                <div className="col-md-6">
                    <label className="form-label" htmlFor="firstName">User name</label>
                    <input type="email" id="userName" name="userName" className="form-control"
                           defaultValue={user.username}/>
                </div>

                <div className="col-md-6">
                    <label className="form-label" htmlFor="phone">Phone</label>
                    <input type="email" id="phone" name="phone" className="form-control"
                           defaultValue={user.phone}/>
                </div>

                <div className="col-md-12">
                    <label className="form-label" htmlFor="email">Email</label>
                    <input type="email" id="email" name="email" className="form-control"
                           defaultValue={user.email}/>
                </div>

                <div className="col-md-12">
                    <label className="form-label" htmlFor="form2Example2">Password</label>
                    <input type="password" id="form2Example2" className="form-control"
                          defaultValue={user.password} />
                </div>
                <div className="col-md-12">
                    <br />
                </div>
                <div className="col-md-6">
                    <button type="button" className="btn btn-danger" onClick={ handleCancel }>Cancel</button>
                </div>
                <div className="col-md-6">
                    <button type="button" className="btn btn-success" onClick={ e => handleSave(user) }>Save</button>
                </div>
            </form>
        </section>);
}

export default Profile;
