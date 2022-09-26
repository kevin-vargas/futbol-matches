import config from '../config/config'

const userService = {

    signup: (user) => {
        const requestOptions = config.endpoints.signup.requestOptions;
        requestOptions.body = JSON.stringify(user);
        const url = `${config.apiHost}${config.endpoints.signup.path}`;

        return fetch(url, requestOptions);
    },

    login: (user, pass) => {
    }
}

export default userService;
