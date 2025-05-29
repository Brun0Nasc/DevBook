$('#unfollow').on('click', unfollow);
$('#follow').on('click', follow);

function unfollow() {
    const userId = $(this).data('user-id');
    $(this).prop('disabled', true);

    $.ajax({
        url: "/users/${userId}/unfollow",
        type: "POST",
    }).done(function() {
        window.location = `/users/${userId}`;
    }).fail(function() {
        Swal.fire("Oops...", "Something went wrong while trying to unfollow this user.", "error");
        $('#unfollow').prop('disabled', false);
    })
}

function follow() {

}