const SignInButton = document.getElementById("BtnSignIn");
const SignOutButton = document.getElementById("BtnSignOut");
const SignUpButton = document.getElementById("BtnSignUp");
const SumButton = document.getElementById("BtnSum");

SignInButton.addEventListener("click", (e) => {
    let User = document.getElementById("InpUserNameSignIn");
    let Pass = document.getElementById("InpPasswordSignIn");
    fetch('http://localhost:8585/signIn', {
        method: 'POST',
        headers: {
            "Content-Type": "application/json; charset=UTF-8",
        },
        body: JSON.stringify({
            "username"   :User.value,
            "password"   :Pass.value,
        }),
    })
        .then(function (response){
            return response.json()
        })
        .then(function (data){
            console.log(data)
            document.getElementById('LBLResult').innerHTML = data.Result;
           // alert(data.Result)
        });
})
SignUpButton.addEventListener("click", (e) => {
    let fName = document.getElementById("fName");
    let lName = document.getElementById("lName");
    let uName = document.getElementById("uName");
    let password = document.getElementById("password");
    let mail = document.getElementById("mail");
    let phone = document.getElementById("phone");
    fetch('http://localhost:8585/signUp', {
        method: 'POST',
        headers: {
            "Content-Type": "application/json; charset=UTF-8",
        },
        body: JSON.stringify({
            "firstname"  :fName.value,
            "lastname"   :lName.value,
            "username"   :uName.value,
            "password"   :password.value,
            "cellNo"     :phone.value,
            "email"      :mail.value
        }),
    })
        .then(function (response){
            return response.json()
        })
        .then(function (data){
            console.log(data)
            //alert(data.Result)
            document.getElementById('LBLResult').innerHTML = data.Result;

        });
})

SignOutButton.addEventListener("click", (e) => {
    let ticket = document.getElementById("ticket");
    fetch('http://localhost:8585/signOut', {
        method: 'POST',
        headers: {
            "Content-Type": "application/json; charset=UTF-8",
        },
        body: JSON.stringify({
            "ticket"   :ticket.value,
        }),
    })
        .then(function (response){
            return response.json()
        })
        .then(function (data){
            console.log(data)
            document.getElementById('LBLResult').innerHTML = data.Result;
        });
})

SumButton.addEventListener("click", (e) => {
    try {
        let Num1 = document.getElementById("Num1");
        let Num2 = document.getElementById("Num2");
        let NewTicket = document.getElementById("NewTicket");
        fetch('http://localhost:8585/calc', {
            method: 'POST',
            headers: {
                "Content-Type": "application/json; charset=UTF-8",
            },
            body: JSON.stringify({
                "number1": parseFloat( Num1.value),
                "number2":  parseFloat(Num2.value),
                "ticket":   NewTicket.value
            }),
        })
            .then(function (response){
                return response.json()
            })
            .then(function (data){
                console.log(data)
                document.getElementById('LBLResult').innerHTML = data.Result;
            });
    }
    catch(err) {
       alert(err)
    }

})