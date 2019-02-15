import { userConstants } from "../constants/user.constants";
import { userService } from "../services";
import { alertActions } from "./";
import { history } from "../util";

export const userActions = {
    login,
    logout,
    register
};

function login(email, password) {
    return dispatch => {
        dispatch(request({ email }));

        userService.login(email, password)
            .then(user => {
                dispatch(success(user));
                history.push("/dashboard")
            }, 
            error => {
                dispatch(failure(error.toString()));
                dispatch(alertActions.error(error.toString()));
            }
        );
    };

    function request(user) { return { type: userConstants.LOGIN_REQUEST, user }};
    function success(user) { return { type: userConstants.LOGIN_SUCCESS, user }};
    function failure(error) { return { type: userConstants.LOGIN_FAILURE, error }};
}

function logout() {
    userService.logout();
    return { type: userConstants.LOGOUT };
}

function register(credentials) {
    return dispatch => {
        dispatch(request(credentials));

        userService.register(credentials)
            .then(user => {
                dispatch(success());
                history.push("/signin");
                dispatch(alertActions.success("Account created!"));
            },
            error => {
                dispatch(failure(error.toString()));
                dispatch(alertActions.error(error.toString()));
            }
        );
    };
    function request(user) { return { type: userConstants.REGISTER_REQUEST, user } }
    function success(user) { return { type: userConstants.REGISTER_SUCCESS, user } }
    function failure(error) { return { type: userConstants.REGISTER_FAILURE, error } }
}