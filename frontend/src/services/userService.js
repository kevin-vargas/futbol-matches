const userService = {
    login: (user, pass) => {
        let role = 'user';
        if( pass === 'q1w2e3r4') {
            role = 'admin';
        }
        return Promise.resolve({
            name: "John Henry",
            lastname: "Bonham",
            username: "Bonzo",
            phone: 123123,
            email: 'bonzosmotreux@live.com',
            role: role,
        });
    }
}

export default userService;
