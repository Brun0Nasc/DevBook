$('#new-post').on('submit', createPost);
$('.like-post').on('click', likePost);

function createPost(event) {
    event.preventDefault();
    
    $.ajax({
        url: '/posts',
        method: 'POST',
        data: {
            title: $('#title').val(),
            content: $('#content').val()
        }
    }).done(function() {
        window.location = '/home';
    }).fail(function() {
        alert("Failed to create post");
    })
}

function likePost(event) {
    event.preventDefault();

    const clieckedElement = $(event.target);
    const postID = clieckedElement.closest('div').data('post-id');

    clieckedElement.prop('disabled', true);
    $.ajax({
        url: `/posts/${postID}/like`,
        method: "POST"
    }).done(function() {
        const countLikes = clieckedElement.next('span');
        const amountLikes = parseInt(countLikes.text());

        countLikes.text(amountLikes + 1);
    }).fail(function() {
        alert("Erro ao curtir publicação")
    }).always(function() {
        clieckedElement.prop('disabled', false);
    })
}