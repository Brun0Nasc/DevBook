$('#form-register').on('submit', createUser)

function createUser(event) {
    event.preventDefault()

    if ($('#password').val() != $('#confirm-password').val()) {
        Swal.fire("Oops!", "Passwords do not match.", "error");
        return;
    }

    $.ajax({
        url: "/users",
        method: "POST",
        data: {
            name: $('#name').val(),
            email: $('#email').val(),
            nick: $('#nick').val(),
            password: $('#password').val()
        }
    }).done(function() {
        Swal.fire("Success!", "User created successfully.", "success").then(() => {
            $.ajax({
                url: "/login",
                method: "POST",
                data: {
                    email: $('#email').val(),
                    password: $('#password').val()
                }
            }).done(function() {
                window.location = "/home";
            }).fail(function() {
                Swal.fire("Oops!", "Invalid email or password.", "error");
            });
        });
    }).fail(function(error) {
        if (error.status == 422) {
            Swal.fire("Oops!", "User already exists.", "error");
        } else {
            Swal.fire("Oops!", "An error occurred while creating the user.", "error");
        }
    });
}