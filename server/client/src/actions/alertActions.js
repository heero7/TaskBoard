import { alertActions, alertConstants } from "../constants";

// All available actions, exported into one object
export const alertActions = {
    success,
    failure,
    clear
};

// Call to create a success toaster notification
function success(message) {
    return { type : alertConstants.SUCCESS };
}

// Call to create a failure toaster notification
function failure(message) {
    return { type : alertConstants.FAILURE };
}

// Call to create a clear toaster notification
function clear(message) {
    return { type : alertConstants.CLEAR };
}