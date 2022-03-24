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


//Переход на другую страницу после авторизации
/*function loginAccess() {
 let button = document.getElementById('submit')
 let login = document.getElementById('Username')
 let password = document.getElementById('Password')
 if ((login = 14299) && (password = '1')) {
 let notif = document.createElement('div')
 document.body.prepend(notif)
 notif.className = 'active'
 notif.textContent = 'Авторизация прошла успешно'
 window.location = '../Home_page/home.html'
 }
 }*/