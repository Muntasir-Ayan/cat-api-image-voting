<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Cat Application</title>
    <link rel="stylesheet" href="static/css/style.css" />
    <link
      href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.0.0-beta3/css/all.min.css"
      rel="stylesheet"
    />
    <style>
      #breed-images {
        display: flex;
        overflow-x: scroll;
        gap: 10px;
      }
      #breed-images .slider-image {
        max-width: 200px;
        height: auto;
        border-radius: 8px;
      }
    </style>
  </head>
  <body>
    <div class="container">
      <div class="nav">
        <a href="http://localhost:8080/custom">Voting</a>
        <a href="#" id="breeds-button"
          ><i class="fa-solid fa-magnifying-glass breeds"> Breeds</i></a
        >
        <a href="#"><i class="fa-regular fa-heart favs">Favs</i></a>
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
        <div id="breed-details">
          <h2 id="breed-name">Breed Name</h2>
          <p id="breed-origin">Origin:</p>
          <p id="breed-id">ID:</p>
          <p id="breed-description">Description will appear here.</p>
          <a id="breed-wikipedia" target="_blank">Wikipedia</a>
        </div>
        <div id="breed-images" class="slider">
          <!-- Images will be dynamically loaded here -->
        </div>
      </div>

      <!-- Footer Section -->
      <div class="footer nav">
        <a href="#" class="favs-down"><i class="fa-regular fa-heart"></i></a>
        <a href="#" class="thumbs-up"
          ><i class="fa-regular fa-thumbs-up"></i
        ></a>
        <a href="#" class="thumbs-down"
          ><i class="fa-regular fa-thumbs-down"></i
        ></a>
      </div>
    </div>

    <!-- JavaScript Section -->
    <script>
      document.addEventListener("DOMContentLoaded", () => {
        const breedsButton = document.querySelector("#breeds-button");
        const breedsSection = document.querySelector("#breeds-section");
        const breedSelect = document.getElementById("breed-select");
        const breedDetails = document.getElementById("breed-details");
        const breedImages = document.getElementById("breed-images");

        // Fetch and display a new random cat image
        breedsButton.addEventListener("click", (event) => {
          event.preventDefault();
          breedsSection.style.display =
            breedsSection.style.display === "none" ? "block" : "none";
          loadBreeds();
        });

        // Load breeds into the dropdown
        const loadBreeds = async () => {
          const response = await fetch("/custom/breeds");
          const breeds = await response.json();

          breedSelect.innerHTML = breeds
            .map(
              (breed) => `<option value="${breed.id}">${breed.name}</option>`
            )
            .join("");
          breedSelect.value = "abys"; // Set initial breed to Abyssinian (or the ID of the breed you want as default)

          loadBreedDetails(breedSelect.value); // Load initial breed details
        };

        // Load breed details including images and information
        // Load breed details including images and information
        const loadBreedDetails = async (breedID) => {
  const response = await fetch(`/custom/breed_images?breed_id=${breedID}`);
  const images = await response.json();

  if (images && images.length > 0) {
    // Extract breed information from the first image's breed data
    const breedInfo = images[0].breeds[0] || {};
    
    // Update breed details section
    breedDetails.querySelector("#breed-name").textContent = breedInfo.name || "N/A";
    breedDetails.querySelector("#breed-origin").textContent = `Origin: ${breedInfo.origin || "Unknown"}`;
    breedDetails.querySelector("#breed-id").textContent = `ID: ${breedID}`;
    breedDetails.querySelector("#breed-description").textContent = breedInfo.description || "No description available.";

    // Handle the Wikipedia link
    const wikiLink = breedInfo.wikipedia || "#"; // Default to "#" if the link is missing
    const wikiText = breedInfo.wikipedia ? "Wikipedia" : "No Wikipedia Link";
    
    breedDetails.querySelector("#breed-wikipedia").href = wikiLink;
    breedDetails.querySelector("#breed-wikipedia").textContent = wikiText;

    // Display breed images in the sliding window
    breedImages.innerHTML = images
      .map(
        (img) => `<img src="${img.url}" alt="${breedInfo.name}" class="slider-image">`
      )
      .join("");
  } else {
    breedDetails.querySelector("#breed-wikipedia").href = "#";
    breedDetails.querySelector("#breed-wikipedia").textContent = "No images found for this breed";
  }
};

        // Handle change in breed selection
        breedSelect.addEventListener("change", (event) => {
          loadBreedDetails(event.target.value);
        });
      });
    </script>
  </body>
</html>
