import React from 'react'

const UserData = (props) => {

    const user = props.user || {}

    return(
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

                <div className="col-md-9">
                    <label className="form-label" htmlFor="form2Example1">Avatar</label>
                    <input type="email" id="avatar" name="avatar" className="form-control" />
                </div>

                <div className="col-md-3">
                    <button type="button" className="btn btn-primary upload">Upload</button>
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
                    <input type="password" id="form2Example2" className="form-control" />
                </div>
                <div className="col-md-12">
                    <br />
                </div>
                <div className="col-md-6">
                    <button to="/" type="button" className="btn btn-danger" onClick={ props.handleCancel }>Cancel</button>
                </div>
                <div className="col-md-6">
                    <button type="button" className="btn btn-success" onClick={ props.handleSave }>Save</button>
                </div>
            </form>
        </section>
    )
}

export default UserData;
