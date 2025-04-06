$("#new-post").on("submit", createPost);
$(document).on("click", ".like-post", likePost);
$(document).on("click", ".dislike-post", dislikePost);
$("#update-post").on("click", updatePost);
$(".delete-post").on("click", deletePost);

function createPost(event) {
  event.preventDefault();

  $.ajax({
    url: "/posts",
    method: "POST",
    data: {
      title: $("#title").val(),
      content: $("#content").val(),
    },
  })
    .done(function () {
      window.location = "/home";
    })
    .fail(function () {
      Swal.fire( "Oops!", "An error occurred while creating the post.", "error");
    });
}

function likePost(event) {
  event.preventDefault();

  const clickedElement = $(event.target);
  const postID = clickedElement.closest("div").data("post-id");

  clickedElement.prop("disabled", true);
  $.ajax({
    url: `/posts/${postID}/like`,
    method: "POST",
  })
    .done(function () {
      const countLikes = clickedElement.next("span");
      const amountLikes = parseInt(countLikes.text());

      countLikes.text(amountLikes + 1);

      clickedElement.addClass("dislike-post");
      clickedElement.addClass("text-danger");
      clickedElement.removeClass("like-post");
    })
    .fail(function () {
      Swal.fire( "Oops!", "An error occurred while liking the post.", "error");
    })
    .always(function () {
      clickedElement.prop("disabled", false);
    });
}

function dislikePost(event) {
  event.preventDefault();

  const clickedElement = $(event.target);
  const postID = clickedElement.closest("div").data("post-id");

  clickedElement.prop("disabled", true);
  $.ajax({
    url: `/posts/${postID}/dislike`,
    method: "POST",
  })
    .done(function () {
      const countLikes = clickedElement.next("span");
      const amountLikes = parseInt(countLikes.text());

      countLikes.text(amountLikes - 1);

      clickedElement.removeClass("dislike-post");
      clickedElement.removeClass("text-danger");
      clickedElement.addClass("like-post");
    })
    .fail(function () {
        Swal.fire( "Oops!", "An error occurred while disliking the post.", "error");
    })
    .always(function () {
      clickedElement.prop("disabled", false);
    });
}

function updatePost() {
  $(this).prop("disabled", true);

  const postID = $(this).data("post-id");

  $.ajax({
    url: `/posts/${postID}`,
    method: "PUT",
    data: {
      title: $("#title").val(),
      content: $("#content").val(),
    },
  })
    .done(function () {
      Swal.fire("Success!", "Post updated successfully.", "success").then(
        function () {
          window.location = "/home";
        }
      );
    })
    .fail(function () {
        Swal.fire("Oops!", "An error occurred while updating the post.", "error");
    })
    .always(function () {
      $("#update-post").prop("disabled", false);
    });
}

function deletePost(event) {
  event.preventDefault();

  Swal.fire({
    title: "Are you sure?",
    text: "You won't be able to revert this!",
    icon: "warning",
    showCancelButton: true,
    confirmButtonColor: "#3085d6",
    cancelButtonColor: "#d33",
    confirmButtonText: "Yes, delete it!",
  }).then(function (confirm) {
    if (!confirm.value) return;

    const clickedElement = $(event.target);
    const post = clickedElement.closest("div");
    const postID = post.data("post-id");

    clickedElement.prop("disabled", true);

    $.ajax({
      url: `/posts/${postID}`,
      method: "DELETE",
    })
      .done(function () {
        post.fadeOut("slow", function () {
          $(this).remove();
        });
      })
      .fail(function () {
        Swal.fire("Oops!", "An error occurred while deleting the post.", "error");
      });
  });
}
