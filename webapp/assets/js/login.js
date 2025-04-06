$('#login').on('submit', doLogin);

function doLogin(event) {
    event.preventDefault();

    $.ajax({
        url: '/login',
        method: 'POST',
        data: {
            email: $('#email').val(),
            password: $('#password').val()
        }
    }).done(function() {
        window.location = "/home";
    }).fail(function() {
        Swal.fire("Oops!", "Invalid email or password.", "error");
    })
}