$('#new-post').on('submit', createPost);
$(document).on('click', '.like-post', likePost)
$(document).on('click', '.dislike-post', dislikePost)

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

        clieckedElement.addClass('dislike-post');
        clieckedElement.addClass('text-danger');
        clieckedElement.removeClass('like-post');
    }).fail(function() {
        alert("Erro ao curtir publicação")
    }).always(function() {
        clieckedElement.prop('disabled', false);
    })
}

function dislikePost(event) {
    event.preventDefault();

    const clieckedElement = $(event.target);
    const postID = clieckedElement.closest('div').data('post-id');

    clieckedElement.prop('disabled', true);
    $.ajax({
        url: `/posts/${postID}/dislike`,
        method: "POST"
    }).done(function() {
        const countLikes = clieckedElement.next('span');
        const amountLikes = parseInt(countLikes.text());

        countLikes.text(amountLikes - 1);

        clieckedElement.removeClass('dislike-post');
        clieckedElement.removeClass('text-danger');
        clieckedElement.addClass('like-post');
    }).fail(function() {
        alert("Erro ao descurtir publicação")
    }).always(function() {
        clieckedElement.prop('disabled', false);
    })
}