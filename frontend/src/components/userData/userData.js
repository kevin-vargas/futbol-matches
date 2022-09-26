import React from 'react'

const UserData = (props) => {

    const user = props.user || {}

    return(
        <section className="container">
            <form className="login-form rounded row">

                <div className="col-md-6">
                    <label className="form-label" htmlFor="firstName">First name</label>
                    <input type="email" id="firstName" name="firstName" className="form-control"
                           ref={user.name}
                    defaultValue={user.name.current}/>
                </div>

                <div className="col-md-6">
                    <label className="form-label" htmlFor="form2Example1">Last name</label>
                    <input type="email" id="lastName" name="lastName" className="form-control"
                           ref={user.lastname}
                           defaultValue={user.lastname.current}/>
                </div>

                <div className="col-md-6">
                    <label className="form-label" htmlFor="firstName">User name</label>
                    <input type="email" id="userName" name="userName" className="form-control"
                           ref={user.username}
                           defaultValue={user.username.current}/>
                </div>

                <div className="col-md-6">
                    <label className="form-label" htmlFor="phone">Phone</label>
                    <input type="email" id="phone" name="phone" className="form-control"
                           ref={user.phone}
                           defaultValue={user.phone.current}/>
                </div>

                <div className="col-md-12">
                    <label className="form-label" htmlFor="email">Email</label>
                    <input type="email" id="email" name="email" className="form-control"
                           ref={user.email}
                           defaultValue={user.email.current}/>
                </div>

                <div className="col-md-12">
                    <label className="form-label" htmlFor="form2Example2">Password</label>
                    <input type="password" id="form2Example2" className="form-control"
                           ref={user.password}/>
                </div>
                <div className="col-md-12">
                    <br />
                </div>
                <div className="col-md-6">
                    <button to="/" type="button" className="btn btn-danger" onClick={ e => props.handleCancel }>Cancel</button>
                </div>
                <div className="col-md-6">
                    <button type="button" className="btn btn-success" onClick={ e => props.handleSave(user) }>Save</button>
                </div>
            </form>
        </section>
    )
}

export default UserData;
