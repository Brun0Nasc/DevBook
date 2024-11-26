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
        alert('User or password incorrect');
    })
}