import axios from "axios";

import { FETCH_USER } from "./types";

export const fetchUser = credentials => async dispatch => {
    const res = await axios.get("http://localhost:8080/signin", credentials);
    dispatch({ type: FETCH_USER, payload: res.data });
};