<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Cat Application</title>
    <link rel="stylesheet" href="/static/css/style.css" />
    <script src="/static/js/script.js"></script>
    <link
      href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.0.0-beta3/css/all.min.css"
      rel="stylesheet"
    />
  </head>
  <body>
    <div class="container">
      <div class="card">
        <div class="nav">
          <button id="voting-button">Voting</button>
          <button href="#" id="breeds-button">
            <i class="fa-solid fa-magnifying-glass breeds"> Breeds</i>
          </button>
          <button id="favs-button"><i class="fa-regular fa-heart favs"> Favs</i></button>
        </div>
  
        <!-- Image Section -->
        <div class="image-container">
          {{if .CatImageURL}}
          <img src="{{.CatImageURL}}" alt="Random Cat Image" class="cat-image" />
          {{else}}
          <p>No image available at the moment.</p>
          {{ end }}
        </div>
  
        <!-- Breeds Section -->
        <div id="breeds-section" style="display: none">
          <select id="breed-select">
            <!-- Options will be dynamically loaded here -->
          </select>
  
          <div id="breed-images" class="slider">
            <!-- Images will be dynamically loaded here -->
          </div>
  
          <div id="breed-details">
            <h2 id="breed-name">Breed Name</h2>
            <p id="breed-origin">Origin:</p>
            <p id="breed-id">ID:</p>
            <p id="breed-description">Description will appear here.</p>
            <a id="breed-wikipedia" target="_blank">Wikipedia</a>
          </div>
        </div>
  
        <div id="favs-section" style="display: none">
          <!-- View toggle buttons -->
          <button class="view-btn grid-btn">
            <svg width="24" height="24" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
              <rect x="3" y="3" width="7" height="7" fill="none" stroke="currentColor" stroke-width="2"/>
              <rect x="14" y="3" width="7" height="7" fill="none" stroke="currentColor" stroke-width="2"/>
              <rect x="3" y="14" width="7" height="7" fill="none" stroke="currentColor" stroke-width="2"/>
              <rect x="14" y="14" width="7" height="7" fill="none" stroke="currentColor" stroke-width="2"/>
            </svg>
          </button>
        
          <button class="view-btn bar-btn">
            <svg width="24" height="24" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
              <line x1="3" y1="6" x2="21" y2="6" stroke="currentColor" stroke-width="2"/>
              <line x1="3" y1="12" x2="21" y2="12" stroke="currentColor" stroke-width="2"/>
              <line x1="3" y1="18" x2="21" y2="18" stroke="currentColor" stroke-width="2"/>
            </svg>
          </button>
        
          <div id="favs-gallery" class="favs-gallery">
            <!-- Favourites images will be dynamically loaded here -->
          </div>
        </div>
  
        <!-- Footer Section -->
        <div class="footer nav">
          <button href="#" class="favs-down">
            <i class="fa-regular fa-heart"></i>
          </button>
          <button href="#" class="thumbs-up">
            <i class="fa-regular fa-thumbs-up"></i>
          </button>
          <button href="#" class="thumbs-down">
            <i class="fa-regular fa-thumbs-down"></i>
          </button>
        </div>
      </div>

    </div>
   
  </body>
</html>
