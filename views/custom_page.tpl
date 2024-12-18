<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Random Cat Image</title>
  <link rel="stylesheet" href="static/css/style.css">
</head>
<body>
  <div class="container">
    <div class="nav">
      <a href="#">Voting</a>
      <a href="#">Breeds</a>
      <a href="#">Favs</a>
    </div>
    <div class="image-container">
      <!-- Display the random cat image -->
      {{if .CatImageURL}}
        <img src="{{.CatImageURL}}" alt="Random Cat Image" class="cat-image">
        <!-- <p>{{.CatImageURL}}</p> -->
      {{else}}
        <p>No image available at the moment.</p>
      {{end}}
    </div>
  </div>
</body>
</html>
