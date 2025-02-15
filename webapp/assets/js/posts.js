$('#new-post').on('submit', CreatePost);

function CreatePost(event) {
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