export function authHeader() {
    // Set the auth header with the JWT token

    let user = JSON.parse(localStorage.getItem("user"));

    if (user && user.token) {
        return { "Authorization" : "Bearer " + user.token };
    }
    return {};
}