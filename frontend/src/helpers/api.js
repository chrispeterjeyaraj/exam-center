export const fetchData = (url, method, body) => {
    return fetch(`http://localhost:1234/${url}`, {
        method: method, // *GET, POST, PUT, DELETE, etc.
        mode: "cors", // no-cors, cors, *same-origin
        body: JSON.stringify(body),
        credentials: "include", // include, *same-origin, omit
        headers: {
            "Content-Type": "application/json; charset=utf-8",
        },
    });
}