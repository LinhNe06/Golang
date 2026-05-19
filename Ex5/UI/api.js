const API = "http://localhost:9999/orders";

function show(data) {
    document.getElementById("result").textContent = JSON.stringify(data, null, 2);
}

async function parseResponse(response) {
    const contentType = response.headers.get("content-type") || "";
    if (contentType.includes("application/json")) {
        return response.json();
    }
    const text = await response.text();
    return { message: text };
}

async function createOrder(data) {
    const response = await fetch(API, {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(data)
    });
    const body = await parseResponse(response);
    show(response.ok ? body : { status: response.status, error: body });
}

async function getOrders() {
    const response = await fetch(API);
    const body = await parseResponse(response);
    show(response.ok ? body : { status: response.status, error: body });
}

async function getOrder(id) {
    const response = await fetch(`${API}/${id}`);
    const body = await parseResponse(response);
    show(response.ok ? body : { status: response.status, error: body });
}

async function updateOrder(id, data) {
    const response = await fetch(`${API}/${id}`, {
        method: "PUT",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(data)
    });
    const body = await parseResponse(response);
    show(response.ok ? body : { status: response.status, error: body });
}

async function deleteOrder(id) {
    const response = await fetch(`${API}/${id}`, {
        method: "DELETE"
    });
    const body = await parseResponse(response);
    show(response.ok ? body : { status: response.status, error: body });
}
