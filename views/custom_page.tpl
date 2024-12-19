<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Random Cat Image</title>
  <link rel="stylesheet" href="static/css/style.css">
  <link href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.0.0-beta3/css/all.min.css" rel="stylesheet">
</head>
<body>
  <div class="container">
    <div class="nav">
      <a href="#">Voting</a>
      <a href="#"><i class="fa-solid fa-magnifying-glass breeds">Breeds</i></a>
      <a href="#"><i class="fa-regular fa-heart favs">Favs</i></a>
    </div>
    <div class="image-container">
      {{if .CatImageURL}}
        <img src="{{.CatImageURL}}" alt="Random Cat Image" class="cat-image">
      {{else}}
        <p>No image available at the moment.</p>
      {{end}}
    </div>
    <div class="footer nav">
      <a href="#" class="favs-down"><i class="fa-regular fa-heart"></i></a>
      <a href="#" class="thumbs-up"><i class="fa-regular fa-thumbs-up"></i></a>
      <a href="#" class="thumbs-down"><i class="fa-regular fa-thumbs-down"></i></a>
    </div>
  </div>

  <script>
    document.addEventListener("DOMContentLoaded", () => {
      const buttons = document.querySelectorAll(".footer a");

      buttons.forEach(button => {
        button.addEventListener("click", event => {
          event.preventDefault();

          fetch("/custom", {
            headers: {
              "X-Requested-With": "XMLHttpRequest"
            }
          })
            .then(response => {
              if (!response.ok) {
                throw new Error("Failed to fetch a new image");
              }
              return response.json();
            })
            .then(data => {
              if (!data.url) {
                throw new Error("Invalid response data");
              }
              const img = document.querySelector(".cat-image");
              if (img) {
                img.src = data.url;
              } else {
                const container = document.querySelector(".image-container");
                container.innerHTML = `<img src="${data.url}" alt="Random Cat Image" class="cat-image">`;
              }
            })
            .catch(error => {
              console.error("Error:", error);
              alert("An error occurred while fetching a new image.");
            });
        });
      });
    });
  </script>
</body>
</html>
