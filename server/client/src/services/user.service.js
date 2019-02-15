import axios from "axios";

export const userService = {
    login,
    logout,
    register
};

function login(email, password) {
    let credentials = {
        email : email,
        password : password
    };
    return axios.post("http://localhost:8080/signin", credentials)
    .then(responseHandler)
    .then(user => {
        // store token in local storage for reference
        // and future requests that require authentication
        localStorage.setItem("user", JSON.stringify(user));
        return user;
    });
}

function logout() {
    localStorage.removeItem("user");
}

function register(credentials) {
    return axios.post("http://localhost:8080/", credentials).then(responseHandler);
}

function responseHandler(response) {
    return response.text().then(text => {
        const data = text && JSON.parse(text); // need to document this
        if (!response.ok) {
            if (response.status === 401) {
                logout();
                window.location.reload(true); // basically resets the page to clear state and such
            }

            const error = (data && data.message) || response.statusText;
            return Promise.reject(error);
        }
        return data;
    });
}