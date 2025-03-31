$('#new-post').on('submit', createPost);
$(document).on('click', '.like-post', likePost)
$(document).on('click', '.dislike-post', dislikePost)
$('#update-post').on('click', updatePost)

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

    const clickedElement = $(event.target);
    const postID = clickedElement.closest('div').data('post-id');

    clickedElement.prop('disabled', true);
    $.ajax({
        url: `/posts/${postID}/like`,
        method: "POST"
    }).done(function() {
        const countLikes = clickedElement.next('span');
        const amountLikes = parseInt(countLikes.text());

        countLikes.text(amountLikes + 1);

        clickedElement.addClass('dislike-post');
        clickedElement.addClass('text-danger');
        clickedElement.removeClass('like-post');
    }).fail(function() {
        alert("Erro ao curtir publicação")
    }).always(function() {
        clickedElement.prop('disabled', false);
    })
}

function dislikePost(event) {
    event.preventDefault();

    const clickedElement = $(event.target);
    const postID = clickedElement.closest('div').data('post-id');

    clickedElement.prop('disabled', true);
    $.ajax({
        url: `/posts/${postID}/dislike`,
        method: "POST"
    }).done(function() {
        const countLikes = clickedElement.next('span');
        const amountLikes = parseInt(countLikes.text());

        countLikes.text(amountLikes - 1);

        clickedElement.removeClass('dislike-post');
        clickedElement.removeClass('text-danger');
        clickedElement.addClass('like-post');
    }).fail(function() {
        alert("Erro ao descurtir publicação")
    }).always(function() {
        clickedElement.prop('disabled', false);
    })
}

function updatePost() {
    $(this).prop('disabled', true);

    const postID = $(this).data('post-id');
    
    $.ajax({
        url: `/posts/${postID}`,
        method: "PUT",
        data: {
            title: $('#title').val(),
            content: $('#content').val()
        }
    }).done(function() {
        alert("Publicação editada com sucesso")
    }).fail(function() {
        alert("Erro ao editar publicação")
    }).always(function() {
        $('#update-post').prop('disabled', false)
    })
}