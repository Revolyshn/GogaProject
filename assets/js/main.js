let authBtn = document.querySelector("#submit")

authBtn.onclick = function () {


    let inputs = document.querySelectorAll("#box > .fields")
    let authData = {};

    for (let i = 0; i < inputs.length; i++) {
        authData[inputs[i].name] = inputs[i].value
    }

    console.log(inputs)
    console.log(authData)

    let xhr = new XMLHttpRequest();

    xhr.open("POST", "/user/authorization");
    xhr.onload = function (e) {
        console.log(e)
    };
    xhr.send(JSON.stringify(authData));

}

function ClearElement(element) {
    if (!element || !(element instanceof HTMLElement)) return;

    while (element.children.length > 0) element.children[0].remove();
}

function Div(props) {
    let div = document.createElement("div");

    if (!props || typeof props !== "object") {
        return div;
    }

    if ("textContent" in props) {
        div.textContent = props.textContent;
    }

    if ("events" in props && typeof props.events === "object") {
        for (let key in props.events) {
            div[key] = props.events[key];
        }
    }

    if ("className" in props) {
        div.className = props.className;
    }

    if ("style" in props) {
        div.style.cssText = props.style;
    }

    if ("id" in props) {
        div.id = props.id;
    }

    if ("dataset" in props) {
        for (let key in props.dataset) {
            div.dataset[key] = props.dataset[key];
        }
    }

    if ("children" in props) {
        div.append(...props.children);
    }

    return div;
}